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
     HandleHelpCommand()
    } else if utils.IsGetMeCommand(receivedCommand) {
     HandleMeCommand(api) // TODO: change error handling
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
