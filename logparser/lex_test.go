package logparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestLex parses a log into fields.
//
// asdfasfafdaf
func TestLex(t *testing.T) {
	as := assert.New(t)

	log := `192.168.0.164 - - [07/Mar/2023:01:25:26 +0800] "GET /ping.txt HTTP/1.1" 400 255 "-" "Go-http-client/1.1" "-" asdf asdf234; wer " dfdfd"`
	fields := Lex(log, 10)

	as.Len(fields, 11)
}
