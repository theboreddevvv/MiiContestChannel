package common

import (
	"bytes"
	"encoding/binary"
	"io"
)

// Writable converts the inherited into a byte representation of itself.
// Structs can either use the preset function, or implement their own.
type Writable interface {
	ToBytes(data any) []byte
}

// ToBytes is the simplest form of writing a struct. Some structs may need to implement their own
// due to nested/unfixed slices, or other writable data types.
func ToBytes(data any) []byte {
	buffer := new(bytes.Buffer)
	WriteBinary(buffer, data)
	return buffer.Bytes()
}

func WriteBinary(writer io.Writer, data any) {
	err := binary.Write(writer, binary.BigEndian, data)
	if err != nil {
		panic(err)
	}
}
