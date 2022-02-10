package main

import (
	"fmt"

	"github.com/eiannone/keyboard"

	fcommand "github.com/goodaye/fakeeyes/protos/command"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

// 使用websocket 控制
func MoveControll(conn *websocket.Conn) error {

	// 读取数据
	go func() {
		defer func() {
			err := recover()
			fmt.Println(err)
		}()
		for {

			mt, message, err := conn.ReadMessage()
			if err != nil {
				panic(err)
			}
			if mt == websocket.CloseMessage {
				break
			}
			if mt == websocket.TextMessage {
				fmt.Println("receive message back :", string(message))
			}
		}
	}()

	// 监听键盘
	if err := keyboard.Open(); err != nil {
		return err
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press ESC to quit")
	var keyname string
	var opcode fcommand.OperateCode
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			return err
		}
		fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)

		if key == keyboard.KeyEsc {
			fmt.Println("Exit!!")
			break
		}

		switch key {
		case keyboard.KeyArrowUp:
			keyname = "Up"
			opcode = fcommand.OperateCode_MoveFront
		case keyboard.KeyArrowDown:
			keyname = "Down"
			opcode = fcommand.OperateCode_MoveBack
		case keyboard.KeyArrowLeft:
			keyname = "Left"
			opcode = fcommand.OperateCode_TurnLeft
		case keyboard.KeyArrowRight:
			keyname = "Right"
			opcode = fcommand.OperateCode_TurnRight
		default:
			keyname = string(char)
		}
		fmt.Println("press key :", keyname)
		if opcode == 0 {
			continue
		}
		op := fcommand.Operation{
			Type: fcommand.Operation_Device,
			Data: &fcommand.Operation_OperateData{
				Opcode: int32(opcode),
			},
		}
		byteop, err := proto.Marshal(&op)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		conn.WriteMessage(websocket.BinaryMessage, byteop)
	}
	return nil
}
