// Package entity
package entity

import "time"

type Image struct {
	ID           int64     `json:"id"`
	InstagramID  string    `json:"instagram_id"`
	DisplaySrc   string    `json:"display_src"`
	ThumbnailSrc string    `json:"thumbnail_src"`
	IsVideo      bool      `json:"is_video"`
	Code         string    `json:"code"`
	Date         time.Time `json:"date"`
	Caption      string    `json:"caption"`
	TagID        int64     `json:"tag_id"`
	TagName      string    `json:"tag_name"`
	Height       int64     `json:"height"`
	Width        int64     `json:"width"`
	Owner        string    `json:"owner"`
	Comments     int64     `json:"comments"`
	Likes        int64     `json:"likes"`
	Created      time.Time `json:"created"`
	Updated      time.Time `json:"updated"`
	IsNew        bool      `json:"is_new"`
}
