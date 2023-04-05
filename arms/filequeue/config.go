package filequeue

import (
	"encoding/binary"
	"os"
	"path/filepath"
)

func Int64ToBytes(i int64) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	if len(buf) < 8 {
		buf = append(make([]byte, 8-len(buf)), buf...)
	}
	return int64(binary.LittleEndian.Uint64(buf))
}

// ReplaceData 替换指定位置之后的数据
func ReplaceData(o []byte, d []byte, i int) {
	for _, item := range d {
		o[i] = item
		i += 1
	}
}

func OpenOrCreateFile(path string) (*os.File, error) {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return nil, err
	}
	return os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
}
