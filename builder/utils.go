package builder

import (
	"bytes"
	"compress/gzip"
)

// Compress to GZ
func gz(data []byte) []byte {
	var gz bytes.Buffer
	w := gzip.NewWriter(&gz)
	w.Write(data)
	w.Close()

 return gz.Bytes()
}
