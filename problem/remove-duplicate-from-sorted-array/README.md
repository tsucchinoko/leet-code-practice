# remove-duplicate-from-sorted-array

## 制約

- 1 <= nums.length <= 3 * 104
- -100 <= nums[i] <= 100
- nums is sorted in non-decreasing order.

## アプローチ

この問題の考え方は、ソートされた配列の中で「重複を見つけて削除する」というより、**重複でない要素だけを前に集める**イメージで進めることです。

## 考え方のポイント

- 配列は最初からソート済みなので、**重複する要素は必ず隣り合っている**ことが保証されている。
- なので、配列の中を順番に見ていき、
  - 前回見たユニークな値と異なる値を見つけたら、次のユニークな場所にその値を置く。
- これにより、**ユニークな要素が配列の先頭から順に並ぶようになる**。

## なぜインプレースでできるか？

- 新しい配列を作らずに、元の配列の中で、
- 「ユニークな要素を詰める場所」を示すポインタと、「次に見る場所」を示すポインタの２つを使う。
- これで重複をまとめて削除したのと同じ状態にできる。

## 視覚的イメージ

例えば

nums = [1][1][1][2][2]

- 「ユニーク書き込み位置」を最初は0番に設定（nums = 0）
- 次の要素を順に見ていき、
- nums[1] = 0 は nums と同じなのでスキップ、
- nums[2] = 1 は違うので、書き込み位置を1に進めて、nums[1]に1を入れる、
- また次の違う値が見つかれば書き込み位置を進めて書き込む、…

結果、重複なしの数だけが配列の先頭に詰められる。

---

この考えを踏まえて、実際にポインタ（あるいはインデックス）を動かしながら重複を除去していきます。

情報源 [1] Remove Duplicates in-place from Sorted Array - takeUforward
https://takeuforward.org/data-structure/remove-duplicates-in-place-from-sorted-array/
[2] Remove Duplicates from Sorted Array LeetCode problem
https://stackoverflow.com/questions/79036220/remove-duplicates-from-sorted-array-leetcode-problem
