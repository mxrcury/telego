package cmd

import (
	"authtg/pkg/telegram"
	"authtg/pkg/utils"
	"fmt"
)

func HandleMeCommand(api * telegram.TelegramAPI) {
  utils.ClearTerminal()
  user, err := api.GetMe()
  if err != nil{
    fmt.Println("error:", err)
    return
   }
   fmt.Printf(
`Username:  %s
Name:      %s %s
ID:        %d
`, user.Username, user.FirstName, user.LastName, user.ID)
   utils.PrintInput()
}
