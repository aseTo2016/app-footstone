package signal

import (
	"fmt"
	"testing"
	"time"
)

func TestListen(t *testing.T) {
	Listen()

	i := 0
	for {
		fmt.Println(i)
		i++
		time.Sleep(time.Second)
	}
}
