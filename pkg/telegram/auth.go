package telegram

import (
	"context"
	"errors"
	"path/filepath"

	"github.com/gotd/td/session"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
	"go.uber.org/zap"
)

// TODO(proxy) part of configs client:
/*
   Device: telegram.DeviceConfig{
     Proxy: tg.InputClientProxy{
       Address: options.Address,
       Port: options.Port,
     },
   },
*/
// TODO: make TelegramAPI as creating session and method from it as struct methods
// but not pass api every time

// FIX: needs to be changed

type SignInRequest struct{
  PhoneNumber string

  Code string 
  CodeHash  string
}

type TelegramAPI struct {
  Client *telegram.Client
  Ctx context.Context
}

type TelegramClientOptions struct {
  Address string
  Port  int
  TdFilesDir  string
  TdDbDir string
  PhoneNumber string
}

type TelegramStorage struct{
  DBDir string
  FileDir string
}

func sessionFolder(phone string) string {
	var out []rune
	for _, r := range phone {
		if r >= '0' && r <= '9' {
			out = append(out, r)
		}
	}
	return "phone-" + string(out)
}

func NewTelegramClient(options *TelegramClientOptions) *telegram.Client {
 sessionPath := filepath.Join("sessions", sessionFolder(options.PhoneNumber))

 return telegram.NewClient(
    API_ID,
    API_HASH,
    telegram.Options{
     SessionStorage: &session.FileStorage{Path: filepath.Join(sessionPath, "session.json")},
      Logger: zap.L(),
    },
  )
}

func NewSession( client *telegram.Client, f func(clientContext context.Context) error) error {
  if err := client.Run(context.Background(), f); err != nil {
    return err
	}
  return nil
}

func NewTelegramAPI(client *telegram.Client, ctx context.Context) *TelegramAPI{
  return &TelegramAPI{
    Client: client,
    Ctx: ctx,
  }
}

func (a *TelegramAPI) GetAuthCode(phoneNumber string) (*tg.AuthSentCode, error) {
    //requestData := &tg.AuthSendCodeRequest{PhoneNumber: phoneNumber, APIID: API_ID, APIHash: API_HASH}
    
    resp, err := a.Client.Auth().SendCode(a.Ctx, phoneNumber, auth.SendCodeOptions{})

    if resp, ok := resp.(*tg.AuthSentCode); ok {
      return resp, nil
    }
		return nil, err
}

func (a *TelegramAPI) SignIn(body *SignInRequest) (interface{}, error) {
  //requestData := &tg.AuthSignInRequest{PhoneNumber: body.PhoneNumber, PhoneCodeHash: body.CodeHash, PhoneCode: body.Code}
  resp, err := a.Client.Auth().SignIn(a.Ctx,body.PhoneNumber, body.Code, body.CodeHash)
  if err != nil {
    return nil, err
  }
  return resp, nil
}

func Is2FAError(err error) bool {
  TwoFAError := errors.New("2FA required")
  return errors.As(err, &TwoFAError)
}
