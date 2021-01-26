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

### Tips

#### 1つのプロジェクトでmainメソッドを持つGoファイルを複数扱う

* フォルダ分ける
* 同列に置くとエラー出る

## アプリ

### 仕様

* DQWの図鑑虹埋めとこころS取得状態を管理
  * 図鑑の銅、銀やこころのS以下の状態は管理しない

### 画面構成

* 一覧画面
* 追加・更新画面

### ストレージ

* SQLite

## コマンド

### プロジェクト作成

```sh
> flutter create アプリ名
```

### 実行可能デバイスの確認

```sh
> flutter devices
```

### 実行

```sh
> flutter run --device-id XXXXX(devicesで取得するID)
```

## Tips

### 画面遷移

* https://github.com/shou6216/flutter_usage/commit/da6e70e5145820a345cf5fbbc336a650d353724f

#### 定義

```dart
return MaterialApp(
    routes: <String, WidgetBuilder>{
        '/home': (BuildContext context) => new MyHomePage(),
        '/new': (BuildContext context) => new NewMonsterPage()
});
```

* `routes`で画面定義名(/home、/new)と画面オブジェクトを定義します

#### 遷移元画面に戻った後の処理

```dart
Navigator.of(context).pushNamed('/new').then((value) {
                  if (value) {
                    setState(() => {print('reloaded search results')});
                  }
                })
```

* (ボタン押下時など)↑のように遷移する。`pushNamed('/new')`の戻り値は`Future`になっていて、遷移先から戻ってきた後に処理をしたい場合は
ここで行う。`value`は遷移先からの戻り値。

#### 遷移元画面にパラメータを渡す

```dart
class _NewMonsterPageState extends State<NewMonsterPage> {

  @override
  Widget build(BuildContext context) {
    return WillPopScope(
        onWillPop: () {
          Navigator.of(context).pop(_shouldReload);
          return Future.value(true);
        },
        child: Scaffold()
    )
  }
```

* `WillPopScope`で遷移元画面に戻る前のイベントを拾える。
そこで、`Navigator.of(context).pop()`に返したいオブジェクトをセットする

#### 遷移先にパラメータを渡す

```dart
Navigator.of(context).pushNamed('/new',arguments: monster)
```

* `arguments`に渡したいオブジェクトをセットする

```dart
final Monster monster = ModalRoute.of(context).settings.arguments;
```

* `settings.arguments`で受け取る

### 非同期描画

* https://github.com/shou6216/flutter_usage/commit/ea76dc0b0dd21a724bdc167d74c00dfd472cab59

* `FutureBuilder`を使う

### スクロール

* https://github.com/shou6216/flutter_usage/commit/42b66d7f7847255ee2428932e874816e5b7e13ff

* `SingleChildScrollView`を使う

### 入力フォーム

* https://github.com/shou6216/flutter_usage/commit/3cd0ad229b89fb30cee6b5388af4374f32c4c9e3

* `TextFormField`を使う

### SQLite

* https://github.com/shou6216/flutter_usage/blob/master/pubspec.yaml#L26-L27
  * `sqflite`使う
  * `path`は、sqliteファイルを配置する場所を指定する際に利用する
    * AndroidとiOSでフォルダ構成が異なる部分を吸収してくれる
* https://github.com/shou6216/flutter_usage/blob/master/lib/database_helper.dart

### 国際化

* https://github.com/shou6216/flutter_usage/blob/master/pubspec.yaml#L28-L30
* https://github.com/shou6216/flutter_usage/commit/6a3ba02eda6f8db1e7f54f8e8ce794f684c476e3

## CodeMagic

* アカウント作成
  * GitHubアカウントで作れる
* リポジトリを登録
* `Start Build`するとビルド開始
* 正常終了するとGitHubアカウントのメールアドレスにapkファイルが添付されたメールが届く