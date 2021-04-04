package conversions

import (
	"bufio"
	"io"
	"os"
)

func FromFile2JsonArray(filename string) ([]byte, error) {

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b := bufio.NewReader(f)
	return fromTest2Json(b)

}

func FromStout2JsonArray(src io.Reader) ([]byte, error) {
	return fromTest2Json(src)
}

func fromTest2Json(src io.Reader) ([]byte, error) {

	fbrace := []byte("[")[0]
	lbrace := []byte("]")[0]
	ccbrace := []byte("}")[0]
	lbreak := []byte("\n")[0]
	ocbrace := []byte("{")[0]
	comma := []byte(",")[0]

	buf, err := io.ReadAll(src)
	if err != nil {
		return nil, err
	}
	blen := len(buf)
	outbuf := make([]byte, 0, blen*2)

	outbuf = append(outbuf, fbrace)
	for i := 1; i < blen; i++ {
		if i != blen-1 {
			if buf[i-1] == ccbrace && buf[i] == lbreak && buf[i+1] == ocbrace {
				outbuf = append(outbuf, buf[i-1])
				outbuf = append(outbuf, comma)
			} else {
				outbuf = append(outbuf, buf[i-1])
			}
		} else {
			outbuf = append(outbuf, buf[i-1])
		}
	}
	outbuf = append(outbuf, lbrace)
	return outbuf, nil
}
