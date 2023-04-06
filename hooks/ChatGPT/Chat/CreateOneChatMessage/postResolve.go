package CreateOneChatMessage

import (
	"custom-go/generated"
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
)

func PostResolve(hook *base.HookRequest, body generated.ChatGPT__Chat__CreateOneChatMessageBody) (res generated.ChatGPT__Chat__CreateOneChatMessageBody, err error) {
	hook.Logger().Info("postResolve")
	type (
		I  = generated.Sqlite__TestInput
		OD = generated.Sqlite__TestResponseData
	)
	input := I{Title: "123"}
	data, err := plugins.ExecuteInternalRequestQueries[I, OD](hook.InternalClient, generated.Sqlite__Test, input)
	hook.Logger().Info(data)
	return body, nil
}
