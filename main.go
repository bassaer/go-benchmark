package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type Param struct {
	num int
	msg string
}

func JoinFmt(p *Param) string {
	return fmt.Sprintf("%d\t%s", p.num, p.msg)
}

func JoinPlus(p *Param) string {
	return strconv.Itoa(p.num) + "\t" + p.msg
}

func JoinBuffer(p *Param) string {
	b := bytes.NewBuffer(make([]byte, 0, 21))
	b.WriteString(strconv.Itoa(p.num))
	b.WriteString("\t")
	b.WriteString(p.msg)
	return b.String()
}

func main() {
	p := &Param{100000000, "abcdefghij"}
	fmt.Println(JoinFmt(p))
	fmt.Println(JoinPlus(p))
	fmt.Println(JoinBuffer(p))
}
