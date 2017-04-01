package gorutracker

import (
	"net/http"
	"encoding/json"
	"github.com/ruslanfedoseenko/gorutracker/internalModels"
	"io/ioutil"
	"github.com/ruslanfedoseenko/gorutracker/Misc"
)

type Client struct {
	apiUrl string
	cache *Misc.MemoryCache
}

func NewClient() Client {
	return Client{
		apiUrl: "http://api.rutracker.org/v1/",
		cache:  Misc.NewMemoryCache(),
	}
}

func (c *Client) GetBulkOperationsObjectsLimit() (limit int, err error) {
	if c.cache.Get("limit", limit) == nil {
		return
	}
	var response *http.Response
	response, err = http.Get(c.apiUrl + "get_limit")
	if err != nil{
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
		c.cache.Put("limit", limit)
		return

	}

}


