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
	"strconv"
	"time"
)

const SystemUsername = "system"

// InitData fills an empty database with what a new forum starts with: the system user, the default categories
// and a pinned welcome topic.
func InitData() {
	count, err := adapter.Engine.Count(&Category{})
	if err != nil {
		panic(err)
	}
	if count > 0 {
		return
	}

	system, err := GetUserByUsername(SystemUsername)
	if err != nil {
		panic(err)
	}
	if system == nil {
		system = &User{Username: SystemUsername, Name: "system", Admin: true, TrustLevel: 4, LastSeenAt: time.Now()}
		_, err = adapter.Engine.Insert(system)
		if err != nil {
			panic(err)
		}
	}

	categories := []*Category{
		{Name: "General", Slug: "general", Color: "3B82F6", TextColor: "FFFFFF", Position: 1,
			Description: "General discussions about 3D printing, making, and technology."},
		{Name: "3D Printers & Hardware", Slug: "hardware", Color: "10B981", TextColor: "FFFFFF", Position: 2,
			Description: "Hardware setups, Creality K1, Snapmaker U1, Bambu Lab, Voron, modifications, and troubleshooting."},
		{Name: "Slicers & Software", Slug: "software", Color: "8B5CF6", TextColor: "FFFFFF", Position: 3,
			Description: "OrcaSlicer, Klipper, Moonraker, OctoPrint, CAD, and slicing tips."},
		{Name: "Models & Showcases", Slug: "showcase", Color: "EC4899", TextColor: "FFFFFF", Position: 4,
			Description: "Share your 3D prints, custom designs, and creative maker projects."},
		{Name: "Site Feedback & Ideas", Slug: "site-feedback", Color: "64748B", TextColor: "FFFFFF", Position: 5,
			Description: "Discussion about MakerSpell and MakerSync apps, features, and roadmap."},
		{Name: "Staff", Slug: "staff", Color: "EF4444", TextColor: "FFFFFF", Position: 6, ReadRestricted: true,
			Description: "Private category for staff discussions. Topics are only visible to admins and moderators."},
	}
	for _, category := range categories {
		_, err = adapter.Engine.Insert(category)
		if err != nil {
			panic(err)
		}
	}
	invalidateCategories()
	err = SetSiteSetting("uncategorized_category_id", strconv.FormatInt(categories[0].Id, 10))
	if err != nil {
		panic(err)
	}
	err = SetSiteSetting("staff_category_id", strconv.FormatInt(categories[5].Id, 10))
	if err != nil {
		panic(err)
	}

	raw := fmt.Sprintf(`Welcome to the **MakerSpell Community**! 🚀

This is the gathering place for 3D printing enthusiasts, makers, hardware hackers, and creators worldwide.

### 🌟 What can you do here?
- **Showcase Your Prints:** Share your latest 3D prints and models in [Models & Showcases](/c/showcase/%d).
- **Hardware & Troubleshooting:** Ask questions and share tips for your Klipper, Creality, Snapmaker, Bambu, or Voron machines in [3D Printers & Hardware](/c/hardware/%d).
- **Slicing & Optimization:** Discuss profiles, slicer settings, and OrcaSlicer integration in [Slicers & Software](/c/software/%d).
- **MakerSpell & MakerSync Feedback:** Give feedback and feature requests in [Site Feedback](/c/site-feedback/%d).

> Visit [MakerSpell.com](https://www.makerspell.com) to search over 26 platforms for 3D models and manage your printers!`, categories[3].Id, categories[1].Id, categories[2].Id, categories[4].Id)
	topic, err := CreateTopic(system, &NewTopic{Title: "Welcome to MakerSpell Community!", Raw: raw, CategoryId: categories[0].Id})
	if err != nil {
		panic(err)
	}
	topic.PinnedAt = time.Now()
	topic.PinnedGlobally = true
	_, err = adapter.Engine.ID(topic.Id).Cols("pinned_at", "pinned_globally").Update(topic)
	if err != nil {
		panic(err)
	}
}
