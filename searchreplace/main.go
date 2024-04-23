package main

import "os"

func main() {
	if len(os.Args) == 4 {
		res := ""
		for _, v := range os.Args[1] {
			if string(v) == os.Args[2] {
				res += os.Args[3]
			} else {
				res += string(v)
			}
		}
		os.Stdout.Write([]byte(res + "\n"))
	}

}
