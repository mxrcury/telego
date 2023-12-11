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
  chatIDs := []int{} // TODO: make it slice and push in first all chats, then take the last element and make pagnation after every last dialog
  for {
    receivedCommand, err := reader.ReadString('\n')
    if err != nil {
      break
    }
       
    if utils.IsHelpCommand(receivedCommand) {
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
    } else if utils.IsGetMeCommand(receivedCommand) {
      utils.ClearTerminal()
      user, err := api.GetMe()
      if err != nil{
        fmt.Println("error:", err)
        continue
      }
      fmt.Printf(
`Username:  %s
Name:      %s %s
ID:        %d
`, user.Username, user.FirstName, user.LastName, user.ID)
      utils.PrintInput()
    } else if utils.IsChatsCommand(receivedCommand){
      err := HandleGetChats(api, chatIDs)
      if err != nil {
        fmt.Println("error:", err)
        continue
      }
    } else if utils.IsLogOutCommand(receivedCommand) {
      break
    } else if utils.IsExitCommand(receivedCommand) {
      os.Exit(1)
    } else if number, isValid := utils.IsChatInfoCommand(receivedCommand); isValid {
      resp, err := api.GetChatInfo(chatIDs[number])
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
