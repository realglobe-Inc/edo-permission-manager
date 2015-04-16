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

package request

import (
	"github.com/realglobe-Inc/edo-permission-manager/database/permission"
)

// アクセス権限の変更要求。
type Element struct {
	id string
	// リソースの識別子。
	rsrc *permission.Identifier
	// アクセス主体 -> アクセス元 TA の ID。
	acss map[string][]string
	// mod
	mod string
	// リクエスト元。
	src *Source
}

func New(id string, rsrc *permission.Identifier, acss map[string][]string, mod string, src *Source) *Element {
	return &Element{
		id:   id,
		rsrc: rsrc,
		acss: acss,
		mod:  mod,
		src:  src,
	}
}

// ID を返す。
func (this *Element) Id() string {
	return this.id
}

// リソースの識別子を返す。
func (this *Element) Resource() *permission.Identifier {
	return this.rsrc
}

// アクセス主体 -> アクセス元 TA の ID を返す。
func (this *Element) Accesors() map[string][]string {
	return this.acss
}

// mod を返す。
func (this *Element) Mode() string {
	return this.mod
}

// リクエスト元を返す。
func (this *Element) Source() *Source {
	return this.src
}
