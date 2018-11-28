package compress

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
)

type Zipper struct{}

func (z *Zipper) Compress(w io.Writer, data []byte) error{
	// Write gzipped data to the client
	gw, err := gzip.NewWriterLevel(w, gzip.BestCompression)
	defer gw.Close()
	gw.Write(data)
	gw.Flush()
	return err
}

func (z *Zipper) Decompress(w io.Writer, data []byte) error {
	// Write gzipped data to the client
	gr, err := gzip.NewReader(bytes.NewBuffer(data))
	defer gr.Close()
	data, err = ioutil.ReadAll(gr)
	if err != nil {
		return err
	}
	w.Write(data)
	return nil
}
