# can-place-flowers

## 問題の性質

配列操作、貪欲法（Greedy）、線形走査。与えられた花壇の左右隣接制約を満たしながら新しく
n 本の花を植えられるかを判定する決定問題で、入力サイズは最大 2×10^4
程度。時間計算量 O(m)、追加メモリ O(1) が望ましい。

## 問題要約

花壇を表す整数配列 flowerbed（0 = 空き、1 =
既に花あり）が与えられる。隣接する区画に花を植えることはできないというルールの下で、新たに
n 本の花を植えられるかを判定して true / false を返す。

例:

- flowerbed = [1,0,0,0,1], n = 1 → true
- flowerbed = [1,0,0,0,1], n = 2 → false

入力は隣接して 1 が並ばないことが保証されている。

## 制約

- 1 <= flowerbed.length <= 2 * 10^4
- flowerbed[i] は 0 または 1
- 初期状態では隣り合う 1 は存在しない
- 0 <= n <= flowerbed.length

時間・空間複雑度目安:

- 時間: O(m)（m = flowerbed.length）
- 追加空間: O(1)

## 考え方

考え方（貪欲）：左から右へ線形に走査し、ある位置 i が 0
かつその左右が空（配列外を空とみなす）であれば、そこに花を植える（1
にする）ことで局所的最適を取る。これにより植えられる本数を最大化できるので、n
が満たせるかを判定できる。

実装上の注意点:

- 配列の端（先頭・末尾）は隣接検査で配列外を空（0）とみなす。
- 検査条件: flowerbed[i] == 0 かつ (i == 0 || flowerbed[i-1] == 0) かつ (i ==
  m-1 || flowerbed[i+1] == 0)
- 位置に植えたら flowerbed[i] = 1 に変更し、カウントを増やす。もしカウントが n
  に達したら早期に true を返す。
- 既に n == 0 の場合は true を返す。

この貪欲アルゴリズムは、局所的に有効な選択が全体最適につながるため正しい。

## キーワード

- 貪欲法
- 線形走査
- インプレース更新
- 境界条件（配列端）
- O(n) 時間, O(1) 追加空間

---

以下に Go 言語での実装とテストを示します。

main.go

```go
package main

import "fmt"

// canPlaceFlowers returns true if n new flowers can be planted
// in flowerbed without violating the no-adjacent-flowers rule.
func canPlaceFlowers(flowerbed []int, n int) bool {
    if n == 0 {
        return true
    }
    m := len(flowerbed)
    for i := 0; i < m; i++ {
        if flowerbed[i] == 0 {
            emptyLeft := (i == 0) || (flowerbed[i-1] == 0)
            emptyRight := (i == m-1) || (flowerbed[i+1] == 0)
            if emptyLeft && emptyRight {
                flowerbed[i] = 1
                n--
                if n == 0 {
                    return true
                }
            }
        }
    }
    return false
}

func main() {
    // 簡単なデモ
    fmt.Println(canPlaceFlowers([]int{1,0,0,0,1}, 1)) // true
    fmt.Println(canPlaceFlowers([]int{1,0,0,0,1}, 2)) // false
}
```

main_test.go

```go
package main

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
    tests := []struct{
        flowerbed []int
        n int
        want bool
    }{
        {[]int{1,0,0,0,1}, 1, true},
        {[]int{1,0,0,0,1}, 2, false},
        {[]int{0}, 1, true},
        {[]int{0,0}, 1, true},
        {[]int{0,0,0}, 2, true},
        {[]int{1,0,0,0,0,1}, 2, false},
        {[]int{0,0,1,0,0}, 2, true},
        {[]int{1,0,1,0,1,0,1}, 1, false},
        {[]int{0,0,0,0,0}, 3, true}, // 植えられる最大は3 (positions 0,2,4)
        {[]int{0,0,0,0,0}, 4, false},
        {[]int{1,0,0,0,0}, 2, true}, // positions 2 and 4 (but 4 is edge)
        {[]int{1,0,0,0,0}, 3, false},
    }

    for i, tt := range tests {
        // テスト関数はスライスを破壊するのでコピーして渡す
        bedCopy := make([]int, len(tt.flowerbed))
        copy(bedCopy, tt.flowerbed)
        got := canPlaceFlowers(bedCopy, tt.n)
        if got != tt.want {
            t.Fatalf("case %d: flowerbed=%v n=%d want=%v got=%v", i, tt.flowerbed, tt.n, tt.want, got)
        }
    }
}
```

補足:

- テストでは canPlaceFlowers
  が入力配列をインプレースで変更するため、元配列を変更しないようコピーして渡しています。実運用で元配列を保持したい場合も同様の対策を取ってください。
