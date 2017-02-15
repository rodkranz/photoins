// Package instagram
package instagram

import "encoding/json"

type Data struct {
	CountryCode  string    `json:"country_code"`
	LanguageCode string    `json:"language_code"`
	Platform     string    `json:"platform"`
	Hostname     string    `json:"hostname"`
	Data         EntryData `json:"entry_data"`
	Config       Config    `json:"config"`
}

func (d *Data) Parser(b []byte) error {
	return json.Unmarshal(b, d)
}

func (d *Data) String() string {
	data, err := json.Marshal(d)
	if err != nil {
		return err.Error()
	}

	return string(data)
}

type Config struct {
	CsrfToken string `json:"csrf_token"`
	Viewer    struct {
		ExternalUrl     string `json:"external_url"`
		ProfilePicUrl   string `json:"profile_pic_url"`
		FullName        string `json:"full_name"`
		Username        string `json:"username"`
		Id              int64  `json:"id"`
		HasProfilePic   bool   `json:"has_profile_pic"`
		ProfilePicUrlHd string `json:"profile_pic_url_hd"`
	} `json:"viewer"`
}

type EntryData struct {
	TagPages []TagPage `json:"TagPage"`
}

type TagPage struct {
	Tags Tag `json:"tag"`
}

type Tag struct {
	Name    string   `json:"name"`
	Media   Media    `json:"media"`
	TopPost TopPosts `json:"top_posts"`
}

type TopPosts struct {
	Nodes []Node `json:"nodes"`
}

type Media struct {
	Nodes    []Node `json:"nodes"`
	Count    int64  `json:"count"`
	PageInfo struct {
		HasNextPage bool   `json:"has_next_page"`
		EndCursor   string `json:"end_cursor"`
	} `json:"page_info"`
}

type Node struct {
	Id               string `json:"id"`
	ThumbnailSrc     string `json:"thumbnail_src"`
	IsVideo          bool   `json:"is_video"`
	Code             string `json:"code"`
	Date             int64  `json:"date"`
	DisplaySrc       string `json:"display_src"`
	Caption          string `json:"caption"`
	CommentsDisabled bool   `json:"comments_disabled"`

	Dimensions struct {
		Height int64 `json:"height"`
		Width  int64 `json:"width"`
	} `json:"dimensions"`

	Owner struct {
		Id string `json:"id"`
	} `json:"owner"`

	Comments struct {
		Count int64 `json:"count"`
	} `json:"comments"`

	Likes struct {
		Count int64 `json:"count"`
	} `json:"likes"`
}
