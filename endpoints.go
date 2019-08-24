package geverse

import (
	"fmt"
	"strings"
)

var (
	endpointCDNBaseURL    = "https://cdn-contents.weverse.io/"
	endpointBaseURL       = "https://weverseapi.weverse.io/"
	endpointStaticBaseURL = endpointBaseURL + "static/"
	endpointAPIBaseURL    = endpointBaseURL + "api/"

	endpointArtistPosts = func(communityID int64) string {
		return endpointAPIBaseURL + fmt.Sprintf("v1/stream/community/%d/artistTab", communityID)
	}
	endpointArtistMomentPosts = func(communityID, artistID int64) string {
		return endpointAPIBaseURL + fmt.Sprintf("v1/stream/community/%d/artist/%d/toFans", communityID, artistID)
	}
	endpointMediaPosts = func(communityID int64) string {
		return endpointAPIBaseURL + fmt.Sprintf("v1/stream/community/%d/mediaTab", communityID)
	}
	endpointNotices = func(communityID int64) string {
		return endpointStaticBaseURL + fmt.Sprintf("communities/%d/notices", communityID)
	}

	endpointMe = endpointAPIBaseURL + "v1/users/me"
)

func absolutify(path string) string {
	if path == "" {
		return " "
	}

	if strings.HasPrefix(path, "http") {
		return path
	}

	return endpointCDNBaseURL + strings.TrimLeft(path, "/")
}
