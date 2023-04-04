package fq

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
)

type Queue struct {
	mu        sync.Mutex
	filename  string
	blockSize int
	maxSize   uint8
	header    []byte // 队列头部数据
}

func NewQueue(filename string, blockSize int, maxSize uint8) (*Queue, error) {
	if blockSize <= 0 {
		return nil, errors.New("invalid block size")
	}

	if maxSize <= 0 {
		return nil, errors.New("invalid max size")
	}

	q := &Queue{
		filename:  filename,
		blockSize: blockSize,
		maxSize:   maxSize,
		header:    make([]byte, 8),
	}

	err := q.Init()
	if err != nil {
		return nil, err
	}

	return q, nil
}

func (q *Queue) Init() error {
	file, err := os.OpenFile(q.filename, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	// 清理文件中已出队的数据
	if err := q.Clean(); err != nil {
		fmt.Println("init clean err")
	}

	// 读取 header
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	fileSize := fileInfo.Size()

	if fileSize < int64(q.blockSize+8) {
		q.header[0] = 1
		q.header[1] = 0
		q.header[2] = 0
		q.header[3] = 0
		q.header[4] = 0
		q.header[5] = 0
		q.header[6] = 0
		q.header[7] = 8
	} else {
		buf := make([]byte, q.blockSize)
		_, err = file.ReadAt(buf, 8)
		if err != nil {
			return err
		}

		q.header[0] = buf[0]
		q.header[1] = buf[1]
		q.header[2] = buf[2]
		q.header[3] = buf[3]
		q.header[4] = buf[4]
		q.header[5] = buf[5]
		q.header[6] = buf[6]
		q.header[7] = buf[7]
	}

	return nil
}

func (q *Queue) Enqueue(element []byte) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	// 扩展文件大小以容纳新的数据
	file, err := os.OpenFile(q.filename, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	fileSize := fileInfo.Size()
	_, err = file.WriteAt(element, fileSize)
	if err != nil {
		return err
	}

	// 检查是否需要清理文件
	if q.header[2]+uint8(len(element)) > q.maxSize {
		if err := q.Clean(); err != nil {
			return err
		}
	}

	// 更新 header
	q.header[3] = uint8(fileSize & 0xff)
	q.header[4] = uint8((fileSize >> 8) & 0xff)
	q.header[5] = uint8((fileSize >> 16) & 0xff)
	q.header[6] = uint8((fileSize >> 24) & 0xff)

	return nil
}

func (q *Queue) Dequeue() ([]byte, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	// 读取队列头部数据
	file, err := os.OpenFile(q.filename, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	fileSize := fileInfo.Size()

	if fileSize < int64(q.blockSize+8) {
		return nil, errors.New("queue is empty")
	}

	buf := make([]byte, q.blockSize)
	_, err = file.ReadAt(buf, int64(q.header[7]))
	if err != nil {
		return nil, err
	}

	// 更新 header，并返回出队元素
	q.header[7] += uint8(q.blockSize)
	q.header[2] -= uint8(q.blockSize)

	return buf, nil
}

func (q *Queue) Clean() error {
	q.mu.Lock()
	defer q.mu.Unlock()

	// 打开文件
	file, err := os.OpenFile(q.filename, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	// 读取 header
	buf := make([]byte, 8)
	_, err = file.ReadAt(buf, 0)
	if err != nil {
		return err
	}

	// 计算需要清理的字节数
	pos := int64(q.header[7])
	fileSize, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}

	cleanSize := pos - 8

	if cleanSize > 0 {
		// 将未出队的数据移动到文件开头
		_, err = file.Seek(pos, io.SeekCurrent)
		if err != nil {
			return err
		}

		batchSize := int64(q.blockSize) * 1024 * 10 // 每次处理 10KB 的数据
		for i := int64(0); i < cleanSize; i += batchSize {
			if i+batchSize > cleanSize {
				batchSize = cleanSize - i
			}

			buf := make([]byte, batchSize)
			_, err = file.ReadAt(buf, pos+i)
			if err != nil {
				return err
			}

			_, err = file.WriteAt(buf, 8+i)
			if err != nil {
				return err
			}
		}

		// 截断文件
		err = file.Truncate(fileSize - cleanSize)
		if err != nil {
			return err
		}

		// 更新 header
		q.header[2] -= uint8(cleanSize)
		q.header[3] = uint8(fileSize & 0xff)
		q.header[4] = uint8((fileSize >> 8) & 0xff)
		q.header[5] = uint8((fileSize >> 16) & 0xff)
		q.header[6] = uint8((fileSize >> 24) & 0xff)

		_, err = file.WriteAt(q.header, 0)
		if err != nil {
			return err
		}
	}

	return nil
}
