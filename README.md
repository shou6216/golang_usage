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

## Tips

### 依存関係

* <https://qiita.com/yokoto/items/13be66b6276e17d9f554>
* Go Modules使う

  ```sh
  > cd ${プロジェクト名}
  > go mod init ${プロジェクト名}
  ```

* go.modファイルができる

## アプリ

### 仕様

//TODO

### 画面構成

//TODO

### ストレージ

//TODO
