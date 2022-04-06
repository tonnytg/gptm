package projects

import (
	"time"
)

type MapProject struct {
	Projects []struct {
		ProjectNumber  string            `json:"projectNumber"`
		ProjectID      string            `json:"projectId"`
		LifecycleState string            `json:"lifecycleState"`
		Name           string            `json:"name"`
		Labels         map[string]string `json:"labels"`
		CreateTime     time.Time         `json:"createTime"`
	} `json:"projects"`
}
