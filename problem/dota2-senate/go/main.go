package main

import (
	"bufio"
	"fmt"
	"os"
)

func predictPartyVictory(senate string) string {
	n := len(senate)
	// キューとしてスライスを使う（先頭ポップはインデックスで管理）
	r := make([]int, 0, n)
	d := make([]int, 0, n)
	for i, ch := range senate {
		if ch == 'R' {
			r = append(r, i)
		} else {
			d = append(d, i)
		}
	}
	ri, di := 0, 0 // 各キューの先頭インデックス
	for ri < len(r) && di < len(d) {
		if r[ri] < d[di] {
			// R が先に行動 -> D の先頭を ban（ポップ）
			// R は次ラウンドに参加する（インデックスに n を加えて末尾に追加）
			r = append(r, r[ri]+n)
			ri++
			di++ // d[di] は ban されたのでスキップ
		} else {
			// D が先に行動
			d = append(d, d[di]+n)
			di++
			ri++
		}
	}
	if ri < len(r) {
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
