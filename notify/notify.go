package notify

import (
	"log/slog"
	"wechat-router/config"

	"github.com/wxpusher/wxpusher-sdk-go"
	"github.com/wxpusher/wxpusher-sdk-go/model"
)

func WxPushNotice(text string) error {
	msg := model.NewMessage(config.Cfg.WxPushAppToken).SetContent(text).AddUId(config.Cfg.WxPushUid)
	msgResult, err := wxpusher.SendMessage(msg)
	slog.Info("result is", "msg_result", msgResult)
	return err
}
