package geverse

import (
	"time"
)

type ArtistPosts struct {
	Posts   []ArtistPost `json:"posts"`
	IsEnded bool         `json:"isEnded"`
}

func (ap *ArtistPosts) absolutify() {
	if ap == nil {
		return
	}

	for i := range ap.Posts {
		ap.Posts[i].absolutify()
	}
}

type ArtistPost struct {
	ID                int64         `json:"id"`
	CommunityUser     CommunityUser `json:"communityUser"`
	Community         Community     `json:"community"`
	CommunityTabID    int64         `json:"communityTabId"`
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

func (ap *ArtistPost) absolutify() {
	if ap == nil {
		return
	}

	ap.Community.absolutify()
	for i := range ap.Photos {
		ap.Photos[i].absolutify()
	}
}

type MediaPosts struct {
	Medias          []MediaPost     `json:"medias"`
	TotalCount      int             `json:"totalCount"`
	IsEnded         bool            `json:"isEnded"`
	LastID          int64           `json:"lastId"`
	MediaCategories []MediaCategory `json:"mediaCategories"`
}

func (mp *MediaPosts) absolutify() {
	if mp == nil {
		return
	}

	for i := range mp.Medias {
		mp.Medias[i].absolutify()
	}
	for i := range mp.MediaCategories {
		mp.MediaCategories[i].absolutify()
	}
}

type MediaPost struct {
	ID                 int64     `json:"id"`
	CommunityID        int64     `json:"communityId"`
	Body               string    `json:"body"`
	Type               string    `json:"type"`
	ThumbnailPath      string    `json:"thumbnailPath"`
	Photos             []Photo   `json:"photos"`
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
	YouTubeID          string    `json:"youtubeId,omitempty"`
}

func (mp *MediaPost) absolutify() {
	if mp == nil {
		return
	}

	mp.ThumbnailPath = absolutify(mp.ThumbnailPath)
	for i := range mp.Photos {
		mp.Photos[i].absolutify()
	}
	mp.Video.absolutify()
}

type MediaCategory struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (mc *MediaCategory) absolutify() {
}

type Notice struct {
	ID    int64
	Link  string
	Label string
	Title string
	Date  time.Time
}

func (n *Notice) absolutify() {
}

type Me struct {
	UserKey       string          `json:"userKey"`
	FirstName     string          `json:"firstName"`
	LastName      string          `json:"lastName"`
	Email         string          `json:"email"`
	MyCommunities []Community     `json:"myCommunities"`
	Profiles      []CommunityUser `json:"profiles"`
	SocialTypes   []interface{}   `json:"socialTypes"` // TODO: no example yet, did not want to connect my SNS accounts…
	LanguageCode  string          `json:"languageCode"`
}

func (m *Me) absolutify() {
	if m == nil {
		return
	}

	for i := range m.MyCommunities {
		m.MyCommunities[i].absolutify()
	}
	for i := range m.Profiles {
		m.Profiles[i].absolutify()
	}
}

type Community struct {
	ID                    int64    `json:"id"`
	Name                  string   `json:"name"`
	Description           string   `json:"description"`
	MemberCount           int      `json:"memberCount"`
	HomeBannerImgPath     string   `json:"homeBannerImgPath"`
	IconImgPath           string   `json:"iconImgPath"`
	BannerImgPath         string   `json:"bannerImgPath"`
	Fullname              []string `json:"fullname"`
	ShowMemberCount       bool     `json:"showMemberCount"`
	FcMember              bool     `json:"fcMember"`
	MembershipBenefitLink string   `json:"membershipBenefitLink"`
}

func (c *Community) absolutify() {
	if c == nil {
		return
	}

	c.HomeBannerImgPath = absolutify(c.HomeBannerImgPath)
	c.IconImgPath = absolutify(c.IconImgPath)
	c.BannerImgPath = absolutify(c.BannerImgPath)
}

type CommunityUser struct {
	ID              int64     `json:"id"`
	CommunityID     int64     `json:"communityId"`
	ArtistID        int64     `json:"artistId"`
	ProfileNickname string    `json:"profileNickname"`
	ProfileImgPath  string    `json:"profileImgPath"`
	Status          string    `json:"status"`
	Grade           string    `json:"grade"`
	Official        bool      `json:"official"`
	FcMember        bool      `json:"fcMember"`
	JoinedAt        time.Time `json:"joinedAt"`
}

func (cu *CommunityUser) absolutify() {
	if cu == nil {
		return
	}

	cu.ProfileImgPath = absolutify(cu.ProfileImgPath)
}

type Video struct {
	ID               int64         `json:"id"`
	ContentIndex     int           `json:"contentIndex"`
	Playtime         int           `json:"playtime"`
	ThumbnailImgPath string        `json:"thumbnailImgPath"`
	ThumbnailWidth   int           `json:"thumbnailWidth"`
	ThumbnailHeight  int           `json:"thumbnailHeight"`
	CommunityID      int64         `json:"communityId"`
	Status           string        `json:"status"`
	HlsPath          string        `json:"hlsPath"`
	DashPath         string        `json:"dashPath"`
	PlayreadyPath    string        `json:"playreadyPath"`
	VideoWidth       int           `json:"videoWidth"`
	VideoHeight      int           `json:"videoHeight"`
	Level            string        `json:"level"`
	IsVertical       bool          `json:"isVertical"`
	CaptionS3Paths   []interface{} `json:"captionS3Paths"` // TODO: example?, probably strings?
}

func (v *Video) absolutify() {
	if v == nil {
		return
	}

	v.ThumbnailImgPath = absolutify(v.ThumbnailImgPath)
	v.HlsPath = absolutify(v.HlsPath)
	v.DashPath = absolutify(v.DashPath)
	v.PlayreadyPath = absolutify(v.PlayreadyPath)
}

type Photo struct {
	ID                  int64  `json:"id"`
	ContentIndex        int    `json:"contentIndex"`
	ThumbnailImgURL     string `json:"thumbnailImgUrl"`
	ThumbnailImgWidth   int    `json:"thumbnailImgWidth"`
	ThumbnailImgHeight  int    `json:"thumbnailImgHeight"`
	OrgImgURL           string `json:"orgImgUrl"`
	OrgImgWidth         int    `json:"orgImgWidth"`
	OrgImgHeight        int    `json:"orgImgHeight"`
	DownloadImgFilename string `json:"downloadImgFilename"`
}

func (p *Photo) absolutify() {
}
