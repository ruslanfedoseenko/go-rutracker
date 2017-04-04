package gorutracker

import (
	"net/http"
	"encoding/json"
	"github.com/ruslanfedoseenko/gorutracker/internalModels"
	"io/ioutil"
	"github.com/ruslanfedoseenko/gorutracker/Misc"
	"strconv"
	"net/url"
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

func (c *Client) GetStatuses() (statuses map[int]string, err error) {
	statuses = make(map[int]string)
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
			if err != nil {
				return nil,err
			}
			statuses[tmp] = v
		}
		c.cache.Put("statuses", &statuses)
		return
	}
}


func (c *Client) GetForumNames(ids []int) (forumNames map[int]string,err error){
	forumNames = make(map[int]string)
	var response *http.Response
	var requestUrl *url.URL
	requestUrl, err = requestUrl.Parse(c.apiUrl + "get_forum_name")
	query := requestUrl.Query()
	query.Add("by", "forum_id")
	query.Add("val", Misc.ArrayToString(Misc.ToInterfaceSlice(ids)))
	requestUrl.RawQuery = query.Encode()
	response, err = http.Get(requestUrl.String())

	if err == nil {
		defer response.Body.Close()
		var forumNamesDto internalModels.KeyValueDto
		buffer, _ := ioutil.ReadAll(response.Body)
		err = json.Unmarshal(buffer, &forumNamesDto)
		if err != nil {
			return
		}
		for k, v := range forumNamesDto.Result {
			var tmp int
			tmp, err = strconv.Atoi(k)
			if err != nil {
				return nil,err
			}
			forumNames[tmp] = v
		}
		return
	}

	return
}


func (c *Client) GetForumData(ids []int) (forumData map[int]internalModels.ForumData,err error) {
	forumData = make(map[int]internalModels.ForumData)
	var response *http.Response
	var requestUrl *url.URL
	requestUrl, err = requestUrl.Parse(c.apiUrl + "get_forum_data")
	query := requestUrl.Query()
	query.Add("by", "forum_id")
	query.Add("val", Misc.ArrayToString(Misc.ToInterfaceSlice(ids)))
	requestUrl.RawQuery = query.Encode()
	response, err = http.Get(requestUrl.String())

	if err == nil {
		defer response.Body.Close()
		var forumDatasDto internalModels.ForumDataDto
		buffer, _ := ioutil.ReadAll(response.Body)
		err = json.Unmarshal(buffer, &forumDatasDto)
		if err != nil {
			return
		}
		for k, v := range forumDatasDto.Result {
			var tmp int
			tmp, err = strconv.Atoi(k)
			if err != nil {
				return nil,err
			}
			forumData[tmp] = v
		}
		return
	}
	return
}

func (c *Client) GetUserName(ids []int) (userNames map[int]string,err error) {
	userNames = make(map[int]string)
	var response *http.Response
	var requestUrl *url.URL
	requestUrl, err = requestUrl.Parse(c.apiUrl + "get_user_name")
	query := requestUrl.Query()
	query.Add("by", "user_id")
	query.Add("val", Misc.ArrayToString(Misc.ToInterfaceSlice(ids)))
	requestUrl.RawQuery = query.Encode()
	response, err = http.Get(requestUrl.String())

	if err == nil {
		defer response.Body.Close()
		var forumNamesDto internalModels.KeyValueDto
		buffer, _ := ioutil.ReadAll(response.Body)
		err = json.Unmarshal(buffer, &forumNamesDto)
		if err != nil {
			return
		}
		for k, v := range forumNamesDto.Result {
			var tmp int
			tmp, err = strconv.Atoi(k)
			if err != nil {
				return nil,err
			}
			userNames[tmp] = v
		}
		return
	}
	return
}

