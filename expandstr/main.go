package main

import "os"

func isvalid(v rune) bool {
	if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || (v >= '0' && v <= '9') {
		return true
	}
	return false
}
func main() {
	if len(os.Args) == 2 {
		hasword := false
		res := ""
		k := 0
		for _, v := range os.Args[1] {
			if isvalid(v) {
				hasword = true
				k = 1
			} else if v == ' ' {
				hasword = false
				if k == 1 {
					res += "   "
					k = 0
				}
			}
			if hasword {
				res += string(v)
			}
		}
		if res[len(res)-1] == ' ' {
			res = res[:len(res)-3]
		}
		os.Stdout.Write([]byte(res + "\n"))
	}
}
