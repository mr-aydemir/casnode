// Copyright 2026 The Casnode Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package object

import (
	"fmt"
	"time"
)

// PrinterBrandModel defines one supported printer model: a third-level category
// (e.g. hardware -> creality -> k1) with a pinned starter topic inside it.
type PrinterBrandModel struct {
	Name        string // model category name / starter topic title, e.g. "Creality K1"
	Slug        string // model category slug, globally unique, e.g. "k1"
	Description string
	Raw         string // markdown body of the pinned starter topic
}

// PrinterBrand defines one printer brand subcategory under the hardware category.
type PrinterBrand struct {
	Name        string
	Slug        string
	Description string
	Color       string
	Position    int
	Models      []*PrinterBrandModel
}

// printerBrands is the MakerSpell printer support matrix (see MakerSpell-Vault
// 02_YAZICI_GUVENLIK_VE_BETA_TESTLERI.md). Verified brands first, then community
// beta (upload-only in the MakerSpell app) brands.
// printerBrands is populated by init() in seed_printers_data.go
var printerBrands []*PrinterBrand

func createSeedTopic(user *User, title, raw string, categoryId int64) (*Topic, error) {
	cooked := CookMarkdown(raw)
	now := time.Now()
	topic := &Topic{
		Title:             title,
		Slug:              Slugify(title),
		UserId:            user.Id,
		Archetype:         ArchetypeRegular,
		Visible:           true,
		CategoryId:        categoryId,
		BumpedAt:          now,
		LastPostedAt:      now,
		LastPostUserId:    user.Id,
		PostsCount:        1,
		HighestPostNumber: 1,
		ParticipantCount:  1,
		PosterIds:         []int64{user.Id},
		WordCount:         countWords(raw),
		Excerpt:           cooked.Excerpt,
		ImageUrl:          cooked.ImageUrl,
		PinnedAt:          now,
		PinnedGlobally:    false,
	}
	if _, err := adapter.Engine.Insert(topic); err != nil {
		return nil, err
	}

	post := &Post{
		TopicId:    topic.Id,
		PostNumber: 1,
		UserId:     user.Id,
		Raw:        raw,
		Cooked:     cooked.Html,
		PostType:   PostTypeRegular,
		Version:    1,
	}
	if _, err := adapter.Engine.Insert(post); err != nil {
		return nil, err
	}

	incrCategory(topic.CategoryId, "topic_count", 1)
	incrCategory(topic.CategoryId, "post_count", 1)
	incrUser(user.Id, "topic_count", 1)
	incrUser(user.Id, "post_count", 1)
	_, _ = adapter.Engine.ID(user.Id).Cols("last_posted_at").Update(&User{LastPostedAt: now})

	if err := SetTopicUser(user.Id, topic.Id, func(tu *TopicUser) {
		tu.Posted = true
		tu.LastReadPostNumber = 1
		tu.NotificationLevel = NotificationLevelWatching
	}); err != nil {
		return nil, err
	}

	PublishLatest(topic, "created")
	return topic, nil
}

// getSeedTopic returns the starter topic for a model title inside a category,
// or nil when it has not been created yet.
func getSeedTopic(title string, categoryId int64) (*Topic, error) {
	topic := &Topic{}
	has, err := adapter.Engine.Where("category_id = ?", categoryId).And("title = ?", title).Get(topic)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return topic, nil
}

// moveSeedTopic moves a starter topic from one category to another (used to
// migrate the old two-level seeds, where model topics lived directly inside the
// brand category, into their new third-level model categories).
func moveSeedTopic(topicId, targetCategoryId int64) error {
	_, err := adapter.Engine.ID(topicId).Cols("category_id").Update(&Topic{CategoryId: targetCategoryId})
	return err
}

// InitPrinterCategories seeds the three-level printer hierarchy under the
// hardware category: hardware -> brand -> model, with a pinned starter topic per
// supported model. It is idempotent: existing slugs are skipped and topics are
// only created once, so re-running (e.g. on every deploy) is safe. Deployments
// that previously seeded model topics directly inside the brand category are
// migrated: those topics are moved into their model categories.
//
// Seeding is best-effort: a failure here must never prevent the forum from
// booting (that used to take the whole site down with a panic loop).
func InitPrinterCategories() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[seed] printer category seeding aborted (non-fatal): %v\n", r)
		}
	}()
	initPrinterCategories()
}

func initPrinterCategories() {
	hardware, err := GetCategoryBySlug("hardware")
	if err != nil {
		panic(err)
	}
	if hardware == nil {
		fmt.Println("[seed] hardware category not found, skipping printer category seeding")
		return
	}

	system, err := GetUserByUsername(SystemUsername)
	if err != nil {
		panic(err)
	}
	if system == nil {
		fmt.Println("[seed] system user not found, skipping printer category seeding")
		return
	}

	createdCategories := 0
	createdTopics := 0
	movedTopics := 0

	for _, brand := range printerBrands {
		category, err := GetCategoryBySlug(brand.Slug)
		if err != nil {
			panic(err)
		}
		if category == nil {
			category = &Category{
				Name:        brand.Name,
				Slug:        brand.Slug,
				Description: brand.Description,
				Color:       brand.Color,
				TextColor:   "FFFFFF",
				ParentId:    hardware.Id,
				Position:    brand.Position,
			}
			err = AddCategory(category)
			if err != nil {
				panic(err)
			}
			createdCategories++
		}

		for i, model := range brand.Models {
			modelCategory, err := GetCategoryBySlug(model.Slug)
			if err != nil {
				panic(err)
			}
			if modelCategory == nil {
				modelCategory = &Category{
					Name:        model.Name,
					Slug:        model.Slug,
					Description: model.Description,
					Color:       brand.Color,
					TextColor:   "FFFFFF",
					ParentId:    category.Id,
					Position:    i + 1,
				}
				err = AddCategory(modelCategory)
				if err != nil {
					panic(err)
				}
				createdCategories++
			}

			// The starter topic may already live in the model category (new
			// installs), in the brand category (old two-level seeds) or not at all.
			existing, err := getSeedTopic(model.Name, modelCategory.Id)
			if err != nil {
				panic(err)
			}
			if existing == nil {
				legacy, err := getSeedTopic(model.Name, category.Id)
				if err != nil {
					panic(err)
				}
				if legacy != nil {
					err = moveSeedTopic(legacy.Id, modelCategory.Id)
					if err != nil {
						fmt.Printf("[seed] skipping topic move %q: %v\n", model.Name, err)
					} else {
						movedTopics++
					}
					continue
				}
				_, err = createSeedTopic(system, model.Name, model.Raw, modelCategory.Id)
				if err != nil {
					fmt.Printf("[seed] skipping topic %q: %v\n", model.Name, err)
					continue
				}
				createdTopics++
			}
		}
	}

	if createdCategories > 0 || createdTopics > 0 || movedTopics > 0 {
		err = RefreshCategoryCounts()
		if err != nil {
			panic(err)
		}
		fmt.Printf("[seed] printer categories: created %d categories, created %d pinned model topics, moved %d topics into model categories\n", createdCategories, createdTopics, movedTopics)
	}
}
