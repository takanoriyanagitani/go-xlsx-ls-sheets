package main

import (
	"log"

	lss "github.com/takanoriyanagitani/go-xlsx-ls-sheets"
)

func main() {
	e := lss.StdinToXlsxToSheetNamesToStdout()
	if nil != e {
		log.Printf("%v\n", e)
	}
}
