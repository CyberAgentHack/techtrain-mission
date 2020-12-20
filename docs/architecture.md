# アーキテクチャ
実装は, `gameapi`以下に置かれています。

## 概要
ユーザからのアクセスをAPIサーバが受け取り, データの永続化のためにDBサーバと通信を行います。その後, 結果をユーザに返却します。
```
ユーザ
|    ^
v    |
APIサーバ
|    ^
v    |
DBサーバ
```

## API仕様
[仕様はこちら(pdf)](./api-specification.pdf)です。

# アプリケーションアーキテクチャ
実装は[こちらの記事](https://nrslib.com/bottomup-ddd/)および[書籍](https://www.amazon.co.jp/dp/B00GRKD6XU/ref=cm_sw_r_tw_dp_x_7BZ3Fb0MMAAVN)を参考に, DDDにしてみました。DDDの選定理由は, 名前には聞いたことがあったのですが, 実際に実装したことがなく, API仕様書を見た時にある程度分割されていたため, 実装しやすいと感じたためです。

## パッケージ構成
```
./gameapi
├── domain     # ドメイン
│   ├── entity   # エンティティ
│   └── service  # ドメインサービス(基本的に利用しない)
├── handler    # HTTPリクエストを捌くためのハンドラ
├── infra      # 永続化と再構築の実装をするインフラ層
├── log        # ロガー
├── repository # 永続化と再構築のためのリポジトリ
└── usecase    # ハンドラの処理を細分化して処理するユースケース
```

# HTTPリクエストからDBアクセスをしてレスポンスを返すまでの流れ
```
main.go
|  ^
V  |
handler/*go
|  ^
V  |
usecase/*.go
|  ^
V  |
infra/
```