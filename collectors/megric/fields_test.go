package megric

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestField_Match(t *testing.T) {
	as := assert.New(t)
	regex := "^\\s*(.*)\\s+login:\\s+\\w+.*$"
	f, err := NewFiledByRegex(regex)
	as.Nil(err)

	fields, ok := f.Match("20230302T12:00:00Z+800 login: usera")
	as.True(ok)
	as.Len(fields, 2)
}

func TestField_Split(t *testing.T) {
	as := assert.New(t)
	f, err := NewFieldByDelimiter(" ")
	as.Nil(err)

	fields, ok := f.Match("20230302T12:00:00Z+800 login: usera")
	as.True(ok)
	as.Len(fields, 3)
}
