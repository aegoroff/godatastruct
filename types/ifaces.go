package types

// Integer defines all integer types interface
type Integer interface {
	int | int64 | int16 | int8 | int32 | uint | uint8 | uint16 | uint32 | uint64
}
