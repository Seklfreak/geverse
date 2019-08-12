package geverse

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (g *Geverse) GetMediaPosts(ctx context.Context, communityID int64) (*MediaPosts, error) {
	req, err := g.newRequest(ctx, http.MethodGet, endpointMediaPosts(communityID), nil)
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

	var posts *MediaPosts
	err = json.Unmarshal(data, &posts)

	return posts, nil
}
