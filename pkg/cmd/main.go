package cmd

import (
	"authtg/pkg/utils"
	"fmt"
)

func HandleHelpCommand() {
  utils.ClearTerminal()
  fmt.Println(
`List of commands:
:chats <Nth> - to get list of your chats by page(every 10 dialogs)
:me - get info about yourself
:chat <Nth> - to get more info of the specific chat in received list
--------
:logout - to logout and end telegram session
:exit - to exit from there
`)
  utils.PrintInput()
}
