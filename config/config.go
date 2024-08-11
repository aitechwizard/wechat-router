package config

import (
	"os"

	"github.com/joho/godotenv"
)

var Cfg Config

type Config struct {
	WxPushAppToken string `json:"wx_push_app_token"`
	WxPushUid      string `json:"wx_push_uid"`
}

func InitConfig() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	token := os.Getenv("WX_PUSH_APP_TOKEN")
	uid := os.Getenv("WX_PUSH_UID")

	Cfg = Config{
		WxPushAppToken: token,
		WxPushUid:      uid,
	}
	return nil
}
