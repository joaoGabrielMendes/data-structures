package main

import (
	"datastructures/array"
)

func main () {
  var a array.Array = *array.NewArray();

  // Push Test
  println("inserindo 1, 2, 3")
  a.Push(1);
  a.Push(2);
  a.Push(3);
  println(a.String())
  println()

  // Pop Test
  println("Retirando ultimo elemento")
  a.Pop()
  println(a.String());
  println()

  // Shift Test
  println("Tirando primeiro elemento")
  a.Shift();
  println(a.String())
  println()

  // Unshift Test
  println("Inserindo no comeco [0, 1]")
  println("(size: ", a.Unshift(0, 1), ")")
  println(a.String())
  println()

  // Concat Test
  b := *array.NewArray()
  c := *array.NewArray()
  var concatBC array.Array

  b.Push(1); b.Push(2); b.Push(3)
  c.Push(4); c.Push(5); c.Push(6)
  concatBC = *b.Concat(&c)

  println("Concatenado b com c ")
  println("B: ",b.String(), "C: ", c.String())
 
  println("Concat: ", concatBC.String())
  println()

  // Slice Test
  slicedArray := *b.Slice(0, 2)
  println("Array slicado b = [1,2,3] -> slice(0, 2)")
  println(slicedArray.String())
  println()

  // Map test
  arr := *array.NewArray()
  arr.Push(1)
  arr.Push(2)
  arr.Push(3)

  // Aplicar uma função que multiplica cada elemento por 2
  newArr := arr.Map(func(element interface{}) interface{} {
    return element.(int) * 2
  })

  println(newArr.String()) // Saída: [2 4 6]
  println()

  // For each test
  arr.ForEach(func(element interface{}) {
    println("element: ")
    println(element.(int) * 2)
  })
}
