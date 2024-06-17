package first

import (
	"bytes"
	"github.com/WiiLink24/MiiContestChannel/common"
	"math"
)

type Addition struct {
	Header       AdditionHeader
	Countries    []CountryField
	Skills       []SkillField
	MarqueeField MarqueeField
}

type AdditionHeader struct {
	Type         [2]byte
	_            [6]byte
	CountryGroup uint32
	_            [12]byte
	Padding      [8]byte
}

type CountryField struct {
	Type        [2]byte
	FieldSize   uint16
	CountryCode uint32
	Text        [192]byte
}

type SkillField struct {
	Type      [2]byte
	FieldSize uint16
	SkillId   uint32
	Text      [96]byte
}

type MarqueeField struct {
	Tag         [2]byte
	SectionSize uint16
	Unknown     uint32
	Text        [1536]byte
}

func (a Addition) ToBytes(any) []byte {
	buffer := new(bytes.Buffer)
	common.WriteBinary(buffer, a.Header)

	for _, country := range a.Countries {
		common.WriteBinary(buffer, country)
	}

	for _, skill := range a.Skills {
		common.WriteBinary(buffer, skill)
	}

	common.WriteBinary(buffer, a.MarqueeField)

	return buffer.Bytes()
}

func MakeAddition() {
	marqueeText := []byte("WiiLink Mii Contest Channel!!!!")
	var actual [1536]byte
	copy(actual[:], marqueeText)

	country := CountryField{
		Type:        [2]byte{'N', 'H'},
		FieldSize:   200,
		CountryCode: 18,
		Text:        [192]byte{'C', 'a', 'n', 'a', 'd', 'a'},
	}

	skill := SkillField{
		Type:      [2]byte{'N', 'J'},
		FieldSize: 104,
		SkillId:   1,
		Text:      [96]byte{'C', 'a', 'n', 'a', 'd', 'a'},
	}

	addition := Addition{
		Header: AdditionHeader{
			Type:         [2]byte{'A', 'D'},
			CountryGroup: 201,
			Padding:      [8]byte{math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8},
		},
		Countries: []CountryField{country},
		Skills:    []SkillField{skill},
		MarqueeField: MarqueeField{
			Tag:         [2]byte{'N', 'W'},
			SectionSize: 1544,
			Unknown:     1,
			Text:        actual,
		},
	}

	common.Write(addition, "addition/201.ces")
}
