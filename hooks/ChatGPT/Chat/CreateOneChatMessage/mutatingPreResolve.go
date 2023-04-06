package CreateOneChatMessage

import (
	"custom-go/generated"
	"custom-go/pkg/base"
)

func MutatingPreResolve(hook *base.HookRequest, body generated.ChatGPT__Chat__CreateOneChatMessageBody) (res generated.ChatGPT__Chat__CreateOneChatMessageBody, err error) {
	hook.Logger().Info("MutatingPreResolve")
	//body.Input.Title = "123123"
	return body, nil
}
