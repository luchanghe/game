package cache

import (
	"fmt"
	"testing"
)

func TestGetUser(t *testing.T) {
	u := GetUser(100000)
	fmt.Println(u)
}
