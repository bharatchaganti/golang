package main

import (
	"fmt"
	"net/http"
	"github.com/mitchellh/mapstructure"
	"github.com/gorilla/websocket"
)

type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}
type Channel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	//fmt.Println("Sample prog")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4000", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello")
	//	var socket *websocket.Conn
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// msgType, msg, err := socket.ReadMessage()
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		var inMessage Message
		if err := socket.ReadJSON(&inMessage); err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%#v\n", inMessage)
		switch inMessage.Name{
		case "channel add ":
			addChannel(inMessage.Data)
		}
		// fmt.Println(string(msg))
		// if err = socket.WriteMessage(msgType, msg); err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
	}
	//upgrader.Upgrade(w,r,nil)

}
func addChannel(data interface{}) ( error) {
	var channel Channel
	//channelMap:=data.(map[string]interface{})
	//channel.Name=channelMap["name"].(string)
	err := mapstructure.Decode(data, &channel)
	if err != nil {
		return err
	}
	channel.Id = "1"
	fmt.Printf("%#v\n", channel)
	return  nil
}
