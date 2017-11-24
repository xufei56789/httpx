package httpx

import "net/http"

func (httpx *Httpx) SetHeader(h *http.Header) {
	httpx.header = h
}