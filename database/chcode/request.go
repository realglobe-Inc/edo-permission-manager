// Copyright 2015 realglobe, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package chcode

import (
	"github.com/realglobe-Inc/edo-permission-manager/database/permission"
)

// アクセス権限の変更対象。
type Request struct {
	// 変更対象タグ。
	tag string
	// リソースの識別子。
	rsrc *permission.Identifier
	// アクセス主体 -> アクセス元 TA の ID。
	acss map[string][]string
	// mod
	mod string
	// 必須かどうか。
	ess bool
	// リソースの存在を検査済みかどうか。
	exist bool
}

func NewRequest(tag string, rsrc *permission.Identifier, acss map[string][]string, mod string, ess, exist bool) *Request {
	return &Request{
		tag:   tag,
		rsrc:  rsrc,
		acss:  acss,
		mod:   mod,
		ess:   ess,
		exist: exist,
	}
}

// 変更対象タグを返す。
func (this *Request) Tag() string {
	return this.tag
}

// リソースの識別子を返す。
func (this *Request) Resource() *permission.Identifier {
	return this.rsrc
}

// アクセス主体 -> アクセス元 TA の ID を返す。
func (this *Request) Accesors() map[string][]string {
	return this.acss
}

// mod を返す。
func (this *Request) Mode() string {
	return this.mod
}

// 必須かどうか。
func (this *Request) Essential() bool {
	return this.ess
}

// リソースの存在を検査済みかどうか。
func (this *Request) Exist() bool {
	return this.exist
}
