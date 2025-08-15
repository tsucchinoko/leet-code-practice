# leet-code-practice

LeetCodeの勉強用リポジトリ

## セットアップ自動化（Taskfile）

このプロジェクトでは[Task](https://taskfile.dev)を使用して、問題のセットアップを自動化しています。

### インストール

```bash
# macOS
brew install go-task/tap/go-task

# その他のOS
# https://taskfile.dev/installation/
```

### 使用方法

```bash
# Go言語の問題を作成
task new-go name=two-sum

# TypeScript/Denoの問題を作成
task new-ts name=two-sum

# 両言語同時に作成
task new name=two-sum

# 作成済み問題の一覧表示
task list
```

上記のコマンドにより、以下が自動的に作成されます：
- ディレクトリ構造（`problem/問題名/go/` または `problem/問題名/ts/`）
- README.md（問題説明用テンプレート）
- Go: `go.mod`の初期化、`main.go`、`main_test.go`の空ファイル
- TypeScript: `deno init`の実行、`main.ts`、`main_test.ts`の作成

## Goで実装する場合（手動）

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

## Denoで実装する場合（手動）

```bash
# ディレクトリ作成と移動
mkdir -p problem/<problem name>/ts
cd problem/<problem name>/ts

# Denoプロジェクトの初期化
deno init

# テスト実行
deno test
```
