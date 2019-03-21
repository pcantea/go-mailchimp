package mailchimp

import (
	"encoding/json"
	"io/ioutil"
)

type BatchOperation struct {
	Method string      `json:"method"`
	Path   string      `json:"path"`
	Body   interface{} `json:"body"`
}

type Batch struct {
	Operations []BatchOperation
}

func CreateBatch() Batch {
	return Batch{
		Operations: make([]BatchOperation, 0),
	}
}

func (b *Batch) AddOperation(o BatchOperation) {
	b.Operations = append(b.Operations, o)
}

func (c *Client) CreateBatch(batch *Batch) (*BatchResponse, error) {
	operations, err := json.Marshal(batch)
	if err != nil {
		return nil, err
	}
	data := make(map[string]interface{})
	data["operations"] = operations

	resp, err := c.do(
		"POST",
		"/batches",
		&data,
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode/100 == 2 {
		batchResponse := new(BatchResponse)
		if err := json.Unmarshal(responseBody, batchResponse); err != nil {
			return nil, err
		}
		return batchResponse, nil
	}

	errorResponse, err := extractError(responseBody)
	if err != nil {
		return nil, err
	}
	return nil, errorResponse
}
