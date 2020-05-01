package solr

import (
	"context"
	"fmt"
	"net/http"
)

type Parameters struct {
	CommitWithin 	int 						`url:"commitWithin,omitempty"`
	Commit 			bool 						`url:"commit,omitempty"`
	Query 			string 						`url:"q,omitempty"`
	Delete 			interface{} 				`url:"delete,omitempty"`
}

type Delete struct {
	Id 				string 						`json:"id,omitempty"`
	Query 			string 						`json:"query,omitempty"`
}

type Document map[string]interface{}

type DocumentAPI struct {
	client *Client
}

func (d *DocumentAPI) Select(ctx context.Context, collection string, query string) (*Response, error) {
	path := fmt.Sprintf("/solr/%s/select", collection)

	req, err := d.client.NewRequest(ctx, http.MethodGet, path, nil, &Parameters{
		Query:	query,
	})
	if err != nil {
		return nil, err
	}

	response, err := d.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, err
}

func (d *DocumentAPI) Update(ctx context.Context, collection string, docs []Document, params *Parameters) (*Response, error) {

	path := fmt.Sprintf("/api/collections/%s/update/json", collection)

	req, err := d.client.NewRequest(ctx, http.MethodPost, path, docs, params)
	if err != nil {
		return nil, err
	}

	response, err := d.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, err
}

func (d *DocumentAPI) Commit(ctx context.Context, collection string) (*Response, error) {

	path := fmt.Sprintf("/api/collections/%s/update/json", collection)

	req, err := d.client.NewRequest(ctx, http.MethodPost, path, nil, Parameters{
		Commit:       true,
	})
	if err != nil {
		return nil, err
	}

	response, err := d.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, err
}

func (d *DocumentAPI) Delete(ctx context.Context, collection string, delete Delete, params *Parameters) (*Response, error) {

	path := fmt.Sprintf("/solr/%s/update", collection)

	req, err := d.client.NewRequest(ctx, http.MethodPost, path, map[string]interface{}{
		"delete": delete,
	}, params)
	if err != nil {
		return nil, err
	}

	response, err := d.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}

	return response, err
}

//
//
//
//
//
//	payload, err := json.Marshal(documents)
//	if err != nil {
//		return nil, err
//	}
//
//	url, err := d.config.
//		getUrlWithQueryStrings(
//			fmt.Sprintf("/api/collections/%s/update/json", collection), params)
//	if err != nil {
//		return nil, err
//	}
//
//	resp, err := d.config.http.Post(url, "application/json", bytes.NewBuffer(payload))
//	if err != nil {
//		return nil, err
//	}
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//
//	var response Response
//	err = json.Unmarshal(body, &response)
//	if err != nil {
//		return nil, err
//	}
//
//	return &response, nil
//}
//
//func (d *DocumentAPI) Commit(collection string) (*Response, error) {
//
//	url, err := d.config.
//		getUrlWithQueryStrings(
//			fmt.Sprintf("/api/collections/%s/update/json", collection),
//			DocumentParameters{Commit: true})
//	if err != nil {
//		return nil, err
//	}
//
//	resp, err := d.config.http.Post(url, "application/json", nil)
//	if err != nil {
//		return nil, err
//	}
//
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//
//	var response Response
//	err = json.Unmarshal(body, &response)
//	if err != nil {
//		return nil, err
//	}
//
//	return &response, nil
//}
//
//func (d *DocumentAPI) Delete(collection string) (*Response, error) {
//	url, err := d.config.getUrl(fmt.Sprintf("/api/collections/%s/update/json", collection))
//	if err != nil {
//		return nil, err
//	}
//
//	resp, err := d.config.http.Get(url)
//	if err != nil {
//		return nil, err
//	}
//
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//
//	var response Response
//	err = json.Unmarshal(body, &response)
//	if err != nil {
//		return nil, err
//	}
//
//	return &response, nil
//}








//func (d *Document) Update(collection string,
//	documents []map[string]interface{}, params DocumentParameters) (*Response, error) {
//
//	payload, err := json.Marshal(documents)
//	if err != nil {
//		return nil, err
//	}
//
//	url, err := d.config.
//		getUrlWithQueryStrings(
//			fmt.Sprintf("/api/collections/%s/update/json", collection), params)
//	if err != nil {
//		return nil, err
//	}
//
//	resp, err := d.config.http.Post(url, "application/json", bytes.NewBuffer(payload))
//	if err != nil {
//		return nil, err
//	}
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//
//	var response Response
//	err = json.Unmarshal(body, &response)
//	if err != nil {
//		return nil, err
//	}
//
//	return &response, nil
//}
//
//func (d *Document) Commit(collection string) (*Response, error) {
//
//	url, err := d.config.
//		getUrlWithQueryStrings(
//			fmt.Sprintf("/api/collections/%s/update/json", collection),
//			DocumentParameters{Commit: true})
//	if err != nil {
//		return nil, err
//	}
//
//	resp, err := d.config.http.Post(url, "application/json", nil)
//	if err != nil {
//		return nil, err
//	}
//
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//
//	var response Response
//	err = json.Unmarshal(body, &response)
//	if err != nil {
//		return nil, err
//	}
//
//	return &response, nil
//}
//
//func (d *Document) Delete(collection string) (*Response, error) {
//	url, err := d.config.getUrl(fmt.Sprintf("/api/collections/%s/update/json", collection))
//	if err != nil {
//		return nil, err
//	}
//
//	resp, err := d.config.http.Get(url)
//	if err != nil {
//		return nil, err
//	}
//
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//
//	var response Response
//	err = json.Unmarshal(body, &response)
//	if err != nil {
//		return nil, err
//	}
//
//	return &response, nil
//}