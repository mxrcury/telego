package cmd

import (
	"authtg/pkg/telegram"
	"authtg/pkg/utils"
	"bufio"
	"fmt"
	"os"
)


func Init(api *telegram.TelegramAPI) error {
  utils.ClearTerminal()
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("You are in interactive telegram API shell") // TODO(text): improve this text
  fmt.Println("Enter ':help' to get list of commands") // TODO(text): improve this text

  for {
    chatIDs := map[int]int{}
    receivedCommand, err := reader.ReadString('\n')
    if err != nil {
      break
    }
       
    if utils.IsHelpCommand(receivedCommand) {
      utils.ClearTerminal()
    fmt.Println(`List of commands:
:chats - to get list of your chats
:me - get info about yourself
:chat <Nth> - to get more info of the specific chat in received list
--------
:logout - to logout and end telegram session
:exit - to exit from there
`)
      utils.PrintInput()
    } else if utils.IsGetMeCommand(receivedCommand) {
      utils.ClearTerminal()
      user, err := telegram.GetMe(api)
      if err != nil{
        fmt.Println("error:", err)
        continue
      }
      fmt.Printf(`Username:  %s
Name:      %s %s
ID:        %d
`, user.Username, user.FirstName, user.LastName, user.ID)
      utils.PrintInput()
    } else if utils.IsChatsCommand(receivedCommand) {
      utils.ClearTerminal()
      chats, err := telegram.GetChats(api)
      if err != nil {
        fmt.Println("error:", err)
        continue
      }
      fmt.Println("CHATS:")
      for i, chat := range chats.Chats {
        fullChat, isFull := chat.AsFull()
        if isFull {
fmt.Printf(`
%d) %s
   ID - %d
`, i +1, fullChat.GetTitle(), fullChat.GetID())
        } else {
fmt.Printf(`
%d) ID - %d
`, i +1, chat.GetID())
        }
        chatIDs[i+1] = int(chat.GetID())
      }
      utils.PrintInput()
    } else if utils.IsLogOutCommand(receivedCommand) {
      break
    } else if utils.IsExitCommand(receivedCommand) {
      os.Exit(1)
    } else if number, isValid := utils.IsChatInfoCommand(receivedCommand); isValid{
      fmt.Printf("CHAT NUMBER %d, CHAT IDS INSIDE M:%d, MAP: %v\n", number, chatIDs[number], chatIDs)
      resp, err := telegram.GetChatInfo(api, chatIDs[number])
      if err != nil {
        fmt.Println("error:", err)
        continue
      }
      fmt.Printf(
`
Chat  %d
About %s
Members %d
`, resp.ChatID, resp.About, resp.MembersLength)
      utils.PrintInput()
    }else {
      fmt.Println("NOTHING HAPPENED")
      /*
      utils.ClearTerminal()
      utils.PrintInput()
      */
    }
  }
  return nil
}
