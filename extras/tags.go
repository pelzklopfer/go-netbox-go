package extras

import (
	"encoding/json"
	"fmt"
	"github.com/sapcc/go-netbox-go/models"
	"io/ioutil"
	"net/http"
)

func (c *Client) ListTags(opts models.ListTagsRequest) (*models.ListTagsResponse, error) {
	request, err := http.NewRequest("GET", c.BaseUrl.String() + basePath + "tags/", nil)
	if err != nil {
		return nil, err
	}
	c.SetAuthToken(&request.Header)
	setListTagsParams(request, opts)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected return code of %d", response.StatusCode)
	}
	resObj := models.ListTagsResponse{}
	byteses, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteses, &resObj)
	if err != nil {
		return nil, err
	}
	return &resObj, nil
}


/*

Permissions

func (c *Client) CreateTag(tag models.Tag) error {
	body, err := json.Marshal(tag)
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", c.BaseUrl.String() + basePath + "tags/", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	c.SetAuthToken(&request.Header)
	response, err := c.HttpClient.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode != 201 {
		return fmt.Errorf("unexpected response code of %d", response.StatusCode)
	}
	return nil
}

 */

func setListTagsParams(req *http.Request, opts models.ListTagsRequest) {
	q := req.URL.Query()
	opts.SetListParams(&q)
	req.URL.RawQuery = q.Encode()
}