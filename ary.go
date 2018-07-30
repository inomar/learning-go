package main

func main() {
  // 配列型
  var a [3]int
  b := [2]string{ "hoge", "fuga" }
  c := [...]int{ 1, 2 }
  d := [3]int{ 1:1, 2:10}
  e := [2]string{0:"hello",1:"world" }
  println(a[0])
  println(b[1]) // fuga
  println(c[0]) // 1
  println(d[0], d[1]) // 1, 10
  println(e[0],e[1],) // hello, world

  println(len(e),len(a)) // 3, 2 
  
  // スライス型
  var f []int // nil
  g := []string{"hoge", "fuga"}
  h := []int{1,2}
  println(g[0],g[1])
  println(h[0],len(f))
}
