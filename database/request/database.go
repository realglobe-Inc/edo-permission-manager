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

import ()

// バックエンドのデータを他のプログラムと共有する前提。

// アクセス権限の変更要求の格納庫。
type Db interface {
	// 取得。
	// 最大で max 個まで。
	GetByOwner(owner string, max int) ([]*Element, error)

	// 数を取得。
	CountOfOwner(owner string) (int, error)

	// 保存。
	Save(elem *Element) error

	// 削除。
	Remove(elem *Element) error

	// 同内容のものが存在するか。
	Exist(elem *Element) (bool, error)
}
