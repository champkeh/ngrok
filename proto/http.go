package proto

import (
	"net/http"
	"time"

	"github.com/champkeh/ngrok/util"
)

type HttpRequest struct {
	*http.Request
	BodyBytes []byte
}

type HttpResponse struct {
	*http.Response
	BodyBytes []byte
}

type HttpTxn struct {
	Req         *HttpRequest
	Resp        *HttpResponse
	Start       time.Time
	Duration    time.Duration
	UserCtx     interface{}
	ConnUserCtx interface{}
}

type Http struct {
	Txns *util.Broadcast
}
