package main

func main() {
  f()
}

func f() {
  defer println("finish")
  println("processing...")

}
