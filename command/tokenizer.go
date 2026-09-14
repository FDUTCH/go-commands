package command

import (
	"bytes"
	"strings"
)

// Tokenizer tokenizes the command line.
type Tokenizer struct {
	buf       *bytes.Buffer
	tokensBuf []string
}

// Tokenize tokenizes the command line into the array of parameters.
func (r *Tokenizer) Tokenize(str string) []string {
	if r.buf == nil {
		r.buf = bytes.NewBuffer(make([]byte, 0, 10))
	}
	buf := r.buf
	str = strings.TrimPrefix(str, "/")
	result := r.tokensBuf
	inQuotes := false

	for i := range str {
		char := str[i]
		switch char {
		case '"':
			inQuotes = !inQuotes
		case ' ':
			if inQuotes {
				buf.WriteByte(char)
			} else if buf.Len() != 0 {
				result = append(result, buf.String())
				buf.Reset()
			}
		default:
			buf.WriteByte(char)
		}
	}

	if buf.Len() != 0 {
		result = append(result, buf.String())
	}

	buf.Reset()
	r.tokensBuf = result[0:]
	return result
}
