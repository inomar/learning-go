package main

func main() {
  println("start")
  sub()
  println("end")
}

func sub() {
  println("sub start")
  panic("panic!!!")
  println("sub end")
}
