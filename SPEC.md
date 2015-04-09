<!--
Copyright 2015 realglobe, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
-->


# edo-permission-manager の仕様（目標）

[PDS 権限変更プロトコル]を担うサーバー。

以降の動作記述において、箇条書きに以下の構造を持たせることがある。

* if
    * then
* else if
    * then
* else


## 1. エンドポイント

|エンドポイント名|初期 URI|機能|
|:--|:--|:--|
|変更要請|/api/chmod|変更コードを発行する|
|変更|/chmod|変更に対する合意処理を始める|
|変更合意|/chmod/agree|変更に対する合意を受け付ける|
|変更合意 UI|/ui/chmod/agree.html|変更に対する合意のための UI を提供する|
|変更対象|/api/target/chmod|UI 用に変更対象を提供する|
|変更要求合意|/request/agree|変更要求に対する合意を受け付ける|
|変更要求合意 UI|/ui/request/agree.html|変更要求に対する合意のための UI を提供する|
|変更要求|/api/ticket|UI 用に変更要求合意チケットを発行する|
|変更要求対象数|/api/target/request/count|UI 用に変更要求の対象数を提供する|
|変更要求対象|/api/target/request|UI 用に変更要求の対象を提供する|
|アカウント情報|/api/info/user|UI 用にアカウント情報を提供する|
|TA 情報|/api/info/ta|UI 用に TA 情報を提供する|


## 2. edo-auth の使用

変更要請エンドポイントは [edo-auth] TA を通すことを前提とする。
変更、変更合意、変更合意 UI、変更要求合意、変更要求合意 UI、変更要求エンドポイントは [edo-auth] user を通すことを前提とする。

変更要求エンドポイントで [edo-auth] user を通す理由は、そこでセッションを発行する際にユーザー認証済みであることを保証するため。


## 3. セッション

変更、変更合意、変更対象、変更要求合意、変更要求、変更要求対象数、変更要求対象、アカウント情報エンドポイントではセッションを利用する。

|Cookie 名|値|
|:--|:--|
|Permission-Manager|セッション ID|

変更、変更合意、変更要求合意、変更要求エンドポイントへのリクエスト時に、セッション ID が通知されなかった場合、セッションを発行する。
セッションの期限に余裕がない場合、設定を引き継いだセッションを発行する。

変更、変更合意、変更要求合意、変更要求エンドポイントからのレスポンス時に、未通知のセッション ID を通知する。

変更要求エンドポイントでもセッションの発行・更新を行う理由は、変更要求合意 UI エンドポイントへ直にアクセスしてきたときにセッションを張るため。


## 4. 変更要請エンドポイント

[PDS 権限変更プロトコル]も参照のこと。

* リクエストに問題がある場合、
    * エラーを返す。
* そうでなく、非明示的に合意できる場合、
    * そのためのエラーを返す。
      [PDS 権限変更プロトコル]を参照。
* そうでなく、`check_exist` が真の変更対象があり、そのリソースが存在しない場合、
    * エラーを返す。
* そうでなければ、変更コードを発行する。
  変更コードに [edo-auth] TA 由来のアカウント情報、要請元 TA の ID、変更内容を紐付ける。
  変更コードを返す。


### 4.1. リソースの存在確認

HEAD メソッドでバックエンドに問い合わせ、レスポンスのステータスが 200 OK 以外なら、存在しないと判断する。

問い合わせる際の URL の生成は、利用者側で置換して運用できるように独立させる。
標準では、[PDS データアクセス API] の URL によるデータ指定の URL になる。

問い合わせる際の HTTP ヘッダは元のリクエストのヘッダを引き継ぐ。


#### 4.1.1. リソースの存在確認リクエスト例

```http
HEAD /api/resource/user/https%3A%2F%2Fwriter.example.org/profile HTTP/1.1
Host: 127.0.0.1:8000
```

[edo-auth] TA による追加ヘッダは省いている。


### 4.2. リクエスト例

```http
POST /chmod HTTP/1.1
Host: ta.example.org
Content-Type: application/json

{
    "chmod": {
        "profile": {
            "user_tag": "user",
            "ta": "https://writer.example.org",
            "path": "/profile",
            "mod": "+r",
            "essential": true
        },
        "diary": {
            "user_tag": "user",
            "ta": "https://writer.example.org",
            "path": "/diary",
            "mod": "+r"
        }
    },
    "redirect_uri": "https://from.example.org/chmod/return",
    "state": "SiuR29g1Iu"
}
```

[edo-auth] TA による追加ヘッダは省いている。


### 4.3. レスポンス例

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
    "code": "SS2EVIbmR7X1IHiHzMInTINxGYi-OY"
}
```


## 5. 変更エンドポイント

* 変更コードが無い場合、
    * エラーを返す。
* そうでなければ、変更合意 UI エンドポイントにリダイレクトさせる。

変更合意 UI エンドポイントへのリダイレクト時には、変更合意チケットを発行する。
変更コードと変更合意チケットをセッションに紐付ける。
変更合意チケットをフラグメントとして付加した変更合意 UI エンドポイントにリダイレクトさせる。


### 5.1. リクエスト例

```http
GET /chmod?code=SS2EVIbmR7X1IHiHzMInTINxGYi-OY HTTP/1.1
Host: ta.example.org
```


### 5.2. レスポンス例

```http
HTTP/1.1 302 Found
Set-Cookie: Permission-Manager=QkjGKwnIIC_m8uG-bSDOkGrcvN6kH1
    Expires=Thu, 09 Apr 2015 11:11:14 GMT; Path=/; Secure; HttpOnly
Location: /ui/chmod/agree.html?target_num=2&request_num=1#th2rEioRZL
```


## 6. 変更合意エンドポイント

* 変更合意チケットと紐付くセッションでない場合、
    * エラーを返す。
* そうでなければ、リクエストから以下のパラメータを取り出す。

|パラメータ名|必要性|値|
|:--|:--|:--|
|**`ticket`**|必須|変更合意チケット|
|**`applied`**|該当するなら必須|適用で合意した変更対象タグの JSON 配列|
|**`forwarded`**|該当するなら必須|転送で合意した変更対象タグの JSON 配列|
|**`denied`**|該当するなら必須|拒否で合意した変更対象タグの JSON 配列|
|**`redirect_to_request`**|任意|続けて変更要求に対する合意を行うかどうか|
|**`locale`**|任意|選択された表示言語|

* チケットがセッションに紐付くものと異なる、または、合意に余分や不足がある場合、
    * エラーを返す。
* そうでなければ、合意を履行する。

* `redirect_to_request` が `true` の場合、
    * セッションに合意結果を紐付ける。
      変更要求合意 UI エンドポイントにリダイレクトさせる。
* そうでなければ、セッションに紐付く変更コードの `redirect_uri` に合意結果を加えてリダイレクトさせる。


### 6.1. 変更の適用

以下に注意。

* パスの部分木に対して適用する。
* 指定されたアクセス元の権限だけを変更する。
* 現状で権限が存在しない場合、まず直上のパスからコピーする。


### 6.2. リクエスト例

```http
POST /chmod/agree HTTP/1.1
Host: ta.example.org
Cookie: Permission-Manager=QkjGKwnIIC_m8uG-bSDOkGrcvN6kH1
Content-Type: application/x-www-form-urlencoded

ticket=th2rEioRZL&applid=%5B%22profile%22%5D&denied=%5B%22diary%22%5D
```


### 6.3. レスポンス例

```http
HTTP/1.1 302 Found
Location: https://from.example.org/chmod/return?applied=%5B%22profile%22%5D
    &denied=%5B%22diary%22%5D&state=SiuR29g1Iu
```


## 7. 変更合意 UI エンドポイント

以下のパラメータを受け付ける。

|パラメータ名|必要性|値|
|:--|:--|:--|
|**`target_num`**|必須|変更対象の数|
|**`request_num`**|任意|変更要求の全対象数|
|**`display`**|任意|[OpenID Connect Core 1.0 Section 3.1.2.1] の `display` と同じもの|
|**`locales`**|任意|[OpenID Connect Core 1.0 Section 3.1.2.1] の `ui_locales` と同じもの|

UI の目的は、変更エンドポイントに POST させること。


## 8. 変更対象エンドポイント

* 変更合意チケットと紐付くセッションでない場合、
    * エラーを返す。
* そうでなければ、リクエストから以下のパラメータを取り出す。

|パラメータ名|必要性|値|
|:--|:--|:--|
|**`ticket`**|必須|変更合意チケット|
|**`target`**|任意|空白区切りの取得する変更対象の番号。無指定なら 0 から対象数 - 1 まで|

* 変更合意チケットがセッションに紐付くものと異なる、または、対象番号が不当な場合、
    * エラーを返す。
* そうでなければ、対象を `target` の順番通りに JSON 配列で返す。

変更対象は以下の最上位要素を含む JSON オブジェクト。

* **`tag`**
    * 変更対象タグ。
* **`user`**
    * [PDS 権限変更プロトコル]の変更要請リクエストパラメータを参照のこと。
* **`ta`**
    * [PDS 権限変更プロトコル]の変更要請リクエストパラメータを参照のこと。
* **`path`**
    * [PDS 権限変更プロトコル]の変更要請リクエストパラメータを参照のこと。
* **`accessor`**
    * [PDS 権限変更プロトコル]の変更要請リクエストパラメータを参照のこと。
* **`mod`**
    * [PDS 権限変更プロトコル]の変更要請リクエストパラメータを参照のこと。
* **`essential`**
    * [PDS 権限変更プロトコル]の変更要請リクエストパラメータを参照のこと。
* **`choices`**
    * 可能な合意の種類の配列。
      以下の合意の種類がある。
        * `apply`
            * 適用。
              ユーザーが変更権限を持つ場合のみ。
        * `forward`
            * 転送。
              ユーザーが変更権限を持たない場合のみ。
        * `deny`
            * 拒否。
* **`requester`**
    * 要請元。
      以下の要素を含むオブジェクト。
        * `user`
            * 要請主体の ID。
        * `ta`
            * 要請元 TA の ID。
* **`exist`**
    * 存在が確認されたかどうか。


### 8.1. リクエスト例

```http
GET /api/target/chmod&ticket=th2rEioRZL HTTP/1.1
Host: ta.example.org
Cookie: Permission-Manager=QkjGKwnIIC_m8uG-bSDOkGrcvN6kH1
```


### 8.2. レスポンス例

```http
HTTP/1.1 200 OK
Content-Type: application/json

[
    {
        "tag": "profile",
        "user": "38BF35F5464C00F9",
        "ta": "https://writer.example.org",
        "path": "/profile",
        "mod": "+r",
        "accessor": {
            "38BF35F5464C00F9": [
                "https://from.example.org"
            ]
        },
        "essential": true,
        "choices": [
            "apply",
            "deny"
        ],
        "requester": {
            "user": "38BF35F5464C00F9",
            "ta": "https://from.example.org"
        }        
    },
    {
        "tag": "diary",
        "user": "38BF35F5464C00F9",
        "ta": "https://writer.example.org",
        "path": "/diary",
        "mod": "+r",
        "accessor": {
            "38BF35F5464C00F9": [
                "https://from.example.org"
            ]
        },
        "choices": [
            "apply",
            "deny"
        ],
        "requester": {
            "user": "38BF35F5464C00F9",
            "ta": "https://from.example.org"
        }        
    }
]
```


## 9. 変更要求合意エンドポイント

* 変更要求合意チケットと紐付くセッションでない場合、
    * エラーを返す。
* そうでなければ、リクエストから以下のパラメータを取り出す。

|パラメータ名|必要性|値|
|:--|:--|:--|
|**`ticket`**|必須|変更要求合意チケット|
|**`applied`**|該当するなら必須|適用で合意した変更対象タグの JSON 配列|
|**`denied`**|該当するなら必須|拒否で合意した変更対象タグの JSON 配列|
|**`postponed`**|任意|合意の無かった変更対象タグの JSON 配列|
|**`locale`**|任意|選択された表示言語|

* 変更要求合意チケットがセッションに紐付くものと異なる、または、正当でない変更対象タグある場合、
    * エラーを返す。
* そうでなければ、合意を履行し、変更要求を削除する。

* 合意結果に紐付くセッションの場合、
    * セッションに紐付く変更コードの `redirect_uri` に合意結果を加えてリダイレクトさせる。
* そうでなければ、変更要求合意 UI エンドポイントにリダイレクトさせる。


### 9.2. リクエスト例

```http
POST /request/agree HTTP/1.1
Host: ta.example.org
Cookie: Permission-Manager=QkjGKwnIIC_m8uG-bSDOkGrcvN6kH1
Content-Type: application/x-www-form-urlencoded

applied=%5B%2248v_2-Jj4C%22%5D&denied=%5B%222Jof_YcIzH%22%5D
```


### 9.3. レスポンス例

変更要求合意 UI エンドポイントへのリダイレクト例。

```http
HTTP/1.1 302 Found
Location: /ui/request/agree.html
```


## 10. 変更要求合意 UI エンドポイント

以下のパラメータを受け付ける。

|パラメータ名|必要性|値|
|:--|:--|:--|
|**`display`**|任意|[OpenID Connect Core 1.0 Section 3.1.2.1] の `display` と同じもの|
|**`locales`**|任意|[OpenID Connect Core 1.0 Section 3.1.2.1] の `ui_locales` と同じもの|

UI の目的は、変更要求合意エンドポイントに POST させること。


## 11. 変更要求エンドポイント

変更要求合意チケットを発行する。
変更要求合意チケットをセッションに紐付ける。
チケットを JSON で返す。


### 11.1. リクエスト例

```http
GET /api/ticket HTTP/1.1
Host: ta.example.org
Cookie: Permission-Manager=QkjGKwnIIC_m8uG-bSDOkGrcvN6kH1
```


### 11.2. レスポンス例

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
    "ticket": "zuMCxS-ySm"
}
```


## 12. 変更要求対象数エンドポイント

* 変更要求合意チケットと紐付くセッションでない場合、
    * エラーを返す。
* そうでなければ、リクエストから以下のパラメータを取り出す。

|パラメータ名|必要性|値|
|:--|:--|:--|
|**`ticket`**|必須|変更要求合意チケット|

* 変更要求合意チケットがセッションに紐付くものと異なる場合、
    * エラーを返す。
* そうでなければ、リソースの保持アカウントの ID から保持リソースに対する変更要求の対象数へのマップを JSON で返す。


### 12.1. リクエスト例

```http
GET /api/target/request/count?ticket=zuMCxS-ySm HTTP/1.1
Host: ta.example.org
Cookie: Permission-Manager=QkjGKwnIIC_m8uG-bSDOkGrcvN6kH1
```


### 12.2. レスポンス例

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
    "22389660E8345308": 1
}
```


## 13. 変更要求対象エンドポイント

* 変更要求合意チケットと紐付くセッションでない場合、
    * エラーを返す。
* そうでなければ、リクエストから以下のパラメータを取り出す。

|パラメータ名|必要性|値|
|:--|:--|:--|
|**`ticket`**|必須|変更要求合意チケット|
|**`holder`**|必須|保持アカウントの ID|
|**`target`**|任意|空白区切りの取得する対象の番号。無指定なら 0 から対象数 -1 まで|

* 変更要求合意チケットがセッションに紐付くものと異なる場合、
    * エラーを返す。
* そうでなければ、対象を `target` の順番通りに JSON 配列で返す。
  対象の番号が正当でない場合は null が入る。

対象は変更対象エンドポイントで返される JSON オブジェクトに以下の変更が加えられたもの。

* 合意の種類に、先送りを意味する `postpone` を加える。
* `requester` に RFC3339 形式で要求日時を示す `date` 要素を追加する。
* `exist` は除く。


## 13. アカウント情報エンドポイント

* 変更合意チケット、または、変更要求合意チケットと紐付くセッションでない場合、
    * エラーを返す。
* そうでなければ、リクエストから以下のパラメータを取り出す。

|パラメータ名|必要性|値|
|:--|:--|:--|
|**`ticket`**|必須|チケット|
|**`users`**|必須|空白区切りのアカウント ID の配列|

* チケットがセッションに紐付くものと異なる場合、
    * エラーを返す。
* そうでなければ、アカウント情報を `users` の順番通りに JSON 配列で返す。

アカウント情報は以下の最上位要素を含む JSON オブジェクト。

* **`preferred_username`**
    * 保存されている場合のみ。
      名前。

アカウント情報は外部データからだけでなく、セッションに紐付く変更コードの [edo-auth] TA 由来のアカウント情報も使う。


### 13.1. リクエスト例

```http
GET /api/info/user?ticket=th2rEioRZL&users=38BF35F5464C00F9%2083AB154986FB1EAE HTTP/1.1
Host: ta.example.org
Cookie: Permission-Manager=QkjGKwnIIC_m8uG-bSDOkGrcvN6kH1
```


### 13.2. レスポンス例

```http
HTTP/1.1 200 OK
Content-Type: application/json

[
    {
        "preferred_username": "俺々"
    },
    {}
]
```


## 14. TA 情報エンドポイント

リクエストから以下のパラメータを取り出す。

|パラメータ名|必要性|値|
|:--|:--|:--|
|**`tas`**|必須|TA の ID の JSON 配列|

* 正当な TA でない場合、
    * エラーを返す。
* そうでなければ、アカウント情報の JSON 配列を返す。

アカウント情報は以下の最上位要素を含む JSON オブジェクト。

* **`friendly_name`**
    * 名前。


### 14.1. リクエスト例

```http
GET /api/info/ta?tas=%5B%22https%3A%2F%2Ffrom.example.org%22%5D HTTP/1.1
Host: ta.example.org
```


### 14.2. レスポンス例

```http
HTTP/1.1 200 OK
Content-Type: application/json

[
    {
        "friendly_name#ja": "何かの TA"
    }
]
```


## 15. エラーレスポンス

[PDS 権限変更プロトコル]を参照のこと。


## 16. 外部データ

以下に分ける。

* 共有データ
    * 他のプログラムと共有する可能性のあるもの。
* 非共有データ
    * 共有するとしてもこのプログラムの別プロセスのみのもの。


### 16.1. 共有データ


#### 16.1.1. 変更権限保持アカウント

権限管理 UI と共有する。

以下を含む。

* 変更権限保持アカウントの ID
* 変更権限保持アカウントによって権限変更されうるリソース保持アカウントの ID

以下の操作が必要。

* 変更権限保持アカウントの ID による取得
* リソース保持アカウントの ID による取得


#### 16.1.2. アクセス権限

以下を含む。

* リソース識別子
    * 保持アカウントの ID
    * 割り当て TA の ID
    * パス
* アカウント・TA ごとの権限

以下の操作が必要。

* 保存
* アカウント・TA ごとの権限の上書き


#### 16.1.3. 変更要求

もっとましな変更要求合意 UI と共有するかもしれない。

以下を含む。

* リソース識別子
    * 保持アカウントの ID
    * 割り当て TA の ID
    * パス
* アクセス元
    * アカウント集
        * ID
        * TA 集
            * ID
    * 要求日時
* `mod`
* 要請元
    * アカウント ID
    * TA の ID

以下の操作が必要。

* 保存
* 保持アカウントの ID による取得
* 保持アカウントの ID による数の取得
* 削除


#### 16.1.4. アカウント情報

以下を含む。

* ID
* 表示名

以下の操作が必要。

* ID による取得


#### 16.1.5. TA 情報

以下を含む。

* ID
* 表示名

以下の操作が必要。

* ID による取得


### 16.2. 非共有データ


#### 16.2.2. 変更コード

以下を含む。

* ID
* 有効 / 無効
* 有効期限
* 処理の主体の ID
* 要請元 TA の ID
* 変更内容
    * `essential`
    * `exist`
    * 他は変更要求と同様。
* `state`
* `redirect_uri`
* [edo-auth] TA から与えられたアカウント情報

以下の操作が必要。

* 保存
* ID による取得
* 無効化
    * 有効でなければ失敗する。


#### 16.2.3. セッション

以下を含む。

* ID \*
* 有効期限 \*
* 変更コード ID
* チケット
* 合意結果

\* は設定を引き継がない。

以下の操作が必要。

* 保存
* ID による取得


<!-- 参照 -->
[OpenID Connect Core 1.0 Section 3.1.2.1]: http://openid.net/specs/openid-connect-core-1_0.html#AuthRequest
[PDS データアクセス API]: https://github.com/realglobe-Inc/edo/blob/master/pds_data.md
[PDS 権限変更プロトコル]: https://github.com/realglobe-Inc/edo/blob/master/pds_change_permission.md
[edo-auth]: https://github.com/realglobe-Inc/edo-auth
