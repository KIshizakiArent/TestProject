# Echo用のテストプロジェクト

## 概要

食べ物の保管ソフト

## テストのやりかた

アプリケーションを実行すると`http://localhost:8000`上で待ち受けます。
ベーシック認証を設定しているので`Authorization`を使用して`UserName`に`username`、`Password`に`password`を指定してください。
`http://localhost:8000/storage/food` に`name`と`type`を`json`形式で`POST`してあげると保管します。
`http://localhost:8000/storage/food` に`GET`してあげると保管されている食べ物リストが表示されます。

jsonサンプル
{
"name" : "potato",
"type" : "vegetable"
}

分からない場合は石崎まで。