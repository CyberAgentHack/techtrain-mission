# データベース設計
デフォルトではMySQL 8.0を使用しています。
ローカルから直接通信する場合は, `$ make docker-up-for-test`を使用してください。
その場合, `$ mysql -h 127.0.0.1 -u ${DB_USER} -p {DB_NAME}`でmysql-clientからアクセスできます。

スキーマは, `mysql/schemas`に置かれています。。

生成されるテーブルは下記の5種類です。
1. users
2. rarities
3. characters
4. usercharacters
5. gacharecords

## 1.users
ユーザの情報を保存するトランザクションテーブルです。

|カラム名|型|制約|
|:-:|:-:|:-:|
|id|int|AUTO_INCREMENT, PRIMARY KEY|
|name|varchar(64)|NOT NULL|
|token|varchar(64)|NOT NULL|

 - `id`および`token`は, `/user/create`が呼ばれた際に割り振られます

## 2.rarities
キャラクタのレアリティを保持するマスタテーブルです。

|カラム名|型|制約|
|:-:|:-:|:-:|
|id|int|AUTO_INCREMENT, PRIMARY KEY|
|value|int|NOT NULL|
|weight|int|NOT NULL|

 - `value`はレア度を保持します
   - 将来的に型が変更される可能性があります
 - `weight`は, レア度の割合を保持するので整数型です
   - 浮動小数点型にすると管理に困る可能性があります

## 3.characters
キャラクタの情報を保持するマスタテーブルです。

|カラム名|型|制約|
|:-:|:-:|:-:|
|id|int|AUTO_INCREMENT, PRIMARY KEY|
|name|varchar(64)|NOT NULL|
|rarity_id|int|NOT NULL, foreign key rarities(id)|

 - `rarity_id`のUPDATEはあり得るのでCASCADEにしました
   - DELETEは許されないのでRESTRICTにしました

## 4.usercharcters
ユーザが保持するキャラクタの情報を保持するトランザクションテーブルです。

|カラム名|型|制約|
|:-:|:-:|:-:|
|id|int|PRIMARY KEY|
|user_id|int|NOT NULL, foreign key users(id)|
|character_id|int|NOT NULL, foreign key characters(id)|
|posessions|int|NOT NULL|

 - 外部キーのUPDATEは許されるのでCASCADEにしました
   - DELETEは許されないのでRESTRICTにしました

## 5.gacharecords
ユーザが引いたガチャの情報を保持するマスタテーブルです。

|カラム名|型|制約|
|:-:|:-:|:-:|
|id|int|PRIMARY KEY|
|user_id|int|NOT NULL, foreign key users(id)|
|character_id|int|NOT NULL, foreign key characters(id)|
|ts|timestamp|NOT NULL|