package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

const chunksSize = 8

type encodingTable map[rune]string

type BinaryChunk string

type BinaryChunks []BinaryChunk

func (bcs BinaryChunks) Bytes() []byte {
	res := make([]byte, 0, len(bcs))

	for _, bc := range bcs {
		res = append(res, bc.Byte())
	}
	return res
}

func (bc BinaryChunk) Byte() byte {
	num, err := strconv.ParseUint(string(bc), 2, chunksSize)
	if err != nil {
		panic("can't parse binary chunk: " + err.Error())
	}
	return byte(num)
}

// Join joins chunks into one line and returns as string
func (bcs BinaryChunks) Join() string {
	var buf strings.Builder
	for _, chunk := range bcs {
		buf.WriteString(string(chunk))
	}
	return buf.String()
}

func NewBinChunks(data []byte) BinaryChunks {
	res := make(BinaryChunks, 0, len(data))
	for _, code := range data {
		res = append(res, NewBinChunk(code))
	}
	return res
}

func NewBinChunk(code byte) BinaryChunk {
	return BinaryChunk(fmt.Sprintf("%08b", code))
}

// splitByChunks split binary string by chunks with given size

func splitByChunks(bStr string, chunksSize int) BinaryChunks {
	strLen := utf8.RuneCountInString(bStr) / chunksSize
	chunksCount := strLen / chunksSize
	// TODO: ПРОВЕРИТЬ
	if strLen%chunksSize != 0 {
		chunksCount++
	}
	res := make(BinaryChunks, 0, chunksCount)

	var buf strings.Builder

	for i, ch := range bStr {
		buf.WriteString(string(ch))

		if (i+1)%chunksSize == 0 {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}
	if buf.Len() != 0 {
		lastChunk := buf.String()
		lastChunk += strings.Repeat("0", chunksSize-len(lastChunk))
		res = append(res, BinaryChunk(lastChunk))
	}

	return res
}
