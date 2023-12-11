package telegram

import (
	"github.com/gotd/td/tg"
)


func (a *TelegramAPI) GetMe() (*tg.User, error){
  status, err := a.Client.Auth().Status(a.Ctx)
  if err != nil {
    return nil, err
  }
  return status.User, nil
}
