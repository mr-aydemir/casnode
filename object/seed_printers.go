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

// PrinterBrandModel defines one supported printer model: a pinned starter topic
// created inside its brand subcategory on the first run.
type PrinterBrandModel struct {
	Name string // topic title, e.g. "Creality K1 Max"
	Raw  string // markdown body of the pinned starter topic
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
// beta (upload-only in the MakerSync app) brands.
var printerBrands = []*PrinterBrand{
	{
		Name: "Creality", Slug: "creality", Color: "F97316", Position: 1,
		Description: "Creality K1, K1C and K1 Max — verified full remote control in MakerSync.",
		Models: []*PrinterBrandModel{
			{
				Name: "Creality K1",
				Raw: `The **Creality K1** is one of the fully verified printers in MakerSync: remote control, monitoring and print management all work out of the box.

**Getting started**
- Connect MakerSync to your K1 over the local network and start printing from your phone or desktop.
- Share your favorite Klipper tweaks and profile settings here.

**Support status:** ✅ Verified — full remote control in MakerSync.

Ask questions, share tuning profiles and post your K1 prints in this topic!`,
			},
			{
				Name: "Creality K1C",
				Raw: `The **Creality K1C** brings carbon-fiber-ready printing to the K1 family and is fully verified in MakerSync.

**Getting started**
- Add the printer in MakerSync for remote control, live monitoring and print queueing.
- Looking for hardened nozzle and CF-filament settings? Post them here.

**Support status:** ✅ Verified — full remote control in MakerSync.

Share your experience, configs and prints!`,
			},
			{
				Name: "Creality K1 Max",
				Raw: `The **Creality K1 Max** is the large-format member of the K1 series with Lidar calibration, fully verified in MakerSync.

**Getting started**
- Connect it in MakerSync for full remote control, camera monitoring and filament management.
- Great place to discuss chamber heating, big-print adhesion and K1 Max specific tuning.

**Support status:** ✅ Verified — full remote control in MakerSync.

Show off your large prints and share your setup!`,
			},
		},
	},
	{
		Name: "Snapmaker", Slug: "snapmaker", Color: "0EA5E9", Position: 2,
		Description: "Snapmaker U1 and High Combo — verified full remote control in MakerSync.",
		Models: []*PrinterBrandModel{
			{
				Name: "Snapmaker U1",
				Raw: `The **Snapmaker U1** is a fully verified printer in MakerSync with complete remote control support.

**Getting started**
- Pair it with MakerSync for remote printing, monitoring and status notifications.
- Discuss U1-specific settings, materials and upgrades here.

**Support status:** ✅ Verified — full remote control in MakerSync.

Share your prints, tips and questions!`,
			},
			{
				Name: "Snapmaker High Combo",
				Raw: `The **Snapmaker High Combo** configuration is fully verified in MakerSync.

**Getting started**
- Add it in MakerSync to manage prints remotely alongside your other machines.
- A good place for module, enclosure and multi-tool discussions.

**Support status:** ✅ Verified — full remote control in MakerSync.

Post your setups and ask anything!`,
			},
		},
	},
	{
		Name: "Bambu Lab", Slug: "bambu-lab", Color: "22C55E", Position: 3,
		Description: "Bambu Lab P1P, P1S, X1C, A1 and A1 mini — community beta (upload-only) in MakerSync.",
		Models: []*PrinterBrandModel{
			{
				Name: "Bambu Lab P1P / P1S",
				Raw: `The **Bambu Lab P1P and P1S** are part of the MakerSync community beta program.

**Support status:** 🧪 Community beta — MakerSync runs in **upload-only mode**: you can transfer files to the printer, but printing must be started from the machine itself, for safety.

**In this topic**
- Connect your P1P/P1S to MakerSync (file upload).
- Share LAN mode tips, firmware notes and print results.

Want full remote control for Bambu Lab? Vote and share your use case here!`,
			},
			{
				Name: "Bambu Lab X1C",
				Raw: `The **Bambu Lab X1 Carbon** is part of the MakerSync community beta program.

**Support status:** 🧪 Community beta — upload-only in MakerSync: files are transferred to the printer, and printing is started on the machine.

**In this topic**
- X1C file upload workflows with MakerSync.
- AMS, filament and chamber tuning discussions.

Tell us how you use your X1C and what full remote control would mean for you!`,
			},
			{
				Name: "Bambu Lab A1 / A1 mini",
				Raw: `The **Bambu Lab A1 and A1 mini** are part of the MakerSync community beta program.

**Support status:** 🧪 Community beta — upload-only in MakerSync: transfer files, then start the print on the printer.

**In this topic**
- A1 family setup with MakerSync.
- Bed-slinger tuning, AMS lite configurations and print quality tips.

Share your benchies and ask questions!`,
			},
		},
	},
	{
		Name: "Voron", Slug: "voron", Color: "EF4444", Position: 4,
		Description: "Voron 2.4 and Trident (Klipper / Moonraker) — community beta (upload-only) in MakerSync.",
		Models: []*PrinterBrandModel{
			{
				Name: "Voron 2.4",
				Raw: `The **Voron 2.4** runs Klipper with Moonraker, and is part of the MakerSync community beta.

**Support status:** 🧪 Community beta — upload-only in MakerSync: send files to the printer via Moonraker, start prints on the machine.

**In this topic**
- Moonraker upload configuration with MakerSync.
- QGL, input shaping and chamber tuning for the 2.4.

Share your build threads and mods!`,
			},
			{
				Name: "Voron Trident",
				Raw: `The **Voron Trident** runs Klipper with Moonraker, and is part of the MakerSync community beta.

**Support status:** 🧪 Community beta — upload-only in MakerSync: transfer files via Moonraker, then start printing from the machine.

**In this topic**
- Trident-specific setup with MakerSync.
- Bed leveling, gantry alignment and mod discussions.

Post your build and tuning results!`,
			},
		},
	},
	{
		Name: "Elegoo", Slug: "elegoo", Color: "3B82F6", Position: 5,
		Description: "Elegoo Neptune 4 and Neptune 4 Pro — community beta (upload-only) in MakerSync.",
		Models: []*PrinterBrandModel{
			{
				Name: "Elegoo Neptune 4 / Neptune 4 Pro",
				Raw: `The **Elegoo Neptune 4 and Neptune 4 Pro** run Klipper and are part of the MakerSync community beta.

**Support status:** 🧪 Community beta — upload-only in MakerSync: files are transferred to the printer; printing starts on the machine.

**In this topic**
- Neptune 4 setup with MakerSync file upload.
- Klipper config tuning, Z-offset and first-layer help.

Share your configs and prints!`,
			},
		},
	},
	{
		Name: "Anycubic", Slug: "anycubic", Color: "A855F7", Position: 6,
		Description: "Anycubic Kobra and other printers — community beta (upload-only) in MakerSync.",
		Models: []*PrinterBrandModel{
			{
				Name: "Anycubic Printers (Kobra series)",
				Raw: `**Anycubic** printers, including the Kobra series, are part of the MakerSync community beta program.

**Support status:** 🧪 Community beta — upload-only in MakerSync: transfer files to the printer, start prints on the machine.

**In this topic**
- Which Anycubic machine do you run? Introduce your setup.
- Firmware, leveling and tuning discussions.

Tell us which models you want supported next!`,
			},
		},
	},
	{
		Name: "Prusa", Slug: "prusa", Color: "F59E0B", Position: 7,
		Description: "Prusa MK4, MK3S+ and MINI+ — community beta (upload-only) in MakerSync.",
		Models: []*PrinterBrandModel{
			{
				Name: "Prusa MK4 / MK3S+ / MINI+",
				Raw: `**Prusa Research** printers — MK4, MK3S+ and MINI+ — are part of the MakerSync community beta.

**Support status:** 🧪 Community beta — upload-only in MakerSync: upload files to your Prusa, then start the print on the printer or via PrusaConnect.

**In this topic**
- PrusaConnect workflows with MakerSync.
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

// getSeedTopic returns the starter topic for a model name inside a brand category,
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

// InitPrinterCategories seeds the printer brand subcategories under the hardware
// category and a pinned starter topic per supported model. It is idempotent:
// existing slugs are skipped, so re-running (e.g. on every deploy) is safe.
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

		for _, model := range brand.Models {
			existing, err := getSeedTopic(model.Name, category.Id)
			if err != nil {
				panic(err)
			}
			if existing != nil {
				continue
			}
			_, err = createSeedTopic(system, model.Name, model.Raw, category.Id)
			if err != nil {
				fmt.Printf("[seed] skipping topic %q: %v\n", model.Name, err)
				continue
			}
			createdTopics++
		}
	}

	if createdCategories > 0 || createdTopics > 0 {
		err = RefreshCategoryCounts()
		if err != nil {
			panic(err)
		}
		fmt.Printf("[seed] printer categories: created %d brand subcategories and %d pinned model topics\n", createdCategories, createdTopics)
	}
}
