# Collection

![Build Status](https://github.com/wilhelm-murdoch/go-collection/actions/workflows/go.yml/badge.svg)
[![GoDoc](https://godoc.org/github.com/wilhelm-murdoch/go-collection?status.svg)](https://pkg.go.dev/github.com/wilhelm-murdoch/go-collection)
[![Go report](https://goreportcard.com/badge/github.com/wilhelm-murdoch/go-collection)](https://goreportcard.com/report/github.com/wilhelm-murdoch/go-collection)

A generic collection for Go with a few convenient methods.
# Install
```
go get github.com/wilhelm-murdoch/go-collection
```
# Usage
Import `go-collection` with the following:
```go
import (
  "github.com/wilhelm-murdoch/lo"
)
```
And use it like so:
```go
fruits := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot")
fmt.Println("Fruits:", fruits.Length())

fruits.Each(func(index int, item string) bool {
  fmt.Println("-", item)
  return false
})
```
Which, using the above example, will yield the following output:
```
Fruits: 6
- apple
- orange
- strawberry
- cherry
- banana
- apricot
```
# Methods
* Filter
* Slice
* Contains
* Shift
* Unshift
* At
* IsEmpty
* Empty
* Find
* FindIndex
* RandomIndex
* Random
* LastIndexOf
* Reduce
* Reverse
* Some
* None
* All
* Push
* Pop
* Length
* Map
* Each
* Concat
* InsertAt
* InsertBefore
* InsertAfter
* AtFirst
* AtLast
* Count
* CountBy

## Filter
Filter returns a new collection with items that have passed predicate check.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
  out := c.Filter(func(item string) bool { 
    return item == "strawberry" || item == "banana"
  })
  // []string{"strawberry, "banana'}
}
```
## Slice
Slice returns a new slice of the current collection starting with `from` and `to` indexes.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## Contains
Contains returns true if an item is present in the current collection.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## Shift
Shift method removes the first item from the current collection, then returns that item.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## Unshift
Unshift method appends one item to the beginning of the current collection, returning the new length of the collection.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## At
At attempts to return the item associated with the specified index for the current collection along with a boolean value stating whether or not an item could be found.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## IsEmpty
IsEmpty returns a boolean value describing the empty state of the current collection.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## Empty
Empty will reset the current collection to zero items.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## Find
Find returns the first item in the provided current collectionthat satisfies the provided testing function. If no items satisfy the testing function, a `nil` value is returned.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## FindIndex
FindIndex returns the index of the first item in the specified collection that satisfies the provided testing function. Otherwise, it returns `-1`, indicating that no element passed the test.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## RandomIndex
RandomIndex returns the index associated with a random item from the current collection.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## Random
Random returns a random item from the current collection.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## LastIndexOf
LastIndexOf returns the last index at which a given item can be found in the current collection, or `-1` if it is not present.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## Reduce
Reduce reduces a collection to a single value. The value is calculated by accumulating the result of running each item in the collection through an accumulator function. Each successive invocation is supplied with the return value returned by the previous call.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## Reverse
Reverse the current collection so that the first item becomes the last, the second item becomes the second to last, and so on.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## Some
Some returns a true value if at least one item within the current collection resolves to true as defined by the predicate function f.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## None
None returns a true value if no items within the current collection resolve to true as defined by the predicate function f.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## All
All returns a true value if all items within the current collection resolve to true as defined by the predicate function f.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## Push
Push method appends one or more items to the end of a collection, returning the new length.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## Pop
Pop method removes the last item from the current collection and then returns that item.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## Length
Length returns number of items associated with the current collection.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## Map
Map method creates to a new collection by using callback invocation result on each array item. On each iteration f is invoked with arguments: index and current item. It should return the new collection.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## Each
Each iterates through the specified list of items executes the specified callback on each item. This method returns the current instance of collection.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## Concat
Concat merges two slices of items. This method returns the current instance collection with the specified slice of items appended to it.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## InsertAt
InsertAt inserts the specified item at the specified index and returns the current collection. If the specified index is less than 0, 0 is used. If an index greater than the size of the collectio nis specified, c.Push is used instead.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## InsertBefore
InsertBefore inserts the specified item before the specified index and returns the current collection. If the specified index is less than 0, 0 is used. If an index greater than the size of the collectio nis specified, c.Push is used instead.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## InsertAfter
InsertAfter inserts the specified item after the specified index and returns the current collection. If the specified index is less than 0, 0 is used. If an index greater than the size of the collectio nis specified, c.Push is used instead.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## AtFirst
AtFirst attempts to return the first item of the collection ago-collectionng with a boolean value stating whether or not an item could be found.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## AtLast
AtLast attempts to return the last item of the collection ago-collectionng with a boolean value stating whether or not an item could be found.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## Count
Count counts the number of items in the collection that compare equal to value.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```
## CountBy
CountBy counts the number of items in the collection for which predicate is true.
```go
package main

import "github.com/wilhelm-murdoch/go-collection"

func main() {
  c := collection.New("apple", "orange", "strawberry", "cherry", "banana", "apricot", "avacado", "beans", "beets")
}
```