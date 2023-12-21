package main

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	"io/ioutil"
	"os"
	"test"
)

var comm *tars.Communicator

func main() {
	file, err := os.Open("./config")
	if err != nil {
		fmt.Println("文件读取错误 err=", err)
	}
	content, _ := ioutil.ReadAll(file)
	file.Close()
	say := new(test.Say)
	comm = tars.NewCommunicator()
	p := tars.NewServantProxy(comm, string(content))
	say.SetServant(p)
	var name string
	_, _ = fmt.Scanln(&name)
	res, err := say.SayHello(test.ReqHello{Name: name})
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println(res.Greeting)
	tars.Run()
}
