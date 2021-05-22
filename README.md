# user-go-server

# ディレクトリ構成
```
├── domains
|   ├──models
|   |   └── user.go				
|   └──repositories
|       └── user_repository.go
├── usecases
|   ├── user_create_usecase.go
|   ├── user_read_usecase.go
|   ├── user_read_all_usecase.go
|   ├── user_update_usecase.go
|   └── user_delete_secase.go
├── interfaces
|   ├── controllers
|   |   ├── user_post_controller.go
|   |   ├── user_get_controller.go
|   |   ├── user_get_all_controller.go
|   |   ├── user_put_controller.go
|   |   ├── user_delete_controller.go
|   |   └── user_translator.go
|   └── gateways
|       └── user_gateway.go
├── infrastructures
|   ├── proto
|   |   ├── server.go
|   |   └── router.go
|   └── database
|       └── postgres.go
├── injector
|   └── injector.go
└── main.go
```
※ わかりやすいようにCRUD操作で命名
※ サーバーサイドの実装なのでPresenterは考えない

# grpcの型を取得
```bash
GOPRIVATE=github.com/* go get github.com/takokun778/user-grpc
```

# Rule

- 変数やメソッド名は基本的にキャメル・パスカルケース
- ファイル名についてはスネークケース

# 参考
アーキテクチャ
- [Go言語でClean Architectureを実現して、gomockでテストしてみた](https://qiita.com/ogady/items/34aae1b2af3080e0fec4)
- [GoにおけるDI](http://inukirom.hatenablog.com/entry/di-in-go)
- [Goはクリーンアーキテクチャの思想を活かせるか？DMMのゲームプラットフォームにGo言語を選んだ理由](https://logmi.jp/tech/articles/323451)
- [Go × Clean Architectureのサンプル実装](http://nakawatch.hatenablog.com/entry/2018/07/11/181453)
- [Go言語を学習し始めて、簡単なCRUDアプリをクリーンアーキテクチャで作成するまで](https://developers.wano.co.jp/1906/)
- [ミラティブのサーバサイドをGo + Clean Architectureに再設計した話](https://tech.mirrativ.stream/entry/2020/11/30/142354)
- [Goとクリーンアーキテクチャとトランザクションと](https://qiita.com/miya-masa/items/316256924a1f0d7374bb)

お作法
- [golangでの自分のためのお作法まとめ](https://qiita.com/tjtjtjtj/items/b3e0508eaf37571839ef)
- [「Effective Go」のチートシート　〜脱Go初心者〜](https://qiita.com/sensuikan1973/items/8a47af7761d002d12260)

テスト
- [golangのテストはじめ](https://qiita.com/tmzkysk/items/8bb37795ac223664d682)
- [Go言語でユニットテスト, テストしやすいコードとモックを書く](https://qiita.com/hiroyky/items/4a9be463e752d5c0c41c)

ほか
- [Go言語開発を便利にするMakefileの書き方](https://qiita.com/yoskeoka/items/317a3afab370155b3ae8)
- [少人数での爆速開発を目指してgolang×GCPの技術選定をした話](https://zenn.dev/sh_komine/articles/35527f84a2be3a)
- [Goでよく見るnewとNewを調べたときのメモ](https://qiita.com/gold-kou/items/4494f8b69b8fa53d5e93)
- [GitHubのプライベートリポジトリをgo getする](https://www.yuyagishita.com/tech/golang/go-get-github-private-repository/)
- [Go のモジュール管理【バージョン 1.16 改訂版】](https://zenn.dev/spiegel/articles/20210223-go-module-aware-mode)