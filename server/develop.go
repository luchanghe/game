package main

import (
	"server/develop"
	"time"
)

func main() {
	develop.CreateModelToPb()
	develop.CreatePb()
	time.Sleep(time.Second)
	develop.CreateAction()
}
