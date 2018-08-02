package main

import "fmt"

var n string = "パッケージ変数" // mainパッケージ内ならどこからでも参照できる
var m int = 100

func main() {
	var (
		i, j int
		s    string
	)
	i = 3
	j = 5
	s = "Hello, GO"
	fmt.Println(i, j, s)

	num := one()

	fmt.Println(num)
	num = 2
	fmt.Println(num)

	fmt.Printf("nは %v\n", n)
	fmt.Printf("m=%d", m)

	zero := 0.0
	pinf := 1.0 / zero
	ninf := -1.0 / zero
	nan := zero / zero
	fmt.Println(pinf, ninf, nan)

	cmp := complex(1.0, 3)
	fmt.Println(cmp, real(cmp), imag(cmp))

	// rune型 シングルクウォート
	rn := '松'
	fmt.Printf("%v", rn)

	rw := `
	    私の
	　　名前は
	　　太郎です
	`
	fmt.Printf("%v", rw)

	ary := [4]int{2, 4, 6, 8}
	fmt.Println(ary[0], ary[3])

	var sa [2]string
	fmt.Println(sa)

	// 要素の省略
	a1 := [...]int{5, 100}
	fmt.Println(a1[1])
	a1[1] = 500
	fmt.Printf("%v\n", a1)

	var in interface{}
	fmt.Println(in)
	in = 3
	in = "太郎"
	in = '山'
	in = 1.34
	fmt.Println(in)

	fmt.Println(plus(2, 4))
	hello()

	q, r := div(5, 2)
	fmt.Println(q, r)
	q1, _ := div(5, 2)
	_, q2 := div(6, 2)
	fmt.Println(q1, q2)

	noname := func(x, y int) int { return y * x }
	fmt.Println(noname(3, 2))
	fmt.Printf("%T\n", func(x, y int) int { return x + y })
	fmt.Printf("%v\n", func(x, y int) int { return x + y }(4, 5))
}

func one() int {
	return 1
}

func plus(x, y int) int {
	return x + y
}

func hello() {
	fmt.Println("Hello")
}

func div(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}
