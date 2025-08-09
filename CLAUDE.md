# CLAUDE.md

このファイルは、このリポジトリでコードを扱う際のClaude Code (claude.ai/code)への指示を提供します。

## リポジトリ概要
これはアルゴリズムとデータ構造を学習するためのLeetCode練習用リポジトリです。各問題は`/problem/`配下の独自ディレクトリに独立して配置され、言語別の実装（GoとTypeScript/Deno）が含まれています。

## プロジェクト構造
```
problem/
├── {problem-name}/
│   ├── go/
│   │   ├── README.md     # 問題の解説（日本語）
│   │   ├── go.mod        # 独立したGoモジュール
│   │   ├── main.go       # ソリューション実装
│   │   └── main_test.go  # ユニットテスト
│   └── ts/
│       ├── README.md     # 問題の解説（日本語）
│       ├── deno.json     # Deno設定ファイル
│       ├── main.ts       # ソリューション実装
│       └── main_test.ts  # ユニットテスト
```

## 共通コマンド

### Go実装の場合

#### ソリューションのテスト
問題のGoディレクトリに移動してから実行:
```bash
cd problem/{problem-name}/go/

# 詳細出力付きでテスト実行
go test -v

# カバレッジ付きでテスト実行
go test -cover

# 特定のテストを実行
go test -v -run TestFunctionName
```

#### ソリューションの実行
```bash
cd problem/{problem-name}/go/
go run main.go
```

#### 新しい問題の追加（Go）
1. ディレクトリ構造を作成: `problem/{problem-name}/go/`
2. Goモジュールを初期化: `go mod init github.com/tsucchinoko/{abbreviated-name}`
3. `main.go`にソリューション、`main_test.go`にテストを作成
4. 必要に応じて`README.md`でアプローチを説明

### Deno/TypeScript実装の場合

#### ソリューションのテスト
問題のTypeScriptディレクトリに移動してから実行:
```bash
cd problem/{problem-name}/ts/

# テスト実行
deno test

# 詳細出力付きでテスト実行
deno test --allow-all

# 特定のテストファイルを実行
deno test main_test.ts
```

#### ソリューションの実行
```bash
cd problem/{problem-name}/ts/
deno run main.ts
```

#### 新しい問題の追加（Deno）
1. ディレクトリ構造を作成: `problem/{problem-name}/ts/`
2. Denoプロジェクトを初期化: `deno init`
3. `main.ts`にソリューション、`main_test.ts`にテストを作成
4. 必要に応じて`README.md`でアプローチを説明

## 命名規則

### Goモジュール
各問題は独立したGoモジュールを使用し、以下の命名パターンに従います:
- モジュール: `github.com/tsucchinoko/{abbreviated-problem-name}`
- 例: `twosum`, `best-buy`, `create-hello-world-fn`

### TypeScript/Denoファイル
- メインファイル: `main.ts`
- テストファイル: `main_test.ts`
- 関数名はLeetCodeの問題名に準拠

## コード規約
- 関数名はLeetCodeの問題名に合わせる（例: Go: `TwoSum`, `MaxProfit` / TypeScript: `twoSum`, `maxProfit`）
- アルゴリズムのロジックを説明する日本語コメントを含める
- 適切な場合はテーブル駆動テストを使用
- 各ソリューションには包括的なテストカバレッジを含める
- コメントで時間計算量と空間計算量を文書化

## 主要なアーキテクチャの決定事項
- **独立モジュール**: 各問題は独立したモジュール（共有依存関係なし）
- **言語分離**: 問題は`/problem/{name}/{language}/`配下で言語別に整理
- **自己完結型**: 各問題には実装、テスト、ドキュメントが含まれる
- **教育重視**: アプローチと計算量を説明する詳細な日本語ドキュメント