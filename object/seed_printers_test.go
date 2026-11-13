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
	"strings"
	"testing"
)

// TestPrinterCatalogStructure verifies the integrity and safety of the
// entire 66-brand, 438-model printer catalog.
func TestPrinterCatalogStructure(t *testing.T) {
	brands := buildPrinterBrands()
	if len(brands) == 0 {
		t.Fatalf("buildPrinterBrands() returned empty slice")
	}

	expectedBrands := 66
	if len(brands) != expectedBrands {
		t.Errorf("expected %d brands, got %d", expectedBrands, len(brands))
	}

	totalModels := 0
	brandSlugs := make(map[string]string)
	modelSlugs := make(map[string]string)

	verifiedHardware := map[string]bool{
		"Creality K1":          true,
		"Creality K1C":         true,
		"Creality K1 Max":      true,
		"Snapmaker U1":         true,
		"Snapmaker High Combo": true,
	}

	for _, b := range brands {
		if strings.TrimSpace(b.Name) == "" {
			t.Errorf("empty brand name encountered: %+v", b)
		}
		if strings.TrimSpace(b.Slug) == "" {
			t.Errorf("empty brand slug encountered for brand %q", b.Name)
		}
		if prev, exists := brandSlugs[b.Slug]; exists {
			t.Errorf("brand slug collision: %q used by both %q and %q", b.Slug, prev, b.Name)
		}
		brandSlugs[b.Slug] = b.Name

		if len(b.Models) == 0 {
			t.Errorf("brand %q has 0 models", b.Name)
		}

		for _, m := range b.Models {
			totalModels++
			if strings.TrimSpace(m.Name) == "" {
				t.Errorf("empty model name in brand %q", b.Name)
			}
			if strings.TrimSpace(m.Slug) == "" {
				t.Errorf("empty model slug for %q in brand %q", m.Name, b.Name)
			}
			if prev, exists := modelSlugs[m.Slug]; exists {
				t.Errorf("model slug collision: %q used by both %q and %q", m.Slug, prev, m.Name)
			}
			modelSlugs[m.Slug] = m.Name

			// Safety Rule: Verified vs Community Beta checks
			if verifiedHardware[m.Name] {
				if !strings.Contains(m.Raw, "Verified") {
					t.Errorf("verified printer %q must have Verified badge in starter topic", m.Name)
				}
			} else {
				// Community beta models must enforce Safe-Mode (Upload-Only)
				if !strings.Contains(m.Raw, "Community beta") || !strings.Contains(m.Raw, "upload-only") {
					t.Errorf("unverified printer %q must enforce Safe-Mode (Community beta / upload-only)", m.Name)
				}
			}

			// Validate markdown cooking does not panic or produce empty output
			cooked := CookMarkdown(m.Raw)
			if strings.TrimSpace(cooked.Html) == "" {
				t.Errorf("CookMarkdown returned empty HTML for model %q", m.Name)
			}
		}
	}

	expectedMinModels := 430
	if totalModels < expectedMinModels {
		t.Errorf("expected at least %d models, got %d", expectedMinModels, totalModels)
	}

	t.Logf("Catalog integrity check passed: %d brands, %d models verified without collisions.", len(brands), totalModels)
}

// TestSlugUniquenessAcrossBrandsAndModels ensures model slugs do not collide with brand slugs.
func TestSlugUniquenessAcrossBrandsAndModels(t *testing.T) {
	brands := buildPrinterBrands()
	allSlugs := make(map[string]string)

	for _, b := range brands {
		if prev, exists := allSlugs[b.Slug]; exists {
			t.Errorf("slug collision: %q used by brand %q and %s", b.Slug, b.Name, prev)
		}
		allSlugs[b.Slug] = "brand: " + b.Name

		for _, m := range b.Models {
			if prev, exists := allSlugs[m.Slug]; exists {
				t.Errorf("slug collision: %q used by model %q and %s", m.Slug, m.Name, prev)
			}
			allSlugs[m.Slug] = "model: " + m.Name
		}
	}
}

// TestSeederTemplates verifies that default templates render model names accurately.
func TestSeederTemplates(t *testing.T) {
	model := "TestPrinter 3000"
	raw := defaultCommunityBetaRaw(model)
	if !strings.Contains(raw, "TestPrinter 3000") {
		t.Errorf("defaultCommunityBetaRaw missing model name")
	}
	if !strings.Contains(raw, "upload-only mode") {
		t.Errorf("defaultCommunityBetaRaw missing upload-only safety requirement")
	}

	desc := defaultCommunityBetaDesc(model)
	if !strings.Contains(desc, "community beta") {
		t.Errorf("defaultCommunityBetaDesc missing community beta tag")
	}
}

// TestSortCategoriesForDisplayHierarchy verifies that 3-level categories
// (hardware -> brand -> model) are properly ordered and NOT dropped.
func TestSortCategoriesForDisplayHierarchy(t *testing.T) {
	categories := []*Category{
		{Id: 1, Name: "General", ParentId: 0, Position: 1},
		{Id: 2, Name: "Hardware", ParentId: 0, Position: 2},
		{Id: 20, Name: "Creality", ParentId: 2, Position: 1},
		{Id: 200, Name: "Creality K1", ParentId: 20, Position: 1},
		{Id: 201, Name: "Creality K1 Max", ParentId: 20, Position: 2},
		{Id: 21, Name: "Snapmaker", ParentId: 2, Position: 2},
		{Id: 210, Name: "Snapmaker U1", ParentId: 21, Position: 1},
		{Id: 3, Name: "Showcase", ParentId: 0, Position: 3},
	}

	sorted := SortCategoriesForDisplay(categories)
	if len(sorted) != len(categories) {
		t.Fatalf("expected %d categories after sorting, got %d (level 3 categories were dropped!)", len(categories), len(sorted))
	}

	expectedOrder := []int64{1, 2, 20, 200, 201, 21, 210, 3}
	for i, cat := range sorted {
		if cat.Id != expectedOrder[i] {
			t.Errorf("at index %d: expected category id %d (%s), got %d (%s)", i, expectedOrder[i], cat.Name, cat.Id, cat.Name)
		}
	}
}
