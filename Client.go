package gorutracker

import (
	"net/http"
	"encoding/json"
	"github.com/ruslanfedoseenko/gorutracker/internalModels"
	"io/ioutil"
	"github.com/ruslanfedoseenko/gorutracker/Misc"
	"github.com/ruslanfedoseenko/gorutracker/Model"
	"strconv"
)

type Client struct {
	apiUrl string
	cache  *Misc.MemoryCache
}

func NewClient() Client {
	return Client{
		apiUrl: "http://api.rutracker.org/v1/",
		cache:  Misc.NewMemoryCache(),
	}
}

func (c *Client) GetBulkOperationsObjectsLimit() (limit int, err error) {
	if c.cache.Get("limit", &limit) == nil {
		return
	}
	var response *http.Response
	response, err = http.Get(c.apiUrl + "get_limit")
	if err != nil {
		return
	} else {
		defer response.Body.Close()
		var limitDto internalModels.LimitDto
		buffer, _ := ioutil.ReadAll(response.Body)
		err = json.Unmarshal(buffer, &limitDto)
		if err != nil {
			return
		}
		limit = limitDto.Result.Limit
		c.cache.Put("limit", &limit)
		return

	}

}

func (c *Client) GetStatuses() (statuses []Model.Status, err error) {
	if c.cache.Get("statuses", &statuses) == nil {
		return
	}
	var response *http.Response
	response, err = http.Get(c.apiUrl + "get_tor_status_titles")
	if err != nil {
		return
	} else {
		defer response.Body.Close()
		var statusesDto internalModels.KeyValueDto
		buffer, _ := ioutil.ReadAll(response.Body)
		err = json.Unmarshal(buffer, &statusesDto)
		if err != nil {
			return
		}
		for k, v := range statusesDto.Result {
			var tmp int
			tmp, err = strconv.Atoi(k)

			statuses = append(statuses, Model.Status{
				Id: tmp,
				Name: v,

			})
		}
		c.cache.Put("statuses", &statuses)
		return
	}
}
