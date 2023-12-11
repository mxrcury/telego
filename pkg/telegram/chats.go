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

func (a *TelegramAPI) GetChats() (*tg.MessagesDialogsSlice, error) {
  resp, err := a.Client.API().MessagesGetDialogs(a.Ctx, &tg.MessagesGetDialogsRequest{Limit: 10, OffsetPeer: &tg.InputPeerChannel{} })
  if resp, ok := resp.(*tg.MessagesDialogsSlice); ok {
    return resp, err
  }
  return nil, err
}

func (a *TelegramAPI) GetChatsWithPagination(lastChatID int) (*tg.MessagesDialogsSlice, error) {
  resp, err := a.Client.API().MessagesGetDialogs(a.Ctx, &tg.MessagesGetDialogsRequest{Limit: 10, OffsetPeer: &tg.InputPeerChannel{}, OffsetID: lastChatID })
  if resp, ok := resp.(*tg.MessagesDialogsSlice); ok {
    return resp, err
  }
  return nil, err
}

func (a *TelegramAPI) GetChatInfo(id int) (*Chat, error) {
  fmt.Println("ID:", id)
  resp, err := a.Client.API().MessagesGetFullChat(a.Ctx, int64(id))
  if err != nil {
    return nil, err
  }
  membersLength := len(resp.GetUsers())
  chat := resp.GetFullChat()

  return &Chat{ChatID: chat.GetID(), About: chat.GetAbout(), MembersLength: membersLength}, nil
}
