package internalModels

import "github.com/ruslanfedoseenko/gorutracker/Misc"

type ForumTree struct {
	Categories map[string]string                `json:"c"`
	Forums     map[string]string                `json:"f"`
	Tree       map[string]map[string][]int      `json:"tree"`
}

type ForumTreeDto struct {
	Format           Misc.JSONString        `json:"format"`
	UpdateUnixTime   int                    `json:"update_time"`
	UpdateTimeString string                 `json:"update_time_humn"`
	Result           ForumTree              `json:"result"`
}
