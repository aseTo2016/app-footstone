package signal

import (
	"github.com/aseTo2016/app-footstone/pkg/runtime"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var receiver = make(chan os.Signal)

func init() {
	signal.Notify(receiver, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
}

func Listen() {
	go func() {
		select {
		case data := <-receiver:
			log.Printf("%v", data)
			log.Println(runtime.Stack())
		}
	}()
}
