package megric

import (
	"fmt"
	"regexp"
	"strings"
)

type By string

const (
	Regex     By = "regex"
	Delimiter By = "delimiter"
)

type Field struct {
	by        By
	regex     *regexp.Regexp
	delimiter string
}

func NewFiledByRegex(regex string) (*Field, error) {
	f := &Field{by: Regex}
	var err error
	f.regex, err = regexp.Compile(regex)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func NewFieldByDelimiter(delimiter string) (*Field, error) {
	f := &Field{by: Delimiter}
	if delimiter == "" {
		return nil, fmt.Errorf("delimiter is empty")
	}

	f.delimiter = delimiter
	return f, nil
}

func (f *Field) Match(line string) ([]string, bool) {
	switch f.by {
	case Delimiter:
		return strings.Split(line, f.delimiter), true
	default:
		fields := f.regex.FindStringSubmatch(line)
		return fields, fields != nil
	}
}
