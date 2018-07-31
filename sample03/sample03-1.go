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
}

func one() int {
	return 1
}
