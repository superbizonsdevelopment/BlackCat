package command

import(
	"fmt"
)

type ConsoleSender struct {}

type CommandSender interface {
	SendMessage(string)
	
	IsUser() bool
	
	HasPermission(int) bool
}


func(sender ConsoleSender) SendMessage(message string) {
	fmt.Println(message)
}

func(sender ConsoleSender) IsUser() {
	return false
}

func(sender ConsoleSender) HasPermission(permissionLvL int) bool{
	return true
}