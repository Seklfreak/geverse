package geverse

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (g *Geverse) GetCommunity(ctx context.Context, communityID int64) (*Community, error) {
	req, err := g.newRequest(ctx, http.MethodGet, endpointCommunity(communityID), nil)
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

	var community *Community
	err = json.Unmarshal(data, &community)

	community.absolutify()

	return community, nil
}
