package main

import (
	"fmt"
	"sort"
	"strings"
)

// gcd は 2 つの整数 a, b の最大公約数をユークリッドの互除法で返します。
// 引数 b が 0 になるまで a, b を入れ替えつつ a % b を繰り返します。
// 正の整数に対して動作しますが、0 を含む場合の扱いにも対応しています。
func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// checkRepeat は文字列 s が文字列 p を k 回繰り返したもの（k >= 0）と等しいかを判定します。
// - p が空文字列の場合は s が空文字列のときのみ true を返します。
// - s の長さが p の長さで割り切れない場合は false を返します。
// - 割り切れる場合は strings.Repeat を用いて p を必要回数繰り返した文字列と s を比較します。
func checkRepeat(s, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	if len(s)%len(p) != 0 {
		return false
	}
	times := len(s) / len(p)
	// build repeated string efficiently using strings.Repeat
	return strings.Repeat(p, times) == s
}

// divisors は正の整数 n の約数を降順のスライスで返します。
// 例: n = 12 の場合は []int{12, 6, 4, 3, 2, 1} を返します。
// 0 以下が与えられた場合は空スライスを返します。
// 実装は i を 1 から sqrt(n) まで走査し、小さい側と対応する大きい側の約数を別々に収集してから結合します。
func divisors(n int) []int {
	if n <= 0 {
		return []int{}
	}
	small := []int{}
	large := []int{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			small = append(small, i)
			if i != n/i {
				large = append(large, n/i)
			}
		}
	}
	// combine large in reverse to get descending order overall
	// collect all and sort descending
	all := append(small, large...)
	sort.Sort(sort.Reverse(sort.IntSlice(all)))
	return all
}

// gcdOfStrings は問題「2 つの文字列の最大公約数文字列」を解きます。
// 手順:
// 1. まず高速チェックとして str1+str2 と str2+str1 を比較し、等しくなければ共通の基底は存在しないため "" を返す。
// 2. str1 と str2 の長さの最大公約数 g を求め、g の約数（＝両方の長さを割る可能性がある長さ）を降順に試す。
// 3. 各候補長 d に対して、str1 の先頭 d 文字を基底 p とし、p が str1 と str2 の両方を繰り返しで構成できるかを checkRepeat で確認する。
// 4. 最初に見つかったものを返し、見つからなければ空文字列を返す。
func gcdOfStrings(str1 string, str2 string) string {
	// まず存在可能性の高速チェック
	// 入れ替えて違う文字 = 余りがでるなら割り切れず、最大公約数ではない
	if str1+str2 != str2+str1 {
		return ""
	}
	// 長さの公約数を全て列挙して長いものから試す
	len1, len2 := len(str1), len(str2)
	// possible divisors are divisors of gcd(len1, len2) or divisors of min(len1,len2)?
	// The length of the common divisor must divide both len1 and len2, so it's a divisor of gcd(len1,len2).
	g := gcd(len1, len2)
	for _, d := range divisors(g) {
		p := str1[:d]
		if checkRepeat(str1, p) && checkRepeat(str2, p) {
			return p
		}
	}
	return ""
}

func main() {
	// 簡単な標準入力/出力のサンプル。実行時には引数を与えるか、stdin から読み取る。
	// ここでは固定サンプルを出力。
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))       // "ABC"
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))      // "AB"
	fmt.Println(gcdOfStrings("LEET", "CODE"))        // ""
	fmt.Println(gcdOfStrings("ABCABCABC", "ABCABC")) // "ABCABC"
	// または stdin から読みたい場合は以下のようにする（必要に応じて有効化）
	// buf := make([]byte, 4096)
	// n, _ := os.Stdin.Read(buf)
	// input := string(buf[:n])
	// io.WriteString(os.Stdout, gcdOfStrings(...))
}
