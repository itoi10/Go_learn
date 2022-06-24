# Go Echo API 

### 実行方法

```
$ go run *.go
```

### 使用方法

エンドポイントをCURLなどでリクエスト


<html>
  <div>
    <h3>ユーザー登録 POST /signup</h3>
    <table border="1">
      <tr>
      <th>Request</td>
      <td>curl -X POST localhost:8080/signup -H "Content-Type: application/json" -d '{"name": "user1", "password": "password"}'</td>
      </tr>
      <tr>
        <th>Resopnse</td>
        <td>{"id":1,"name":"user1","password":""}</td>
      </tr>
    </table>
  </div> 

  <div>
    <h3>ログイン POST /login</h3>
    <table border="1">
      <tr>
      <th>Request</td>
      <td>curl -X POST localhost:8080/login -H "Content-Type: application/json" -d '{"name": "user1", "password": "password"}'</td>
      </tr>
      <tr>
        <th>Resopnse</td>
        <td>{"token":"(JWTトークン)"}</td>
      </tr>
    </table>
  </div>

  <div>
    <h3>todo登録 POST /api/todos</h3>
    <table border="1">
      <tr>
      <th>Request</td>
      <td>curl -X POST localhost:8080/api/todos -H "Authorization: Bearer (JWTトークン)" -H "Content-Type: application/json" -d '{"name": "todo1"}'</td>
      </tr>
      <tr>
        <th>Resopnse</td>
        <td>{"uid":1,"id":1,"name":"todo1","completed":false}</td>
      </tr>
    </table>
  </div>

  <div>
    <h3>todo一覧取得 GET /api/todos</h3>
    <table border="1">
      <tr>
      <th>Request</td>
      <td>curl -X GET localhost:8080/api/todos -H "Authorization: Bearer (JWTトークン)"</td>
      </tr>
      <tr>
        <th>Resopnse</td>
        <td>[{"uid":1,"id":1,"name":"todo1","completed":false}]</td>
      </tr>
    </table>
  </div>

  <div>
    <h3>todo更新 PUT /api/todos/:id/completed </h3>
    <table border="1">
      <tr>
      <th>Request</td>
      <td>curl -X PUT localhost:8080/api/todos/1/completed -H "Authorization: Bearer (JWTトークン)"</td>
      </tr>
      <tr>
        <th>Resopnse</td>
        <td>(No Content)</td>
      </tr>
    </table>
  </div>

  <div>
    <h3>todo削除 DELETE /api/todos/:id </h3>
    <table border="1">
      <tr>
      <th>Request</td>
      <td>curl -X DELETE localhost:8080/api/todos/1 -H "Authorization: Bearer (JWTトークン)"</td>
      </tr>
      <tr>
        <th>Resopnse</td>
        <td>(No Content)</td>
      </tr>
    </table>
  </div>
</html>


### 参考記事

[Go言語でEchoを用いて認証付きWebアプリの作成]( https://qiita.com/x-color/items/24ff2491751f55e866cf#%E3%83%87%E3%83%BC%E3%82%BF%E3%83%99%E3%83%BC%E3%82%B9%E5%87%A6%E7%90%86 )