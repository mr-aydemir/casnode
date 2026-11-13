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

import "fmt"

func defaultCommunityBetaRaw(modelName string) string {
	return fmt.Sprintf(`The **%s** is part of the MakerSpell community beta program.

**Support status:** 🧪 Community beta — MakerSpell runs in **upload-only mode**: you can transfer files to the printer, but printing must be started from the machine itself, for hardware safety.

**In this topic**
- Share connection tips, profile tweaks and slicing configurations.
- Discuss firmware, mods and maintenance.
- Community troubleshooting and user-to-user help.

Own a **%s**? Share your setup, your prints, and vote for full remote control verification!`, modelName, modelName)
}

func defaultCommunityBetaDesc(modelName string) string {
	return fmt.Sprintf("%s — community beta (upload-only) in MakerSpell.", modelName)
}

func defaultBrandDesc(brandName string) string {
	return fmt.Sprintf("%s 3D printers — community beta (upload-only) in MakerSpell.", brandName)
}

// buildPrinterBrands constructs the full list of printer brands and models sourced
// from Orca Slicer profiles (verified hardware first, then community beta).
func buildPrinterBrands() []*PrinterBrand {
	var brands []*PrinterBrand

	// Brand: Creality (53 models)
	{
		brand := &PrinterBrand{
			Name:        "Creality",
			Slug:        "creality",
			Description: "Creality K1 series and Ender/CR printers — verified and community beta support in MakerSpell.",
			Color:       "F97316",
			Position:    1,
			Models:      make([]*PrinterBrandModel, 0, 53),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality CR-10",
			Slug:        "creality-cr-10",
			Description: defaultCommunityBetaDesc("Creality CR-10"),
			Raw:         defaultCommunityBetaRaw("Creality CR-10"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality CR-10 Max",
			Slug:        "creality-cr-10-max",
			Description: defaultCommunityBetaDesc("Creality CR-10 Max"),
			Raw:         defaultCommunityBetaRaw("Creality CR-10 Max"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality CR-10 SE",
			Slug:        "creality-cr-10-se",
			Description: defaultCommunityBetaDesc("Creality CR-10 SE"),
			Raw:         defaultCommunityBetaRaw("Creality CR-10 SE"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality CR-10 Smart",
			Slug:        "creality-cr-10-smart",
			Description: defaultCommunityBetaDesc("Creality CR-10 Smart"),
			Raw:         defaultCommunityBetaRaw("Creality CR-10 Smart"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality CR-10 Smart Pro",
			Slug:        "creality-cr-10-smart-pro",
			Description: defaultCommunityBetaDesc("Creality CR-10 Smart Pro"),
			Raw:         defaultCommunityBetaRaw("Creality CR-10 Smart Pro"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality CR-10 V2",
			Slug:        "creality-cr-10-v2",
			Description: defaultCommunityBetaDesc("Creality CR-10 V2"),
			Raw:         defaultCommunityBetaRaw("Creality CR-10 V2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality CR-10 V3",
			Slug:        "creality-cr-10-v3",
			Description: defaultCommunityBetaDesc("Creality CR-10 V3"),
			Raw:         defaultCommunityBetaRaw("Creality CR-10 V3"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality CR-10S Pro V2",
			Slug:        "creality-cr-10s-pro-v2",
			Description: defaultCommunityBetaDesc("Creality CR-10S Pro V2"),
			Raw:         defaultCommunityBetaRaw("Creality CR-10S Pro V2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality CR-20 Pro",
			Slug:        "creality-cr-20-pro",
			Description: defaultCommunityBetaDesc("Creality CR-20 Pro"),
			Raw:         defaultCommunityBetaRaw("Creality CR-20 Pro"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality CR-6 Max",
			Slug:        "creality-cr-6-max",
			Description: defaultCommunityBetaDesc("Creality CR-6 Max"),
			Raw:         defaultCommunityBetaRaw("Creality CR-6 Max"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality CR-6 SE",
			Slug:        "creality-cr-6-se",
			Description: defaultCommunityBetaDesc("Creality CR-6 SE"),
			Raw:         defaultCommunityBetaRaw("Creality CR-6 SE"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality CR-M4",
			Slug:        "creality-cr-m4",
			Description: defaultCommunityBetaDesc("Creality CR-M4"),
			Raw:         defaultCommunityBetaRaw("Creality CR-M4"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality CR-M4 SE",
			Slug:        "creality-cr-m4-se",
			Description: defaultCommunityBetaDesc("Creality CR-M4 SE"),
			Raw:         defaultCommunityBetaRaw("Creality CR-M4 SE"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Ender-2 Pro",
			Slug:        "creality-ender-2-pro",
			Description: defaultCommunityBetaDesc("Creality Ender-2 Pro"),
			Raw:         defaultCommunityBetaRaw("Creality Ender-2 Pro"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Ender-3",
			Slug:        "creality-ender-3",
			Description: defaultCommunityBetaDesc("Creality Ender-3"),
			Raw:         defaultCommunityBetaRaw("Creality Ender-3"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Ender-3 Max",
			Slug:        "creality-ender-3-max",
			Description: defaultCommunityBetaDesc("Creality Ender-3 Max"),
			Raw:         defaultCommunityBetaRaw("Creality Ender-3 Max"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Ender-3 Max Neo",
			Slug:        "creality-ender-3-max-neo",
			Description: defaultCommunityBetaDesc("Creality Ender-3 Max Neo"),
			Raw:         defaultCommunityBetaRaw("Creality Ender-3 Max Neo"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Ender-3 Neo",
			Slug:        "creality-ender-3-neo",
			Description: defaultCommunityBetaDesc("Creality Ender-3 Neo"),
			Raw:         defaultCommunityBetaRaw("Creality Ender-3 Neo"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Ender-3 Pro",
			Slug:        "creality-ender-3-pro",
			Description: defaultCommunityBetaDesc("Creality Ender-3 Pro"),
			Raw:         defaultCommunityBetaRaw("Creality Ender-3 Pro"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Ender-3 S1",
			Slug:        "creality-ender-3-s1",
			Description: defaultCommunityBetaDesc("Creality Ender-3 S1"),
			Raw:         defaultCommunityBetaRaw("Creality Ender-3 S1"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Ender-3 S1 Plus",
			Slug:        "creality-ender-3-s1-plus",
			Description: defaultCommunityBetaDesc("Creality Ender-3 S1 Plus"),
			Raw:         defaultCommunityBetaRaw("Creality Ender-3 S1 Plus"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Ender-3 S1 Pro",
			Slug:        "creality-ender-3-s1-pro",
			Description: defaultCommunityBetaDesc("Creality Ender-3 S1 Pro"),
			Raw:         defaultCommunityBetaRaw("Creality Ender-3 S1 Pro"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Ender-3 S1 Sonic",
			Slug:        "creality-ender-3-s1-sonic",
			Description: defaultCommunityBetaDesc("Creality Ender-3 S1 Sonic"),
			Raw:         defaultCommunityBetaRaw("Creality Ender-3 S1 Sonic"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Ender-3 V2",
			Slug:        "creality-ender-3-v2",
			Description: defaultCommunityBetaDesc("Creality Ender-3 V2"),
			Raw:         defaultCommunityBetaRaw("Creality Ender-3 V2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Ender-3 V2 Neo",
			Slug:        "creality-ender-3-v2-neo",
			Description: defaultCommunityBetaDesc("Creality Ender-3 V2 Neo"),
			Raw:         defaultCommunityBetaRaw("Creality Ender-3 V2 Neo"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Ender-3 V3",
			Slug:        "creality-ender-3-v3",
			Description: defaultCommunityBetaDesc("Creality Ender-3 V3"),
			Raw:         defaultCommunityBetaRaw("Creality Ender-3 V3"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Ender-3 V3 KE",
			Slug:        "creality-ender-3-v3-ke",
			Description: defaultCommunityBetaDesc("Creality Ender-3 V3 KE"),
			Raw:         defaultCommunityBetaRaw("Creality Ender-3 V3 KE"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Ender-3 V3 Plus",
			Slug:        "creality-ender-3-v3-plus",
			Description: defaultCommunityBetaDesc("Creality Ender-3 V3 Plus"),
			Raw:         defaultCommunityBetaRaw("Creality Ender-3 V3 Plus"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Ender-3 V3 SE",
			Slug:        "creality-ender-3-v3-se",
			Description: defaultCommunityBetaDesc("Creality Ender-3 V3 SE"),
			Raw:         defaultCommunityBetaRaw("Creality Ender-3 V3 SE"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Ender-3 V4",
			Slug:        "creality-ender-3-v4",
			Description: defaultCommunityBetaDesc("Creality Ender-3 V4"),
			Raw:         defaultCommunityBetaRaw("Creality Ender-3 V4"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Ender-5",
			Slug:        "creality-ender-5",
			Description: defaultCommunityBetaDesc("Creality Ender-5"),
			Raw:         defaultCommunityBetaRaw("Creality Ender-5"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Ender-5 Max",
			Slug:        "creality-ender-5-max",
			Description: defaultCommunityBetaDesc("Creality Ender-5 Max"),
			Raw:         defaultCommunityBetaRaw("Creality Ender-5 Max"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Ender-5 Plus",
			Slug:        "creality-ender-5-plus",
			Description: defaultCommunityBetaDesc("Creality Ender-5 Plus"),
			Raw:         defaultCommunityBetaRaw("Creality Ender-5 Plus"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Ender-5 Pro (2019)",
			Slug:        "creality-ender-5-pro-2019",
			Description: defaultCommunityBetaDesc("Creality Ender-5 Pro (2019)"),
			Raw:         defaultCommunityBetaRaw("Creality Ender-5 Pro (2019)"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Ender-5 S1",
			Slug:        "creality-ender-5-s1",
			Description: defaultCommunityBetaDesc("Creality Ender-5 S1"),
			Raw:         defaultCommunityBetaRaw("Creality Ender-5 S1"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Ender-5 S1 Sonic",
			Slug:        "creality-ender-5-s1-sonic",
			Description: defaultCommunityBetaDesc("Creality Ender-5 S1 Sonic"),
			Raw:         defaultCommunityBetaRaw("Creality Ender-5 S1 Sonic"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Ender-5S",
			Slug:        "creality-ender-5s",
			Description: defaultCommunityBetaDesc("Creality Ender-5S"),
			Raw:         defaultCommunityBetaRaw("Creality Ender-5S"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Ender-6",
			Slug:        "creality-ender-6",
			Description: defaultCommunityBetaDesc("Creality Ender-6"),
			Raw:         defaultCommunityBetaRaw("Creality Ender-6"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Hi",
			Slug:        "creality-hi",
			Description: defaultCommunityBetaDesc("Creality Hi"),
			Raw:         defaultCommunityBetaRaw("Creality Hi"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality K1",
			Slug:        "k1",
			Description: "Creality K1 — verified full remote control in MakerSpell.",
			Raw: `The **Creality K1** is one of the fully verified printers in MakerSpell: remote control, monitoring and print management all work out of the box.

**Getting started**
- Connect MakerSpell to your K1 over the local network and start printing from your phone or desktop.
- Share your favorite Klipper tweaks and profile settings here.

**Support status:** ✅ Verified — full remote control in MakerSpell.

Ask questions, share tuning profiles and post your K1 prints in this topic!`,
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality K1 Max",
			Slug:        "k1-max",
			Description: "Creality K1 Max — verified full remote control in MakerSpell.",
			Raw: `The **Creality K1 Max** is the large-format member of the K1 series with Lidar calibration, fully verified in MakerSpell.

**Getting started**
- Connect it in MakerSpell for full remote control, camera monitoring and filament management.
- Great place to discuss chamber heating, big-print adhesion and K1 Max specific tuning.

**Support status:** ✅ Verified — full remote control in MakerSpell.

Show off your large prints and share your setup!`,
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality K1 SE",
			Slug:        "creality-k1-se",
			Description: defaultCommunityBetaDesc("Creality K1 SE"),
			Raw:         defaultCommunityBetaRaw("Creality K1 SE"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality K1C",
			Slug:        "k1c",
			Description: "Creality K1C — verified full remote control in MakerSpell.",
			Raw: `The **Creality K1C** brings carbon-fiber-ready printing to the K1 family and is fully verified in MakerSpell.

**Getting started**
- Add the printer in MakerSpell for remote control, live monitoring and print queueing.
- Looking for hardened nozzle and CF-filament settings? Post them here.

**Support status:** ✅ Verified — full remote control in MakerSpell.

Share your experience, configs and prints!`,
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality K2",
			Slug:        "creality-k2",
			Description: defaultCommunityBetaDesc("Creality K2"),
			Raw:         defaultCommunityBetaRaw("Creality K2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality K2 Plus",
			Slug:        "creality-k2-plus",
			Description: defaultCommunityBetaDesc("Creality K2 Plus"),
			Raw:         defaultCommunityBetaRaw("Creality K2 Plus"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality K2 Pro",
			Slug:        "creality-k2-pro",
			Description: defaultCommunityBetaDesc("Creality K2 Pro"),
			Raw:         defaultCommunityBetaRaw("Creality K2 Pro"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality K2 SE",
			Slug:        "creality-k2-se",
			Description: defaultCommunityBetaDesc("Creality K2 SE"),
			Raw:         defaultCommunityBetaRaw("Creality K2 SE"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Sermoon D3",
			Slug:        "creality-sermoon-d3",
			Description: defaultCommunityBetaDesc("Creality Sermoon D3"),
			Raw:         defaultCommunityBetaRaw("Creality Sermoon D3"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Sermoon D3 Pro",
			Slug:        "creality-sermoon-d3-pro",
			Description: defaultCommunityBetaDesc("Creality Sermoon D3 Pro"),
			Raw:         defaultCommunityBetaRaw("Creality Sermoon D3 Pro"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Sermoon M300",
			Slug:        "creality-sermoon-m300",
			Description: defaultCommunityBetaDesc("Creality Sermoon M300"),
			Raw:         defaultCommunityBetaRaw("Creality Sermoon M300"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Sermoon V1",
			Slug:        "creality-sermoon-v1",
			Description: defaultCommunityBetaDesc("Creality Sermoon V1"),
			Raw:         defaultCommunityBetaRaw("Creality Sermoon V1"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality Sermoon V1 Pro",
			Slug:        "creality-sermoon-v1-pro",
			Description: defaultCommunityBetaDesc("Creality Sermoon V1 Pro"),
			Raw:         defaultCommunityBetaRaw("Creality Sermoon V1 Pro"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Creality SPARKX i7",
			Slug:        "creality-sparkx-i7",
			Description: defaultCommunityBetaDesc("Creality SPARKX i7"),
			Raw:         defaultCommunityBetaRaw("Creality SPARKX i7"),
		})
		brands = append(brands, brand)
	}

	// Brand: Snapmaker (20 models)
	{
		brand := &PrinterBrand{
			Name:        "Snapmaker",
			Slug:        "snapmaker",
			Description: "Snapmaker U1 and High Combo — verified full remote control in MakerSpell.",
			Color:       "0EA5E9",
			Position:    2,
			Models:      make([]*PrinterBrandModel, 0, 20),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Snapmaker A250",
			Slug:        "snapmaker-a250",
			Description: defaultCommunityBetaDesc("Snapmaker A250"),
			Raw:         defaultCommunityBetaRaw("Snapmaker A250"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Snapmaker A250 BKit",
			Slug:        "snapmaker-a250-bkit",
			Description: defaultCommunityBetaDesc("Snapmaker A250 BKit"),
			Raw:         defaultCommunityBetaRaw("Snapmaker A250 BKit"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Snapmaker A250 Dual",
			Slug:        "snapmaker-a250-dual",
			Description: defaultCommunityBetaDesc("Snapmaker A250 Dual"),
			Raw:         defaultCommunityBetaRaw("Snapmaker A250 Dual"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Snapmaker A250 Dual BKit",
			Slug:        "snapmaker-a250-dual-bkit",
			Description: defaultCommunityBetaDesc("Snapmaker A250 Dual BKit"),
			Raw:         defaultCommunityBetaRaw("Snapmaker A250 Dual BKit"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Snapmaker A250 Dual QS+B Kit",
			Slug:        "snapmaker-a250-dual-qsb-kit",
			Description: defaultCommunityBetaDesc("Snapmaker A250 Dual QS+B Kit"),
			Raw:         defaultCommunityBetaRaw("Snapmaker A250 Dual QS+B Kit"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Snapmaker A250 Dual QSKit",
			Slug:        "snapmaker-a250-dual-qskit",
			Description: defaultCommunityBetaDesc("Snapmaker A250 Dual QSKit"),
			Raw:         defaultCommunityBetaRaw("Snapmaker A250 Dual QSKit"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Snapmaker A250 QS+B Kit",
			Slug:        "snapmaker-a250-qsb-kit",
			Description: defaultCommunityBetaDesc("Snapmaker A250 QS+B Kit"),
			Raw:         defaultCommunityBetaRaw("Snapmaker A250 QS+B Kit"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Snapmaker A250 QSKit",
			Slug:        "snapmaker-a250-qskit",
			Description: defaultCommunityBetaDesc("Snapmaker A250 QSKit"),
			Raw:         defaultCommunityBetaRaw("Snapmaker A250 QSKit"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Snapmaker A350",
			Slug:        "snapmaker-a350",
			Description: defaultCommunityBetaDesc("Snapmaker A350"),
			Raw:         defaultCommunityBetaRaw("Snapmaker A350"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Snapmaker A350 BKit",
			Slug:        "snapmaker-a350-bkit",
			Description: defaultCommunityBetaDesc("Snapmaker A350 BKit"),
			Raw:         defaultCommunityBetaRaw("Snapmaker A350 BKit"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Snapmaker A350 Dual",
			Slug:        "snapmaker-a350-dual",
			Description: defaultCommunityBetaDesc("Snapmaker A350 Dual"),
			Raw:         defaultCommunityBetaRaw("Snapmaker A350 Dual"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Snapmaker A350 Dual BKit",
			Slug:        "snapmaker-a350-dual-bkit",
			Description: defaultCommunityBetaDesc("Snapmaker A350 Dual BKit"),
			Raw:         defaultCommunityBetaRaw("Snapmaker A350 Dual BKit"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Snapmaker A350 Dual QS+B Kit",
			Slug:        "snapmaker-a350-dual-qsb-kit",
			Description: defaultCommunityBetaDesc("Snapmaker A350 Dual QS+B Kit"),
			Raw:         defaultCommunityBetaRaw("Snapmaker A350 Dual QS+B Kit"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Snapmaker A350 Dual QSKit",
			Slug:        "snapmaker-a350-dual-qskit",
			Description: defaultCommunityBetaDesc("Snapmaker A350 Dual QSKit"),
			Raw:         defaultCommunityBetaRaw("Snapmaker A350 Dual QSKit"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Snapmaker A350 QS+B Kit",
			Slug:        "snapmaker-a350-qsb-kit",
			Description: defaultCommunityBetaDesc("Snapmaker A350 QS+B Kit"),
			Raw:         defaultCommunityBetaRaw("Snapmaker A350 QS+B Kit"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Snapmaker A350 QSKit",
			Slug:        "snapmaker-a350-qskit",
			Description: defaultCommunityBetaDesc("Snapmaker A350 QSKit"),
			Raw:         defaultCommunityBetaRaw("Snapmaker A350 QSKit"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Snapmaker Artisan",
			Slug:        "snapmaker-artisan",
			Description: defaultCommunityBetaDesc("Snapmaker Artisan"),
			Raw:         defaultCommunityBetaRaw("Snapmaker Artisan"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Snapmaker J1",
			Slug:        "snapmaker-j1",
			Description: defaultCommunityBetaDesc("Snapmaker J1"),
			Raw:         defaultCommunityBetaRaw("Snapmaker J1"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Snapmaker U1",
			Slug:        "u1",
			Description: "Snapmaker U1 — verified full remote control in MakerSpell.",
			Raw: `The **Snapmaker U1** is a fully verified printer in MakerSpell with complete remote control support.

**Getting started**
- Pair it with MakerSpell for remote printing, monitoring and status notifications.
- Discuss U1-specific settings, materials and upgrades here.

**Support status:** ✅ Verified — full remote control in MakerSpell.

Share your prints, tips and questions!`,
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Snapmaker High Combo",
			Slug:        "high-combo",
			Description: "Snapmaker High Combo — verified full remote control in MakerSpell.",
			Raw: `The **Snapmaker High Combo** configuration is fully verified in MakerSpell.

**Getting started**
- Add it in MakerSpell to manage prints remotely alongside your other machines.
- A good place for module, enclosure and multi-tool discussions.

**Support status:** ✅ Verified — full remote control in MakerSpell.

Post your setups and ask anything!`,
		})
		brands = append(brands, brand)
	}

	// Brand: Bambu Lab (16 models)
	{
		brand := &PrinterBrand{
			Name:        "Bambu Lab",
			Slug:        "bambu-lab",
			Description: "Bambu Lab P1P, P1S, X1C, A1 series and more — community beta (upload-only) in MakerSpell.",
			Color:       "22C55E",
			Position:    3,
			Models:      make([]*PrinterBrandModel, 0, 16),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Bambu Lab A1 / A1 mini",
			Slug:        "a1-a1-mini",
			Description: "Bambu Lab A1 and A1 mini — community beta (upload-only) in MakerSpell.",
			Raw: `The **Bambu Lab A1 and A1 mini** are part of the MakerSpell community beta program.

**Support status:** 🧪 Community beta — upload-only in MakerSpell: transfer files, then start the print on the printer.

**In this topic**
- A1 family setup with MakerSpell.
- Bed-slinger tuning, AMS lite configurations and print quality tips.

Share your benchies and ask questions!`,
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Bambu Lab P1P / P1S",
			Slug:        "p1-p1s",
			Description: "Bambu Lab P1P and P1S — community beta (upload-only) in MakerSpell.",
			Raw: `The **Bambu Lab P1P and P1S** are part of the MakerSpell community beta program.

**Support status:** 🧪 Community beta — MakerSpell runs in **upload-only mode**: you can transfer files to the printer, but printing must be started from the machine itself, for safety.

**In this topic**
- Connect your P1P/P1S to MakerSpell (file upload).
- Share LAN mode tips, firmware notes and print results.

Want full remote control for Bambu Lab? Vote and share your use case here!`,
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Bambu Lab A1",
			Slug:        "bambu-lab-a1",
			Description: defaultCommunityBetaDesc("Bambu Lab A1"),
			Raw:         defaultCommunityBetaRaw("Bambu Lab A1"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Bambu Lab A1 mini",
			Slug:        "bambu-lab-a1-mini",
			Description: defaultCommunityBetaDesc("Bambu Lab A1 mini"),
			Raw:         defaultCommunityBetaRaw("Bambu Lab A1 mini"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Bambu Lab A2L",
			Slug:        "bambu-lab-a2l",
			Description: defaultCommunityBetaDesc("Bambu Lab A2L"),
			Raw:         defaultCommunityBetaRaw("Bambu Lab A2L"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Bambu Lab H2C",
			Slug:        "bambu-lab-h2c",
			Description: defaultCommunityBetaDesc("Bambu Lab H2C"),
			Raw:         defaultCommunityBetaRaw("Bambu Lab H2C"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Bambu Lab H2D",
			Slug:        "bambu-lab-h2d",
			Description: defaultCommunityBetaDesc("Bambu Lab H2D"),
			Raw:         defaultCommunityBetaRaw("Bambu Lab H2D"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Bambu Lab H2D Pro",
			Slug:        "bambu-lab-h2d-pro",
			Description: defaultCommunityBetaDesc("Bambu Lab H2D Pro"),
			Raw:         defaultCommunityBetaRaw("Bambu Lab H2D Pro"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Bambu Lab H2S",
			Slug:        "bambu-lab-h2s",
			Description: defaultCommunityBetaDesc("Bambu Lab H2S"),
			Raw:         defaultCommunityBetaRaw("Bambu Lab H2S"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Bambu Lab P1P",
			Slug:        "bambu-lab-p1p",
			Description: defaultCommunityBetaDesc("Bambu Lab P1P"),
			Raw:         defaultCommunityBetaRaw("Bambu Lab P1P"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Bambu Lab P1S",
			Slug:        "bambu-lab-p1s",
			Description: defaultCommunityBetaDesc("Bambu Lab P1S"),
			Raw:         defaultCommunityBetaRaw("Bambu Lab P1S"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Bambu Lab P2S",
			Slug:        "bambu-lab-p2s",
			Description: defaultCommunityBetaDesc("Bambu Lab P2S"),
			Raw:         defaultCommunityBetaRaw("Bambu Lab P2S"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Bambu Lab X1",
			Slug:        "bambu-lab-x1",
			Description: defaultCommunityBetaDesc("Bambu Lab X1"),
			Raw:         defaultCommunityBetaRaw("Bambu Lab X1"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Bambu Lab X1 Carbon",
			Slug:        "bambu-lab-x1-carbon",
			Description: defaultCommunityBetaDesc("Bambu Lab X1 Carbon"),
			Raw:         defaultCommunityBetaRaw("Bambu Lab X1 Carbon"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Bambu Lab X1E",
			Slug:        "bambu-lab-x1e",
			Description: defaultCommunityBetaDesc("Bambu Lab X1E"),
			Raw:         defaultCommunityBetaRaw("Bambu Lab X1E"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Bambu Lab X2D",
			Slug:        "bambu-lab-x2d",
			Description: defaultCommunityBetaDesc("Bambu Lab X2D"),
			Raw:         defaultCommunityBetaRaw("Bambu Lab X2D"),
		})
		brands = append(brands, brand)
	}

	// Brand: Voron (10 models)
	{
		brand := &PrinterBrand{
			Name:        "Voron",
			Slug:        "voron",
			Description: "Voron 2.4, Trident and Switchwire (Klipper / Moonraker) — community beta (upload-only) in MakerSpell.",
			Color:       "EF4444",
			Position:    4,
			Models:      make([]*PrinterBrandModel, 0, 10),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Voron Trident",
			Slug:        "voron-trident",
			Description: "Voron Trident (Klipper / Moonraker) — community beta (upload-only) in MakerSpell.",
			Raw: `The **Voron Trident** runs Klipper with Moonraker, and is part of the MakerSpell community beta.

**Support status:** 🧪 Community beta — upload-only in MakerSpell: transfer files via Moonraker, then start printing from the machine.

**In this topic**
- Trident-specific setup with MakerSpell.
- Bed leveling, gantry alignment and mod discussions.

Post your build and tuning results!`,
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Voron 2.4",
			Slug:        "voron-2-4",
			Description: "Voron 2.4 (Klipper / Moonraker) — community beta (upload-only) in MakerSpell.",
			Raw: `The **Voron 2.4** runs Klipper with Moonraker, and is part of the MakerSpell community beta.

**Support status:** 🧪 Community beta — upload-only in MakerSpell: send files to the printer via Moonraker, start prints on the machine.

**In this topic**
- Moonraker upload configuration with MakerSpell.
- QGL, input shaping and chamber tuning for the 2.4.

Share your build threads and mods!`,
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Voron 0.1",
			Slug:        "voron-01",
			Description: defaultCommunityBetaDesc("Voron 0.1"),
			Raw:         defaultCommunityBetaRaw("Voron 0.1"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Voron 2.4 250",
			Slug:        "voron-24-250",
			Description: defaultCommunityBetaDesc("Voron 2.4 250"),
			Raw:         defaultCommunityBetaRaw("Voron 2.4 250"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Voron 2.4 300",
			Slug:        "voron-24-300",
			Description: defaultCommunityBetaDesc("Voron 2.4 300"),
			Raw:         defaultCommunityBetaRaw("Voron 2.4 300"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Voron 2.4 350",
			Slug:        "voron-24-350",
			Description: defaultCommunityBetaDesc("Voron 2.4 350"),
			Raw:         defaultCommunityBetaRaw("Voron 2.4 350"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Voron Switchwire 250",
			Slug:        "voron-switchwire-250",
			Description: defaultCommunityBetaDesc("Voron Switchwire 250"),
			Raw:         defaultCommunityBetaRaw("Voron Switchwire 250"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Voron Trident 250",
			Slug:        "voron-trident-250",
			Description: defaultCommunityBetaDesc("Voron Trident 250"),
			Raw:         defaultCommunityBetaRaw("Voron Trident 250"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Voron Trident 300",
			Slug:        "voron-trident-300",
			Description: defaultCommunityBetaDesc("Voron Trident 300"),
			Raw:         defaultCommunityBetaRaw("Voron Trident 300"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Voron Trident 350",
			Slug:        "voron-trident-350",
			Description: defaultCommunityBetaDesc("Voron Trident 350"),
			Raw:         defaultCommunityBetaRaw("Voron Trident 350"),
		})
		brands = append(brands, brand)
	}

	// Brand: Elegoo (19 models)
	{
		brand := &PrinterBrand{
			Name:        "Elegoo",
			Slug:        "elegoo",
			Description: "Elegoo Neptune, Centauri and OrangeStorm series — community beta (upload-only) in MakerSpell.",
			Color:       "3B82F6",
			Position:    5,
			Models:      make([]*PrinterBrandModel, 0, 19),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Elegoo Neptune 4 / Neptune 4 Pro",
			Slug:        "neptune-4",
			Description: "Elegoo Neptune 4 and Neptune 4 Pro — community beta (upload-only) in MakerSpell.",
			Raw: `The **Elegoo Neptune 4 and Neptune 4 Pro** run Klipper and are part of the MakerSpell community beta.

**Support status:** 🧪 Community beta — upload-only in MakerSpell: files are transferred to the printer; printing starts on the machine.

**In this topic**
- Neptune 4 setup with MakerSpell file upload.
- Klipper config tuning, Z-offset and first-layer help.

Share your configs and prints!`,
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Elegoo Centauri",
			Slug:        "elegoo-centauri",
			Description: defaultCommunityBetaDesc("Elegoo Centauri"),
			Raw:         defaultCommunityBetaRaw("Elegoo Centauri"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Elegoo Centauri 2",
			Slug:        "elegoo-centauri-2",
			Description: defaultCommunityBetaDesc("Elegoo Centauri 2"),
			Raw:         defaultCommunityBetaRaw("Elegoo Centauri 2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Elegoo Centauri Carbon",
			Slug:        "elegoo-centauri-carbon",
			Description: defaultCommunityBetaDesc("Elegoo Centauri Carbon"),
			Raw:         defaultCommunityBetaRaw("Elegoo Centauri Carbon"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Elegoo Centauri Carbon 2",
			Slug:        "elegoo-centauri-carbon-2",
			Description: defaultCommunityBetaDesc("Elegoo Centauri Carbon 2"),
			Raw:         defaultCommunityBetaRaw("Elegoo Centauri Carbon 2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Elegoo Neptune",
			Slug:        "elegoo-neptune",
			Description: defaultCommunityBetaDesc("Elegoo Neptune"),
			Raw:         defaultCommunityBetaRaw("Elegoo Neptune"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Elegoo Neptune 2",
			Slug:        "elegoo-neptune-2",
			Description: defaultCommunityBetaDesc("Elegoo Neptune 2"),
			Raw:         defaultCommunityBetaRaw("Elegoo Neptune 2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Elegoo Neptune 2D",
			Slug:        "elegoo-neptune-2d",
			Description: defaultCommunityBetaDesc("Elegoo Neptune 2D"),
			Raw:         defaultCommunityBetaRaw("Elegoo Neptune 2D"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Elegoo Neptune 2S",
			Slug:        "elegoo-neptune-2s",
			Description: defaultCommunityBetaDesc("Elegoo Neptune 2S"),
			Raw:         defaultCommunityBetaRaw("Elegoo Neptune 2S"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Elegoo Neptune 3",
			Slug:        "elegoo-neptune-3",
			Description: defaultCommunityBetaDesc("Elegoo Neptune 3"),
			Raw:         defaultCommunityBetaRaw("Elegoo Neptune 3"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Elegoo Neptune 3 Max",
			Slug:        "elegoo-neptune-3-max",
			Description: defaultCommunityBetaDesc("Elegoo Neptune 3 Max"),
			Raw:         defaultCommunityBetaRaw("Elegoo Neptune 3 Max"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Elegoo Neptune 3 Plus",
			Slug:        "elegoo-neptune-3-plus",
			Description: defaultCommunityBetaDesc("Elegoo Neptune 3 Plus"),
			Raw:         defaultCommunityBetaRaw("Elegoo Neptune 3 Plus"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Elegoo Neptune 3 Pro",
			Slug:        "elegoo-neptune-3-pro",
			Description: defaultCommunityBetaDesc("Elegoo Neptune 3 Pro"),
			Raw:         defaultCommunityBetaRaw("Elegoo Neptune 3 Pro"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Elegoo Neptune 4",
			Slug:        "elegoo-neptune-4",
			Description: defaultCommunityBetaDesc("Elegoo Neptune 4"),
			Raw:         defaultCommunityBetaRaw("Elegoo Neptune 4"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Elegoo Neptune 4 Max",
			Slug:        "elegoo-neptune-4-max",
			Description: defaultCommunityBetaDesc("Elegoo Neptune 4 Max"),
			Raw:         defaultCommunityBetaRaw("Elegoo Neptune 4 Max"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Elegoo Neptune 4 Plus",
			Slug:        "elegoo-neptune-4-plus",
			Description: defaultCommunityBetaDesc("Elegoo Neptune 4 Plus"),
			Raw:         defaultCommunityBetaRaw("Elegoo Neptune 4 Plus"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Elegoo Neptune 4 Pro",
			Slug:        "elegoo-neptune-4-pro",
			Description: defaultCommunityBetaDesc("Elegoo Neptune 4 Pro"),
			Raw:         defaultCommunityBetaRaw("Elegoo Neptune 4 Pro"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Elegoo Neptune X",
			Slug:        "elegoo-neptune-x",
			Description: defaultCommunityBetaDesc("Elegoo Neptune X"),
			Raw:         defaultCommunityBetaRaw("Elegoo Neptune X"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Elegoo OrangeStorm Giga",
			Slug:        "elegoo-orangestorm-giga",
			Description: defaultCommunityBetaDesc("Elegoo OrangeStorm Giga"),
			Raw:         defaultCommunityBetaRaw("Elegoo OrangeStorm Giga"),
		})
		brands = append(brands, brand)
	}

	// Brand: Anycubic (22 models)
	{
		brand := &PrinterBrand{
			Name:        "Anycubic",
			Slug:        "anycubic",
			Description: "Anycubic Kobra series and Chiron / Mega printers — community beta (upload-only) in MakerSpell.",
			Color:       "A855F7",
			Position:    6,
			Models:      make([]*PrinterBrandModel, 0, 22),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Anycubic Kobra series",
			Slug:        "kobra-series",
			Description: "Anycubic Kobra series — community beta (upload-only) in MakerSpell.",
			Raw: `**Anycubic** printers, including the Kobra series, are part of the MakerSpell community beta program.

**Support status:** 🧪 Community beta — upload-only in MakerSpell: transfer files to the printer, start prints on the machine.

**In this topic**
- Which Anycubic machine do you run? Introduce your setup.
- Firmware, leveling and tuning discussions.

Tell us which models you want supported next!`,
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Anycubic 4Max Pro",
			Slug:        "anycubic-4max-pro",
			Description: defaultCommunityBetaDesc("Anycubic 4Max Pro"),
			Raw:         defaultCommunityBetaRaw("Anycubic 4Max Pro"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Anycubic 4Max Pro 2",
			Slug:        "anycubic-4max-pro-2",
			Description: defaultCommunityBetaDesc("Anycubic 4Max Pro 2"),
			Raw:         defaultCommunityBetaRaw("Anycubic 4Max Pro 2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Anycubic Chiron",
			Slug:        "anycubic-chiron",
			Description: defaultCommunityBetaDesc("Anycubic Chiron"),
			Raw:         defaultCommunityBetaRaw("Anycubic Chiron"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Anycubic i3 Mega S",
			Slug:        "anycubic-i3-mega-s",
			Description: defaultCommunityBetaDesc("Anycubic i3 Mega S"),
			Raw:         defaultCommunityBetaRaw("Anycubic i3 Mega S"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Anycubic Kobra",
			Slug:        "anycubic-kobra",
			Description: defaultCommunityBetaDesc("Anycubic Kobra"),
			Raw:         defaultCommunityBetaRaw("Anycubic Kobra"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Anycubic Kobra 2",
			Slug:        "anycubic-kobra-2",
			Description: defaultCommunityBetaDesc("Anycubic Kobra 2"),
			Raw:         defaultCommunityBetaRaw("Anycubic Kobra 2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Anycubic Kobra 2 Max",
			Slug:        "anycubic-kobra-2-max",
			Description: defaultCommunityBetaDesc("Anycubic Kobra 2 Max"),
			Raw:         defaultCommunityBetaRaw("Anycubic Kobra 2 Max"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Anycubic Kobra 2 Neo",
			Slug:        "anycubic-kobra-2-neo",
			Description: defaultCommunityBetaDesc("Anycubic Kobra 2 Neo"),
			Raw:         defaultCommunityBetaRaw("Anycubic Kobra 2 Neo"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Anycubic Kobra 2 Plus",
			Slug:        "anycubic-kobra-2-plus",
			Description: defaultCommunityBetaDesc("Anycubic Kobra 2 Plus"),
			Raw:         defaultCommunityBetaRaw("Anycubic Kobra 2 Plus"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Anycubic Kobra 2 Pro",
			Slug:        "anycubic-kobra-2-pro",
			Description: defaultCommunityBetaDesc("Anycubic Kobra 2 Pro"),
			Raw:         defaultCommunityBetaRaw("Anycubic Kobra 2 Pro"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Anycubic Kobra 3",
			Slug:        "anycubic-kobra-3",
			Description: defaultCommunityBetaDesc("Anycubic Kobra 3"),
			Raw:         defaultCommunityBetaRaw("Anycubic Kobra 3"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Anycubic Kobra 3 Max",
			Slug:        "anycubic-kobra-3-max",
			Description: defaultCommunityBetaDesc("Anycubic Kobra 3 Max"),
			Raw:         defaultCommunityBetaRaw("Anycubic Kobra 3 Max"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Anycubic Kobra 3 V2",
			Slug:        "anycubic-kobra-3-v2",
			Description: defaultCommunityBetaDesc("Anycubic Kobra 3 V2"),
			Raw:         defaultCommunityBetaRaw("Anycubic Kobra 3 V2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Anycubic Kobra Max",
			Slug:        "anycubic-kobra-max",
			Description: defaultCommunityBetaDesc("Anycubic Kobra Max"),
			Raw:         defaultCommunityBetaRaw("Anycubic Kobra Max"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Anycubic Kobra Neo",
			Slug:        "anycubic-kobra-neo",
			Description: defaultCommunityBetaDesc("Anycubic Kobra Neo"),
			Raw:         defaultCommunityBetaRaw("Anycubic Kobra Neo"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Anycubic Kobra Plus",
			Slug:        "anycubic-kobra-plus",
			Description: defaultCommunityBetaDesc("Anycubic Kobra Plus"),
			Raw:         defaultCommunityBetaRaw("Anycubic Kobra Plus"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Anycubic Kobra S1",
			Slug:        "anycubic-kobra-s1",
			Description: defaultCommunityBetaDesc("Anycubic Kobra S1"),
			Raw:         defaultCommunityBetaRaw("Anycubic Kobra S1"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Anycubic Kobra S1 Max",
			Slug:        "anycubic-kobra-s1-max",
			Description: defaultCommunityBetaDesc("Anycubic Kobra S1 Max"),
			Raw:         defaultCommunityBetaRaw("Anycubic Kobra S1 Max"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Anycubic Kobra X",
			Slug:        "anycubic-kobra-x",
			Description: defaultCommunityBetaDesc("Anycubic Kobra X"),
			Raw:         defaultCommunityBetaRaw("Anycubic Kobra X"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Anycubic Predator",
			Slug:        "anycubic-predator",
			Description: defaultCommunityBetaDesc("Anycubic Predator"),
			Raw:         defaultCommunityBetaRaw("Anycubic Predator"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Anycubic Vyper",
			Slug:        "anycubic-vyper",
			Description: defaultCommunityBetaDesc("Anycubic Vyper"),
			Raw:         defaultCommunityBetaRaw("Anycubic Vyper"),
		})
		brands = append(brands, brand)
	}

	// Brand: Prusa (19 models)
	{
		brand := &PrinterBrand{
			Name:        "Prusa",
			Slug:        "prusa",
			Description: "Prusa MK4, MK3S+, MINI+, XL and CORE series — community beta (upload-only) in MakerSpell.",
			Color:       "F59E0B",
			Position:    7,
			Models:      make([]*PrinterBrandModel, 0, 19),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Prusa MK4 / MK3S+ / MINI+",
			Slug:        "prusa-mk4-mk3s-mini",
			Description: "Prusa MK4, MK3S+ and MINI+ — community beta (upload-only) in MakerSpell.",
			Raw: `**Prusa Research** printers — MK4, MK3S+ and MINI+ — are part of the MakerSpell community beta.

**Support status:** 🧪 Community beta — upload-only in MakerSpell: upload files to your Prusa, then start the print on the printer or via PrusaConnect.

**In this topic**
- PrusaConnect workflows with MakerSpell.
- Filament profiles, first-layer calibration and reliability tips.

Share your experience with your Prusa!`,
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Prusa CORE One",
			Slug:        "prusa-core-one",
			Description: defaultCommunityBetaDesc("Prusa CORE One"),
			Raw:         defaultCommunityBetaRaw("Prusa CORE One"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Prusa CORE One HF",
			Slug:        "prusa-core-one-hf",
			Description: defaultCommunityBetaDesc("Prusa CORE One HF"),
			Raw:         defaultCommunityBetaRaw("Prusa CORE One HF"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Prusa CORE One INDX 4T",
			Slug:        "prusa-core-one-indx-4t",
			Description: defaultCommunityBetaDesc("Prusa CORE One INDX 4T"),
			Raw:         defaultCommunityBetaRaw("Prusa CORE One INDX 4T"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Prusa CORE One INDX 8T",
			Slug:        "prusa-core-one-indx-8t",
			Description: defaultCommunityBetaDesc("Prusa CORE One INDX 8T"),
			Raw:         defaultCommunityBetaRaw("Prusa CORE One INDX 8T"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Prusa CORE One L",
			Slug:        "prusa-core-one-l",
			Description: defaultCommunityBetaDesc("Prusa CORE One L"),
			Raw:         defaultCommunityBetaRaw("Prusa CORE One L"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Prusa CORE One L HF",
			Slug:        "prusa-core-one-l-hf",
			Description: defaultCommunityBetaDesc("Prusa CORE One L HF"),
			Raw:         defaultCommunityBetaRaw("Prusa CORE One L HF"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Prusa CORE One MMU3",
			Slug:        "prusa-core-one-mmu3",
			Description: defaultCommunityBetaDesc("Prusa CORE One MMU3"),
			Raw:         defaultCommunityBetaRaw("Prusa CORE One MMU3"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Prusa MINI",
			Slug:        "prusa-mini",
			Description: defaultCommunityBetaDesc("Prusa MINI"),
			Raw:         defaultCommunityBetaRaw("Prusa MINI"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Prusa MINI IS",
			Slug:        "prusa-mini-is",
			Description: defaultCommunityBetaDesc("Prusa MINI IS"),
			Raw:         defaultCommunityBetaRaw("Prusa MINI IS"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Prusa MK3.5",
			Slug:        "prusa-mk35",
			Description: defaultCommunityBetaDesc("Prusa MK3.5"),
			Raw:         defaultCommunityBetaRaw("Prusa MK3.5"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Prusa MK3S",
			Slug:        "prusa-mk3s",
			Description: defaultCommunityBetaDesc("Prusa MK3S"),
			Raw:         defaultCommunityBetaRaw("Prusa MK3S"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Prusa MK4",
			Slug:        "prusa-mk4",
			Description: defaultCommunityBetaDesc("Prusa MK4"),
			Raw:         defaultCommunityBetaRaw("Prusa MK4"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Prusa MK4 MMU3",
			Slug:        "prusa-mk4-mmu3",
			Description: defaultCommunityBetaDesc("Prusa MK4 MMU3"),
			Raw:         defaultCommunityBetaRaw("Prusa MK4 MMU3"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Prusa MK4 MMU3 HF",
			Slug:        "prusa-mk4-mmu3-hf",
			Description: defaultCommunityBetaDesc("Prusa MK4 MMU3 HF"),
			Raw:         defaultCommunityBetaRaw("Prusa MK4 MMU3 HF"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Prusa MK4S",
			Slug:        "prusa-mk4s",
			Description: defaultCommunityBetaDesc("Prusa MK4S"),
			Raw:         defaultCommunityBetaRaw("Prusa MK4S"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Prusa MK4S HF",
			Slug:        "prusa-mk4s-hf",
			Description: defaultCommunityBetaDesc("Prusa MK4S HF"),
			Raw:         defaultCommunityBetaRaw("Prusa MK4S HF"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Prusa XL",
			Slug:        "prusa-xl",
			Description: defaultCommunityBetaDesc("Prusa XL"),
			Raw:         defaultCommunityBetaRaw("Prusa XL"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Prusa XL 5T",
			Slug:        "prusa-xl-5t",
			Description: defaultCommunityBetaDesc("Prusa XL 5T"),
			Raw:         defaultCommunityBetaRaw("Prusa XL 5T"),
		})
		brands = append(brands, brand)
	}

	// Brand: Qidi (12 models)
	{
		brand := &PrinterBrand{
			Name:        "Qidi",
			Slug:        "qidi",
			Description: defaultBrandDesc("Qidi"),
			Color:       "14B8A6",
			Position:    8,
			Models:      make([]*PrinterBrandModel, 0, 12),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Qidi Q1 Pro",
			Slug:        "qidi-q1-pro",
			Description: defaultCommunityBetaDesc("Qidi Q1 Pro"),
			Raw:         defaultCommunityBetaRaw("Qidi Q1 Pro"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Qidi Q2",
			Slug:        "qidi-q2",
			Description: defaultCommunityBetaDesc("Qidi Q2"),
			Raw:         defaultCommunityBetaRaw("Qidi Q2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Qidi Q2C",
			Slug:        "qidi-q2c",
			Description: defaultCommunityBetaDesc("Qidi Q2C"),
			Raw:         defaultCommunityBetaRaw("Qidi Q2C"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Qidi X-CF Pro",
			Slug:        "qidi-x-cf-pro",
			Description: defaultCommunityBetaDesc("Qidi X-CF Pro"),
			Raw:         defaultCommunityBetaRaw("Qidi X-CF Pro"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Qidi X-Max",
			Slug:        "qidi-x-max",
			Description: defaultCommunityBetaDesc("Qidi X-Max"),
			Raw:         defaultCommunityBetaRaw("Qidi X-Max"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Qidi X-Max 3",
			Slug:        "qidi-x-max-3",
			Description: defaultCommunityBetaDesc("Qidi X-Max 3"),
			Raw:         defaultCommunityBetaRaw("Qidi X-Max 3"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Qidi X-Max 4",
			Slug:        "qidi-x-max-4",
			Description: defaultCommunityBetaDesc("Qidi X-Max 4"),
			Raw:         defaultCommunityBetaRaw("Qidi X-Max 4"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Qidi X-Plus",
			Slug:        "qidi-x-plus",
			Description: defaultCommunityBetaDesc("Qidi X-Plus"),
			Raw:         defaultCommunityBetaRaw("Qidi X-Plus"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Qidi X-Plus 3",
			Slug:        "qidi-x-plus-3",
			Description: defaultCommunityBetaDesc("Qidi X-Plus 3"),
			Raw:         defaultCommunityBetaRaw("Qidi X-Plus 3"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Qidi X-Plus 4",
			Slug:        "qidi-x-plus-4",
			Description: defaultCommunityBetaDesc("Qidi X-Plus 4"),
			Raw:         defaultCommunityBetaRaw("Qidi X-Plus 4"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Qidi X-Plus 5",
			Slug:        "qidi-x-plus-5",
			Description: defaultCommunityBetaDesc("Qidi X-Plus 5"),
			Raw:         defaultCommunityBetaRaw("Qidi X-Plus 5"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Qidi X-Smart 3",
			Slug:        "qidi-x-smart-3",
			Description: defaultCommunityBetaDesc("Qidi X-Smart 3"),
			Raw:         defaultCommunityBetaRaw("Qidi X-Smart 3"),
		})
		brands = append(brands, brand)
	}

	// Brand: Flashforge (13 models)
	{
		brand := &PrinterBrand{
			Name:        "Flashforge",
			Slug:        "flashforge",
			Description: defaultBrandDesc("Flashforge"),
			Color:       "EC4899",
			Position:    9,
			Models:      make([]*PrinterBrandModel, 0, 13),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Flashforge AD5X",
			Slug:        "flashforge-ad5x",
			Description: defaultCommunityBetaDesc("Flashforge AD5X"),
			Raw:         defaultCommunityBetaRaw("Flashforge AD5X"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Flashforge Adventurer 3 Series",
			Slug:        "flashforge-adventurer-3-series",
			Description: defaultCommunityBetaDesc("Flashforge Adventurer 3 Series"),
			Raw:         defaultCommunityBetaRaw("Flashforge Adventurer 3 Series"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Flashforge Adventurer 4 Series",
			Slug:        "flashforge-adventurer-4-series",
			Description: defaultCommunityBetaDesc("Flashforge Adventurer 4 Series"),
			Raw:         defaultCommunityBetaRaw("Flashforge Adventurer 4 Series"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Flashforge Adventurer 5M",
			Slug:        "flashforge-adventurer-5m",
			Description: defaultCommunityBetaDesc("Flashforge Adventurer 5M"),
			Raw:         defaultCommunityBetaRaw("Flashforge Adventurer 5M"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Flashforge Adventurer 5M Pro",
			Slug:        "flashforge-adventurer-5m-pro",
			Description: defaultCommunityBetaDesc("Flashforge Adventurer 5M Pro"),
			Raw:         defaultCommunityBetaRaw("Flashforge Adventurer 5M Pro"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Flashforge Adventurer A5",
			Slug:        "flashforge-adventurer-a5",
			Description: defaultCommunityBetaDesc("Flashforge Adventurer A5"),
			Raw:         defaultCommunityBetaRaw("Flashforge Adventurer A5"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Flashforge Artemis",
			Slug:        "flashforge-artemis",
			Description: defaultCommunityBetaDesc("Flashforge Artemis"),
			Raw:         defaultCommunityBetaRaw("Flashforge Artemis"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Flashforge Creator 5",
			Slug:        "flashforge-creator-5",
			Description: defaultCommunityBetaDesc("Flashforge Creator 5"),
			Raw:         defaultCommunityBetaRaw("Flashforge Creator 5"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Flashforge Creator 5 Pro",
			Slug:        "flashforge-creator-5-pro",
			Description: defaultCommunityBetaDesc("Flashforge Creator 5 Pro"),
			Raw:         defaultCommunityBetaRaw("Flashforge Creator 5 Pro"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Flashforge Guider 2s",
			Slug:        "flashforge-guider-2s",
			Description: defaultCommunityBetaDesc("Flashforge Guider 2s"),
			Raw:         defaultCommunityBetaRaw("Flashforge Guider 2s"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Flashforge Guider 3 Ultra",
			Slug:        "flashforge-guider-3-ultra",
			Description: defaultCommunityBetaDesc("Flashforge Guider 3 Ultra"),
			Raw:         defaultCommunityBetaRaw("Flashforge Guider 3 Ultra"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Flashforge Guider4",
			Slug:        "flashforge-guider4",
			Description: defaultCommunityBetaDesc("Flashforge Guider4"),
			Raw:         defaultCommunityBetaRaw("Flashforge Guider4"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Flashforge Guider4 Pro",
			Slug:        "flashforge-guider4-pro",
			Description: defaultCommunityBetaDesc("Flashforge Guider4 Pro"),
			Raw:         defaultCommunityBetaRaw("Flashforge Guider4 Pro"),
		})
		brands = append(brands, brand)
	}

	// Brand: Sovol (13 models)
	{
		brand := &PrinterBrand{
			Name:        "Sovol",
			Slug:        "sovol",
			Description: defaultBrandDesc("Sovol"),
			Color:       "8B5CF6",
			Position:    10,
			Models:      make([]*PrinterBrandModel, 0, 13),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Sovol SV01",
			Slug:        "sovol-sv01",
			Description: defaultCommunityBetaDesc("Sovol SV01"),
			Raw:         defaultCommunityBetaRaw("Sovol SV01"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Sovol SV01 Pro",
			Slug:        "sovol-sv01-pro",
			Description: defaultCommunityBetaDesc("Sovol SV01 Pro"),
			Raw:         defaultCommunityBetaRaw("Sovol SV01 Pro"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Sovol SV02",
			Slug:        "sovol-sv02",
			Description: defaultCommunityBetaDesc("Sovol SV02"),
			Raw:         defaultCommunityBetaRaw("Sovol SV02"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Sovol SV05",
			Slug:        "sovol-sv05",
			Description: defaultCommunityBetaDesc("Sovol SV05"),
			Raw:         defaultCommunityBetaRaw("Sovol SV05"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Sovol SV06",
			Slug:        "sovol-sv06",
			Description: defaultCommunityBetaDesc("Sovol SV06"),
			Raw:         defaultCommunityBetaRaw("Sovol SV06"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Sovol SV06 ACE",
			Slug:        "sovol-sv06-ace",
			Description: defaultCommunityBetaDesc("Sovol SV06 ACE"),
			Raw:         defaultCommunityBetaRaw("Sovol SV06 ACE"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Sovol SV06 Plus",
			Slug:        "sovol-sv06-plus",
			Description: defaultCommunityBetaDesc("Sovol SV06 Plus"),
			Raw:         defaultCommunityBetaRaw("Sovol SV06 Plus"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Sovol SV06 Plus ACE",
			Slug:        "sovol-sv06-plus-ace",
			Description: defaultCommunityBetaDesc("Sovol SV06 Plus ACE"),
			Raw:         defaultCommunityBetaRaw("Sovol SV06 Plus ACE"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Sovol SV07",
			Slug:        "sovol-sv07",
			Description: defaultCommunityBetaDesc("Sovol SV07"),
			Raw:         defaultCommunityBetaRaw("Sovol SV07"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Sovol SV07 Plus",
			Slug:        "sovol-sv07-plus",
			Description: defaultCommunityBetaDesc("Sovol SV07 Plus"),
			Raw:         defaultCommunityBetaRaw("Sovol SV07 Plus"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Sovol SV08",
			Slug:        "sovol-sv08",
			Description: defaultCommunityBetaDesc("Sovol SV08"),
			Raw:         defaultCommunityBetaRaw("Sovol SV08"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Sovol SV08 MAX",
			Slug:        "sovol-sv08-max",
			Description: defaultCommunityBetaDesc("Sovol SV08 MAX"),
			Raw:         defaultCommunityBetaRaw("Sovol SV08 MAX"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Sovol Zero",
			Slug:        "sovol-zero",
			Description: defaultCommunityBetaDesc("Sovol Zero"),
			Raw:         defaultCommunityBetaRaw("Sovol Zero"),
		})
		brands = append(brands, brand)
	}

	// Brand: Ratrig (21 models)
	{
		brand := &PrinterBrand{
			Name:        "Ratrig",
			Slug:        "ratrig",
			Description: defaultBrandDesc("Ratrig"),
			Color:       "6366F1",
			Position:    11,
			Models:      make([]*PrinterBrandModel, 0, 21),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "RatRig V-Cast",
			Slug:        "ratrig-v-cast",
			Description: defaultCommunityBetaDesc("RatRig V-Cast"),
			Raw:         defaultCommunityBetaRaw("RatRig V-Cast"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "RatRig V-Core 3 200",
			Slug:        "ratrig-v-core-3-200",
			Description: defaultCommunityBetaDesc("RatRig V-Core 3 200"),
			Raw:         defaultCommunityBetaRaw("RatRig V-Core 3 200"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "RatRig V-Core 3 300",
			Slug:        "ratrig-v-core-3-300",
			Description: defaultCommunityBetaDesc("RatRig V-Core 3 300"),
			Raw:         defaultCommunityBetaRaw("RatRig V-Core 3 300"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "RatRig V-Core 3 400",
			Slug:        "ratrig-v-core-3-400",
			Description: defaultCommunityBetaDesc("RatRig V-Core 3 400"),
			Raw:         defaultCommunityBetaRaw("RatRig V-Core 3 400"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "RatRig V-Core 3 500",
			Slug:        "ratrig-v-core-3-500",
			Description: defaultCommunityBetaDesc("RatRig V-Core 3 500"),
			Raw:         defaultCommunityBetaRaw("RatRig V-Core 3 500"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "RatRig V-Core 4 300",
			Slug:        "ratrig-v-core-4-300",
			Description: defaultCommunityBetaDesc("RatRig V-Core 4 300"),
			Raw:         defaultCommunityBetaRaw("RatRig V-Core 4 300"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "RatRig V-Core 4 400",
			Slug:        "ratrig-v-core-4-400",
			Description: defaultCommunityBetaDesc("RatRig V-Core 4 400"),
			Raw:         defaultCommunityBetaRaw("RatRig V-Core 4 400"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "RatRig V-Core 4 500",
			Slug:        "ratrig-v-core-4-500",
			Description: defaultCommunityBetaDesc("RatRig V-Core 4 500"),
			Raw:         defaultCommunityBetaRaw("RatRig V-Core 4 500"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "RatRig V-Core 4 HYBRID 300",
			Slug:        "ratrig-v-core-4-hybrid-300",
			Description: defaultCommunityBetaDesc("RatRig V-Core 4 HYBRID 300"),
			Raw:         defaultCommunityBetaRaw("RatRig V-Core 4 HYBRID 300"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "RatRig V-Core 4 HYBRID 400",
			Slug:        "ratrig-v-core-4-hybrid-400",
			Description: defaultCommunityBetaDesc("RatRig V-Core 4 HYBRID 400"),
			Raw:         defaultCommunityBetaRaw("RatRig V-Core 4 HYBRID 400"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "RatRig V-Core 4 HYBRID 500",
			Slug:        "ratrig-v-core-4-hybrid-500",
			Description: defaultCommunityBetaDesc("RatRig V-Core 4 HYBRID 500"),
			Raw:         defaultCommunityBetaRaw("RatRig V-Core 4 HYBRID 500"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "RatRig V-Core 4 IDEX 300",
			Slug:        "ratrig-v-core-4-idex-300",
			Description: defaultCommunityBetaDesc("RatRig V-Core 4 IDEX 300"),
			Raw:         defaultCommunityBetaRaw("RatRig V-Core 4 IDEX 300"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "RatRig V-Core 4 IDEX 300 COPY MODE",
			Slug:        "ratrig-v-core-4-idex-300-copy-mode",
			Description: defaultCommunityBetaDesc("RatRig V-Core 4 IDEX 300 COPY MODE"),
			Raw:         defaultCommunityBetaRaw("RatRig V-Core 4 IDEX 300 COPY MODE"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "RatRig V-Core 4 IDEX 300 MIRROR MODE",
			Slug:        "ratrig-v-core-4-idex-300-mirror-mode",
			Description: defaultCommunityBetaDesc("RatRig V-Core 4 IDEX 300 MIRROR MODE"),
			Raw:         defaultCommunityBetaRaw("RatRig V-Core 4 IDEX 300 MIRROR MODE"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "RatRig V-Core 4 IDEX 400",
			Slug:        "ratrig-v-core-4-idex-400",
			Description: defaultCommunityBetaDesc("RatRig V-Core 4 IDEX 400"),
			Raw:         defaultCommunityBetaRaw("RatRig V-Core 4 IDEX 400"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "RatRig V-Core 4 IDEX 400 COPY MODE",
			Slug:        "ratrig-v-core-4-idex-400-copy-mode",
			Description: defaultCommunityBetaDesc("RatRig V-Core 4 IDEX 400 COPY MODE"),
			Raw:         defaultCommunityBetaRaw("RatRig V-Core 4 IDEX 400 COPY MODE"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "RatRig V-Core 4 IDEX 400 MIRROR MODE",
			Slug:        "ratrig-v-core-4-idex-400-mirror-mode",
			Description: defaultCommunityBetaDesc("RatRig V-Core 4 IDEX 400 MIRROR MODE"),
			Raw:         defaultCommunityBetaRaw("RatRig V-Core 4 IDEX 400 MIRROR MODE"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "RatRig V-Core 4 IDEX 500",
			Slug:        "ratrig-v-core-4-idex-500",
			Description: defaultCommunityBetaDesc("RatRig V-Core 4 IDEX 500"),
			Raw:         defaultCommunityBetaRaw("RatRig V-Core 4 IDEX 500"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "RatRig V-Core 4 IDEX 500 COPY MODE",
			Slug:        "ratrig-v-core-4-idex-500-copy-mode",
			Description: defaultCommunityBetaDesc("RatRig V-Core 4 IDEX 500 COPY MODE"),
			Raw:         defaultCommunityBetaRaw("RatRig V-Core 4 IDEX 500 COPY MODE"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "RatRig V-Core 4 IDEX 500 MIRROR MODE",
			Slug:        "ratrig-v-core-4-idex-500-mirror-mode",
			Description: defaultCommunityBetaDesc("RatRig V-Core 4 IDEX 500 MIRROR MODE"),
			Raw:         defaultCommunityBetaRaw("RatRig V-Core 4 IDEX 500 MIRROR MODE"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "RatRig V-Minion",
			Slug:        "ratrig-v-minion",
			Description: defaultCommunityBetaDesc("RatRig V-Minion"),
			Raw:         defaultCommunityBetaRaw("RatRig V-Minion"),
		})
		brands = append(brands, brand)
	}

	// Brand: UltiMaker (4 models)
	{
		brand := &PrinterBrand{
			Name:        "UltiMaker",
			Slug:        "ultimaker",
			Description: defaultBrandDesc("UltiMaker"),
			Color:       "10B981",
			Position:    12,
			Models:      make([]*PrinterBrandModel, 0, 4),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "UltiMaker 2",
			Slug:        "ultimaker-2",
			Description: defaultCommunityBetaDesc("UltiMaker 2"),
			Raw:         defaultCommunityBetaRaw("UltiMaker 2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "UltiMaker 2 extended+",
			Slug:        "ultimaker-2-extended",
			Description: defaultCommunityBetaDesc("UltiMaker 2 extended+"),
			Raw:         defaultCommunityBetaRaw("UltiMaker 2 extended+"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "UltiMaker 2+",
			Slug:        "ultimaker-ultimaker-2",
			Description: defaultCommunityBetaDesc("UltiMaker 2+"),
			Raw:         defaultCommunityBetaRaw("UltiMaker 2+"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "UltiMaker S5",
			Slug:        "ultimaker-s5",
			Description: defaultCommunityBetaDesc("UltiMaker S5"),
			Raw:         defaultCommunityBetaRaw("UltiMaker S5"),
		})
		brands = append(brands, brand)
	}

	// Brand: Anker (3 models)
	{
		brand := &PrinterBrand{
			Name:        "Anker",
			Slug:        "anker",
			Description: defaultBrandDesc("Anker"),
			Color:       "F43F5E",
			Position:    13,
			Models:      make([]*PrinterBrandModel, 0, 3),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Anker M5",
			Slug:        "anker-m5",
			Description: defaultCommunityBetaDesc("Anker M5"),
			Raw:         defaultCommunityBetaRaw("Anker M5"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Anker M5 All-Metal Hot End",
			Slug:        "anker-m5-all-metal-hot-end",
			Description: defaultCommunityBetaDesc("Anker M5 All-Metal Hot End"),
			Raw:         defaultCommunityBetaRaw("Anker M5 All-Metal Hot End"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Anker M5C",
			Slug:        "anker-m5c",
			Description: defaultCommunityBetaDesc("Anker M5C"),
			Raw:         defaultCommunityBetaRaw("Anker M5C"),
		})
		brands = append(brands, brand)
	}

	// Brand: Artillery (10 models)
	{
		brand := &PrinterBrand{
			Name:        "Artillery",
			Slug:        "artillery",
			Description: defaultBrandDesc("Artillery"),
			Color:       "D97706",
			Position:    14,
			Models:      make([]*PrinterBrandModel, 0, 10),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Artillery Genius",
			Slug:        "artillery-genius",
			Description: defaultCommunityBetaDesc("Artillery Genius"),
			Raw:         defaultCommunityBetaRaw("Artillery Genius"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Artillery Genius Pro",
			Slug:        "artillery-genius-pro",
			Description: defaultCommunityBetaDesc("Artillery Genius Pro"),
			Raw:         defaultCommunityBetaRaw("Artillery Genius Pro"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Artillery Hornet",
			Slug:        "artillery-hornet",
			Description: defaultCommunityBetaDesc("Artillery Hornet"),
			Raw:         defaultCommunityBetaRaw("Artillery Hornet"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Artillery M1 Pro",
			Slug:        "artillery-m1-pro",
			Description: defaultCommunityBetaDesc("Artillery M1 Pro"),
			Raw:         defaultCommunityBetaRaw("Artillery M1 Pro"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Artillery Sidewinder X1",
			Slug:        "artillery-sidewinder-x1",
			Description: defaultCommunityBetaDesc("Artillery Sidewinder X1"),
			Raw:         defaultCommunityBetaRaw("Artillery Sidewinder X1"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Artillery Sidewinder X2",
			Slug:        "artillery-sidewinder-x2",
			Description: defaultCommunityBetaDesc("Artillery Sidewinder X2"),
			Raw:         defaultCommunityBetaRaw("Artillery Sidewinder X2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Artillery Sidewinder X3 Plus",
			Slug:        "artillery-sidewinder-x3-plus",
			Description: defaultCommunityBetaDesc("Artillery Sidewinder X3 Plus"),
			Raw:         defaultCommunityBetaRaw("Artillery Sidewinder X3 Plus"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Artillery Sidewinder X3 Pro",
			Slug:        "artillery-sidewinder-x3-pro",
			Description: defaultCommunityBetaDesc("Artillery Sidewinder X3 Pro"),
			Raw:         defaultCommunityBetaRaw("Artillery Sidewinder X3 Pro"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Artillery Sidewinder X4 Plus",
			Slug:        "artillery-sidewinder-x4-plus",
			Description: defaultCommunityBetaDesc("Artillery Sidewinder X4 Plus"),
			Raw:         defaultCommunityBetaRaw("Artillery Sidewinder X4 Plus"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Artillery Sidewinder X4 Pro",
			Slug:        "artillery-sidewinder-x4-pro",
			Description: defaultCommunityBetaDesc("Artillery Sidewinder X4 Pro"),
			Raw:         defaultCommunityBetaRaw("Artillery Sidewinder X4 Pro"),
		})
		brands = append(brands, brand)
	}

	// Brand: BIQU (3 models)
	{
		brand := &PrinterBrand{
			Name:        "BIQU",
			Slug:        "biqu",
			Description: defaultBrandDesc("BIQU"),
			Color:       "06B6D4",
			Position:    15,
			Models:      make([]*PrinterBrandModel, 0, 3),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "BIQU B1",
			Slug:        "biqu-b1",
			Description: defaultCommunityBetaDesc("BIQU B1"),
			Raw:         defaultCommunityBetaRaw("BIQU B1"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "BIQU BX",
			Slug:        "biqu-bx",
			Description: defaultCommunityBetaDesc("BIQU BX"),
			Raw:         defaultCommunityBetaRaw("BIQU BX"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "BIQU Hurakan",
			Slug:        "biqu-hurakan",
			Description: defaultCommunityBetaDesc("BIQU Hurakan"),
			Raw:         defaultCommunityBetaRaw("BIQU Hurakan"),
		})
		brands = append(brands, brand)
	}

	// Brand: Blocks (3 models)
	{
		brand := &PrinterBrand{
			Name:        "Blocks",
			Slug:        "blocks",
			Description: defaultBrandDesc("Blocks"),
			Color:       "84CC16",
			Position:    16,
			Models:      make([]*PrinterBrandModel, 0, 3),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "BLOCKS Pro S100",
			Slug:        "blocks-pro-s100",
			Description: defaultCommunityBetaDesc("BLOCKS Pro S100"),
			Raw:         defaultCommunityBetaRaw("BLOCKS Pro S100"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "BLOCKS RD50 V2",
			Slug:        "blocks-rd50-v2",
			Description: defaultCommunityBetaDesc("BLOCKS RD50 V2"),
			Raw:         defaultCommunityBetaRaw("BLOCKS RD50 V2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "BLOCKS RF50",
			Slug:        "blocks-rf50",
			Description: defaultCommunityBetaDesc("BLOCKS RF50"),
			Raw:         defaultCommunityBetaRaw("BLOCKS RF50"),
		})
		brands = append(brands, brand)
	}

	// Brand: Chuanying (1 models)
	{
		brand := &PrinterBrand{
			Name:        "Chuanying",
			Slug:        "chuanying",
			Description: defaultBrandDesc("Chuanying"),
			Color:       "64748B",
			Position:    17,
			Models:      make([]*PrinterBrandModel, 0, 1),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Chuanying X1",
			Slug:        "chuanying-x1",
			Description: defaultCommunityBetaDesc("Chuanying X1"),
			Raw:         defaultCommunityBetaRaw("Chuanying X1"),
		})
		brands = append(brands, brand)
	}

	// Brand: Co Print (1 models)
	{
		brand := &PrinterBrand{
			Name:        "Co Print",
			Slug:        "co-print",
			Description: defaultBrandDesc("Co Print"),
			Color:       "F97316",
			Position:    18,
			Models:      make([]*PrinterBrandModel, 0, 1),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Co Print ChromaSet",
			Slug:        "co-print-chromaset",
			Description: defaultCommunityBetaDesc("Co Print ChromaSet"),
			Raw:         defaultCommunityBetaRaw("Co Print ChromaSet"),
		})
		brands = append(brands, brand)
	}

	// Brand: CoLiDo (5 models)
	{
		brand := &PrinterBrand{
			Name:        "CoLiDo",
			Slug:        "colido",
			Description: defaultBrandDesc("CoLiDo"),
			Color:       "0EA5E9",
			Position:    19,
			Models:      make([]*PrinterBrandModel, 0, 5),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "CoLiDo 160 V2",
			Slug:        "colido-160-v2",
			Description: defaultCommunityBetaDesc("CoLiDo 160 V2"),
			Raw:         defaultCommunityBetaRaw("CoLiDo 160 V2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "CoLiDo DIY 4.0",
			Slug:        "colido-diy-40",
			Description: defaultCommunityBetaDesc("CoLiDo DIY 4.0"),
			Raw:         defaultCommunityBetaRaw("CoLiDo DIY 4.0"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "CoLiDo DIY 4.0 V2",
			Slug:        "colido-diy-40-v2",
			Description: defaultCommunityBetaDesc("CoLiDo DIY 4.0 V2"),
			Raw:         defaultCommunityBetaRaw("CoLiDo DIY 4.0 V2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "CoLiDo SR1",
			Slug:        "colido-sr1",
			Description: defaultCommunityBetaDesc("CoLiDo SR1"),
			Raw:         defaultCommunityBetaRaw("CoLiDo SR1"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "CoLiDo X16",
			Slug:        "colido-x16",
			Description: defaultCommunityBetaDesc("CoLiDo X16"),
			Raw:         defaultCommunityBetaRaw("CoLiDo X16"),
		})
		brands = append(brands, brand)
	}

	// Brand: Comgrow (2 models)
	{
		brand := &PrinterBrand{
			Name:        "Comgrow",
			Slug:        "comgrow",
			Description: defaultBrandDesc("Comgrow"),
			Color:       "22C55E",
			Position:    20,
			Models:      make([]*PrinterBrandModel, 0, 2),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Comgrow T300",
			Slug:        "comgrow-t300",
			Description: defaultCommunityBetaDesc("Comgrow T300"),
			Raw:         defaultCommunityBetaRaw("Comgrow T300"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Comgrow T500",
			Slug:        "comgrow-t500",
			Description: defaultCommunityBetaDesc("Comgrow T500"),
			Raw:         defaultCommunityBetaRaw("Comgrow T500"),
		})
		brands = append(brands, brand)
	}

	// Brand: CONSTRUCT3D (2 models)
	{
		brand := &PrinterBrand{
			Name:        "CONSTRUCT3D",
			Slug:        "construct3d",
			Description: defaultBrandDesc("CONSTRUCT3D"),
			Color:       "EF4444",
			Position:    21,
			Models:      make([]*PrinterBrandModel, 0, 2),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Construct 1",
			Slug:        "construct-1",
			Description: defaultCommunityBetaDesc("Construct 1"),
			Raw:         defaultCommunityBetaRaw("Construct 1"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Construct 1 XL",
			Slug:        "construct-1-xl",
			Description: defaultCommunityBetaDesc("Construct 1 XL"),
			Raw:         defaultCommunityBetaRaw("Construct 1 XL"),
		})
		brands = append(brands, brand)
	}

	// Brand: Cubicon (3 models)
	{
		brand := &PrinterBrand{
			Name:        "Cubicon",
			Slug:        "cubicon",
			Description: defaultBrandDesc("Cubicon"),
			Color:       "3B82F6",
			Position:    22,
			Models:      make([]*PrinterBrandModel, 0, 3),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Cubicon xCeler-I",
			Slug:        "cubicon-xceler-i",
			Description: defaultCommunityBetaDesc("Cubicon xCeler-I"),
			Raw:         defaultCommunityBetaRaw("Cubicon xCeler-I"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Cubicon xCeler-Mini",
			Slug:        "cubicon-xceler-mini",
			Description: defaultCommunityBetaDesc("Cubicon xCeler-Mini"),
			Raw:         defaultCommunityBetaRaw("Cubicon xCeler-Mini"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Cubicon xCeler-Plus",
			Slug:        "cubicon-xceler-plus",
			Description: defaultCommunityBetaDesc("Cubicon xCeler-Plus"),
			Raw:         defaultCommunityBetaRaw("Cubicon xCeler-Plus"),
		})
		brands = append(brands, brand)
	}

	// Brand: DeltaMaker (3 models)
	{
		brand := &PrinterBrand{
			Name:        "DeltaMaker",
			Slug:        "deltamaker",
			Description: defaultBrandDesc("DeltaMaker"),
			Color:       "A855F7",
			Position:    23,
			Models:      make([]*PrinterBrandModel, 0, 3),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "DeltaMaker 2",
			Slug:        "deltamaker-2",
			Description: defaultCommunityBetaDesc("DeltaMaker 2"),
			Raw:         defaultCommunityBetaRaw("DeltaMaker 2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "DeltaMaker 2T",
			Slug:        "deltamaker-2t",
			Description: defaultCommunityBetaDesc("DeltaMaker 2T"),
			Raw:         defaultCommunityBetaRaw("DeltaMaker 2T"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "DeltaMaker 2XT",
			Slug:        "deltamaker-2xt",
			Description: defaultCommunityBetaDesc("DeltaMaker 2XT"),
			Raw:         defaultCommunityBetaRaw("DeltaMaker 2XT"),
		})
		brands = append(brands, brand)
	}

	// Brand: Dremel (3 models)
	{
		brand := &PrinterBrand{
			Name:        "Dremel",
			Slug:        "dremel",
			Description: defaultBrandDesc("Dremel"),
			Color:       "F59E0B",
			Position:    24,
			Models:      make([]*PrinterBrandModel, 0, 3),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Dremel 3D20",
			Slug:        "dremel-3d20",
			Description: defaultCommunityBetaDesc("Dremel 3D20"),
			Raw:         defaultCommunityBetaRaw("Dremel 3D20"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Dremel 3D40",
			Slug:        "dremel-3d40",
			Description: defaultCommunityBetaDesc("Dremel 3D40"),
			Raw:         defaultCommunityBetaRaw("Dremel 3D40"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Dremel 3D45",
			Slug:        "dremel-3d45",
			Description: defaultCommunityBetaDesc("Dremel 3D45"),
			Raw:         defaultCommunityBetaRaw("Dremel 3D45"),
		})
		brands = append(brands, brand)
	}

	// Brand: Eryone (3 models)
	{
		brand := &PrinterBrand{
			Name:        "Eryone",
			Slug:        "eryone",
			Description: defaultBrandDesc("Eryone"),
			Color:       "14B8A6",
			Position:    25,
			Models:      make([]*PrinterBrandModel, 0, 3),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Eryone ER20",
			Slug:        "eryone-er20",
			Description: defaultCommunityBetaDesc("Eryone ER20"),
			Raw:         defaultCommunityBetaRaw("Eryone ER20"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Eryone ER20 Klipper",
			Slug:        "eryone-er20-klipper",
			Description: defaultCommunityBetaDesc("Eryone ER20 Klipper"),
			Raw:         defaultCommunityBetaRaw("Eryone ER20 Klipper"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Thinker X400",
			Slug:        "thinker-x400",
			Description: defaultCommunityBetaDesc("Thinker X400"),
			Raw:         defaultCommunityBetaRaw("Thinker X400"),
		})
		brands = append(brands, brand)
	}

	// Brand: FELIX Printers (7 models)
	{
		brand := &PrinterBrand{
			Name:        "FELIX Printers",
			Slug:        "felix-printers",
			Description: defaultBrandDesc("FELIX Printers"),
			Color:       "EC4899",
			Position:    26,
			Models:      make([]*PrinterBrandModel, 0, 7),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "FELIX Food Double",
			Slug:        "felix-food-double",
			Description: defaultCommunityBetaDesc("FELIX Food Double"),
			Raw:         defaultCommunityBetaRaw("FELIX Food Double"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "FELIX Food Single",
			Slug:        "felix-food-single",
			Description: defaultCommunityBetaDesc("FELIX Food Single"),
			Raw:         defaultCommunityBetaRaw("FELIX Food Single"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "FELIX Food Switch",
			Slug:        "felix-food-switch",
			Description: defaultCommunityBetaDesc("FELIX Food Switch"),
			Raw:         defaultCommunityBetaRaw("FELIX Food Switch"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "FELIX Pro 3",
			Slug:        "felix-pro-3",
			Description: defaultCommunityBetaDesc("FELIX Pro 3"),
			Raw:         defaultCommunityBetaRaw("FELIX Pro 3"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "FELIX Pro L 3.0",
			Slug:        "felix-pro-l-30",
			Description: defaultCommunityBetaDesc("FELIX Pro L 3.0"),
			Raw:         defaultCommunityBetaRaw("FELIX Pro L 3.0"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "FELIX Pro XL 3.0",
			Slug:        "felix-pro-xl-30",
			Description: defaultCommunityBetaDesc("FELIX Pro XL 3.0"),
			Raw:         defaultCommunityBetaRaw("FELIX Pro XL 3.0"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "FELIX Tec 4",
			Slug:        "felix-tec-4",
			Description: defaultCommunityBetaDesc("FELIX Tec 4"),
			Raw:         defaultCommunityBetaRaw("FELIX Tec 4"),
		})
		brands = append(brands, brand)
	}

	// Brand: FLSun (6 models)
	{
		brand := &PrinterBrand{
			Name:        "FLSun",
			Slug:        "flsun",
			Description: defaultBrandDesc("FLSun"),
			Color:       "8B5CF6",
			Position:    27,
			Models:      make([]*PrinterBrandModel, 0, 6),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "FLSun Q5",
			Slug:        "flsun-q5",
			Description: defaultCommunityBetaDesc("FLSun Q5"),
			Raw:         defaultCommunityBetaRaw("FLSun Q5"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "FLSun QQ-S Pro",
			Slug:        "flsun-qq-s-pro",
			Description: defaultCommunityBetaDesc("FLSun QQ-S Pro"),
			Raw:         defaultCommunityBetaRaw("FLSun QQ-S Pro"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "FLSun S1",
			Slug:        "flsun-s1",
			Description: defaultCommunityBetaDesc("FLSun S1"),
			Raw:         defaultCommunityBetaRaw("FLSun S1"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "FLSun Super Racer (SR)",
			Slug:        "flsun-super-racer-sr",
			Description: defaultCommunityBetaDesc("FLSun Super Racer (SR)"),
			Raw:         defaultCommunityBetaRaw("FLSun Super Racer (SR)"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "FLSun T1",
			Slug:        "flsun-t1",
			Description: defaultCommunityBetaDesc("FLSun T1"),
			Raw:         defaultCommunityBetaRaw("FLSun T1"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "FLSun V400",
			Slug:        "flsun-v400",
			Description: defaultCommunityBetaDesc("FLSun V400"),
			Raw:         defaultCommunityBetaRaw("FLSun V400"),
		})
		brands = append(brands, brand)
	}

	// Brand: FlyingBear (4 models)
	{
		brand := &PrinterBrand{
			Name:        "FlyingBear",
			Slug:        "flyingbear",
			Description: defaultBrandDesc("FlyingBear"),
			Color:       "6366F1",
			Position:    28,
			Models:      make([]*PrinterBrandModel, 0, 4),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "FlyingBear Ghost 6",
			Slug:        "flyingbear-ghost-6",
			Description: defaultCommunityBetaDesc("FlyingBear Ghost 6"),
			Raw:         defaultCommunityBetaRaw("FlyingBear Ghost 6"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "FlyingBear Ghost7",
			Slug:        "flyingbear-ghost7",
			Description: defaultCommunityBetaDesc("FlyingBear Ghost7"),
			Raw:         defaultCommunityBetaRaw("FlyingBear Ghost7"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "FlyingBear Reborn3",
			Slug:        "flyingbear-reborn3",
			Description: defaultCommunityBetaDesc("FlyingBear Reborn3"),
			Raw:         defaultCommunityBetaRaw("FlyingBear Reborn3"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "FlyingBear S1",
			Slug:        "flyingbear-s1",
			Description: defaultCommunityBetaDesc("FlyingBear S1"),
			Raw:         defaultCommunityBetaRaw("FlyingBear S1"),
		})
		brands = append(brands, brand)
	}

	// Brand: Folgertech (3 models)
	{
		brand := &PrinterBrand{
			Name:        "Folgertech",
			Slug:        "folgertech",
			Description: defaultBrandDesc("Folgertech"),
			Color:       "10B981",
			Position:    29,
			Models:      make([]*PrinterBrandModel, 0, 3),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Folgertech FT-5",
			Slug:        "folgertech-ft-5",
			Description: defaultCommunityBetaDesc("Folgertech FT-5"),
			Raw:         defaultCommunityBetaRaw("Folgertech FT-5"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Folgertech FT-6",
			Slug:        "folgertech-ft-6",
			Description: defaultCommunityBetaDesc("Folgertech FT-6"),
			Raw:         defaultCommunityBetaRaw("Folgertech FT-6"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Folgertech i3",
			Slug:        "folgertech-i3",
			Description: defaultCommunityBetaDesc("Folgertech i3"),
			Raw:         defaultCommunityBetaRaw("Folgertech i3"),
		})
		brands = append(brands, brand)
	}

	// Brand: Geeetech (16 models)
	{
		brand := &PrinterBrand{
			Name:        "Geeetech",
			Slug:        "geeetech",
			Description: defaultBrandDesc("Geeetech"),
			Color:       "F43F5E",
			Position:    30,
			Models:      make([]*PrinterBrandModel, 0, 16),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Geeetech A10 M",
			Slug:        "geeetech-a10-m",
			Description: defaultCommunityBetaDesc("Geeetech A10 M"),
			Raw:         defaultCommunityBetaRaw("Geeetech A10 M"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Geeetech A10 Pro",
			Slug:        "geeetech-a10-pro",
			Description: defaultCommunityBetaDesc("Geeetech A10 Pro"),
			Raw:         defaultCommunityBetaRaw("Geeetech A10 Pro"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Geeetech A10 T",
			Slug:        "geeetech-a10-t",
			Description: defaultCommunityBetaDesc("Geeetech A10 T"),
			Raw:         defaultCommunityBetaRaw("Geeetech A10 T"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Geeetech A20",
			Slug:        "geeetech-a20",
			Description: defaultCommunityBetaDesc("Geeetech A20"),
			Raw:         defaultCommunityBetaRaw("Geeetech A20"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Geeetech A20 M",
			Slug:        "geeetech-a20-m",
			Description: defaultCommunityBetaDesc("Geeetech A20 M"),
			Raw:         defaultCommunityBetaRaw("Geeetech A20 M"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Geeetech A20 T",
			Slug:        "geeetech-a20-t",
			Description: defaultCommunityBetaDesc("Geeetech A20 T"),
			Raw:         defaultCommunityBetaRaw("Geeetech A20 T"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Geeetech A30 M",
			Slug:        "geeetech-a30-m",
			Description: defaultCommunityBetaDesc("Geeetech A30 M"),
			Raw:         defaultCommunityBetaRaw("Geeetech A30 M"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Geeetech A30 Pro",
			Slug:        "geeetech-a30-pro",
			Description: defaultCommunityBetaDesc("Geeetech A30 Pro"),
			Raw:         defaultCommunityBetaRaw("Geeetech A30 Pro"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Geeetech A30 T",
			Slug:        "geeetech-a30-t",
			Description: defaultCommunityBetaDesc("Geeetech A30 T"),
			Raw:         defaultCommunityBetaRaw("Geeetech A30 T"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Geeetech M1",
			Slug:        "geeetech-m1",
			Description: defaultCommunityBetaDesc("Geeetech M1"),
			Raw:         defaultCommunityBetaRaw("Geeetech M1"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Geeetech Mizar",
			Slug:        "geeetech-mizar",
			Description: defaultCommunityBetaDesc("Geeetech Mizar"),
			Raw:         defaultCommunityBetaRaw("Geeetech Mizar"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Geeetech Mizar M",
			Slug:        "geeetech-mizar-m",
			Description: defaultCommunityBetaDesc("Geeetech Mizar M"),
			Raw:         defaultCommunityBetaRaw("Geeetech Mizar M"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Geeetech Mizar Max",
			Slug:        "geeetech-mizar-max",
			Description: defaultCommunityBetaDesc("Geeetech Mizar Max"),
			Raw:         defaultCommunityBetaRaw("Geeetech Mizar Max"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Geeetech Mizar Pro",
			Slug:        "geeetech-mizar-pro",
			Description: defaultCommunityBetaDesc("Geeetech Mizar Pro"),
			Raw:         defaultCommunityBetaRaw("Geeetech Mizar Pro"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Geeetech Mizar S",
			Slug:        "geeetech-mizar-s",
			Description: defaultCommunityBetaDesc("Geeetech Mizar S"),
			Raw:         defaultCommunityBetaRaw("Geeetech Mizar S"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Geeetech Thunder",
			Slug:        "geeetech-thunder",
			Description: defaultCommunityBetaDesc("Geeetech Thunder"),
			Raw:         defaultCommunityBetaRaw("Geeetech Thunder"),
		})
		brands = append(brands, brand)
	}

	// Brand: Ginger Additive (1 models)
	{
		brand := &PrinterBrand{
			Name:        "Ginger Additive",
			Slug:        "ginger-additive",
			Description: defaultBrandDesc("Ginger Additive"),
			Color:       "D97706",
			Position:    31,
			Models:      make([]*PrinterBrandModel, 0, 1),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Ginger G1",
			Slug:        "ginger-g1",
			Description: defaultCommunityBetaDesc("Ginger G1"),
			Raw:         defaultCommunityBetaRaw("Ginger G1"),
		})
		brands = append(brands, brand)
	}

	// Brand: Goofoo (3 models)
	{
		brand := &PrinterBrand{
			Name:        "Goofoo",
			Slug:        "goofoo",
			Description: defaultBrandDesc("Goofoo"),
			Color:       "06B6D4",
			Position:    32,
			Models:      make([]*PrinterBrandModel, 0, 3),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Goofoo GP1",
			Slug:        "goofoo-gp1",
			Description: defaultCommunityBetaDesc("Goofoo GP1"),
			Raw:         defaultCommunityBetaRaw("Goofoo GP1"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Goofoo TurboM",
			Slug:        "goofoo-turbom",
			Description: defaultCommunityBetaDesc("Goofoo TurboM"),
			Raw:         defaultCommunityBetaRaw("Goofoo TurboM"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Goofoo Velo+",
			Slug:        "goofoo-velo",
			Description: defaultCommunityBetaDesc("Goofoo Velo+"),
			Raw:         defaultCommunityBetaRaw("Goofoo Velo+"),
		})
		brands = append(brands, brand)
	}

	// Brand: IdeaFormer (1 models)
	{
		brand := &PrinterBrand{
			Name:        "IdeaFormer",
			Slug:        "ideaformer",
			Description: defaultBrandDesc("IdeaFormer"),
			Color:       "84CC16",
			Position:    33,
			Models:      make([]*PrinterBrandModel, 0, 1),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "IdeaFormer IR3 V2",
			Slug:        "ideaformer-ir3-v2",
			Description: defaultCommunityBetaDesc("IdeaFormer IR3 V2"),
			Raw:         defaultCommunityBetaRaw("IdeaFormer IR3 V2"),
		})
		brands = append(brands, brand)
	}

	// Brand: InfiMech (4 models)
	{
		brand := &PrinterBrand{
			Name:        "InfiMech",
			Slug:        "infimech",
			Description: defaultBrandDesc("InfiMech"),
			Color:       "64748B",
			Position:    34,
			Models:      make([]*PrinterBrandModel, 0, 4),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "InfiMech EX",
			Slug:        "infimech-ex",
			Description: defaultCommunityBetaDesc("InfiMech EX"),
			Raw:         defaultCommunityBetaRaw("InfiMech EX"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "InfiMech EX+APS",
			Slug:        "infimech-exaps",
			Description: defaultCommunityBetaDesc("InfiMech EX+APS"),
			Raw:         defaultCommunityBetaRaw("InfiMech EX+APS"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "InfiMech TX",
			Slug:        "infimech-tx",
			Description: defaultCommunityBetaDesc("InfiMech TX"),
			Raw:         defaultCommunityBetaRaw("InfiMech TX"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "InfiMech TX Hardened Steel Nozzle",
			Slug:        "infimech-tx-hardened-steel-nozzle",
			Description: defaultCommunityBetaDesc("InfiMech TX Hardened Steel Nozzle"),
			Raw:         defaultCommunityBetaRaw("InfiMech TX Hardened Steel Nozzle"),
		})
		brands = append(brands, brand)
	}

	// Brand: iQ (2 models)
	{
		brand := &PrinterBrand{
			Name:        "iQ",
			Slug:        "iq",
			Description: defaultBrandDesc("iQ"),
			Color:       "F97316",
			Position:    35,
			Models:      make([]*PrinterBrandModel, 0, 2),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "TiQ2",
			Slug:        "tiq2",
			Description: defaultCommunityBetaDesc("TiQ2"),
			Raw:         defaultCommunityBetaRaw("TiQ2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "TiQ8",
			Slug:        "tiq8",
			Description: defaultCommunityBetaDesc("TiQ8"),
			Raw:         defaultCommunityBetaRaw("TiQ8"),
		})
		brands = append(brands, brand)
	}

	// Brand: Kingroon (5 models)
	{
		brand := &PrinterBrand{
			Name:        "Kingroon",
			Slug:        "kingroon",
			Description: defaultBrandDesc("Kingroon"),
			Color:       "0EA5E9",
			Position:    36,
			Models:      make([]*PrinterBrandModel, 0, 5),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Kingroon KLP1",
			Slug:        "kingroon-klp1",
			Description: defaultCommunityBetaDesc("Kingroon KLP1"),
			Raw:         defaultCommunityBetaRaw("Kingroon KLP1"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Kingroon KP3S 3.0",
			Slug:        "kingroon-kp3s-30",
			Description: defaultCommunityBetaDesc("Kingroon KP3S 3.0"),
			Raw:         defaultCommunityBetaRaw("Kingroon KP3S 3.0"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Kingroon KP3S PRO S1",
			Slug:        "kingroon-kp3s-pro-s1",
			Description: defaultCommunityBetaDesc("Kingroon KP3S PRO S1"),
			Raw:         defaultCommunityBetaRaw("Kingroon KP3S PRO S1"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Kingroon KP3S PRO V2",
			Slug:        "kingroon-kp3s-pro-v2",
			Description: defaultCommunityBetaDesc("Kingroon KP3S PRO V2"),
			Raw:         defaultCommunityBetaRaw("Kingroon KP3S PRO V2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Kingroon KP3S V1",
			Slug:        "kingroon-kp3s-v1",
			Description: defaultCommunityBetaDesc("Kingroon KP3S V1"),
			Raw:         defaultCommunityBetaRaw("Kingroon KP3S V1"),
		})
		brands = append(brands, brand)
	}

	// Brand: LH (2 models)
	{
		brand := &PrinterBrand{
			Name:        "LH",
			Slug:        "lh",
			Description: defaultBrandDesc("LH"),
			Color:       "22C55E",
			Position:    37,
			Models:      make([]*PrinterBrandModel, 0, 2),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "LH Stinger",
			Slug:        "lh-stinger",
			Description: defaultCommunityBetaDesc("LH Stinger"),
			Raw:         defaultCommunityBetaRaw("LH Stinger"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "LH Stinger MMU",
			Slug:        "lh-stinger-mmu",
			Description: defaultCommunityBetaDesc("LH Stinger MMU"),
			Raw:         defaultCommunityBetaRaw("LH Stinger MMU"),
		})
		brands = append(brands, brand)
	}

	// Brand: LONGER (2 models)
	{
		brand := &PrinterBrand{
			Name:        "LONGER",
			Slug:        "longer",
			Description: defaultBrandDesc("LONGER"),
			Color:       "EF4444",
			Position:    38,
			Models:      make([]*PrinterBrandModel, 0, 2),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "LONGER LK10",
			Slug:        "longer-lk10",
			Description: defaultCommunityBetaDesc("LONGER LK10"),
			Raw:         defaultCommunityBetaRaw("LONGER LK10"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "LONGER LK10 Plus",
			Slug:        "longer-lk10-plus",
			Description: defaultCommunityBetaDesc("LONGER LK10 Plus"),
			Raw:         defaultCommunityBetaRaw("LONGER LK10 Plus"),
		})
		brands = append(brands, brand)
	}

	// Brand: Lulzbot (7 models)
	{
		brand := &PrinterBrand{
			Name:        "Lulzbot",
			Slug:        "lulzbot",
			Description: defaultBrandDesc("Lulzbot"),
			Color:       "3B82F6",
			Position:    39,
			Models:      make([]*PrinterBrandModel, 0, 7),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Lulzbot Mini 1",
			Slug:        "lulzbot-mini-1",
			Description: defaultCommunityBetaDesc("Lulzbot Mini 1"),
			Raw:         defaultCommunityBetaRaw("Lulzbot Mini 1"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Lulzbot Taz 4 or 5",
			Slug:        "lulzbot-taz-4-or-5",
			Description: defaultCommunityBetaDesc("Lulzbot Taz 4 or 5"),
			Raw:         defaultCommunityBetaRaw("Lulzbot Taz 4 or 5"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Lulzbot Taz 6",
			Slug:        "lulzbot-taz-6",
			Description: defaultCommunityBetaDesc("Lulzbot Taz 6"),
			Raw:         defaultCommunityBetaRaw("Lulzbot Taz 6"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Lulzbot Taz Mini 2",
			Slug:        "lulzbot-taz-mini-2",
			Description: defaultCommunityBetaDesc("Lulzbot Taz Mini 2"),
			Raw:         defaultCommunityBetaRaw("Lulzbot Taz Mini 2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Lulzbot Taz Pro Dual",
			Slug:        "lulzbot-taz-pro-dual",
			Description: defaultCommunityBetaDesc("Lulzbot Taz Pro Dual"),
			Raw:         defaultCommunityBetaRaw("Lulzbot Taz Pro Dual"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Lulzbot Taz Pro S",
			Slug:        "lulzbot-taz-pro-s",
			Description: defaultCommunityBetaDesc("Lulzbot Taz Pro S"),
			Raw:         defaultCommunityBetaRaw("Lulzbot Taz Pro S"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Lulzbot Taz Workhorse",
			Slug:        "lulzbot-taz-workhorse",
			Description: defaultCommunityBetaDesc("Lulzbot Taz Workhorse"),
			Raw:         defaultCommunityBetaRaw("Lulzbot Taz Workhorse"),
		})
		brands = append(brands, brand)
	}

	// Brand: M3D (2 models)
	{
		brand := &PrinterBrand{
			Name:        "M3D",
			Slug:        "m3d",
			Description: defaultBrandDesc("M3D"),
			Color:       "A855F7",
			Position:    40,
			Models:      make([]*PrinterBrandModel, 0, 2),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "M3D Enabler D7500",
			Slug:        "m3d-enabler-d7500",
			Description: defaultCommunityBetaDesc("M3D Enabler D7500"),
			Raw:         defaultCommunityBetaRaw("M3D Enabler D7500"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "M3D Enabler D8500",
			Slug:        "m3d-enabler-d8500",
			Description: defaultCommunityBetaDesc("M3D Enabler D8500"),
			Raw:         defaultCommunityBetaRaw("M3D Enabler D8500"),
		})
		brands = append(brands, brand)
	}

	// Brand: MagicMaker (5 models)
	{
		brand := &PrinterBrand{
			Name:        "MagicMaker",
			Slug:        "magicmaker",
			Description: defaultBrandDesc("MagicMaker"),
			Color:       "F59E0B",
			Position:    41,
			Models:      make([]*PrinterBrandModel, 0, 5),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "MM BoneKing",
			Slug:        "mm-boneking",
			Description: defaultCommunityBetaDesc("MM BoneKing"),
			Raw:         defaultCommunityBetaRaw("MM BoneKing"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "MM hj SK",
			Slug:        "mm-hj-sk",
			Description: defaultCommunityBetaDesc("MM hj SK"),
			Raw:         defaultCommunityBetaRaw("MM hj SK"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "MM hqs hj",
			Slug:        "mm-hqs-hj",
			Description: defaultCommunityBetaDesc("MM hqs hj"),
			Raw:         defaultCommunityBetaRaw("MM hqs hj"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "MM hqs SF",
			Slug:        "mm-hqs-sf",
			Description: defaultCommunityBetaDesc("MM hqs SF"),
			Raw:         defaultCommunityBetaRaw("MM hqs SF"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "MM slb",
			Slug:        "mm-slb",
			Description: defaultCommunityBetaDesc("MM slb"),
			Raw:         defaultCommunityBetaRaw("MM slb"),
		})
		brands = append(brands, brand)
	}

	// Brand: Mellow (1 models)
	{
		brand := &PrinterBrand{
			Name:        "Mellow",
			Slug:        "mellow",
			Description: defaultBrandDesc("Mellow"),
			Color:       "14B8A6",
			Position:    42,
			Models:      make([]*PrinterBrandModel, 0, 1),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "M1",
			Slug:        "m1",
			Description: defaultCommunityBetaDesc("M1"),
			Raw:         defaultCommunityBetaRaw("M1"),
		})
		brands = append(brands, brand)
	}

	// Brand: Meltingplot (5 models)
	{
		brand := &PrinterBrand{
			Name:        "Meltingplot",
			Slug:        "meltingplot",
			Description: defaultBrandDesc("Meltingplot"),
			Color:       "EC4899",
			Position:    43,
			Models:      make([]*PrinterBrandModel, 0, 5),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Meltingplot CHX 350",
			Slug:        "meltingplot-chx-350",
			Description: defaultCommunityBetaDesc("Meltingplot CHX 350"),
			Raw:         defaultCommunityBetaRaw("Meltingplot CHX 350"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Meltingplot MBL 133",
			Slug:        "meltingplot-mbl-133",
			Description: defaultCommunityBetaDesc("Meltingplot MBL 133"),
			Raw:         defaultCommunityBetaRaw("Meltingplot MBL 133"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Meltingplot MBL 136",
			Slug:        "meltingplot-mbl-136",
			Description: defaultCommunityBetaDesc("Meltingplot MBL 136"),
			Raw:         defaultCommunityBetaRaw("Meltingplot MBL 136"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Meltingplot MBL 308",
			Slug:        "meltingplot-mbl-308",
			Description: defaultCommunityBetaDesc("Meltingplot MBL 308"),
			Raw:         defaultCommunityBetaRaw("Meltingplot MBL 308"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Meltingplot MBL 480",
			Slug:        "meltingplot-mbl-480",
			Description: defaultCommunityBetaDesc("Meltingplot MBL 480"),
			Raw:         defaultCommunityBetaRaw("Meltingplot MBL 480"),
		})
		brands = append(brands, brand)
	}

	// Brand: OpenEYE (1 models)
	{
		brand := &PrinterBrand{
			Name:        "OpenEYE",
			Slug:        "openeye",
			Description: defaultBrandDesc("OpenEYE"),
			Color:       "8B5CF6",
			Position:    44,
			Models:      make([]*PrinterBrandModel, 0, 1),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "OpenEYE Peacock V2",
			Slug:        "openeye-peacock-v2",
			Description: defaultCommunityBetaDesc("OpenEYE Peacock V2"),
			Raw:         defaultCommunityBetaRaw("OpenEYE Peacock V2"),
		})
		brands = append(brands, brand)
	}

	// Brand: Peopoly (1 models)
	{
		brand := &PrinterBrand{
			Name:        "Peopoly",
			Slug:        "peopoly",
			Description: defaultBrandDesc("Peopoly"),
			Color:       "6366F1",
			Position:    45,
			Models:      make([]*PrinterBrandModel, 0, 1),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Peopoly Magneto X",
			Slug:        "peopoly-magneto-x",
			Description: defaultCommunityBetaDesc("Peopoly Magneto X"),
			Raw:         defaultCommunityBetaRaw("Peopoly Magneto X"),
		})
		brands = append(brands, brand)
	}

	// Brand: Phrozen (1 models)
	{
		brand := &PrinterBrand{
			Name:        "Phrozen",
			Slug:        "phrozen",
			Description: defaultBrandDesc("Phrozen"),
			Color:       "10B981",
			Position:    46,
			Models:      make([]*PrinterBrandModel, 0, 1),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Phrozen Arco",
			Slug:        "phrozen-arco",
			Description: defaultCommunityBetaDesc("Phrozen Arco"),
			Raw:         defaultCommunityBetaRaw("Phrozen Arco"),
		})
		brands = append(brands, brand)
	}

	// Brand: Positron3D (1 models)
	{
		brand := &PrinterBrand{
			Name:        "Positron3D",
			Slug:        "positron3d",
			Description: defaultBrandDesc("Positron3D"),
			Color:       "F43F5E",
			Position:    47,
			Models:      make([]*PrinterBrandModel, 0, 1),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "The Positron",
			Slug:        "the-positron",
			Description: defaultCommunityBetaDesc("The Positron"),
			Raw:         defaultCommunityBetaRaw("The Positron"),
		})
		brands = append(brands, brand)
	}

	// Brand: Pragostroj (2 models)
	{
		brand := &PrinterBrand{
			Name:        "Pragostroj",
			Slug:        "pragostroj",
			Description: defaultBrandDesc("Pragostroj"),
			Color:       "D97706",
			Position:    48,
			Models:      make([]*PrinterBrandModel, 0, 2),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "KINARB 1HB",
			Slug:        "kinarb-1hb",
			Description: defaultCommunityBetaDesc("KINARB 1HB"),
			Raw:         defaultCommunityBetaRaw("KINARB 1HB"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "KINARB 2HB",
			Slug:        "kinarb-2hb",
			Description: defaultCommunityBetaDesc("KINARB 2HB"),
			Raw:         defaultCommunityBetaRaw("KINARB 2HB"),
		})
		brands = append(brands, brand)
	}

	// Brand: Printcepts (1 models)
	{
		brand := &PrinterBrand{
			Name:        "Printcepts",
			Slug:        "printcepts",
			Description: defaultBrandDesc("Printcepts"),
			Color:       "06B6D4",
			Position:    49,
			Models:      make([]*PrinterBrandModel, 0, 1),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "BabyBelt Pro",
			Slug:        "babybelt-pro",
			Description: defaultCommunityBetaDesc("BabyBelt Pro"),
			Raw:         defaultCommunityBetaRaw("BabyBelt Pro"),
		})
		brands = append(brands, brand)
	}

	// Brand: Raise3D (2 models)
	{
		brand := &PrinterBrand{
			Name:        "Raise3D",
			Slug:        "raise3d",
			Description: defaultBrandDesc("Raise3D"),
			Color:       "84CC16",
			Position:    50,
			Models:      make([]*PrinterBrandModel, 0, 2),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Raise3D Pro3",
			Slug:        "raise3d-pro3",
			Description: defaultCommunityBetaDesc("Raise3D Pro3"),
			Raw:         defaultCommunityBetaRaw("Raise3D Pro3"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Raise3D Pro3 Plus",
			Slug:        "raise3d-pro3-plus",
			Description: defaultCommunityBetaDesc("Raise3D Pro3 Plus"),
			Raw:         defaultCommunityBetaRaw("Raise3D Pro3 Plus"),
		})
		brands = append(brands, brand)
	}

	// Brand: re3D (6 models)
	{
		brand := &PrinterBrand{
			Name:        "re3D",
			Slug:        "re3d",
			Description: defaultBrandDesc("re3D"),
			Color:       "64748B",
			Position:    51,
			Models:      make([]*PrinterBrandModel, 0, 6),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "re3D Gigabot 4",
			Slug:        "re3d-gigabot-4",
			Description: defaultCommunityBetaDesc("re3D Gigabot 4"),
			Raw:         defaultCommunityBetaRaw("re3D Gigabot 4"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "re3D Gigabot 4 XLT",
			Slug:        "re3d-gigabot-4-xlt",
			Description: defaultCommunityBetaDesc("re3D Gigabot 4 XLT"),
			Raw:         defaultCommunityBetaRaw("re3D Gigabot 4 XLT"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "re3D GigabotX 2",
			Slug:        "re3d-gigabotx-2",
			Description: defaultCommunityBetaDesc("re3D GigabotX 2"),
			Raw:         defaultCommunityBetaRaw("re3D GigabotX 2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "re3D GigabotX 2 XLT",
			Slug:        "re3d-gigabotx-2-xlt",
			Description: defaultCommunityBetaDesc("re3D GigabotX 2 XLT"),
			Raw:         defaultCommunityBetaRaw("re3D GigabotX 2 XLT"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "re3D Terabot 4",
			Slug:        "re3d-terabot-4",
			Description: defaultCommunityBetaDesc("re3D Terabot 4"),
			Raw:         defaultCommunityBetaRaw("re3D Terabot 4"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "re3D TerabotX 2",
			Slug:        "re3d-terabotx-2",
			Description: defaultCommunityBetaDesc("re3D TerabotX 2"),
			Raw:         defaultCommunityBetaRaw("re3D TerabotX 2"),
		})
		brands = append(brands, brand)
	}

	// Brand: RH3D (2 models)
	{
		brand := &PrinterBrand{
			Name:        "RH3D",
			Slug:        "rh3d",
			Description: defaultBrandDesc("RH3D"),
			Color:       "F97316",
			Position:    52,
			Models:      make([]*PrinterBrandModel, 0, 2),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "E3NG v1.2S",
			Slug:        "e3ng-v12s",
			Description: defaultCommunityBetaDesc("E3NG v1.2S"),
			Raw:         defaultCommunityBetaRaw("E3NG v1.2S"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "VIRTU E3",
			Slug:        "virtu-e3",
			Description: defaultCommunityBetaDesc("VIRTU E3"),
			Raw:         defaultCommunityBetaRaw("VIRTU E3"),
		})
		brands = append(brands, brand)
	}

	// Brand: RolohaunDesign (2 models)
	{
		brand := &PrinterBrand{
			Name:        "RolohaunDesign",
			Slug:        "rolohaundesign",
			Description: defaultBrandDesc("RolohaunDesign"),
			Color:       "0EA5E9",
			Position:    53,
			Models:      make([]*PrinterBrandModel, 0, 2),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Rolohaun Delta Flyer Refit",
			Slug:        "rolohaun-delta-flyer-refit",
			Description: defaultCommunityBetaDesc("Rolohaun Delta Flyer Refit"),
			Raw:         defaultCommunityBetaRaw("Rolohaun Delta Flyer Refit"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Rook MK1 LDO",
			Slug:        "rook-mk1-ldo",
			Description: defaultCommunityBetaDesc("Rook MK1 LDO"),
			Raw:         defaultCommunityBetaRaw("Rook MK1 LDO"),
		})
		brands = append(brands, brand)
	}

	// Brand: SecKit (2 models)
	{
		brand := &PrinterBrand{
			Name:        "SecKit",
			Slug:        "seckit",
			Description: defaultBrandDesc("SecKit"),
			Color:       "22C55E",
			Position:    54,
			Models:      make([]*PrinterBrandModel, 0, 2),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Seckit Go3",
			Slug:        "seckit-go3",
			Description: defaultCommunityBetaDesc("Seckit Go3"),
			Raw:         defaultCommunityBetaRaw("Seckit Go3"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "SecKit SK-Tank",
			Slug:        "seckit-sk-tank",
			Description: defaultCommunityBetaDesc("SecKit SK-Tank"),
			Raw:         defaultCommunityBetaRaw("SecKit SK-Tank"),
		})
		brands = append(brands, brand)
	}

	// Brand: SeeMeCNC (7 models)
	{
		brand := &PrinterBrand{
			Name:        "SeeMeCNC",
			Slug:        "seemecnc",
			Description: defaultBrandDesc("SeeMeCNC"),
			Color:       "EF4444",
			Position:    55,
			Models:      make([]*PrinterBrandModel, 0, 7),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "SeeMeCNC Artemis 300",
			Slug:        "seemecnc-artemis-300",
			Description: defaultCommunityBetaDesc("SeeMeCNC Artemis 300"),
			Raw:         defaultCommunityBetaRaw("SeeMeCNC Artemis 300"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "SeeMeCNC BOSSdelta 300",
			Slug:        "seemecnc-bossdelta-300",
			Description: defaultCommunityBetaDesc("SeeMeCNC BOSSdelta 300"),
			Raw:         defaultCommunityBetaRaw("SeeMeCNC BOSSdelta 300"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "SeeMeCNC BOSSdelta 500 0505",
			Slug:        "seemecnc-bossdelta-500-0505",
			Description: defaultCommunityBetaDesc("SeeMeCNC BOSSdelta 500 0505"),
			Raw:         defaultCommunityBetaRaw("SeeMeCNC BOSSdelta 500 0505"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "SeeMeCNC BOSSdelta 500 0510",
			Slug:        "seemecnc-bossdelta-500-0510",
			Description: defaultCommunityBetaDesc("SeeMeCNC BOSSdelta 500 0510"),
			Raw:         defaultCommunityBetaRaw("SeeMeCNC BOSSdelta 500 0510"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "SeeMeCNC BOSSdelta 500 0521",
			Slug:        "seemecnc-bossdelta-500-0521",
			Description: defaultCommunityBetaDesc("SeeMeCNC BOSSdelta 500 0521"),
			Raw:         defaultCommunityBetaRaw("SeeMeCNC BOSSdelta 500 0521"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "SeeMeCNC RostockMAX v3.2",
			Slug:        "seemecnc-rostockmax-v32",
			Description: defaultCommunityBetaDesc("SeeMeCNC RostockMAX v3.2"),
			Raw:         defaultCommunityBetaRaw("SeeMeCNC RostockMAX v3.2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "SeeMeCNC RostockMAX v4",
			Slug:        "seemecnc-rostockmax-v4",
			Description: defaultCommunityBetaDesc("SeeMeCNC RostockMAX v4"),
			Raw:         defaultCommunityBetaRaw("SeeMeCNC RostockMAX v4"),
		})
		brands = append(brands, brand)
	}

	// Brand: Tiertime (4 models)
	{
		brand := &PrinterBrand{
			Name:        "Tiertime",
			Slug:        "tiertime",
			Description: defaultBrandDesc("Tiertime"),
			Color:       "3B82F6",
			Position:    56,
			Models:      make([]*PrinterBrandModel, 0, 4),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Tiertime UP300 HS",
			Slug:        "tiertime-up300-hs",
			Description: defaultCommunityBetaDesc("Tiertime UP300 HS"),
			Raw:         defaultCommunityBetaRaw("Tiertime UP300 HS"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Tiertime UP310 Pro",
			Slug:        "tiertime-up310-pro",
			Description: defaultCommunityBetaDesc("Tiertime UP310 Pro"),
			Raw:         defaultCommunityBetaRaw("Tiertime UP310 Pro"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Tiertime UP400 Pro",
			Slug:        "tiertime-up400-pro",
			Description: defaultCommunityBetaDesc("Tiertime UP400 Pro"),
			Raw:         defaultCommunityBetaRaw("Tiertime UP400 Pro"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Tiertime UP600 HS",
			Slug:        "tiertime-up600-hs",
			Description: defaultCommunityBetaDesc("Tiertime UP600 HS"),
			Raw:         defaultCommunityBetaRaw("Tiertime UP600 HS"),
		})
		brands = append(brands, brand)
	}

	// Brand: Tronxy (1 models)
	{
		brand := &PrinterBrand{
			Name:        "Tronxy",
			Slug:        "tronxy",
			Description: defaultBrandDesc("Tronxy"),
			Color:       "A855F7",
			Position:    57,
			Models:      make([]*PrinterBrandModel, 0, 1),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Tronxy X5SA 400 Marlin Firmware",
			Slug:        "tronxy-x5sa-400-marlin-firmware",
			Description: defaultCommunityBetaDesc("Tronxy X5SA 400 Marlin Firmware"),
			Raw:         defaultCommunityBetaRaw("Tronxy X5SA 400 Marlin Firmware"),
		})
		brands = append(brands, brand)
	}

	// Brand: TwoTrees (2 models)
	{
		brand := &PrinterBrand{
			Name:        "TwoTrees",
			Slug:        "twotrees",
			Description: defaultBrandDesc("TwoTrees"),
			Color:       "F59E0B",
			Position:    58,
			Models:      make([]*PrinterBrandModel, 0, 2),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "TwoTrees SK1",
			Slug:        "twotrees-sk1",
			Description: defaultCommunityBetaDesc("TwoTrees SK1"),
			Raw:         defaultCommunityBetaRaw("TwoTrees SK1"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "TwoTrees SP-5 Klipper",
			Slug:        "twotrees-sp-5-klipper",
			Description: defaultCommunityBetaDesc("TwoTrees SP-5 Klipper"),
			Raw:         defaultCommunityBetaRaw("TwoTrees SP-5 Klipper"),
		})
		brands = append(brands, brand)
	}

	// Brand: Vivedino (6 models)
	{
		brand := &PrinterBrand{
			Name:        "Vivedino",
			Slug:        "vivedino",
			Description: defaultBrandDesc("Vivedino"),
			Color:       "14B8A6",
			Position:    59,
			Models:      make([]*PrinterBrandModel, 0, 6),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Troodon 2.0 - Klipper",
			Slug:        "troodon-20-klipper",
			Description: defaultCommunityBetaDesc("Troodon 2.0 - Klipper"),
			Raw:         defaultCommunityBetaRaw("Troodon 2.0 - Klipper"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Troodon 2.0 - RRF",
			Slug:        "troodon-20-rrf",
			Description: defaultCommunityBetaDesc("Troodon 2.0 - RRF"),
			Raw:         defaultCommunityBetaRaw("Troodon 2.0 - RRF"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Xplorer Dual Gantry",
			Slug:        "xplorer-dual-gantry",
			Description: defaultCommunityBetaDesc("Xplorer Dual Gantry"),
			Raw:         defaultCommunityBetaRaw("Xplorer Dual Gantry"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Xplorer IDEX",
			Slug:        "xplorer-idex",
			Description: defaultCommunityBetaDesc("Xplorer IDEX"),
			Raw:         defaultCommunityBetaRaw("Xplorer IDEX"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Xplorer IQEX",
			Slug:        "xplorer-iqex",
			Description: defaultCommunityBetaDesc("Xplorer IQEX"),
			Raw:         defaultCommunityBetaRaw("Xplorer IQEX"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Xplorer Single",
			Slug:        "xplorer-single",
			Description: defaultCommunityBetaDesc("Xplorer Single"),
			Raw:         defaultCommunityBetaRaw("Xplorer Single"),
		})
		brands = append(brands, brand)
	}

	// Brand: Volumic (21 models)
	{
		brand := &PrinterBrand{
			Name:        "Volumic",
			Slug:        "volumic",
			Description: defaultBrandDesc("Volumic"),
			Color:       "EC4899",
			Position:    60,
			Models:      make([]*PrinterBrandModel, 0, 21),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "EXO42",
			Slug:        "exo42",
			Description: defaultCommunityBetaDesc("EXO42"),
			Raw:         defaultCommunityBetaRaw("EXO42"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "EXO42 IDRE",
			Slug:        "exo42-idre",
			Description: defaultCommunityBetaDesc("EXO42 IDRE"),
			Raw:         defaultCommunityBetaRaw("EXO42 IDRE"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "EXO42 Performance",
			Slug:        "exo42-performance",
			Description: defaultCommunityBetaDesc("EXO42 Performance"),
			Raw:         defaultCommunityBetaRaw("EXO42 Performance"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "EXO42 Stage 2",
			Slug:        "exo42-stage-2",
			Description: defaultCommunityBetaDesc("EXO42 Stage 2"),
			Raw:         defaultCommunityBetaRaw("EXO42 Stage 2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "EXO65",
			Slug:        "exo65",
			Description: defaultCommunityBetaDesc("EXO65"),
			Raw:         defaultCommunityBetaRaw("EXO65"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "EXO65 IDRE",
			Slug:        "exo65-idre",
			Description: defaultCommunityBetaDesc("EXO65 IDRE"),
			Raw:         defaultCommunityBetaRaw("EXO65 IDRE"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "EXO65 Performance",
			Slug:        "exo65-performance",
			Description: defaultCommunityBetaDesc("EXO65 Performance"),
			Raw:         defaultCommunityBetaRaw("EXO65 Performance"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "EXO65 Stage 2",
			Slug:        "exo65-stage-2",
			Description: defaultCommunityBetaDesc("EXO65 Stage 2"),
			Raw:         defaultCommunityBetaRaw("EXO65 Stage 2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "SH65",
			Slug:        "sh65",
			Description: defaultCommunityBetaDesc("SH65"),
			Raw:         defaultCommunityBetaRaw("SH65"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "SH65 IDRE",
			Slug:        "sh65-idre",
			Description: defaultCommunityBetaDesc("SH65 IDRE"),
			Raw:         defaultCommunityBetaRaw("SH65 IDRE"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "SH65 Performance",
			Slug:        "sh65-performance",
			Description: defaultCommunityBetaDesc("SH65 Performance"),
			Raw:         defaultCommunityBetaRaw("SH65 Performance"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "SH65 Stage 2",
			Slug:        "sh65-stage-2",
			Description: defaultCommunityBetaDesc("SH65 Stage 2"),
			Raw:         defaultCommunityBetaRaw("SH65 Stage 2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "VS20MK2",
			Slug:        "vs20mk2",
			Description: defaultCommunityBetaDesc("VS20MK2"),
			Raw:         defaultCommunityBetaRaw("VS20MK2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "VS30MK2",
			Slug:        "vs30mk2",
			Description: defaultCommunityBetaDesc("VS30MK2"),
			Raw:         defaultCommunityBetaRaw("VS30MK2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "VS30MK3",
			Slug:        "vs30mk3",
			Description: defaultCommunityBetaDesc("VS30MK3"),
			Raw:         defaultCommunityBetaRaw("VS30MK3"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "VS30MK3 Stage 2",
			Slug:        "vs30mk3-stage-2",
			Description: defaultCommunityBetaDesc("VS30MK3 Stage 2"),
			Raw:         defaultCommunityBetaRaw("VS30MK3 Stage 2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "VS30SC",
			Slug:        "vs30sc",
			Description: defaultCommunityBetaDesc("VS30SC"),
			Raw:         defaultCommunityBetaRaw("VS30SC"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "VS30SC2",
			Slug:        "vs30sc2",
			Description: defaultCommunityBetaDesc("VS30SC2"),
			Raw:         defaultCommunityBetaRaw("VS30SC2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "VS30SC2 Performance",
			Slug:        "vs30sc2-performance",
			Description: defaultCommunityBetaDesc("VS30SC2 Performance"),
			Raw:         defaultCommunityBetaRaw("VS30SC2 Performance"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "VS30SC2 Stage 2",
			Slug:        "vs30sc2-stage-2",
			Description: defaultCommunityBetaDesc("VS30SC2 Stage 2"),
			Raw:         defaultCommunityBetaRaw("VS30SC2 Stage 2"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "VS30ULTRA",
			Slug:        "vs30ultra",
			Description: defaultCommunityBetaDesc("VS30ULTRA"),
			Raw:         defaultCommunityBetaRaw("VS30ULTRA"),
		})
		brands = append(brands, brand)
	}

	// Brand: Voxelab (1 models)
	{
		brand := &PrinterBrand{
			Name:        "Voxelab",
			Slug:        "voxelab",
			Description: defaultBrandDesc("Voxelab"),
			Color:       "8B5CF6",
			Position:    61,
			Models:      make([]*PrinterBrandModel, 0, 1),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Voxelab Aquila X2",
			Slug:        "voxelab-aquila-x2",
			Description: defaultCommunityBetaDesc("Voxelab Aquila X2"),
			Raw:         defaultCommunityBetaRaw("Voxelab Aquila X2"),
		})
		brands = append(brands, brand)
	}

	// Brand: Vzbot (2 models)
	{
		brand := &PrinterBrand{
			Name:        "Vzbot",
			Slug:        "vzbot",
			Description: defaultBrandDesc("Vzbot"),
			Color:       "6366F1",
			Position:    62,
			Models:      make([]*PrinterBrandModel, 0, 2),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Vzbot 235 AWD",
			Slug:        "vzbot-235-awd",
			Description: defaultCommunityBetaDesc("Vzbot 235 AWD"),
			Raw:         defaultCommunityBetaRaw("Vzbot 235 AWD"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Vzbot 330 AWD",
			Slug:        "vzbot-330-awd",
			Description: defaultCommunityBetaDesc("Vzbot 330 AWD"),
			Raw:         defaultCommunityBetaRaw("Vzbot 330 AWD"),
		})
		brands = append(brands, brand)
	}

	// Brand: Wanhao (19 models)
	{
		brand := &PrinterBrand{
			Name:        "Wanhao",
			Slug:        "wanhao",
			Description: defaultBrandDesc("Wanhao"),
			Color:       "10B981",
			Position:    63,
			Models:      make([]*PrinterBrandModel, 0, 19),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Wanhao D12-300",
			Slug:        "wanhao-d12-300",
			Description: defaultCommunityBetaDesc("Wanhao D12-300"),
			Raw:         defaultCommunityBetaRaw("Wanhao D12-300"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Wanhao D12 230 PRO M2 DIRECT",
			Slug:        "wanhao-d12-230-pro-m2-direct",
			Description: defaultCommunityBetaDesc("Wanhao D12 230 PRO M2 DIRECT"),
			Raw:         defaultCommunityBetaRaw("Wanhao D12 230 PRO M2 DIRECT"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Wanhao D12 230 PRO M2 MONO DUAL",
			Slug:        "wanhao-d12-230-pro-m2-mono-dual",
			Description: defaultCommunityBetaDesc("Wanhao D12 230 PRO M2 MONO DUAL"),
			Raw:         defaultCommunityBetaRaw("Wanhao D12 230 PRO M2 MONO DUAL"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Wanhao D12 230 PRO M2 MONO DUAL PoopTool",
			Slug:        "wanhao-d12-230-pro-m2-mono-dual-pooptool",
			Description: defaultCommunityBetaDesc("Wanhao D12 230 PRO M2 MONO DUAL PoopTool"),
			Raw:         defaultCommunityBetaRaw("Wanhao D12 230 PRO M2 MONO DUAL PoopTool"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Wanhao D12 230 PRO SMARTPAD DIRECT",
			Slug:        "wanhao-d12-230-pro-smartpad-direct",
			Description: defaultCommunityBetaDesc("Wanhao D12 230 PRO SMARTPAD DIRECT"),
			Raw:         defaultCommunityBetaRaw("Wanhao D12 230 PRO SMARTPAD DIRECT"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Wanhao D12 230 PRO SMARTPAD MONO DUAL",
			Slug:        "wanhao-d12-230-pro-smartpad-mono-dual",
			Description: defaultCommunityBetaDesc("Wanhao D12 230 PRO SMARTPAD MONO DUAL"),
			Raw:         defaultCommunityBetaRaw("Wanhao D12 230 PRO SMARTPAD MONO DUAL"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Wanhao D12 230 PRO SMARTPAD MONO DUAL PoopTool",
			Slug:        "wanhao-d12-230-pro-smartpad-mono-dual-pooptool",
			Description: defaultCommunityBetaDesc("Wanhao D12 230 PRO SMARTPAD MONO DUAL PoopTool"),
			Raw:         defaultCommunityBetaRaw("Wanhao D12 230 PRO SMARTPAD MONO DUAL PoopTool"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Wanhao D12 300 PRO M2 DIRECT",
			Slug:        "wanhao-d12-300-pro-m2-direct",
			Description: defaultCommunityBetaDesc("Wanhao D12 300 PRO M2 DIRECT"),
			Raw:         defaultCommunityBetaRaw("Wanhao D12 300 PRO M2 DIRECT"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Wanhao D12 300 PRO M2 MONO DUAL",
			Slug:        "wanhao-d12-300-pro-m2-mono-dual",
			Description: defaultCommunityBetaDesc("Wanhao D12 300 PRO M2 MONO DUAL"),
			Raw:         defaultCommunityBetaRaw("Wanhao D12 300 PRO M2 MONO DUAL"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Wanhao D12 300 PRO M2 MONO DUAL PoopTool",
			Slug:        "wanhao-d12-300-pro-m2-mono-dual-pooptool",
			Description: defaultCommunityBetaDesc("Wanhao D12 300 PRO M2 MONO DUAL PoopTool"),
			Raw:         defaultCommunityBetaRaw("Wanhao D12 300 PRO M2 MONO DUAL PoopTool"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Wanhao D12 300 PRO SMARTPAD DIRECT",
			Slug:        "wanhao-d12-300-pro-smartpad-direct",
			Description: defaultCommunityBetaDesc("Wanhao D12 300 PRO SMARTPAD DIRECT"),
			Raw:         defaultCommunityBetaRaw("Wanhao D12 300 PRO SMARTPAD DIRECT"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Wanhao D12 300 PRO SMARTPAD MONO DUAL",
			Slug:        "wanhao-d12-300-pro-smartpad-mono-dual",
			Description: defaultCommunityBetaDesc("Wanhao D12 300 PRO SMARTPAD MONO DUAL"),
			Raw:         defaultCommunityBetaRaw("Wanhao D12 300 PRO SMARTPAD MONO DUAL"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Wanhao D12 300 PRO SMARTPAD MONO DUAL PoopTool",
			Slug:        "wanhao-d12-300-pro-smartpad-mono-dual-pooptool",
			Description: defaultCommunityBetaDesc("Wanhao D12 300 PRO SMARTPAD MONO DUAL PoopTool"),
			Raw:         defaultCommunityBetaRaw("Wanhao D12 300 PRO SMARTPAD MONO DUAL PoopTool"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Wanhao D12 500 PRO M2 DIRECT",
			Slug:        "wanhao-d12-500-pro-m2-direct",
			Description: defaultCommunityBetaDesc("Wanhao D12 500 PRO M2 DIRECT"),
			Raw:         defaultCommunityBetaRaw("Wanhao D12 500 PRO M2 DIRECT"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Wanhao D12 500 PRO M2 MONO DUAL",
			Slug:        "wanhao-d12-500-pro-m2-mono-dual",
			Description: defaultCommunityBetaDesc("Wanhao D12 500 PRO M2 MONO DUAL"),
			Raw:         defaultCommunityBetaRaw("Wanhao D12 500 PRO M2 MONO DUAL"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Wanhao D12 500 PRO M2 MONO DUAL PoopTool",
			Slug:        "wanhao-d12-500-pro-m2-mono-dual-pooptool",
			Description: defaultCommunityBetaDesc("Wanhao D12 500 PRO M2 MONO DUAL PoopTool"),
			Raw:         defaultCommunityBetaRaw("Wanhao D12 500 PRO M2 MONO DUAL PoopTool"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Wanhao D12 500 PRO SMARTPAD DIRECT",
			Slug:        "wanhao-d12-500-pro-smartpad-direct",
			Description: defaultCommunityBetaDesc("Wanhao D12 500 PRO SMARTPAD DIRECT"),
			Raw:         defaultCommunityBetaRaw("Wanhao D12 500 PRO SMARTPAD DIRECT"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Wanhao D12 500 PRO SMARTPAD MONO DUAL",
			Slug:        "wanhao-d12-500-pro-smartpad-mono-dual",
			Description: defaultCommunityBetaDesc("Wanhao D12 500 PRO SMARTPAD MONO DUAL"),
			Raw:         defaultCommunityBetaRaw("Wanhao D12 500 PRO SMARTPAD MONO DUAL"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Wanhao D12 500 PRO SMARTPAD MONO DUAL PoopTool",
			Slug:        "wanhao-d12-500-pro-smartpad-mono-dual-pooptool",
			Description: defaultCommunityBetaDesc("Wanhao D12 500 PRO SMARTPAD MONO DUAL PoopTool"),
			Raw:         defaultCommunityBetaRaw("Wanhao D12 500 PRO SMARTPAD MONO DUAL PoopTool"),
		})
		brands = append(brands, brand)
	}

	// Brand: WEMAKE3D (2 models)
	{
		brand := &PrinterBrand{
			Name:        "WEMAKE3D",
			Slug:        "wemake3d",
			Description: defaultBrandDesc("WEMAKE3D"),
			Color:       "F43F5E",
			Position:    64,
			Models:      make([]*PrinterBrandModel, 0, 2),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "WEMAKE3D PhoenixProV1",
			Slug:        "wemake3d-phoenixprov1",
			Description: defaultCommunityBetaDesc("WEMAKE3D PhoenixProV1"),
			Raw:         defaultCommunityBetaRaw("WEMAKE3D PhoenixProV1"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "WEMAKE3D TinyBotV1",
			Slug:        "wemake3d-tinybotv1",
			Description: defaultCommunityBetaDesc("WEMAKE3D TinyBotV1"),
			Raw:         defaultCommunityBetaRaw("WEMAKE3D TinyBotV1"),
		})
		brands = append(brands, brand)
	}

	// Brand: WonderMaker (3 models)
	{
		brand := &PrinterBrand{
			Name:        "WonderMaker",
			Slug:        "wondermaker",
			Description: defaultBrandDesc("WonderMaker"),
			Color:       "D97706",
			Position:    65,
			Models:      make([]*PrinterBrandModel, 0, 3),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "WonderMaker ZR",
			Slug:        "wondermaker-zr",
			Description: defaultCommunityBetaDesc("WonderMaker ZR"),
			Raw:         defaultCommunityBetaRaw("WonderMaker ZR"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "WonderMaker ZR Ultra",
			Slug:        "wondermaker-zr-ultra",
			Description: defaultCommunityBetaDesc("WonderMaker ZR Ultra"),
			Raw:         defaultCommunityBetaRaw("WonderMaker ZR Ultra"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "WonderMaker ZR Ultra S",
			Slug:        "wondermaker-zr-ultra-s",
			Description: defaultCommunityBetaDesc("WonderMaker ZR Ultra S"),
			Raw:         defaultCommunityBetaRaw("WonderMaker ZR Ultra S"),
		})
		brands = append(brands, brand)
	}

	// Brand: Z-Bolt (9 models)
	{
		brand := &PrinterBrand{
			Name:        "Z-Bolt",
			Slug:        "z-bolt",
			Description: defaultBrandDesc("Z-Bolt"),
			Color:       "06B6D4",
			Position:    66,
			Models:      make([]*PrinterBrandModel, 0, 9),
		}
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Z-Bolt S1000",
			Slug:        "z-bolt-s1000",
			Description: defaultCommunityBetaDesc("Z-Bolt S1000"),
			Raw:         defaultCommunityBetaRaw("Z-Bolt S1000"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Z-Bolt S1000 Dual",
			Slug:        "z-bolt-s1000-dual",
			Description: defaultCommunityBetaDesc("Z-Bolt S1000 Dual"),
			Raw:         defaultCommunityBetaRaw("Z-Bolt S1000 Dual"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Z-Bolt S300",
			Slug:        "z-bolt-s300",
			Description: defaultCommunityBetaDesc("Z-Bolt S300"),
			Raw:         defaultCommunityBetaRaw("Z-Bolt S300"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Z-Bolt S300 Dual",
			Slug:        "z-bolt-s300-dual",
			Description: defaultCommunityBetaDesc("Z-Bolt S300 Dual"),
			Raw:         defaultCommunityBetaRaw("Z-Bolt S300 Dual"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Z-Bolt S400",
			Slug:        "z-bolt-s400",
			Description: defaultCommunityBetaDesc("Z-Bolt S400"),
			Raw:         defaultCommunityBetaRaw("Z-Bolt S400"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Z-Bolt S400 Dual",
			Slug:        "z-bolt-s400-dual",
			Description: defaultCommunityBetaDesc("Z-Bolt S400 Dual"),
			Raw:         defaultCommunityBetaRaw("Z-Bolt S400 Dual"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Z-Bolt S600",
			Slug:        "z-bolt-s600",
			Description: defaultCommunityBetaDesc("Z-Bolt S600"),
			Raw:         defaultCommunityBetaRaw("Z-Bolt S600"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Z-Bolt S600 Dual",
			Slug:        "z-bolt-s600-dual",
			Description: defaultCommunityBetaDesc("Z-Bolt S600 Dual"),
			Raw:         defaultCommunityBetaRaw("Z-Bolt S600 Dual"),
		})
		brand.Models = append(brand.Models, &PrinterBrandModel{
			Name:        "Z-Bolt S800 Dual",
			Slug:        "z-bolt-s800-dual",
			Description: defaultCommunityBetaDesc("Z-Bolt S800 Dual"),
			Raw:         defaultCommunityBetaRaw("Z-Bolt S800 Dual"),
		})
		brands = append(brands, brand)
	}

	return brands
}

func init() {
	printerBrands = buildPrinterBrands()
}
