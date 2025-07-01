# go-collections

Little collection of generic collections in golang.

## heap

```go
type CustomObject struct {
    ID int	
}

// max heap based on the ID field
h := NewHeap(func(a, b CustomObject) bool {
	return a.ID > b.ID
})

// for ordered default types there is a wrapper
h := NewOrderedHeap[int]()

```

## queue

```go
var q Queue[int]

q.Push(42)
q.Push(43)

pop := q.Pop() // returns 42
```

## set

```go
// a set can be created empty
s := NewSet[int]()

// or with initial values
s := NewSet(42, 43)

if s.Contains(42) { 
    // do something
}
```
## stack


```go
var s Stack[int]

s.Push(42)
s.Push(43)

pop := s.Pop() // returns 43
```
