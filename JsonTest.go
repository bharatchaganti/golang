package main

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}
type Channel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

/*
type Speaker interface{
	Speak()
}
func (m Message) Speak(){
	fmt.Println("I am a "+m.Name+" Event")
}
func SomeFun(speaker Speaker){
	speaker.Speak()
}
*/

func main() {
	recRawMsg := []byte(`{"name":"channel add",` + `"data":{"name":"Hardware support"}}`)
	var recMessage Message
	err := json.Unmarshal(recRawMsg, &recMessage)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", recMessage)
	if recMessage.Name == "channel add" {
		channel, err := addChannel(recMessage.Data)
		var sendMessage Message
		sendMessage.Name = "channel add"
		sendMessage.Data = channel
		sendRawMsg, err := json.Marshal(sendMessage)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf(string(sendRawMsg))

	}
}
func addChannel(data interface{}) (Channel, error) {
	var channel Channel
	//channelMap:=data.(map[string]interface{})
	//channel.Name=channelMap["name"].(string)
	err := mapstructure.Decode(data, &channel)
	if err != nil {
		return channel, err
	}
	channel.Id = "1"
	fmt.Printf("%#v\n", channel)
	return channel, nil
}
