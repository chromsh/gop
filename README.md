# gop
Converting golang primitive data types to pointers.

# Usage

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