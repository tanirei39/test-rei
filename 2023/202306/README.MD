# 今月のゴール：コードを綺麗にする&編集機能をつける
- 0604
    - セットアップ
    - githubの使い方
- 0610
    - MySQL使って投稿機能の実装
- 0617
    - 一覧表示/削除機能の実装

# 0604
## settup
- install brew:`/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`
- install go:`brew install go`


## github
### install git
```
brew install git
```
### Clone repository
```
git clone git@github.com:tanirei39/test-rei.git
```
### basic command
1. add
2. commit
3. push

- パスフレーズ不要にする：https://fumidzuki.com/knowledge/3401/

```
git add .
git commit  -m "コミットメッセージ"
git push origin head
```

- branch:
```
//新しいブランチ
$ git branch <branchname>
//ブランチの切り替え
$ git checkout <branchname>
```

### optional command
- stash: 一時的にファイルを避難させるためのコマンド
```
git stash -u //待避
git stash pop//取り出
```
- chaery-pick: 

## how to
### go
- tutorial: https://go.dev/tour/welcome/1
- http package: https://pkg.go.dev/net/http

### Docker
- image build: https://docs.docker.com/language/golang/build-images/

### Cloud Run
- https://qiita.com/massie_g/items/5a9ce514eaa7c460b5e3
- https://cloud.google.com/run/docs/quickstarts/build-and-deploy/deploy-go-service?hl=ja


### vscode
- command shift p：検索
- command ファイルをダブルクリック：画面分割表示


# 0610
## Mysql（今回使う方）
- オープンソースのデータベースで無料で使用できる（商用の場合、有料。）
- インストール：https://chiritsumo-blog.com/mac-mysql-install/
    - MySQLログイン：`mysql -uroot -ppassword`
- データベースの作成：https://prog-8.com/docs/mysql-database-setup
- GoでMySQLに接続(登録/削除etc...)：https://www.chuken-engineer.com/entry/2021/09/24/162120
- 入力フォーム：https://leben.mobi/go/form-database/practice/web/

- 後で読んで見る：https://leben.mobi/go/mysql-connect/practice/

## Firestore
- クラウドサービス上にあるデータベース

# 0617
## Mysql
### 一覧表示
- リダイレクト処理：https://uokada.hatenablog.jp/entry/2015/04/11/205649
- MySQLのリスト表示：https://leben.mobi/go/template-loop/practice/web/

### 削除機能
- Path Parameterのパース：https://rennnosukesann.hatenablog.com/entry/2020/03/29/000000
- HTMLのボタンイベントをGo側で受け取る方法：http://pineplanter.moo.jp/non-it-salaryman/2021/02/21/go-websocket/

