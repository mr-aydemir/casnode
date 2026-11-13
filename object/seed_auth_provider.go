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
)

// InitMakerSpellAuthProvider seeds the central MakerSpell ID single sign-on
// provider (id.makerspell.com) so the login and signup buttons on the forum
// redirect through it, the same as the main website. It is idempotent: when a
// provider called "makerspell" already exists it is left alone, so anything
// changed later in the admin UI (/admin/logins) is kept on re-deploys.
//
// The provider is OAuth 2.0 with PKCE (no client secret is validated by
// id.makerspell.com), and the user info carries the standard OpenID Connect
// claims (sub, preferred_username, name, email, picture).
//
// Seeding is best-effort: a failure here must never prevent the forum from
// booting.
func InitMakerSpellAuthProvider() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[seed] MakerSpell ID auth provider seeding aborted (non-fatal): %v\n", r)
		}
	}()

	existing, err := GetAuthProvider("makerspell")
	if err != nil {
		panic(err)
	}
	if existing != nil {
		return
	}

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

	_, err = adapter.Engine.Insert(provider)
	if err != nil {
		panic(err)
	}

	fmt.Println("[seed] MakerSpell ID auth provider created (central login via id.makerspell.com)")
}
