&{[0xc0001c30b0 0xc0002733c8]}
# Collection

![Build Status](https://github.com/wilhelm-murdoch/go-collection/actions/workflows/go.yml/badge.svg)
[![GoDoc](https://godoc.org/github.com/wilhelm-murdoch/go-collection?status.svg)](https://pkg.go.dev/github.com/wilhelm-murdoch/go-collection)
[![Go report](https://goreportcard.com/badge/github.com/wilhelm-murdoch/go-collection)](https://goreportcard.com/report/github.com/wilhelm-murdoch/go-collection)

A generic collection for Go with a few convenient methods.
# Install
```
go get github.com/wilhelm-murdoch/go-collection
```
### Func New
* `func New[T any](items ...T) *Collection[T]` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L17)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:17-21` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L17-L21)

New returns a new collection of type T containing the specifieditems and their types. ( Chainable )

### Func Items
* `func (c *Collection[T]) Items() []T` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L24)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:24-26` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L24-L26)

Items returns the current collection's set of items.

### Func Filter
* `func (c *Collection[T]) Filter(f func(T) bool) (out Collection[T])` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L30)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:30-38` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L30-L38)

Filter returns a new collection with items that have passed predicate check.( Chainable )

### Func Slice
* `func (c *Collection[T]) Slice(from, to int) *Collection[T]` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L42)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:42-52` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L42-L52)

Slice returns a new collection containing a slice of the current collectionstarting with `from` and `to` indexes. ( Chainable )

### Func Contains
* `func (c *Collection[T]) Contains(item T) (found bool)` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L55)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:55-63` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L55-L63)

Contains returns true if an item is present in the current collection.

### Func PushDistinct
* `func (c *Collection[T]) PushDistinct(items ...T) int` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L69)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:69-77` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L69-L77)

PushDistinct method appends one or more distinct items to the currentcollection, returning the new length. Items that already exist within thecurrent collection will be ignored. You can check for this by comparing oldv.s. new collection lengths.

### Func Shift
* `func (c *Collection[T]) Shift() T` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L81)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:81-86` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L81-L86)

Shift method removes the first item from the current collection, thenreturns that item.

### Func Unshift
* `func (c *Collection[T]) Unshift(item T) int` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L90)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:90-93` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L90-L93)

Unshift method appends one item to the beginning of the current collection,returning the new length of the collection.

### Func At
* `func (c *Collection[T]) At(index int) (T, bool)` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L98)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:98-105` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L98-L105)

At attempts to return the item associated with the specified index for thecurrent collection along with a boolean value stating whether or not an itemcould be found.

### Func IsEmpty
* `func (c *Collection[T]) IsEmpty() bool` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L109)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:109-115` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L109-L115)

IsEmpty returns a boolean value describing the empty state of the currentcollection.

### Func Empty
* `func (c *Collection[T]) Empty() *Collection[T]` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L118)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:118-122` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L118-L122)

Empty will reset the current collection to zero items. ( Chainable )

### Func Find
* `func (c *Collection[T]) Find(f func(i int, item T) bool) (item T)` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L127)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:127-135` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L127-L135)

Find returns the first item in the provided current collectionthat satisfiesthe provided testing function. If no items satisfy the testing function,a <nil> value is returned.

### Func FindIndex
* `func (c *Collection[T]) FindIndex(f func(i int, item T) bool) int` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L140)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:140-148` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L140-L148)

FindIndex returns the index of the first item in the specified collectionthat satisfies the provided testing function. Otherwise, it returns -1,indicating that no element passed the test.

### Func RandomIndex
* `func (c *Collection[T]) RandomIndex() int` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L152)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:152-155` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L152-L155)

RandomIndex returns the index associated with a random item from the currentcollection.

### Func Random
* `func (c *Collection[T]) Random() (T, bool)` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L158)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:158-161` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L158-L161)

Random returns a random item from the current collection.

### Func LastIndexOf
* `func (c *Collection[T]) LastIndexOf(item T) int` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L165)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:165-174` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L165-L174)

LastIndexOf returns the last index at which a given item can be found in thecurrent collection, or -1 if it is not present.

### Func Reduce
* `func (c *Collection[T]) Reduce(f func(i int, item, accumulator T) T) (out T)` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L180)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:180-186` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L180-L186)

Reduce reduces a collection to a single value. The value is calculated byaccumulating the result of running each item in the collection through anaccumulator function. Each successive invocation is supplied with the returnvalue returned by the previous call.

### Func Reverse
* `func (c *Collection[T]) Reverse() *Collection[T]` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L190)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:190-195` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L190-L195)

Reverse the current collection so that the first item becomes the last, thesecond item becomes the second to last, and so on. ( Chainable )

### Func Some
* `func (c *Collection[T]) Some(f func(i int, item T) bool) bool` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L199)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:199-207` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L199-L207)

Some returns a true value if at least one item within the current collectionresolves to true as defined by the predicate function f.

### Func None
* `func (c *Collection[T]) None(f func(i int, item T) bool) bool` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L211)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:211-220` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L211-L220)

None returns a true value if no items within the current collection resolve totrue as defined by the predicate function f.

### Func All
* `func (c *Collection[T]) All(f func(i int, item T) bool) bool` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L224)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:224-233` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L224-L233)

All returns a true value if all items within the current collection resolve totrue as defined by the predicate function f.

### Func Push
* `func (c *Collection[T]) Push(items ...T) int` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L237)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:237-240` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L237-L240)

Push method appends one or more items to the end of a collection, returningthe new length.

### Func Pop
* `func (c *Collection[T]) Pop() (out T, found bool)` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L244)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:244-253` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L244-L253)

Pop method removes the last item from the current collection and thenreturns that item.

### Func Length
* `func (c *Collection[T]) Length() int` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L256)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:256-258` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L256-L258)

Length returns number of items associated with the current collection.

### Func Map
* `func (c *Collection[T]) Map(f func(int, T) T) (out Collection[T])` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L263)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:263-269` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L263-L269)

Map method creates to a new collection by using callback invocation resulton each array item. On each iteration f is invoked with arguments: index andcurrent item. It should return the new collection. ( Chainable )

### Func Each
* `func (c *Collection[T]) Each(f func(int, T) bool) *Collection[T]` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L274)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:274-282` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L274-L282)

Each iterates through the specified list of items executes the specifiedcallback on each item. This method returns the current instance ofcollection. ( Chainable )

### Func Concat
* `func (c *Collection[T]) Concat(items []T) *Collection[T]` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L286)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:286-289` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L286-L289)

Concat merges two slices of items. This method returns the current instancecollection with the specified slice of items appended to it. ( Chainable )

### Func InsertAt
* `func (c *Collection[T]) InsertAt(item T, index int) *Collection[T]` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L295)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:295-310` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L295-L310)

InsertAt inserts the specified item at the specified index and returns thecurrent collection. If the specified index is less than 0, 0 is used. If anindex greater than the size of the collectio nis specified, c.Push is usedinstead. ( Chainable )

### Func InsertBefore
* `func (c *Collection[T]) InsertBefore(item T, index int) *Collection[T]` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L316)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:316-318` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L316-L318)

InsertBefore inserts the specified item before the specified index andreturns the current collection. If the specified index is less than 0,c.Unshift is used. If an index greater than the size of the collection isspecified, c.Push is used instead. ( Chainable )

### Func InsertAfter
* `func (c *Collection[T]) InsertAfter(item T, index int) *Collection[T]` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L324)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:324-326` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L324-L326)

InsertAfter inserts the specified item after the specified index and returnsthe current collection. If the specified index is less than 0, 0 is used. Ifan index greater than the size of the collectio nis specified, c.Push is usedinstead. ( Chainable )

### Func AtFirst
* `func (c *Collection[T]) AtFirst() (T, bool)` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L330)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:330-332` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L330-L332)

AtFirst attempts to return the first item of the collection along with aboolean value stating whether or not an item could be found.

### Func AtLast
* `func (c *Collection[T]) AtLast() (T, bool)` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L336)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:336-338` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L336-L338)

AtLast attempts to return the last item of the collection along with aboolean value stating whether or not an item could be found.

### Func Count
* `func (c *Collection[T]) Count(item T) (count int)` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L341)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:341-348` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L341-L348)

Count counts the number of items in the collection that compare equal to value.

### Func CountBy
* `func (c *Collection[T]) CountBy(f func(T) bool) (count int)` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L351)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:351-358` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L351-L358)

CountBy counts the number of items in the collection for which predicate is true.

### Func MarshalJSON
* `func (c *Collection[T]) MarshalJSON() ([]byte, error)` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L362)
* `/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go:362-371` [#](/Users/wilhelm/Development/github.com/wilhelm-murdoch/go-collection/collection.go#L362-L371)

MarshalJSON implements the Marshaler interface so the current collection'sitems can be marshalled into valid JSON.


# License
Copyright © 2022 [Wilhelm Murdoch](https://wilhelm.codes).

This project is [MIT](./LICENSE) licensed.
