package main

import (
	"bytes"
	"fmt"
)

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello World"))

	var pw Writer = &PointerConsoleWriter{}
	pw.Write([]byte("\nHello Pointer\n"))

	//
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i <= 5; i++ {
		fmt.Println(inc.Increment())
	}

	//
	var wc WriterCloser = NewBufferWriterCloser()
	wc.Write([]byte("\nXin chao cac ban minh la Hieu developer"))
	wc.Close()

	var obj interface{} = NewBufferWriterCloser()
	owc, err := obj.(WriterCloser)
	if err == false {
		fmt.Println("Da co loi xay ra khi chuyen kieu")
	}
	owc.(WriterCloser).Write([]byte("\nXin chao cac ban minh la Nuew\n"))
	owc.(WriterCloser).Close()

	//
	var obj2 interface{} = 2.1
	switch obj2.(type) {
	case int:
		fmt.Println("So nguyen")
	case float64:
		fmt.Println("So thuc")
	default:
		fmt.Println("Khong ton tai")

	}
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type PointerConsoleWriter struct{}

func (cw *PointerConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

//
type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

//
type WriterSecond interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	WriterSecond
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
