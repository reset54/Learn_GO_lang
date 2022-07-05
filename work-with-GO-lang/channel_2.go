package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan bool)

	go func() {
		fmt.Println(1)
		time.Sleep(time.Second * 5)
		close(c)
	}()

	// Останавливает выполнение до получения сообщения из канала или его закрытия.
	<-c

	fmt.Println("Done")
}