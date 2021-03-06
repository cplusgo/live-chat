package main

import (
	"fmt"
	"github.com/cplusgo/go-library"
)

type SampleWorker struct {

}

func (this *SampleWorker) Try() {
	fmt.Println("我在执行")
	fmt.Println("准备抛出异常")
	panic("抛出异常")
}

func (this *SampleWorker) Catch(err interface{})  {
	fmt.Println("异常已经被捕获")
}

func main()  {
	worker := &SampleWorker{}
	go_library.Run(worker)
	fmt.Println("虽然有异常，我还在执行")
}
