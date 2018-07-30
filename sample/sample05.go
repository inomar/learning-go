package main

func main() {
  var fnc = f()
  println(fnc())
  println(fnc())
  println(fnc())

}

func f() func()int {
  l := 0
  return func() int {
    l++
    return l
  }
}
