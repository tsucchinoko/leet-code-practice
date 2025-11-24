# dota2-senate

## 問題の性質

- ゲーム理論／シミュレーション問題（ターン制）
- 文字列処理、キューやカウント管理によるシミュレーションで解ける
- 最適戦略を仮定した競合シミュレーション（双方とも最適に行動する）

## 問題要約

与えられた文字列 senate（各文字は 'R' または 'D'）は順に並

んだ上院議員を示す。ラウンドごとに生きている（権利がある）議員は順に行動する。各議員は以下のいずれかを行う：

- 相手側の任意の議員を「禁止（ban）」して以降の全ラウンドで行動させないようにする
- 現在行動可能な議員が全員同じ党なら勝利を宣言する

双方は最善手を選ぶと仮定したとき、最終的に勝利する党（"Radiant" または
"Dire"）を返す。

## 制約

- 1 <= n <= 10^4（senate の長さ）
- senate[i] は 'R' または 'D'
- 時間制約に対して O(n)～O(n log n) 程度の解法が望ましい
- メモリは入力サイズに対して線形まで許容される

## 考え方

直感：

- 各議員は相手の議員を1人禁止できるため、早く行動する側が有利。
- 重要なのは「次に行動できる自党の議員の位置（順序）」で、禁止対象は次に行動する相手側の議員（最も近い順に）を狙うのが最適。

効率的なアルゴリズム（キューを使ったシミュレーション）：

1. 'R' のインデックスをキュー Rq に、'D' のインデックスをキュー Dq に格納する（0
   から n-1）。
2. 先に行動するのはインデックスが小さい方の先頭。先頭のインデックスが小さいキューの議員が行動し、相手の先頭議員を「ban」する（相手キューから先頭を
   pop する）。
3. 行動した議員は次のラウンドでも末尾に戻ってくるが、サイクルを表現するためにインデックスに
   n を加えてキューの末尾に push する（これで次回の順序を表す）。
4. 片方のキューが空になるまで続け、残った側が勝者。
5. 最終出力は "Radiant"（Rq が残る）または "Dire"（Dq が残

る）。

正当性の要点：

- 各行動は常に相手キューの最前（最も早く行動する相手）を封じることが最も合理的。
- インデックスに n を足して戻すことでラウンドが進み、順序が維持される。

計算量：

- 各議員は ban されるか ban
  するまでに少なくとも一度キューから処理されるため、O(n)
  回のキュー操作。したがって O(n) 時間、O(n) 空間。

## キーワード

- シミュレーション
- キュー（FIFO）
- ゲーム理論（双方最適）
- 循環（ラウンド）
- O(n)

---

以下に Go 言語での具体的な実装を示します。ファイルは main.go と main_test.go
です。

main.go

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func predictPartyVictory(senate string) string {
    n := len(senate)
    rq := make([]int, 0, n)
    dq := make([]int, 0, n)

    for i, ch := range senate {
        if ch == 'R' {
            rq = append(rq, i)
        } else {
            dq = append(dq, i)
        }
    }

    // simulate using two queues implemented by slices with head indices
    ri, di := 0, 0
    for ri < len(rq) && di < len(dq) {
        rpos := rq[ri]
        dpos := dq[di]
        if rpos < dpos {
            // R acts and bans this D
            // push R to back with index + n
            rq = append(rq, rpos+n)
            ri++
            di++ // D at front is banned
        } else {
            // D acts and bans this R
            dq = append(dq, dpos+n)
            di++
            ri++
        }
    }

    if ri < len(rq) {
        return "Radiant"
    }
    return "Dire"
}

func main() {
    in := bufio.NewReader(os.Stdin)
    var s string
    if _, err := fmt.Fscan(in, &s); err != nil {
        return
    }
    fmt.Println(predictPartyVictory(s))
}
```

main_test.go

```go
package main

import "testing"

func TestPredictPartyVictory(t *testing.T) {
    tests := []struct {
        senate string
        want   string
    }{
        {"RD", "Radiant"},
        {"RDD", "Dire"},
        {"R", "Radiant"},
        {"D", "Dire"},
        {"RRDDD", "Dire"},
        {"RDRDR", "Radiant"},
        {"DRRDR", "Radiant"},
        {"DDRRR", "Radiant"},
        {"DR", "Dire"},
    }

    for _, tt := range tests {
        got := predictPartyVictory(tt.senate)
        if got != tt.want {
            t.Fatalf("senate=%q: got %q, want %q", tt.senate, got, tt.want)
        }
    }
}
```

実行方法（例）：

- go test でテストを実行
- 単一実行なら echo "RDD" | go run main.go として標準入力から文字列を渡

す
