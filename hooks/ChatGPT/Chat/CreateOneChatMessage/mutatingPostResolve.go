package CreateOneChatMessage

import (
	"custom-go/generated"
	"custom-go/pkg/base"
)

func MutatingPostResolve(hook *base.HookRequest, body generated.ChatGPT__Chat__CreateOneChatMessageBody) (res generated.ChatGPT__Chat__CreateOneChatMessageBody, err error) {
	hook.Logger().Info("MutatingPostResolve")
	if body.Response != nil {
		body.Response.Data.Data.Sqlite_createOneUntitled1.CreatedAt = ""
	}
	return body, nil
}
