package myshuffle

// ホコリがまとまるみたいに、小さい文字列を結合してまとめる
import (
	"fmt"
	"os"
	"strings"

	lib2 "github.com/intelfike/wshuffle/proc/myshuffle/lib"
)

// 主目的
// ギリギリ理解できるかもしれない程度に文章をシャッフル
// limLen = 最低の単語長
func MyShuffle(s string, limLen int, joiner string) string {
	if len(s) < limLen {
		return ""
	}
	ss := TypeSplit(s)
	merged := DustMerge(ss, limLen)
	shuffled := StringsShuffle(merged, limLen)
	// fmt.Println(ss, "\n", merged)
	return strings.Join(shuffled, joiner)
}

func DustMerge(ss []string, limLen int) []string {
	if len(ss) == 0 {
		fmt.Println("第一引数の長さが短いです")
		os.Exit(1)
	}
	result := make([]string, 0)
	word := ss[0]
	for _, v := range ss[1:] {
		if len([]rune(word)) >= limLen {
			result = append(result, word)
			word = ""
		}
		word += v
	}
	result = append(result, word)
	return result
}

// 文字種別に応じて分割する
// getCharType依存
func TypeSplit(s string) []string {
	rs := []rune(s)
	ss := make([]string, 0)
	prevCt := lib2.GetCharType(rs[0])
	sGroup := make([]rune, 0)

	sGroup = append(sGroup, rs[0])
	for _, v := range rs[1:] {
		ct := lib2.GetCharType(v)
		// ー(伸ばし棒)は直前の文字を引き継ぐ
		if prevCt == 4 || prevCt == 5 {
			if v == 'ー' {
				ct = prevCt
			}
		}
		if prevCt != ct {
			ss = append(ss, string(sGroup))
			sGroup = make([]rune, 0)
		}
		sGroup = append(sGroup, v)
		prevCt = ct
	}
	ss = append(ss, string(sGroup))
	return ss
}

// 特定の長さ以上の文字列のみシャッフルする
func StringsShuffle(ss []string, limSize int) []string {
	for n, v := range ss {
		if len(v) >= limSize {
			ss[n] = lib2.MiddleShuffle(v)
		}
	}
	return ss
}
