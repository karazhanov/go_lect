package concurrent

import (
	"fmt"
	"time"
)

func say(s string)  {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Second)
	}
}

func TestConcurrent() {
	say("Hello")
	say("World")
}