package NewProfile1

import (
	"custom-go/pkg/base"
	"custom-go/pkg/plugins"
)

func PreUpload(request *base.PreUploadHookRequest, body *plugins.UploadBody[any]) (*base.UploadHookResponse, error) {
	return &base.UploadHookResponse{FileKey: body.File.Name}, nil
}
