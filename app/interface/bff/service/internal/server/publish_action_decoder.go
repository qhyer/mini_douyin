package server

import (
	"bytes"
	"io"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"

	v1 "douyin/api/bff"
)

func PublishActionDecoder(r *http.Request, v interface{}) error {
	if (r.URL.Path == "/douyin/publish/action" || r.URL.Path == "/douyin/publish/action/") && strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
		paRequest := v.(*v1.PublishActionRequest)
		var maxMemory int64 = 32 << 20
		err := r.ParseMultipartForm(maxMemory)
		if err != nil {
			return errors.BadRequest("CODEC", err.Error())
		}
		paRequest.Title = r.FormValue("title")
		paRequest.Token = r.FormValue("token")
		file, _, err := r.FormFile("data")
		if err != nil {
			return errors.BadRequest("CODEC", err.Error())
		}
		var buf bytes.Buffer
		_, err = io.Copy(&buf, file)
		if err != nil {
			return errors.BadRequest("CODEC", err.Error())
		}
		paRequest.Data = buf.Bytes()
		log.Info("app upload success")
		return nil
	}
	return nil
}
