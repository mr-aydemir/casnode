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
var printerBrands = []*PrinterBrand{
	{
		Name: "Creality", Slug: "creality", Color: "F97316", Position: 1,
		Description: "Creality K1, K1C and K1 Max — verified full remote control in MakerSpell.",
		Models: []*PrinterBrandModel{
			{
				Name: "Creality K1", Slug: "k1",
				Description: "Creality K1 — verified full remote control in MakerSpell.",
				Raw: `The **Creality K1** is one of the fully verified printers in MakerSpell: remote control, monitoring and print management all work out of the box.

**Getting started**
- Connect MakerSpell to your K1 over the local network and start printing from your phone or desktop.
- Share your favorite Klipper tweaks and profile settings here.

**Support status:** ✅ Verified — full remote control in MakerSpell.

Ask questions, share tuning profiles and post your K1 prints in this topic!`,
			},
			{
				Name: "Creality K1C", Slug: "k1c",
				Description: "Creality K1C — verified full remote control in MakerSpell.",
				Raw: `The **Creality K1C** brings carbon-fiber-ready printing to the K1 family and is fully verified in MakerSpell.

**Getting started**
- Add the printer in MakerSpell for remote control, live monitoring and print queueing.
- Looking for hardened nozzle and CF-filament settings? Post them here.

**Support status:** ✅ Verified — full remote control in MakerSpell.

Share your experience, configs and prints!`,
			},
			{
				Name: "Creality K1 Max", Slug: "k1-max",
				Description: "Creality K1 Max — verified full remote control in MakerSpell.",
				Raw: `The **Creality K1 Max** is the large-format member of the K1 series with Lidar calibration, fully verified in MakerSpell.

**Getting started**
- Connect it in MakerSpell for full remote control, camera monitoring and filament management.
- Great place to discuss chamber heating, big-print adhesion and K1 Max specific tuning.

**Support status:** ✅ Verified — full remote control in MakerSpell.

Show off your large prints and share your setup!`,
			},
		},
	},
	{
		Name: "Snapmaker", Slug: "snapmaker", Color: "0EA5E9", Position: 2,
		Description: "Snapmaker U1 and High Combo — verified full remote control in MakerSpell.",
		Models: []*PrinterBrandModel{
			{
				Name: "Snapmaker U1", Slug: "u1",
				Description: "Snapmaker U1 — verified full remote control in MakerSpell.",
				Raw: `The **Snapmaker U1** is a fully verified printer in MakerSpell with complete remote control support.

**Getting started**
- Pair it with MakerSpell for remote printing, monitoring and status notifications.
- Discuss U1-specific settings, materials and upgrades here.

**Support status:** ✅ Verified — full remote control in MakerSpell.

Share your prints, tips and questions!`,
			},
			{
				Name: "Snapmaker High Combo", Slug: "high-combo",
				Description: "Snapmaker High Combo — verified full remote control in MakerSpell.",
				Raw: `The **Snapmaker High Combo** configuration is fully verified in MakerSpell.

**Getting started**
- Add it in MakerSpell to manage prints remotely alongside your other machines.
- A good place for module, enclosure and multi-tool discussions.

**Support status:** ✅ Verified — full remote control in MakerSpell.

Post your setups and ask anything!`,
			},
		},
	},
	{
		Name: "Bambu Lab", Slug: "bambu-lab", Color: "22C55E", Position: 3,
		Description: "Bambu Lab P1P, P1S, X1C, A1 and A1 mini — community beta (upload-only) in MakerSpell.",
		Models: []*PrinterBrandModel{
			{
				Name: "Bambu Lab P1P / P1S", Slug: "p1-p1s",
				Description: "Bambu Lab P1P and P1S — community beta (upload-only) in MakerSpell.",
				Raw: `The **Bambu Lab P1P and P1S** are part of the MakerSpell community beta program.

**Support status:** 🧪 Community beta — MakerSpell runs in **upload-only mode**: you can transfer files to the printer, but printing must be started from the machine itself, for safety.

**In this topic**
- Connect your P1P/P1S to MakerSpell (file upload).
- Share LAN mode tips, firmware notes and print results.

Want full remote control for Bambu Lab? Vote and share your use case here!`,
			},
			{
				Name: "Bambu Lab X1C", Slug: "x1c",
				Description: "Bambu Lab X1 Carbon — community beta (upload-only) in MakerSpell.",
				Raw: `The **Bambu Lab X1 Carbon** is part of the MakerSpell community beta program.

**Support status:** 🧪 Community beta — upload-only in MakerSpell: files are transferred to the printer, and printing is started on the machine.

**In this topic**
- X1C file upload workflows with MakerSpell.
- AMS, filament and chamber tuning discussions.

Tell us how you use your X1C and what full remote control would mean for you!`,
			},
			{
				Name: "Bambu Lab A1 / A1 mini", Slug: "a1-a1-mini",
				Description: "Bambu Lab A1 and A1 mini — community beta (upload-only) in MakerSpell.",
				Raw: `The **Bambu Lab A1 and A1 mini** are part of the MakerSpell community beta program.

**Support status:** 🧪 Community beta — upload-only in MakerSpell: transfer files, then start the print on the printer.

**In this topic**
- A1 family setup with MakerSpell.
- Bed-slinger tuning, AMS lite configurations and print quality tips.

Share your benchies and ask questions!`,
			},
		},
	},
	{
		Name: "Voron", Slug: "voron", Color: "EF4444", Position: 4,
		Description: "Voron 2.4 and Trident (Klipper / Moonraker) — community beta (upload-only) in MakerSpell.",
		Models: []*PrinterBrandModel{
			{
				Name: "Voron 2.4", Slug: "voron-2-4",
				Description: "Voron 2.4 (Klipper / Moonraker) — community beta (upload-only) in MakerSpell.",
				Raw: `The **Voron 2.4** runs Klipper with Moonraker, and is part of the MakerSpell community beta.

**Support status:** 🧪 Community beta — upload-only in MakerSpell: send files to the printer via Moonraker, start prints on the machine.

**In this topic**
- Moonraker upload configuration with MakerSpell.
- QGL, input shaping and chamber tuning for the 2.4.

Share your build threads and mods!`,
			},
			{
				Name: "Voron Trident", Slug: "voron-trident",
				Description: "Voron Trident (Klipper / Moonraker) — community beta (upload-only) in MakerSpell.",
				Raw: `The **Voron Trident** runs Klipper with Moonraker, and is part of the MakerSpell community beta.

**Support status:** 🧪 Community beta — upload-only in MakerSpell: transfer files via Moonraker, then start printing from the machine.

**In this topic**
- Trident-specific setup with MakerSpell.
- Bed leveling, gantry alignment and mod discussions.

Post your build and tuning results!`,
			},
		},
	},
	{
		Name: "Elegoo", Slug: "elegoo", Color: "3B82F6", Position: 5,
		Description: "Elegoo Neptune 4 and Neptune 4 Pro — community beta (upload-only) in MakerSpell.",
		Models: []*PrinterBrandModel{
			{
				Name: "Elegoo Neptune 4 / Neptune 4 Pro", Slug: "neptune-4",
				Description: "Elegoo Neptune 4 and Neptune 4 Pro — community beta (upload-only) in MakerSpell.",
				Raw: `The **Elegoo Neptune 4 and Neptune 4 Pro** run Klipper and are part of the MakerSpell community beta.

**Support status:** 🧪 Community beta — upload-only in MakerSpell: files are transferred to the printer; printing starts on the machine.

**In this topic**
- Neptune 4 setup with MakerSpell file upload.
- Klipper config tuning, Z-offset and first-layer help.

Share your configs and prints!`,
			},
		},
	},
	{
		Name: "Anycubic", Slug: "anycubic", Color: "A855F7", Position: 6,
		Description: "Anycubic Kobra and other printers — community beta (upload-only) in MakerSpell.",
		Models: []*PrinterBrandModel{
			{
				Name: "Anycubic Kobra series", Slug: "kobra-series",
				Description: "Anycubic Kobra series — community beta (upload-only) in MakerSpell.",
				Raw: `**Anycubic** printers, including the Kobra series, are part of the MakerSpell community beta program.

**Support status:** 🧪 Community beta — upload-only in MakerSpell: transfer files to the printer, start prints on the machine.

**In this topic**
- Which Anycubic machine do you run? Introduce your setup.
- Firmware, leveling and tuning discussions.

Tell us which models you want supported next!`,
			},
		},
	},
	{
		Name: "Prusa", Slug: "prusa", Color: "F59E0B", Position: 7,
		Description: "Prusa MK4, MK3S+ and MINI+ — community beta (upload-only) in MakerSpell.",
		Models: []*PrinterBrandModel{
			{
				Name: "Prusa MK4 / MK3S+ / MINI+", Slug: "prusa-mk4-mk3s-mini",
				Description: "Prusa MK4, MK3S+ and MINI+ — community beta (upload-only) in MakerSpell.",
				Raw: `**Prusa Research** printers — MK4, MK3S+ and MINI+ — are part of the MakerSpell community beta.

**Support status:** 🧪 Community beta — upload-only in MakerSpell: upload files to your Prusa, then start the print on the printer or via PrusaConnect.

**In this topic**
- PrusaConnect workflows with MakerSpell.
- Filament profiles, first-layer calibration and reliability tips.

Share your experience with your Prusa!`,
			},
		},
	},
}

// createSeedTopic inserts a pinned starter topic for the seeder. It deliberately
// bypasses CreateTopic's user-facing validation (min_topic_title_length etc.),
// because model names such as "Creality K1" are shorter than the site's minimum
// title length while still being the intended topic titles. Panicking there would
// crash the whole forum on every boot.
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
