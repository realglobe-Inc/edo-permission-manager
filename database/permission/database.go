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

// バックエンドのデータを他のプログラムと共有する前提。

// アクセス権限の格納庫。
type Db interface {
	// 取得。
	// アクセス主体による権限のフィルタリングが可能。
	Get(id *Identifier, acsAcnts []string) (*Element, error)

	// 保存。
	// 既にリソースのアクセス権限が存在する場合、elem に設定されている権限のみ上書きする。
	Save(elem *Element) error
}
