PHOTO-AWARDで使用するサーバーはここに集約している。

frontendからも使用するため、[パスやレスポンスの定義](../swagger.yaml)は[親フォルダ](..)に配置している。

## 使い方

### go-swaggerのインストール

```
go install github.com/go-swagger/go-swagger/cmd/swagger@latest
```

### swagger定義変更時にやること

generate
```
make generate
```

をすることで定義に従いコードが生成される。

これらを用いて `./handler`　に各種処理を実装する。

なお、このときに `./gen/restapi/configure_photo_award.go` は更新されない。

`./handler`に作成した各種エンドポイント用の関数をこのファイルでサーバーに結び付ける。