package geverse

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (g *Geverse) GetMe(ctx context.Context) (*Me, error) {
	req, err := g.newRequest(ctx, http.MethodGet, endpointMe, nil)
	if err != nil {
		return nil, err
	}

	resp, err := g.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received unexpected status code: expected 200, received %d", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var me *Me
	err = json.Unmarshal(data, &me)

	return me, nil
}
