package solr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type DocumentParameters struct {
	CommitWithin 	int 	`url:"commitWithin,omitempty"`
	Commit 			bool 	`url:"commit,omitempty"`
}

type Document struct {
	config *Config
}

func NewDocument(config *Config) Document {
	return Document{config: config}
}

func (d *Document) Select() (*Response, error) {
	return nil, nil
}

func (d *Document) Update(collection string,
	documents []map[string]interface{}, params DocumentParameters) (*Response, error) {

	payload, err := json.Marshal(documents)
	if err != nil {
		return nil, err
	}

	url, err := d.config.
		getUrlWithQueryStrings(
			fmt.Sprintf("/api/collections/%s/update/json", collection), params)
	if err != nil {
		return nil, err
	}

	resp, err := d.config.http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (d *Document) Commit(collection string) (*Response, error) {

	url, err := d.config.
		getUrlWithQueryStrings(
			fmt.Sprintf("/api/collections/%s/update/json", collection),
			DocumentParameters{Commit: true})
	if err != nil {
		return nil, err
	}

	resp, err := d.config.http.Post(url, "application/json", nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (d *Document) Delete(collection string) (*Response, error) {
	url, err := d.config.getUrl(fmt.Sprintf("/api/collections/%s/update/json", collection))
	if err != nil {
		return nil, err
	}

	resp, err := d.config.http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}