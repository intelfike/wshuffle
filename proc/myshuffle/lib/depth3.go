package lib

import "math/rand"

// type
// 数字=0　英語=1　記号=3　ひらがな=4　カタカナ=5　漢字=6　その他=7
func GetCharType(r rune) int {
	if '0' <= r && r <= '9' {
		return 0
	}
	if '０' <= r && r <= '９' {
		return 0
	}

	if 'a' <= r && r <= 'z' {
		return 1
	}
	if 'ａ' <= r && r <= 'ｚ' {
		return 1
	}
	if 'A' <= r && r <= 'Z' {
		return 1
	}
	if 'Ａ' <= r && r <= 'Ｚ' {
		return 1
	}

	if '!' <= r && r <= '~' {
		return 3
	}
	if 'ぁ' <= r && r <= 'ゔ' {
		return 4
	}
	if 'ァ' <= r && r <= 'ヴ' {
		return 5
	}
	if '亜' <= r && r <= '㿿' {
		return 6
	}
	return -1
}

// 文字列を完全ランダムにシャッフルする
func StringShuffle(s string) string {
	rs := []rune(s)
	for n, _ := range rs {
		ran := int(rand.Int31()) % len(rs)
		rs[n], rs[ran] = rs[ran], rs[n]
	}
	return string(rs)
}

// 先頭と末尾以外シャッフル
func MiddleShuffle(s string) string {
	rs := []rune(s)
	shuffled := StringShuffle(string(rs[1 : len(rs)-1]))
	result := make([]rune, 0)
	result = append(result, rs[0])
	result = append(result, []rune(shuffled)...)
	result = append(result, rs[len(rs)-1])
	return string(result)
}
