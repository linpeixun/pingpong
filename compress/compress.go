package compress

type CompressType byte

const (
	Nozip CompressType = iota
	Gzip
)

type Compressor interface {
	Zip(b []byte) ([]byte, error)

	Unzip(b []byte) ([]byte, error)
}

var (
	Compressors = map[CompressType]Compressor{
		Nozip: &DefaultCompressor{},
		Gzip:  &GzipCompressor{},
	}
)
