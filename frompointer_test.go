package gop

import (
	"testing"
)

func TestBoolVal(t *testing.T) {
	type args struct {
		v *bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"nil", args{nil}, false},
		{"true", args{Bool(true)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoolVal(tt.args.v); got != tt.want {
				t.Errorf("BoolVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringVal(t *testing.T) {
	type args struct {
		v *string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"nil", args{nil}, ""},
		{"`ok`", args{String(`ok`)}, "ok"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringVal(tt.args.v); got != tt.want {
				t.Errorf("StringVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntVal(t *testing.T) {
	type args struct {
		v *int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"nil", args{nil}, 0},
		{"1", args{Int(1)}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntVal(tt.args.v); got != tt.want {
				t.Errorf("IntVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8Val(t *testing.T) {
	type args struct {
		v *int8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{"nil", args{nil}, 0},
		{"1", args{Int8(1)}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int8Val(tt.args.v); got != tt.want {
				t.Errorf("Int8Val() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16Val(t *testing.T) {
	type args struct {
		v *int16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{"nil", args{nil}, 0},
		{"1", args{Int16(1)}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int16Val(tt.args.v); got != tt.want {
				t.Errorf("Int16Val() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32Val(t *testing.T) {
	type args struct {
		v *int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{"nil", args{nil}, 0},
		{"1", args{Int32(1)}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32Val(tt.args.v); got != tt.want {
				t.Errorf("Int32Val() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64Val(t *testing.T) {
	type args struct {
		v *int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"nil", args{nil}, 0},
		{"1", args{Int64(1)}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64Val(tt.args.v); got != tt.want {
				t.Errorf("Int64Val() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUintVal(t *testing.T) {
	type args struct {
		v *uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{"nil", args{nil}, 0},
		{"1", args{Uint(1)}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UintVal(tt.args.v); got != tt.want {
				t.Errorf("UintVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8Val(t *testing.T) {
	type args struct {
		v *uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{"nil", args{nil}, 0},
		{"1", args{Uint8(1)}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint8Val(tt.args.v); got != tt.want {
				t.Errorf("Uint8Val() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16Val(t *testing.T) {
	type args struct {
		v *uint16
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{"nil", args{nil}, 0},
		{"1", args{Uint16(1)}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint16Val(tt.args.v); got != tt.want {
				t.Errorf("Uint16Val() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32Val(t *testing.T) {
	type args struct {
		v *uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"nil", args{nil}, 0},
		{"1", args{Uint32(1)}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint32Val(tt.args.v); got != tt.want {
				t.Errorf("Uint32Val() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64Val(t *testing.T) {
	type args struct {
		v *uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"nil", args{nil}, 0},
		{"1", args{Uint64(1)}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64Val(tt.args.v); got != tt.want {
				t.Errorf("Uint64Val() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUintptrVal(t *testing.T) {
	type args struct {
		v *uintptr
	}
	tests := []struct {
		name string
		args args
		want uintptr
	}{
		{"nil", args{nil}, 0},
		{"1", args{Uintptr(1)}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UintptrVal(tt.args.v); got != tt.want {
				t.Errorf("UintptrVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteVal(t *testing.T) {
	type args struct {
		v *byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{"nil", args{nil}, 0},
		{"1", args{Byte(1)}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteVal(tt.args.v); got != tt.want {
				t.Errorf("ByteVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRuneVal(t *testing.T) {
	type args struct {
		v *rune
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{"nil", args{nil}, 0},
		{"1", args{Rune(1)}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RuneVal(tt.args.v); got != tt.want {
				t.Errorf("RuneVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32Val(t *testing.T) {
	type args struct {
		v *float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{"nil", args{nil}, 0},
		{"1", args{Float32(1)}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32Val(tt.args.v); got != tt.want {
				t.Errorf("Float32Val() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64Val(t *testing.T) {
	type args struct {
		v *float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"nil", args{nil}, 0},
		{"1", args{Float64(1)}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64Val(tt.args.v); got != tt.want {
				t.Errorf("Float64Val() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64Val(t *testing.T) {
	type args struct {
		v *complex64
	}
	tests := []struct {
		name string
		args args
		want complex64
	}{
		{"nil", args{nil}, 0},
		{"1", args{Complex64(1)}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Complex64Val(tt.args.v); got != tt.want {
				t.Errorf("Complex64Val() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128Val(t *testing.T) {
	type args struct {
		v *complex128
	}
	tests := []struct {
		name string
		args args
		want complex128
	}{
		{"nil", args{nil}, 0},
		{"1", args{Complex128(1)}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Complex128Val(tt.args.v); got != tt.want {
				t.Errorf("Complex128Val() = %v, want %v", got, tt.want)
			}
		})
	}
}
