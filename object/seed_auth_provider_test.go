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

func TestMakerSpellAuthProviderSpec(t *testing.T) {
	provider := &AuthProvider{
		Name:        "makerspell",
		DisplayName: "MakerSpell ID",
		Type:        "oauth2",
		ClientId:    "casnode",
		AuthUrl:     "https://id.makerspell.com/api/oauth/authorize",
		TokenUrl:    "https://id.makerspell.com/api/oauth/token",
		UserInfoUrl: "https://id.makerspell.com/api/oauth/userinfo",
		Scopes:      "openid profile email",
		IdField:     "sub",
		AllowSignup: true,
		Enabled:     true,
		SortOrder:   1,
	}

	if provider.Name != "makerspell" {
		t.Errorf("expected provider name 'makerspell', got %q", provider.Name)
	}
	if !strings.Contains(provider.AuthUrl, "id.makerspell.com") {
		t.Errorf("AuthUrl must point to id.makerspell.com, got %q", provider.AuthUrl)
	}
	if !provider.Enabled {
		t.Errorf("MakerSpell auth provider must be enabled by default")
	}
	if !provider.AllowSignup {
		t.Errorf("MakerSpell auth provider must allow signup")
	}
}
