package geverse

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func (g *Geverse) GetArtistPosts(ctx context.Context, communityID int64) (*ArtistPosts, error) {
	endpoint := endpointArtistPosts(communityID)

	req, err := g.newRequest(ctx, http.MethodGet, endpoint, nil)
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

	return posts, nil
}

type ArtistPosts struct {
	Posts   []ArtistPost `json:"posts"`
	IsEnded bool         `json:"isEnded"`
}

type ArtistPost struct {
	ID                int64         `json:"id"`
	CommunityUser     CommunityUser `json:"communityUser"`
	Community         Community     `json:"community"`
	CommunityTabID    int           `json:"communityTabId"`
	Type              string        `json:"type"`
	Body              string        `json:"body"`
	CommentCount      int           `json:"commentCount"`
	LikeCount         int           `json:"likeCount"`
	MaxCommentCount   int64         `json:"maxCommentCount"`
	HasMyLike         bool          `json:"hasMyLike"`
	CreatedAt         time.Time     `json:"createdAt"`
	UpdatedAt         time.Time     `json:"updatedAt"`
	IsLocked          bool          `json:"isLocked"`
	IsBlind           bool          `json:"isBlind"`
	IsActive          bool          `json:"isActive"`
	IsPrivate         bool          `json:"isPrivate"`
	Photos            []Photo       `json:"photos,omitempty"`
	IsHotTrendingPost bool          `json:"isHotTrendingPost"`
}

type CommunityUser struct {
	ID              int       `json:"id"`
	CommunityID     int       `json:"communityId"`
	ArtistID        int       `json:"artistId"`
	ProfileNickname string    `json:"profileNickname"`
	ProfileImgPath  string    `json:"profileImgPath"`
	Status          string    `json:"status"`
	Official        bool      `json:"official"`
	FcMember        bool      `json:"fcMember"`
	JoinedAt        time.Time `json:"joinedAt"`
}

type Community struct {
	ID                int      `json:"id"`
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	MemberCount       int      `json:"memberCount"`
	HomeBannerImgPath string   `json:"homeBannerImgPath"`
	IconImgPath       string   `json:"iconImgPath"`
	BannerImgPath     string   `json:"bannerImgPath"`
	Fullname          []string `json:"fullname"`
	ShowMemberCount   bool     `json:"showMemberCount"`
	FcMember          bool     `json:"fcMember"`
}

type Photo struct {
	ID                  int    `json:"id"`
	ContentIndex        int    `json:"contentIndex"`
	ThumbnailImgURL     string `json:"thumbnailImgUrl"`
	ThumbnailImgWidth   int    `json:"thumbnailImgWidth"`
	ThumbnailImgHeight  int    `json:"thumbnailImgHeight"`
	OrgImgURL           string `json:"orgImgUrl"`
	OrgImgWidth         int    `json:"orgImgWidth"`
	OrgImgHeight        int    `json:"orgImgHeight"`
	DownloadImgFilename string `json:"downloadImgFilename"`
}
