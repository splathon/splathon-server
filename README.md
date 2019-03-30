# splathon-api
Splathonで使用するAPIをSwaggerで作る

## Swagger UIへのアクセス
https://splathon.github.io/splathon-api/dist/index.html

## ざっくり今後やる必要があることメモ(Issue管理するほどでもなかったので)
 - [ ] swagger-codegen.jarをgit管理下に追加する(https://github.com/swagger-api/swagger-codegen)
 - [ ] CI環境を構築する(CircleCI?)
 - [ ] swagger-code-gen.jarを使って各サーバー、クライアント用のAPIを作成する用のスクリプトを用意する
 - [ ] 各API作成用のconfigファイルを用意する
 - [ ] ↑で作ったスクリプトとconfigを使ってCI環境に連携させて、swagger.yaml更新されたら自動でAPIも更新するようにする
 - [ ] swagger.yamlの中身を書く

## ディレクトリ構成
```
<root>
   - docs
       - dist -> Swagger UIの表示用ディレクトリ
       - index.html -> Swagger UI表示用のindexファイル
       - swagger.yaml -> 本体
   - README.md -> これ
   - swagger.yaml(シンボリックリンク) -> 本体へのリンク
```

## 注意事項
**`docs/swagger.yaml`の本体は移動させないこと**
