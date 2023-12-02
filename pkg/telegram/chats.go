package telegram

import (
	"fmt"

	"github.com/gotd/td/tg"
)

type Chat struct{
  ChatID  int64 
  About   string
  MembersLength int
}

func GetChats(api *TelegramAPI) (*tg.MessagesDialogsSlice, error) {
  resp, err := api.Client.API().MessagesGetDialogs(api.Ctx, &tg.MessagesGetDialogsRequest{Limit: 10, OffsetPeer: &tg.InputPeerChannel{}})
  if resp, ok := resp.(*tg.MessagesDialogsSlice); ok {
    return resp, err
  }
  return nil, err
}
func GetChatInfo(api *TelegramAPI, id int) (*Chat, error) {
  resp, err := api.Client.API().MessagesGetFullChat(api.Ctx, int64(id))
  fmt.Println("RESPP:", resp)
  if err != nil {
    return nil, err
  }
  membersLength := len(resp.GetUsers())
  chat := resp.GetFullChat()

  return &Chat{ChatID: chat.GetID(), About: chat.GetAbout(), MembersLength: membersLength}, nil
}
