package auth

import (
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
)

func MutatingPostAuthentication(hook *base.AuthenticationHookRequest) (*plugins.AuthenticationResponse, error) {
	return nil, nil
}
