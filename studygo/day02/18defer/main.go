package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y)) // 走到这里，登记defer calc AA，所以要先把calc A算出来 =3，然后入栈calc(AA,1,3)
	x = 0
	defer calc("BB", x, calc("B", x, y)) // 同上，算出calc(B,x,y)=2，登记defer calc(BB,0,2)，入栈
	y = 1
}
