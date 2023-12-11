package cmd

import (
	"authtg/pkg/telegram"
	"authtg/pkg/utils"
	"fmt"
)

func HandleGetChats(api *telegram.TelegramAPI, chatsIDs []int) (error){
utils.ClearTerminal()
      chats, err := api.GetChats()
      if err != nil {
        return err
      }
      for i, chat := range chats.Chats {
        fullChat, isFull := chat.AsFull()
        if isFull {
fmt.Printf(
`
%d) %s
   ID - %d
`, i +1, fullChat.GetTitle(), fullChat.GetID())
        } else {
fmt.Printf(
`
%d) ID - %d
`, i +1, chat.GetID())
        }
        chatsIDs = append(chatsIDs, int(chat.GetID()))
      }
      utils.PrintInput()
return nil
}
