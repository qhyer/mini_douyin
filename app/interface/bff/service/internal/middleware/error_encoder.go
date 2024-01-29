package middleware

import (
	nethttp "net/http"

	"github.com/go-kratos/kratos/v2/transport/http"

	"douyin/common/ecode"
)

func ErrorEncoder(w nethttp.ResponseWriter, r *nethttp.Request, err error) {
	se := ecode.ConvertErr(err)
	codec, _ := http.CodecForRequest(r, "Accept")
	b, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(nethttp.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/"+codec.Name())
	w.Write(b)
}
