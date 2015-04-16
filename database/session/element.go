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

package session

import (
	"time"
)

// セッション。
type Element struct {
	id string
	// 有効期限。
	exp time.Time
	// 変更コード。
	cod string
	// チケット。
	tic string
	// 合意結果。
	applied   []string
	forwarded []string
	denied    []string
	// 最後に選択された表示言語。
	lang string
}

func New(id string, exp time.Time) *Element {
	return &Element{
		id:  id,
		exp: exp,
	}
}

// 履歴を引き継いだセッションを作成する。
func (this *Element) New(id string, exp time.Time) *Element {
	elem := &Element{
		id:   id,
		exp:  exp,
		lang: this.lang,
	}
	return elem
}

// ID を返す。
func (this *Element) Id() string {
	return this.id
}

// 有効期限を返す。
func (this *Element) ExpiresIn() time.Time {
	return this.exp
}

// 変更コードを返す。
func (this *Element) Code() string {
	return this.cod
}

// 変更コードを保存する。
func (this *Element) SetCode(cod string) {
	this.cod = cod
}

// チケットを返す。
func (this *Element) Ticket() string {
	return this.tic
}

// チケットを保存する。
func (this *Element) SetTicket(tic string) {
	this.tic = tic
}

// 合意結果を返す。
func (this *Element) ConsentResult() (applied, forwarded, denied []string) {
	return this.applied, this.forwarded, this.denied
}

// 合意結果を保存する。
func (this *Element) SetConsentResult(applied, forwarded, denied []string) {
	this.applied = applied
	this.forwarded = forwarded
	this.denied = denied
}

// 最後に選択された表示言語を返す。
func (this *Element) Language() string {
	return this.lang
}

// 表示言語を保存する。
func (this *Element) SetLanguage(lang string) {
	this.lang = lang
}

// 一時データを消す。
func (this *Element) Clear() {
	this.cod = ""
	this.tic = ""
	this.applied = nil
	this.forwarded = nil
	this.denied = nil
}
