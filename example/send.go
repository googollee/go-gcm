package main

import (
	"fmt"
	"github.com/googollee/go-gcm"
)

func main() {
	client := gcm.New("your Google APIs key")

	load := gcm.NewMessage("your device id")
	load.AddRecipient("abc")
	load.SetPayload("data", "1")
	load.CollapseKey = "demo"
	load.DelayWhileIdle = true
	load.TimeToLive = 10

	resp, err := client.Send(load)

	fmt.Printf("id: %+v\n", resp)
	fmt.Println("err:", err)
	fmt.Println("err index:", resp.ErrorIndexes())
	fmt.Println("reg index:", resp.RefreshIndexes())
}
