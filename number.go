package faker

import "math"

func (f *Faker) RandomInt8() int8 {
	return int8(f.r.Intn(math.MaxInt8))
}

func (f *Faker) RandomInt16() int16 {
	return int16(f.r.Intn(math.MaxInt16))
}

func (f *Faker) RandomInt32() int32 {
	return f.r.Int31()
}

func (f *Faker) RandomInt64() int64 {
	return f.r.Int63()
}

func (f *Faker) RandomUint8() uint8 {
	return uint8(f.r.Intn(math.MaxUint8))
}

func (f *Faker) RandomUint16() uint16 {
	return uint16(f.r.Intn(math.MaxUint16))
}

func (f *Faker) RandomUint32() uint32 {
	return f.r.Uint32()
}

func (f *Faker) RandomUint64() uint64 {
	return f.r.Uint64()
}

func (f *Faker) RandomInt() int {
	return f.r.Int()
}

func (f *Faker) RandomUint() uint {
	return uint(f.r.Intn(math.MaxInt))
}

func (f *Faker) RandomFloat32() float32 {
	return f.r.Float32()
}

func (f *Faker) RandomFloat64() float64 {
	return f.r.Float64()
}

func (f *Faker) RandomBetween(min, max int) int {
	if max <= min {
		panic("max must be greater than min")
	}
	return f.r.Intn(max-min) + min
}
