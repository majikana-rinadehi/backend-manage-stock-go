# templatesディレクトリについて

- golangのテストコードを自動生成するためのテンプレートファイルです。
- [cwell/gotests](https://github.com/cweill/gotests)の、タグ`v1.6.0`にチェックアウトし、`gotests/internal/render/template`のファイルをここにコピー・配置してあります。
- 下記コマンドで、このtemplateファイルをもとにテストコードが自動生成されます。

```go
$ gotests -template_dir util/test/templates -w -all path/to/the/targetfile.go
```