# pointer
Abstractions for creating pointers to non-addressable Go primitives.

For example,
```go
func somefunc() int64 {
  return 1
}
val := &somefunc()
```
```
./README.md:9:9: cannot take the address of somefunc()
```
Instead,
```go
func somefunc() int64 {
  return 1
}
val := pointer.Int64(somefunc())
```
