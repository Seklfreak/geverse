package geverse

import (
	"fmt"
)

var (
	endpointBaseURL       = "https://weverseapi.weverse.io/"
	endpointStaticBaseURL = endpointBaseURL + "static/"
	endpointAPIBaseURL    = endpointBaseURL + "api/"

	endpointArtistPosts = func(communityID int64) string {
		return endpointAPIBaseURL + fmt.Sprintf("v1/stream/community/%d/artistTab", communityID)
	}
	endpointMediaPosts = func(communityID int64) string {
		return endpointAPIBaseURL + fmt.Sprintf("v1/stream/community/%d/mediaTab", communityID)
	}
	endpointNotices = func(communityID int64) string {
		return endpointStaticBaseURL + fmt.Sprintf("communities/%d/notices", communityID)
	}

	endpointMe = endpointAPIBaseURL + "v1/users/me"
)
