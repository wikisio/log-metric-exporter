package megric

import (
	"bytes"
)

// Lex scan and split the line into fields.
// 1. spaces is the basic delimiter.
// 2. text between bracket [] considered as one field.
// 3. text between double quotes "" considered as one field. No escape char taken into account.
// 4. after maxFields parsed, the rest will be considered as one fields.
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
