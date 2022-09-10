package main

import (
	"errors"
	"fmt"
	"unicode"
)

var (
	src       []byte
	lookAhead = 0
)

func expr() error {
	err := term()
	if err != nil {
		return err
	}
	for {
		if lookAhead >= len(src) {
			return nil
		}
		t := rune(src[lookAhead])
		switch t {
		case '+':
			err := match(t)
			if err != nil {
				return err
			}

			err = term()
			if err != nil {
				return err
			}
			fmt.Print("+")
		case '-':
			err := match(t)
			if err != nil {
				return err
			}

			err = term()
			if err != nil {
				return err
			}
			fmt.Print("-")
		default:
			return nil
		}
	}
}

func term() error {
	t := rune(src[lookAhead])
	if unicode.IsDigit(t) {
		fmt.Print(string(t))
		match(t)
		return nil
	} else {
		return errors.New("解析错误，遇到不是数字的term")
	}
}

func match(t rune) error {
	if src[lookAhead] == byte(t) {
		lookAhead++
		return nil
	} else {
		msg := fmt.Sprintf("match不匹配，位置 %d 字符串：%c", lookAhead, src[lookAhead])
		return errors.New(msg)
	}
}
func main() {
	fmt.Println("请输入普通的算术表达式，只能单数字和加减符号,如： 3+2-2")
	var line string
	fmt.Scanf("%s", &line)
	src = []byte(line)
	fmt.Println("输入的算术表达式为： ", string(src))
	fmt.Println("-------下面分析的后缀表达式-------")
	err := expr()
	if err != nil {
		fmt.Println("错误：", err)
	}
}
