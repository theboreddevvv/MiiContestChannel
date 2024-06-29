//go:build ignore

// Generates a JSON with skills and countries from an already generated addition file.

package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	_ "embed"
	"encoding/binary"
	"encoding/json"
	"github.com/wii-tools/lzx/lz10"
	"io"
	"os"
)

//go:generate go run generate_addition_json.go

//go:embed 201.ces
var addition []byte

type Root struct {
	Countries []Child `json:"countries"`
	Skills    []Child `json:"skills"`
}

type Child struct {
	Code uint32 `json:"code"`
	Name string `json:"name"`
}

func main() {
	block, err := aes.NewCipher([]byte{0x8D, 0x22, 0xA3, 0xD8, 0x08, 0xD5, 0xD0, 0x72, 0x02, 0x74, 0x36, 0xB6, 0x30, 0x3C, 0x5B, 0x50})
	if err != nil {
		panic(err)
	}

	mode := cipher.NewCBCDecrypter(block, []byte{0xBE, 0x5E, 0x54, 0x89, 0x25, 0xAC, 0xDD, 0x3C, 0xD5, 0x34, 0x2E, 0x08, 0xFB, 0x8A, 0xBF, 0xEC})

	dst := make([]byte, len(addition)-24)
	mode.CryptBlocks(dst, addition[24:])

	decompressed, err := lz10.Decompress(dst)
	if err != nil {
		panic(err)
	}

	buffer := bytes.NewReader(decompressed)
	_ = read[[32]byte](buffer)

	var root Root
	for buffer.Len() != 0 {
		tag := read[[2]byte](buffer)
		if tag == [2]byte{'N', 'H'} {
			// Country
			_ = read[[2]byte](buffer)

			code := read[uint32](buffer)
			name := read[[192]byte](buffer)

			root.Countries = append(root.Countries, Child{
				Code: code,
				Name: string(bytes.Trim(name[:], "\x00")),
			})
		} else if tag == [2]byte{'N', 'J'} {
			// Skills
			_ = read[[2]byte](buffer)

			code := read[uint32](buffer)
			name := read[[96]byte](buffer)

			root.Skills = append(root.Skills, Child{
				Code: code,
				Name: string(bytes.Trim(name[:], "\x00")),
			})
		} else if tag == [2]byte{'N', 'W'} {
			// Marquee, skip this.
			_ = read[[1542]byte](buffer)
		}
	}

	data, err := json.MarshalIndent(root, "", "\t")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("addition.json", data, 0644)
	if err != nil {
		panic(err)
	}
}

func read[T comparable](reader io.Reader) T {
	var r T
	err := binary.Read(reader, binary.BigEndian, &r)
	if err != nil {
		panic(err)
	}

	return r
}
