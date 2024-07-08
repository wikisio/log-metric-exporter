package logparser

import (
	"bytes"
)

// Lex scan and split the line into fields according to Common Log Format and Extended Log Format.
//
//	[Common Log Format]:https://wikipedea.org
//	1. Spaces are the basic delimiter.
//	2. Text between bracket [] considered as one field.
//	3. Text between double quotes "" considered as one field.
//	4. After maxFields parsed, the rest will be considered as one fields. the total return fields may be maxFields + 1
//	5. Escape char such as "\"", or embedded delimiter such as [[abc]]
func Lex(line string, maxFields int) []string {
	var fields []string
	fieldBuf := bytes.Buffer{}
	inBracket := false
	inQuote := false
	for i, c := range line {
		if len(fields) >= maxFields {
			fieldBuf.WriteString(line[i:])
			break
		}

		switch c {
		case ' ':
			if !inBracket && !inQuote {
				addField(&fieldBuf, &fields)
			} else {
				fieldBuf.WriteRune(c)
			}
		case '[':
			inBracket = true
			addField(&fieldBuf, &fields)
		case ']':
			inBracket = false
			addField(&fieldBuf, &fields)
		case '"':
			if !inQuote {
				inQuote = true
				addField(&fieldBuf, &fields)
			} else {
				inQuote = false
				addField(&fieldBuf, &fields)
			}
		default:
			fieldBuf.WriteRune(c)
		}
	}

	addField(&fieldBuf, &fields)

	return fields
}

func addField(fieldBuf *bytes.Buffer, fields *[]string) {
	if fieldBuf.Len() > 0 {
		*fields = append(*fields, fieldBuf.String())
		fieldBuf.Reset()
	}
}
