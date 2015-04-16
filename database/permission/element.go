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

package permission

import ()

// アクセス権限。
type Element struct {
	// リソースの識別子。
	id *Identifier

	// アクセス主体 -> アクセス元 TA の ID -> 権限。
	perm map[string]map[string]string
}

// リソースの識別子を返す。
func (this *Element) Id() *Identifier {
	return this.id
}

// アクセス主体 -> アクセス元 TA の ID -> 権限を返す。
func (this *Element) Permissions() map[string]map[string]string {
	return this.perm
}
