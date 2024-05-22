package main

import "datastructures/array"

func main () {
  var a array.Array = *array.NewArray();

  println("inserindo 1, 2, 3")
  a.Push(1);
  a.Push(2);
  a.Push(3);
  println(a.String())
  println()

  println("Retirando ultimo elemento")
  a.Pop()
  println(a.String());
  println()

  println("Tirando primeiro elemento")
  a.Shift();
  println(a.String())
  println()
}
