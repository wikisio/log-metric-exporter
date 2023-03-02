package main

import (
	"fmt"

	"github.com/wikisio/collectors/megric"
)

func main() {
	e := megric.NewExporter()
	fmt.Println(e)
}
