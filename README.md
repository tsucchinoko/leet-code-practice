# leet-code-practice
LeetCodeの勉強用リポジトリ


## Goで実装する場合
```bash
# ディレクトリ作成と移動
mkdir -p problem/<problem name>/go
cd problem/<problem name>/go

# Goモジュールの初期化
go mod init github.com/tsucchinoko/<problem name>

# ファイル作成
touch main.go main_test.go

# テスト実行
go test -v
```

## Denoで実装する場合
```bash
# ディレクトリ作成と移動
mkdir -p problem/<problem name>/ts
cd problem/<problem name>/ts

# Denoプロジェクトの初期化
deno init

# テスト実行
deno test
```
