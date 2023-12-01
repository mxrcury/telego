package cmd

import (
	"authtg/pkg/telegram"
	"authtg/pkg/utils"
	"bufio"
	"context"
	"fmt"
	"os"

	tg "github.com/gotd/td/telegram"
)


func Init(client *tg.Client, ctx context.Context) error {
  utils.ClearTerminal()
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("You are in interactive telegram API shell") // TODO(text): improve this text
  fmt.Println("Enter ':help' to get list of commands") // TODO(text): improve this text

  for {
    receivedCommand, err := reader.ReadString('\n')
    if err != nil {
      break
    }
       
    if utils.IsHelpCommand(receivedCommand) {
      utils.ClearTerminal()
    fmt.Println(`List of commands:
:chats - to get list of your chats
:me - get info about yourself
--------
:logout - to logout and end telegram session
:exit - to exit from there
`)
    }
    if utils.IsGetMeCommand(receivedCommand) {
      utils.ClearTerminal()
      user, err := telegram.GetMe(client, ctx)
      if err != nil {
        fmt.Println("error:", err)
        continue
      }
      fmt.Printf(`Username:  %s
Name:      %s %s
ID:        %s\n`, user.Username, user.FirstName, user.LastName, user.ID)
    }
    if utils.IsChatsCommand(receivedCommand) {
      panic("gettings chats not implemented yet. 8====D")
    }
    if utils.IsLogOutCommand(receivedCommand) {
      break
    }
    if utils.IsExitCommand(receivedCommand) {
      os.Exit(1)
     }
  }
  return nil
}
