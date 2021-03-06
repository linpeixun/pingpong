package util

import "unsafe"

// 无拷贝，减少内存使用
func SliceByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// 无拷贝，减少内存使用
func StringToSliceByte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// 无拷贝，减少内存使用
func CopyMeta(src, dst map[string]string) {
	if dst == nil {
		return
	}
	for k, v := range src {
		dst[k] = v
	}
}
