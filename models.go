package geverse

import (
	"time"
)

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

type MediaPosts struct {
	Medias          []MediaPost     `json:"medias"`
	TotalCount      int             `json:"totalCount"`
	IsEnded         bool            `json:"isEnded"`
	LastID          int             `json:"lastId"`
	MediaCategories []MediaCategory `json:"mediaCategories"`
}

type MediaPost struct {
	ID                 int       `json:"id"`
	CommunityID        int       `json:"communityId"`
	Body               string    `json:"body"`
	Type               string    `json:"type"`
	ThumbnailPath      string    `json:"thumbnailPath"`
	Photos             Photo     `json:"photos"`
	Video              Video     `json:"video,omitempty"`
	Title              string    `json:"title"`
	Level              string    `json:"level"`
	TotalPlaytime      int       `json:"totalPlaytime"`
	IsExtVideoVertical bool      `json:"isExtVideoVertical"`
	ViewCount          int       `json:"viewCount"`
	LikeCount          int       `json:"likeCount"`
	PlayCount          int       `json:"playCount"`
	HasMyLike          bool      `json:"hasMyLike"`
	CreatedAt          time.Time `json:"createdAt"`
	ExtVideoPath       string    `json:"extVideoPath,omitempty"`
	YoutubeID          string    `json:"youtubeId,omitempty"`
}

type MediaCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
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

type Video struct {
	ID               int           `json:"id"`
	ContentIndex     int           `json:"contentIndex"`
	Playtime         int           `json:"playtime"`
	ThumbnailImgPath string        `json:"thumbnailImgPath"`
	ThumbnailWidth   int           `json:"thumbnailWidth"`
	ThumbnailHeight  int           `json:"thumbnailHeight"`
	CommunityID      int           `json:"communityId"`
	Status           string        `json:"status"`
	HlsPath          string        `json:"hlsPath"`
	DashPath         string        `json:"dashPath"`
	PlayreadyPath    string        `json:"playreadyPath"`
	VideoWidth       int           `json:"videoWidth"`
	VideoHeight      int           `json:"videoHeight"`
	Level            string        `json:"level"`
	IsVertical       bool          `json:"isVertical"`
	CaptionS3Paths   []interface{} `json:"captionS3Paths"`
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
