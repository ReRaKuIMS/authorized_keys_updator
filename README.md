
# 使い方

```
$ authorized_keys_updator -f 'path/to/authrorized_keys' -s 'https://example.com/your_public_keys_file_url'
```

と実行すると、

以下のルールに従って `-f` に指定された `authrorized_keys` の内容を書き変えます。

- `#### PUBKEY_MASTER EDIT START ####`と`#### PUBKEY_MASTER EDIT END ####`に囲まれた部分がある場合
  - 囲まれた部分の内容を `-s` で指定したURLのファイルの内容に書き換えます
- `#### PUBKEY_MASTER EDIT START ####`と`#### PUBKEY_MASTER EDIT END ####`に囲まれた部分がない場合
  - `-s` で指定したURLのファイルの内容を、`#### PUBKEY_MASTER EDIT START ####`と`#### PUBKEY_MASTER EDIT END ####`で囲んで、末尾に追記します

# Install

[github releases](https://github.com/ReRaKuIMS/authorized_keys_updator/releases)からバイナリをサーバーにぽいと置く。

cronなどで定期的に`authorized_keys`をudpateしてあげると吉でしょう。

# リリース方法

## 依存パッケージ
### gox

```
$ go get github.com/mitchellh/gox
```

### ghr

```
$ go get -u github.com/tcnksm/ghr
```

## リリースコマンド

```
$ export GITHUB_TOKEN="....."
$ ./release.sh
```
