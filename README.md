ブラウザでファイルアップロードを行う際、HTTPリクエストに付くContent-Typeヘッダーの値がどうなるのか気になったので、実際にやってみるためのリポジトリです。
作った後に、Windows Registryの`HKEY_CLASSES_ROOT`にファイル拡張子毎に、Content Typeが定義されていて、どうやらブラウザもこれを参照しているらしいと分かり、わざわざ動かしてみるまでもなくなりましたが、一応置き残しておきます。

## Usage

```
go run main.go
go run main.go --port 8080
```
