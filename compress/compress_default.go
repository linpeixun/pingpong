package compress

//没有压缩
type DefaultCompressor struct {
}

func (c *DefaultCompressor) Zip(b []byte) ([]byte, error) {
	return b, nil
}

func (c *DefaultCompressor) Unzip(b []byte) ([]byte, error) {
	return b, nil
}
