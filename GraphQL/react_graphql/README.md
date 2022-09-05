# GraphQL React サンプル

GraphQLサーバから情報を取得し一覧表示する

## 起動方法

client_react
```
npm start
```

server_express
```
node ./gqlServer.js
```

### コード

client_react

```js
// データ取得用関数
src/api/fetchGraphQL.js
// データ取得と表示
src/components/BooksTable.js 
src/components/BooksRow

```

server_express

```js
// GraphQLサーバ
gqlServer.js
```


### 参考記事

[3 時間で React 入門 ～ GraphQL と Relay で作る Web アプリケーション～｜研修コースに参加してみた](https://www.seplus.jp/dokushuzemi/blog/2021/06/quick_start_react_with_graphql.html)
