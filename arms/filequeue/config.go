package filequeue

import "encoding/binary"

const (
	// header
	headOffset = 0
	// versionOffset 版本号在文件中下标
	versionOffset = headOffset
	// blockLenConfigOffset  数据库在文件中下标
	blockLenConfigOffset = 8
	// offsetConfigOffset 偏移量在文件中的下标
	offsetConfigOffset = 16
	// headLen head 长度 文件前 xB 的数据为header 的存储空间
	headLen int64 = 64
)

func Int64ToBytes(i int64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	if len(buf) < 8 {
		buf = append(make([]byte, 8-len(buf)), buf...)
	}
	return int64(binary.BigEndian.Uint64(buf))
}

func ReplaceData(o []byte, d []byte, i int) {
	for _, item := range d {
		o[i] = item
		i += 1
	}
}
