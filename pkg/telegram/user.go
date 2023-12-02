package telegram

import (
	"github.com/gotd/td/tg"
)


func GetMe(api *TelegramAPI) (*tg.User, error){
  status, err := api.Client.Auth().Status(api.Ctx)
  if err != nil {
    return nil, err
  }
  return status.User, nil
}
