package compress

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

//gzip
type GzipCompressor struct {
}

func (c *GzipCompressor) Zip(b []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	_, err := w.Write(b)
	if err != nil {
		return nil, err
	}
	err = w.Flush()
	if err != nil {
		return nil, err
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c *GzipCompressor) Unzip(b []byte) ([]byte, error) {
	gr, err := gzip.NewReader(bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	defer gr.Close()
	b, err = ioutil.ReadAll(gr)
	if err != nil {
		return nil, err
	}
	return b, err
}
