package Models

import "github.com/ruslanfedoseenko/go-rutracker/Misc"

type forumTree struct {
	Categories 	 map[string]string    		`json:"c"`
	Forums 		 map[string]string    		`json:"f"`
	Tree		 map[string]map[string][]int	`json:"tree"`
}

type forumTreeResponse struct {
	Format 		 Misc.JSONString	`json:"format"`
	UpdateUnixTime 	 int			`json:"update_time"`
	UpdateTimeString string 		`json:"update_time_humn"`
	Result forumTree			`json:"result"`
}
