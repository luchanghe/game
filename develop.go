package main

import (
	"game/develop"
	"time"
)

func main() {
	develop.CreateModelToPb()
	develop.CreatePb()
	time.Sleep(time.Second)
	develop.CreateAction()
}
