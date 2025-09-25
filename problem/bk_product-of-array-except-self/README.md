## 問題の性質

配列の各要素について、その要素を除いた配列全体の積を求める問題。除算を使わずに
O(n)
時間で解く必要がある。出力配列は結果を保持するための領域であり、追加の定数（O(1)）領域での解法が問われる（出力配列は余分な空間に含めない）。

## 問題要約

与えられた整数配列 nums に対して、各インデックス i について answer[i] を nums
のすべての要素の積から nums[i] を除いた値（すなわち nums の要素のうち nums[i]
を除いた積）にする配列を返す。除算は使用不可。時間計算量 O(n)
を満たすこと。出力の各要素は 32-bit 整数に収まると保証されている。

## 制約

- 2 <= nums.length <= 100,000
- -30 <= nums[i] <= 30
- 任意の prefix または suffix の積は 32-bit
  整数に収まる（つまりオーバーフローしない保証）
- 除算演算子は使用禁止
- 必須: O(n) の時間計算量
- 追加課題: 出力配列を除けば O(1) の追加空間で実装可能

## 考え方

基本アイデアは「左からの累積積」と「右からの累積積」を掛け合わせること。

1. 左側の積（prefix product）: left[i] を nums[0]..nums[i-1] の積とする（i
   の前までの積）。 left[0] = 1。
2. 右側の積（suffix product）: right[i] を nums[i+1]..nums[n-1] の積とする（i
   の後ろの積）。 right[n-1] = 1。
3. 各 i に対して answer[i] = left[i] * right[i]。

これを直接行うと追加 O(n) の配列 left, right
を使うが、右の累積を出力配列に統合することで追加空間を O(1)
にできる（出力配列自体は除外）。

具体的手順（出力配列を利用して O(1) 追加空間にする）:

1. answer を長さ n で確保する。
2. 一度左から走査して answer[i] に left[i] を格納（左累積を直接 answer
   に入れる）。初期 running_left = 1。ループ内で answer[i] = running_left;
   running_left *= nums[i]。
3. 次に右から走査し、running_right を初期 1 として各 i で answer[i] *=
   running_right; その後 running_right *= nums[i]。これで answer[i] は
   left[i]*right[i] となる。

時間 O(n)、追加の定数空間（running_left, running_right 等のみ）。

注意点:

- ゼロを含む場合にもこの方法は正しく動作する（ゼロを含むと prefix/suffix が 0
  になる箇所が自動的に反映される）。
- 乗算の中間結果が 32-bit
  に収まることが問題で保証されているため、実装言語の型に注意（Zig では i32
  を使うのが妥当）。

## キーワード

- prefix product, suffix product
- 積を除算せずに計算
- スキャンアルゴリズム（2-pass）
- O(n) time, O(1) extra space（出力配列を除く）
