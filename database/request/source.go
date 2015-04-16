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
	"time"
)

// 変更要求のリクエスト元。
type Source struct {
	// 要求主体の ID。
	acnt string
	// 要求元 TA の ID。
	ta string
	// 要求日時。
	date time.Time
}

func NewSource(acnt, ta string) *Source {
	return &Source{
		acnt: acnt,
		ta:   ta,
		date: time.Now(),
	}
}

// 要求主体の ID を返す。
func (this *Source) Account() string {
	return this.acnt
}

// 要求元 TA の ID を返す。
func (this *Source) Ta() string {
	return this.ta
}

// 要求日時を返す。
func (this *Source) Date() time.Time {
	return this.date
}
