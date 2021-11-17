# gop
Converting golang primitive data types to pointers. And vice versa.

# Usage

## to pointer

```go
import (
	"github.com/chromsh/gop"
)

type SomeStruct struct {
	Name *string
	Age *int
}

func main() {
	stringp := gop.String("foo")
	intp := gop.Int(10)
	bytep := gop.Byte(0x00)

	s := SomeStruct{
		Name: gop.String("chrom"),
		Age: gop.Int(100),
	}
	...
}
```

## from pointer

```go
import (
	"github.com/chromsh/gop"
)

type Request struct {
	Title *string
	Size *int
}

func main() {
	...
	req := Request{}
	err := json.Unmarshal(&req, body)
	if err != nil {
		panic(err)
	}
	title := gop.StringVal(req.Title)
	size := gop.IntVal(req.Size)
}
```

# Supported types

All Basic types are supported.

- bool
- string
- int
- int8
- int16
- int32
- int64
- uint
- uint8
- uint16
- uint32
- uint64
- uintptr
- byte
- rune
- float32
- float64
- complex64
- complex128