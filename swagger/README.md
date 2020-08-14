## 概要
- `Swagger Spec`
-> REST APIに対して Swagger の仕様に準じたドキュメント
- `Swagger Editer`
-> Swagger Spec の設計書を記載するためのエディタ
- `Swagger UI`
-> Swagger Spec で記載された設計からドキュメントをHTML形式で自動生成するツール
- `Swagger Codegen`
-> コマンドラインツール
-> Swagger Spec で記載された設計からAPIのスタブを自動生成


### 1. Docker-compose.yamlでコンテナ起動
```
$ docker-compose up -d
```
### 2. Swagger EditorからYAMLファイルをダウンロード

### 3. npmライブラリでYAML->HTMLに変換して作成
```
$ npm install -g bootprint
$ npm install -g bootprint-openapi

$ bootprint openapi [出力したYAMLかJSONファイル] [出力先ディレクトリ]
```
