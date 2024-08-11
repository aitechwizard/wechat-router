package main

import (
	"log/slog"
	"os"
	"wechat-router/config"
	"wechat-router/notify"

	"wechat-router/bot"
	"wechat-router/utils"

	"github.com/eatmoreapple/openwechat"
)

func init() {
	// 1. log init
	f, _ := os.OpenFile("run.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	logger := slog.New(slog.NewJSONHandler(f, nil))
	slog.SetDefault(logger)

	// 2. config init
	err := config.InitConfig()
	if err != nil {
		slog.Error("config init failed", "err", err)
		os.Exit(1)
	}

	// 3. WeChat bot init
	if err = bot.Init(); err != nil {
		slog.Error("wechat init login failed", "err", err)
		os.Exit(1)
	}
	slog.Info("wechat login success!")
}

func main() {
	// 获取登陆的用户
	self, err := bot.Bot.GetCurrentUser()
	if err != nil {
		slog.Error("GetCurrentUser failed", "err", err)
		return
	}
	slog.Info("self info", "info", utils.MarshalAnyToString(self))

	bot.Bot.MessageHandler = MessageHandler // 微信消息回调注册
	bot.Bot.Block()

	// 示警
	notify.WxPushNotice("微信已退出")
}

func MessageHandler(msg *openwechat.Message) {
	// 目前暂且支持文字或者图片消息
	if !msg.IsText() && !msg.IsPicture() {
		slog.Warn("msg type not support", "msg_type", msg.MsgType)
		return
	}

	err := notify.WxPushNotice("你有新的消息：" + msg.Content)
	if err != nil {
		slog.Error("WxPushNotice failed", "err", err)
	} else {
		slog.Info("WxPushNotice success!")
	}
}
