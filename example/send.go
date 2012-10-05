package main

import (
	"fmt"
	"github.com/googollee/go-gcm"
)

func main() {
	c2dm := New("your Google APIs key")

	load := NewMessage("your device id")
	load.AddRecipient("abc")
	load.SetPayload("data", 1)
	load.CollapseKey = "demo"
	load.DelayWhileIdle = true
	load.TimeToLive = 10

	resp, err := c2dm.Send(load)

	fmt.Printf("id: %+v\n", resp)
	fmt.Println("err:", err)
	fmt.Println("err index:", resp.ErrorIndexes())
	fmt.Println("reg index:", resp.RefreshIndexes())
}
