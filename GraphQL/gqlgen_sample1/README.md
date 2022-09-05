# gqlgen

init
```
$ go run github.com/99designs/gqlgen init
```

schema変更後コード再生成
```
$ go run github.com/99designs/gqlgen generate
```

生成ファイル
- ./graph/generated/generated.go *generateで再生成
- ./graph/model/models_gen.go    *generateで再生成
- ./graph/schema.resolvers.go    *generateで再生成
- ./graph/resolver.go
- ./graph/schema.graphqls
- ./gqlgen.yml


起動
```sh
$ go run server.go
```


### 参考

- Getting Started Building GraphQL servers in golang
https://gqlgen.com/getting-started/

