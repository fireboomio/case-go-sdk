package server

import (
    "custom-go/global"
    
    "custom-go/auth"
    
            
    "custom-go/generated"
        
    hooks_ChatGPT__Chat__CreateOneChatMessage "custom-go/hooks/ChatGPT/Chat/CreateOneChatMessage"
    
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
	"custom-go/pkg/types"
    
    "custom-go/customize"
)

func init() {
    types.WdgHooksAndServerConfig = types.WunderGraphHooksAndServerConfig{
    	Hooks: types.HooksConfiguration{
            Global: plugins.GlobalConfiguration{
                HttpTransport: plugins.HttpTransportHooks{
                    OnOriginRequest:  global.OnOriginRequest,
                    OnOriginResponse: global.OnOriginResponse,
                },
                WsTransport: plugins.WsTransportHooks{
                    OnConnectionInit: nil,
                },
            },
                                            
            Authentication: plugins.AuthenticationConfiguration{
                Revalidate:                 auth.Revalidate,
                MutatingPostAuthentication: auth.MutatingPostAuthentication,
                PostAuthentication:         auth.PostAuthentication,
                PostLogout:                 auth.PostLogout,
            },
                                            
            Queries: base.OperationHooks{
            },
                                
            Mutations: base.OperationHooks{
                "ChatGPT/Chat/CreateOneChatMessage": {
                    MutatingPostResolve: plugins.ConvertBodyFunc[generated.ChatGPT__Chat__CreateOneChatMessageInput, generated.ChatGPT__Chat__CreateOneChatMessageResponse](hooks_ChatGPT__Chat__CreateOneChatMessage.MutatingPostResolve),
                    MutatingPreResolve:  plugins.ConvertBodyFunc[generated.ChatGPT__Chat__CreateOneChatMessageInput, generated.ChatGPT__Chat__CreateOneChatMessageResponse](hooks_ChatGPT__Chat__CreateOneChatMessage.MutatingPreResolve),
                    PostResolve:         plugins.ConvertBodyFunc[generated.ChatGPT__Chat__CreateOneChatMessageInput, generated.ChatGPT__Chat__CreateOneChatMessageResponse](hooks_ChatGPT__Chat__CreateOneChatMessage.PostResolve),
                },
            },
        
    		Uploads: map[string]plugins.UploadHooks{
                "T": {

                },
    		},
    	},
    	GraphqlServers: []plugins.GraphQLServerConfig{
                {
                ApiNamespace:          "chatGPT",
                ServerName:            "chatGPT",
                EnableGraphQLEndpoint: true,
                Schema:                customize.ChatGPT_schema,
            },                        
    	},
    }
}