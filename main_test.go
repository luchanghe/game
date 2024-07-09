package main

import (
	"fmt"
	"game/pb"
	"testing"
)

func TestGo(t *testing.T) {
	fmt.Println(pb.RouteMap(1001))
	fmt.Println(pb.RouteMap(1002))
}
