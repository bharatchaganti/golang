package main

import (
	"bytes"
	"fmt"
)

func main() {
	st := []string{"geskfeees", "reeelaee"}
	slv := solved(st)
	fmt.Println(slv)

}
func solved(strings []string) []string {
	strs := make([]string, len(strings))
	for i := 0; i < len(strings); i++ {
		strs[i] = solve(strings[i])
	}
	return strs
}

func solve(s string) string {
	var buf bytes.Buffer
	var last rune
	for i, r := range s {
		if r != last || i == 0 {
			buf.WriteRune(r)
			last = r
		}
	}
	return buf.String()
}
