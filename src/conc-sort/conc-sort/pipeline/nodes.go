package pipeline

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

// 临时变量，记录开始时间
var startTime time.Time

// 初始化开始状态
func Init() {
	startTime = time.Now()
}

// 打印开始时间计时
func PrintTime(event string) {
	fmt.Println(event + ": " + time.Now().Sub(startTime).String())
}

// 将数组转换成一个 chan
func ArraySource(a ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()

	return out
}

// 对 chan 的数据进行内存排序（调用 sort.Ints() ）
func InMemSort(in <-chan int) <-chan int {
	out := make(chan int, 1024)

	go func() {
		// Read into memory
		a := []int{}
		for v := range in {
			a = append(a, v)
		}
		PrintTime("Read done")

		// Sort
		sort.Ints(a)
		PrintTime("InMemSort done")

		// Output
		for _, v := range a {
			out <- v
		}

		close(out)
	}()

	return out
}

// 合并两个排序完成的 chan（归并两个）
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 1024)

	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2

		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}

		PrintTime("Merge done")
		close(out)
	}()

	return out
}

// 读取文件为 chan
func ReaderSource(
	reader io.Reader, chunkSize int) <-chan int {
	out := make(chan int, 1024)

	go func() {
		buffer := make([]byte, 8)
		bytesRead := 0
		for {
			n, err := reader.Read(buffer)
			bytesRead += 8
			if n > 0 {
				v := int(
					binary.BigEndian.Uint64(buffer))
				out <- v
			}
			if err != nil ||
				(chunkSize != -1 &&
					bytesRead >= chunkSize) {
				break
			}
		}

		close(out)
	}()

	return out
}

// 写文件
func WriterSink(writter io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(
			buffer, uint64(v))
		writter.Write(buffer)
	}
}

// 生成随数
func RandomSource(count int) <-chan int {
	out := make(chan int, 1024)

	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()

	return out
}

// 将 N 个 chan 进行合并（递归两两合并）
func MergeN(inputs ...<-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}
	m := len(inputs) / 2
	// merge inputs[0..m) and inputs,m..end)
	return Merge(
		MergeN(inputs[:m]...),
		MergeN(inputs[m:]...),
	)
}
