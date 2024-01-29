package controllers

import "net/http"

type IPingPongController interface {
	Pong(w http.ResponseWriter, r *http.Request)
}

type PingPongController struct{}

func (*PingPongController) Pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func NewPingPongController() *PingPongController {
	return &PingPongController{}
}
