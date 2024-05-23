package array

import (
  "errors"
  "strings"
  "fmt"
)

//import "datastructures/factory"

type Array struct {
  elements []interface{}
  size        int
  capacity    int 
}

func NewArray() *Array {
  capacity := 10
  return &Array{
    elements: make([]interface{}, capacity),
    size: 0,
    capacity: capacity,
  }
}

func (a *Array) Push(element interface{}) {
  if a.size == a.capacity {
    a.resize(a.capacity * 2)
  }
  a.elements[a.size] = element
  a.size ++
}

func (a *Array) Pop() (interface{}, error) {
  if a.size == 0 {
    return nil, errors.New("Array is empty")
  }
  last_element := a.elements[a.size -1]
  a.elements[a.size - 1] = nil
  a.size --
  return last_element, nil
}

func (a *Array) Shift() (interface{}, error) {
  if (a.size == 0) {
    return nil, errors.New("Array is empty")
  }
  first_element := a.elements[0]
  for i := 0; i < a.size -1 ; i ++{
    a.elements[i] = a.elements[i + 1]
  }
  a.size --
  a.elements[a.size] = nil
  return first_element, nil
}

func (a *Array) Unshift(elements ... interface{}) int {
  numNewElements := len(elements)
  newSize := a.size + numNewElements

  if newSize > a.capacity {
    newCapacity := a.capacity
    for newCapacity < newSize {
      newCapacity *= 2
    }

    a.resize(newCapacity)
  } 
  // Shift Right
  for i := a.size; i >= 0; i -- {
    a.elements[i + numNewElements] = a.elements[i] 
  }

  // Insert elements in the start
  for i, element := range elements {
    a.elements[i] = element
  }

  a.size = newSize
  return a.size - 1

}

func (a *Array) Concat(arrays ... *Array) *Array {
  // calc size of original array 
  totalSize := a.size

  // calc size of news arrays
  for _, arr := range arrays {
    totalSize += arr.size
  }

  // Create a new arrays with the new capacity
  result := &Array{
    elements: make([]interface{}, totalSize),
    size: totalSize,
    capacity: totalSize * 2,
  }

  // Copy the elements of the orinal arrays
  copy(result.elements, a.elements[:a.size])

  // Copy the elements of the new array(S)
  offset := a.size
  for _, arr := range arrays {
    copy(result.elements[offset:], arr.elements[:arr.size])
    offset += arr.size
  }
  return result
} 

func (a *Array) resize(newCapacity int) {
  newElements := make([]interface{}, newCapacity)
  copy(newElements, a.elements)
  a.elements = newElements
  a.capacity = newCapacity
}

func (a *Array) String() string {
    var sb strings.Builder
    sb.WriteString("[")
    for i, elem := range a.elements[:a.size] {
        if i > 0 {
            sb.WriteString(", ")
        }
        sb.WriteString(fmt.Sprintf("%v", elem))
    }
    sb.WriteString("]")
    return sb.String()
}
