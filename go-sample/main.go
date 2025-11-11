package main // パッケージ宣言（必須）

import (
	"fmt"
	"math/rand"
) // インポート宣言

// n 個の乱数を返す関数
func randomArray(n int) []int {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	return arr
}

func main() { // エントリーポイント
	fmt.Println("Hello, Go!")

	// 乱数の取得 rand.Intn(n)
	var a int = rand.Intn(100)
	fmt.Println(a)

	//var randoms [10]int
	// 配列 randoms に 0〜99 の乱数を10個設定
	//for i := range randoms {
	//	randoms[i] = rand.Intn(100)
	//}

	randoms := randomArray(10)

	//var input [3]int
	//配列 input に適当な数字 (0～99) を3個設定する
	//for i := range input {
	//	input[i] = rand.Intn(100)
	//}

	input := randomArray(3)

	// randoms, input の各配列の値と randoms に input の数字が何個含まれているかを出力する
	fmt.Println("randoms:", randoms)
	fmt.Println("input  :", input)

	//count := 0
	//for _, in := range input {
	//	for _, r := range randoms {
	//		if in == r {
	//			count++
	//			break // 同じ数字が複数あっても1回カウント
	//		}
	//	}
	//}
	//fmt.Printf("→ input の数字のうち、randoms に含まれているのは %d 個です。\n", count)

	attempts := 0
	for {
		attempts++
		allFound := true
		for _, v := range input {
			found := false
			for _, r := range randoms {
				if v == r {
					found = true
					break
				}
			}
			if !found {
				allFound = false
				// 見つからない場合は randoms に新しい乱数を追加（または置き換え）
				randoms = append(randoms, rand.Intn(100))
			}
		}
		if allFound {
			break
		}
	}

	fmt.Printf("input の数字が全部 randoms に含まれるまでの回数: %d\n", attempts)
	fmt.Println("最終 randoms:", randoms)

}
