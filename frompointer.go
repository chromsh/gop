package gop

func BoolVal(v *bool) bool {
	if v == nil {
		return false
	}
	return *v
}

func StringVal(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func IntVal(v *int) int {
	if v == nil {
		return 0
	}
	return *v
}

func Int8Val(v *int8) int8 {
	if v == nil {
		return 0
	}
	return *v
}

func Int16Val(v *int16) int16 {
	if v == nil {
		return 0
	}
	return *v
}

func Int32Val(v *int32) int32 {
	if v == nil {
		return 0
	}
	return *v
}

func Int64Val(v *int64) int64 {
	if v == nil {
		return 0
	}
	return *v
}

func UintVal(v *uint) uint {
	if v == nil {
		return 0
	}
	return *v
}

func Uint8Val(v *uint8) uint8 {
	if v == nil {
		return 0
	}
	return *v
}

func Uint16Val(v *uint16) uint16 {
	if v == nil {
		return 0
	}
	return *v
}

func Uint32Val(v *uint32) uint32 {
	if v == nil {
		return 0
	}
	return *v
}

func Uint64Val(v *uint64) uint64 {
	if v == nil {
		return 0
	}
	return *v
}

func UintptrVal(v *uintptr) uintptr {
	if v == nil {
		return 0
	}
	return *v
}

func ByteVal(v *byte) byte {
	if v == nil {
		return 0
	}
	return *v
}

func RuneVal(v *rune) rune {
	if v == nil {
		return 0
	}
	return *v
}

func Float32Val(v *float32) float32 {
	if v == nil {
		return 0
	}
	return *v
}

func Float64Val(v *float64) float64 {
	if v == nil {
		return 0
	}
	return *v
}

func Complex64Val(v *complex64) complex64 {
	if v == nil {
		return 0
	}
	return *v
}

func Complex128Val(v *complex128) complex128 {
	if v == nil {
		return 0
	}
	return *v
}
