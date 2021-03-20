# Golang学習用プロジェクト

## 目的

* Golangアプリ開発の全体像把握

## 開発環境

* Windows 10
* VSC 1.52.1
* Go 1.15.6
* Git 2.27.0.windows.1

## 構築手順

### Go

1. [ここ](https://golang.org/dl/)のインストーラをダウンロード
2. インストーラ実行（全部デフォルトで良い）
3. コマンドプロンプト起動
4. 以下のコマンドでGoのversionが表示すれば良い

    ```sh
    > go version
    ```

5. インストーラが`%GOPATH`にGoのパスを設定する

### パッケージ取得

1. godocをダウンロードする

    ```sh
    > go get golang.org/x/tools/cmd/godoc
    ```

2. godocの使い方例

    ```sh
    > go doc fmt
    ```

### VCS

1. `Go`で検索して、`golang.go`のプラグインをインストール
2. コマンドパレット(Ctrl+Shift+P)を開く
3. `GO: Install/Update tools`で検索
4. 全ツールを選択し、OKボタンを押下してインストール
5. delveをインストール
   * VCSの「実行とデバッグ」ボタンを押すとインストールされる

## ドキュメント

### VCSのプラグイン

* <https://code.visualstudio.com/docs/languages/go>

### パッケージ

* <https://golang.org/pkg/>

## コマンド

### 実行

```sh
> go run ${Goファイル}
```

### コード整形

```sh
> gofmt ${Goファイル} // 整形の提案表示
> gofmt -w ${Goファイル} // 整形の提案通り反映
```

### テスト

```sh
> go test -v ./...
```

-v : 実行結果表示
./... : カレントディレクトリ以下全てからテストを探して実行する

### godoc

```sh
> godoc -http=localhost:6060
```

<http://localhost:6060>でgodocを閲覧する

## Tips

### 依存関係

* <https://qiita.com/yokoto/items/13be66b6276e17d9f554>
* Go Modules使う

  ```sh
  > cd ${プロジェクト名}
  > go mod init ${プロジェクト名}
  ```

* go.modファイルができる

### 単体テスト

* Goはテスト対象のGoファイルと同列にテストファイルを作成する
* lib/math.goの単体テストは、lib/math_test.goとして作成する
* `testing`をimportしてテスト用ライブラリを利用する
* その他のテスティングフレームワークには、Ginkgo、Gomegaがある

### ローカルドキュメント

* 基本javadocのように記載する
  * lib/math.go参照
* godocのexamplesは、テストファイルで`Example関数名`で関数作るとそれが例になる
  * lib/math_test.goのExampleAverageとExamplePerson_Say参照

### 設定ファイル

* TODO

### SQLite

* https://www.sqlite.org/download.html
  * sqlite-tools-win32-x86-3340100.zip
  * PATH通す
* https://jmeubank.github.io/tdm-gcc/
  * tdm-gcc-webdl.exe
* go get github.com/mattn/go-sqlite3


## アプリ

### 仕様

* 預金残高の管理
* 残高推移を表すグラフを表示

### ライブラリ

* gotron
  * node.js(npm)の環境が必要
    * https://github.com/Equanox/gotron
    * bin/gotron-builder
      * プロジェクトの直下でexe起動しないとだめ
* jQuery
  * http://tk2-207-13211.vs.sakura.ne.jp/2016/02/749/
  * https://noitalog.tokyo/electron-jquery/#toc2
* i18n
  * gotron自体に機能なし。
  * 自分で辞書作って`window.send`でオブジェクト渡せばできるけどなんか違う
  * https://www.i18next.com/overview/getting-started
* koazee
  * https://github.com/wesovilabs/koazee
* guregu/null
  * https://qiita.com/isao_e_dev/items/55a9ad66f00ef031ac6c
### 画面構成

//TODO

### ストレージ


* https://qvault.io/2019/10/21/golang-constant-maps-slices/
