package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/intelfike/wshuffle/io/input"
	"github.com/intelfike/wshuffle/io/output"
	"github.com/intelfike/wshuffle/proc/myshuffle"
)

func main() {
	s, err := input.ReadString()
	if err != nil {
		fmt.Println("入力エラー")
		os.Exit(1)
	}
	t := time.Now()
	rand.Seed(int64(t.Nanosecond()))
	shuffled := myshuffle.MyShuffle(s, 4, "")
	output.WriteString(shuffled)
}
