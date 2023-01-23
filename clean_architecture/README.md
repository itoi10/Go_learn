# Goクリーンアーキテクチャ

Go言語でクリーンアーキテクチャの練習


# 起動方法
1. firebaseで生成した秘密鍵に```serviceAccountKey.json```という名前をつけてmain.goと同じディレクトリに配置する。

2. 実行
```sh
$ go run .
```

# エンドポイント
```sh
$ curl -XPOST -H "Content-Type: application/json" -d '{"name" : "anya" , "age" : 6, "address": "Berlint"}' localhost:8080/users
[{"name":"anya","age":6,"address":"Berlint"}]
```

```sh
$ curl -XGET localhost:8080/users
[{"name":"anya","age":6,"address":"Berlint"}]
```

### 参考記事

[【Go言語】クリーンアーキテクチャで作るREST API](https://rightcode.co.jp/blog/information-technology/golang-clean-architecture-rest-api-syain)
