package telegram

import (
	"context"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/tg"
)


func GetMe(client *telegram.Client, ctx context.Context) (*tg.User, error){
  status, err := client.Auth().Status(ctx)
  if err != nil {
    return nil, err
  }
  return status.User, nil
}
