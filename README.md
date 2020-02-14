# nullable_json

# 概要
- `gin`でnullableな`json`を返却するためのサンプルコードです。
- [こちらの記事の内容です](https://qiita.com/whisky-shusuky/items/a72627973beba36f90ca)

# 使用方法
- `go run main.go`によりginが起動します
- `http://0.0.0.0:8080/`にアクセスすると以下のようなjsonが返却されます

```
{"nullable_id_init":2,"nullable_id_not_init":null,"not_null_id_init":1,"not_null_id_not_init":0}
```

- `nullable_id_init`については現在時刻の分が偶数であればその分が入り、奇数ならばnilが返ります。
- `nullable_id_not_init`はポインタであり値を代入していないのでnullが返ります。
- `not_null_id_init`は値1を明示的に代入しているので1が返ります。
- `not_null_id_not_init`は値を代入しておらずint型であるため初期値として0が返ります。
