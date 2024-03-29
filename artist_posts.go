package geverse

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (g *Geverse) GetArtistPosts(ctx context.Context, communityID int64) (*ArtistPosts, error) {
	req, err := g.newRequest(ctx, http.MethodGet, endpointArtistPosts(communityID), nil)
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

	var posts *ArtistPosts
	err = json.Unmarshal(data, &posts)

	posts.absolutify()

	return posts, nil
}

func (g *Geverse) GetArtistMoments(ctx context.Context, communityID, artistID int64) (*ArtistPosts, error) {
	req, err := g.newRequest(ctx, http.MethodGet, endpointArtistMomentPosts(communityID, artistID), nil)
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

	var posts *ArtistPosts
	err = json.Unmarshal(data, &posts)

	posts.absolutify()

	return posts, nil
}
