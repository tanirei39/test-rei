## 今日のゴール

- vscode の使い方マスター
- github の使い方マスター

## vscode

- command shift p：検索
- command ファイルをダブルクリック：画面分割表示

## github
### basic command
1. add
2. commit
3. push

- パスフレーズ不要にする：httpsg://fumidzuki.com/knowledge/3401/

```
git add .

git commit  -m "コミットメッセージ"
s
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
git stash -u

git stash pop
```

- chaery-pick: 