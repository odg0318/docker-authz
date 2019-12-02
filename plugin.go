package main

import (
	"github.com/docker/go-plugins-helpers/authorization"
)

type Plugin struct {
}

func NewPlugin() (*Plugin, error) {
	return &Plugin{}, nil
}

func (p *Plugin) AuthZReq(req authorization.Request) authorization.Response {
	/*
		r := a.Allow(req)
		if r.Error != "" {
			return authorization.Response{Err: r.Error}
		}
		if !r.Allow {
			return authorization.Response{Msg: r.Msg["text"]}
		}
	*/

	return authorization.Response{Allow: true}
}

func (p *Plugin) AuthZRes(req authorization.Request) authorization.Response {
	return authorization.Response{Allow: true}
}
