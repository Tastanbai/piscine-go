package main

import (
	"ascii-art/functions"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	w, _, errr := getTerminalSize()
	if errr != nil {
		fmt.Println(errr)
		return
	}
	text := ""
	style := "standard.txt"
	if len(os.Args) > 1 {
		text = os.Args[1]
	} else {
		return
	}
	if len(os.Args) > 2 {
		style = os.Args[2] + ".txt"
	}
	res := ""
	err := ""
	text = functions.SingleLine(text)
	textArr := strings.Split(text, "\n")
	var resArr []string
	var EmpetyTexts bool = false
	for i := 0; i < len(textArr); i++ {
		if len(textArr[i]) == 0 {
			EmpetyTexts = true
		} else {
			EmpetyTexts = false
			break
		}
	}
	if EmpetyTexts == true {
		for i := 0; i < len(textArr)-1; i++ {
			fmt.Println()
		}
		return
	}
	for i := 0; i < len(textArr); i++ {
		res, err = functions.AsciiReturner(textArr[i], style)
		if err != "" {
			fmt.Println(err)
			return
		}
		resArr = append(resArr, res)
		res = ""
		res += "\n"
	}
	res = res[:len(res)-1]
	for i := 0; i < len(resArr); i++ {
		if len(resArr[i]) == 0 {
			fmt.Print(textArr[i])
		} else {
			if w > len(resArr[i])/8 {
				fmt.Print(resArr[i])
			} else {
				fmt.Println("Too Long string. The weight is", w, ", your strings len is", len(resArr[i])/8)
			}
		}
	}
}
func getTerminalSize() (width, height int, err error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}
	size := strings.Split(strings.TrimSpace(string(out)), " ")
	if len(size) != 2 {
		return 0, 0, fmt.Errorf("неверный формат вывода stty size")
	}
	fmt.Sscanf(size[0], "%d", &height)
	fmt.Sscanf(size[1], "%d", &width)
	return width, height, nil
}
