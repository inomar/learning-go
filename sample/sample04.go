package main


func main() {
  var fnc func(int,int) int

  fnc = add
  println(fnc(2,1))
  fnc = mul
  println(fnc(3,2))

}

func add(a int, b int) int {
  return a + b
}

func mul(a int, b int) int {
  return a * b
}
