package lss

import (
	"bufio"
	"fmt"
	"io"
	"os"

	xpkg "github.com/xuri/excelize/v2"
)

type Xfile struct{ *xpkg.File }

func (f Xfile) Close() error         { return f.File.Close() }
func (f Xfile) SheetNames() []string { return f.File.GetSheetList() }

type Reader struct{ io.Reader }

func (r Reader) ToBuffered() Reader { return Reader{bufio.NewReader(r.Reader)} }

func (r Reader) ToXfile() (Xfile, error) {
	f, e := xpkg.OpenReader(r.Reader)
	return Xfile{File: f}, e
}

func (r Reader) ToSheetNames() ([]string, error) {
	x, e := r.ToXfile()
	if nil != e {
		return nil, e
	}

	return x.SheetNames(), nil
}

func (r Reader) SheetNamesToWriter(w Writer) error {
	names, e := r.ToSheetNames()
	if nil != e {
		return e
	}
	return w.WriteStrings(names)
}

var ReaderStdin Reader = Reader{os.Stdin}.ToBuffered()

type Writer struct{ *bufio.Writer }

func (w Writer) Flush() error { return w.Writer.Flush() }

func (w Writer) WriteStrings(strs []string) error {
	for _, s := range strs {
		_, e := fmt.Fprintln(w.Writer, s)
		if nil != e {
			return e
		}
	}
	return w.Flush()
}

var WriterStdout Writer = Writer{bufio.NewWriter(os.Stdout)}

func StdinToXlsxToSheetNamesToStdout() error {
	return ReaderStdin.SheetNamesToWriter(WriterStdout)
}
