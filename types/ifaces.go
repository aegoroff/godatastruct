package types

// Integer defines all integer types interface
type Integer interface {
	Signed | Unsigned
}

// Signed is a constraint that permits any signed integer type.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a constraint that permits any unsigned integer type.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}
