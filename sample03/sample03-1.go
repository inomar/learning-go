package main

import (
	"fmt"
)

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
	callFunction(func() {
		fmt.Println("I'm a function")
	})

	f2 := later()

	fmt.Println(f2("Golang"))
	fmt.Println(f2("is"))
	fmt.Println(f2("awesome"))

	ints := integers()

	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	x, y := ones()
	fmt.Printf("x=%d, y=%d", x, y)
	switch_num()
	switch_case()

	map_sample()
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

func callFunction(f func()) {
	fmt.Println("before I'm a function")
	f()
}

// クロージャ
func later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

// ジェネレータ
func integers() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

const X = 1

func ones() (int, int) {

	const Y = 2
	return X, Y

}

func switch_num() {
	switch n := 2; n {
	case 1, 3, 5, 7, 9:
		fmt.Printf("%d is odd\n", n)
	case 2, 4, 6, 8, 10:
		fmt.Printf("%d is even\n", n)
	}
}

func switch_case() {
	n := 4
	switch {
	case n > 0 && n < 3:
		fmt.Println("1 ~ 2")
	case n > 3 && n < 6:
		fmt.Println("3 ~ 5")
	}
}

func map_sample() {
	m := make(map[int]string)
	m2 := make(map[string]string)

	m[1] = "US"
	m[3] = "Japan"
	m2["US"] = "アメリカ"
	m2["Japan"] = "日本"

	fmt.Println(m, m2)

	m3 := map[int]string{1: "Taro", 2: "Hanako", 3: "Jiro"}
	fmt.Println(m3)

	m4 := map[string][]int{
		"one":   []int{1},
		"two":   []int{1, 2},
		"three": []int{1, 2, 3}, // <= この書き方の場合は最後にカンマが必要
	}
	fmt.Println(m4)

	m5 := map[int][]int{
		1: {1},
		2: {1, 2},
		3: {1, 2, 3},
	}
	fmt.Println(m5)

	m6 := map[int]map[string]int{
		1: {"one": 1, "two": 2},
		2: {"おはよう": 10, "こんにちは": 20, "こんばんは": 30},
		3: {"羊": 100},
	}
	fmt.Println(m6)
}
