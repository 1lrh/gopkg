package codec

import (
	"errors"
	"strconv"
)

var ShortIDSwitch = false

var AlphanumericSet = []rune{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
}

var (
	ErrCharacterSetEmpty = errors.New("empty character set")
)

type NameGenerator struct {
	characterSet []rune
	characterMap map[rune]int
	size         uint64
}

func NewNameGenerator(characterSet []rune) (*NameGenerator, error) {
	if len(characterSet) == 0 {
		return nil, ErrCharacterSetEmpty
	}

	var array []rune
	set := make(map[rune]struct{})

	// get unique character array
	for _, c := range characterSet {
		if _, ok := set[c]; !ok {
			array = append(array, c)
			set[c] = struct{}{}
		}
	}

	g := &NameGenerator{
		characterSet: array,
		characterMap: make(map[rune]int, len(array)),
		size:         uint64(len(array)),
	}

	for i, r := range g.characterSet {
		g.characterMap[r] = i
	}

	return g, nil
}

// Encode
// you should store salt for decoding later
// salt should be a large value to avoid a small id
// bug: two same result
func (g *NameGenerator) Encode(id uint64, salt uint64) string {
	id = id + salt
	var code []rune
	for id > 0 {
		idx := id % g.size
		code = append(code, g.characterSet[idx])
		id = id / g.size
	}
	return string(code)
}

func (g *NameGenerator) Decode(code string, salt uint64) uint64 {
	var id uint64
	runes := []rune(code)
	for i := len(runes) - 1; i >= 0; i-- {
		r := runes[i]
		index := g.characterMap[r]
		id = id*g.size + uint64(index)
	}
	id = id - salt
	return id
}

const salt = int64(100)

// NumToShortID num to string
func (g *NameGenerator) NumToShortID(id int64) string {
	sid := strconv.FormatInt(id, 10)
	if len(sid) < 17 {
		return ""
	}
	sTypeCode := sid[1:4]
	sid = sid[4:int32(len(sid))]
	id, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		return ""
	}
	typeCode, err := strconv.ParseInt(sTypeCode, 10, 64)
	if err != nil {
		return ""
	}
	code := g.Encode(uint64(id), uint64(salt))
	tcode := g.Encode(uint64(typeCode), uint64(salt))
	return tcode + code
}

// ShortIDToNum string to num
func (g *NameGenerator) ShortIDToNum(code string) int64 {
	if len(code) < 2 {
		return 0
	}
	scodeType := code[0:2]
	code = code[2:int32(len(code))]

	id := g.Decode(code, uint64(salt))
	codeType := g.Decode(scodeType, uint64(salt))
	return 10000000000000000 + int64(codeType)*10000000000000 + int64(id)
}

func (g *NameGenerator) EnShortID(id string) string {
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return id
	}
	return g.NumToShortID(num)
}

func (g *NameGenerator) DeShortID(sid string) string {
	num, err := strconv.ParseInt(sid, 10, 64)
	if err != nil {
		return strconv.FormatInt(g.ShortIDToNum(sid), 10)
	}
	if num < 10000000000000000 {
		return strconv.FormatInt(g.ShortIDToNum(sid), 10)
	}
	return sid
}

func IsShortID(id string) bool {
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return true
	}
	if num < 10000000000000000 {
		return true
	}
	return false
}
