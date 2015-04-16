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
	"net/url"
	"time"
)

// 変更コード情報。
type Element struct {
	id string
	// 無効にされたかどうか。
	inv bool
	// 有効期限。
	exp time.Time
	// 変更主体の ID。
	acnt string
	// 要求元 TA の ID。
	ta string
	// 変更内容。
	reqs []*Request
	// リクエストの state。
	stat string
	// リクエストの redirect_uri。
	rediUri *url.URL

	// 更新日時。
	date time.Time
}

func New(id string, exp time.Time, acnt, ta string, reqs []*Request, stat string, rediUri *url.URL) *Element {
	return &Element{
		id:      id,
		exp:     exp,
		acnt:    acnt,
		ta:      ta,
		reqs:    reqs,
		stat:    stat,
		rediUri: rediUri,
		date:    time.Now(),
	}
}

// ID を返す。
func (this *Element) Id() string {
	return this.id
}

// 無効にされているかどうか。
func (this *Element) Invalid() bool {
	return this.inv
}

// 無効にする。
func (this *Element) Invalidate() {
	this.inv = true
	this.date = time.Now()
}

// 有効期限を返す。
func (this *Element) ExpiresIn() time.Time {
	return this.exp
}

// 変更主体の ID を返す。
func (this *Element) Account() string {
	return this.acnt
}

// 要求元 TA の ID。
func (this *Element) Ta() string {
	return this.ta
}

// 変更内容を返す。
func (this *Element) Requests() []*Request {
	return this.reqs
}

// リクエストの state を返す。
func (this *Element) State() string {
	return this.stat
}

// リクエストの redirect_uri を返す。
func (this *Element) RedirectUri() *url.URL {
	return this.rediUri
}

// 更新日時を返す。
func (this *Element) Date() time.Time {
	return this.date
}
