package init

import "go-scheduler/pkg/web/conf"

type (
	Controller struct {
	}
	PostRequest struct {
		Etcd     conf.Etcd     `json:"etcd"`
		Database conf.Database `json:"database"`
		User     conf.User     `json:"user"`
		Auth     conf.Auth     `json:"auth"`
	}
)
