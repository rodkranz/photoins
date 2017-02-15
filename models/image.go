// Package models
package models

import (
	"time"

	"fmt"
	"github.com/go-xorm/xorm"
	"strings"
)

type Image struct {
	ID          int64  `xorm:"pk autoincr"`
	InstagramID string `xorm:"VARCHAR(255) INDEX NOT NULL"`

	DisplaySrc   string `xorm:"VARCHAR(255)"`
	ThumbnailSrc string `xorm:"VARCHAR(255)"`

	IsVideo bool   `xorm:"BOOL"`
	Code    string `xorm:"VARCHAR(100)"`

	Date     time.Time `xorm:"-"`
	DateUnix int64

	Caption string `xorm:"TEXT"`

	TagID   int64  `xorm:"INDEX NOT NULL"`
	TagName string `xorm:"VARCHAR(255) NOT NULL"`

	Height int64 `xorm:"INT"`
	Width  int64 `xorm:"INT"`

	Owner    string `xorm:"VARCHAR(255)"`
	Comments int64  `xorm:"INT"`
	Likes    int64  `xorm:"INT"`

	IsNew bool `xorm:"-"`

	Created     time.Time `xorm:"-"`
	CreatedUnix int64

	Updated     time.Time `xorm:"-"`
	UpdatedUnix int64
}

func (i *Image) BeforeInsert() {
	i.CreatedUnix = time.Now().UTC().Unix()
}

func (i *Image) BeforeUpdate() {
	i.UpdatedUnix = time.Now().UTC().Unix()
}

func (i *Image) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "created_unix":
		i.Created = time.Unix(i.CreatedUnix, 0).Local()
	case "updated_unix":
		i.Updated = time.Unix(i.UpdatedUnix, 0).Local()
	}
}

func CreateImage(i *Image) (err error) {
	sess := x.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return err
	}

	if _, err = sess.Insert(i); err != nil {
		sess.Rollback()
		return err
	}

	return sess.Commit()
}

func UpdateImage(i *Image) error {
	_, err := x.Id(i.ID).AllCols().Update(i)
	return err
}

func GetImageByID(id int64) (*Image, error) {
	t := new(Image)
	has, err := x.Id(id).Get(t)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrImageNotExist{id, ""}
	}
	return t, nil
}

func GetImageByInstagramID(id string) (*Image, error) {
	t := new(Image)
	has, err := x.Where("instagram_id = ?", id).Get(t)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrImageNotExist{0, id}
	}
	return t, nil
}

func IsImageExist(uid int64, InstagramID string) (bool, error) {
	return x.Where("id!=?", uid).Get(&Image{InstagramID: InstagramID})
}

func CountImages() int64 {
	count, _ := x.Count(new(Image))
	return count
}

func Images(page, pageSize int) ([]*Image, error) {
	images := make([]*Image, 0, pageSize)
	return images, x.Limit(pageSize, (page-1)*pageSize).Desc("id").Find(&images)
}

type SearchImageOptions struct {
	Keyword  string
	Tag      string
	Page     int
	PageSize int
	OrderBy  string
}

func SearchImage(opts *SearchImageOptions) (images []*Image, _ int64, _ error) {
	if opts.PageSize <= 0 {
		opts.PageSize = 10
	}

	if opts.Page <= 0 {
		opts.Page = 1
	}

	if len(opts.Keyword) == 0 && len(opts.Tag) == 0 {
		images, err := Images(opts.Page, opts.PageSize)
		return images, CountImages(), err
	}

	opts.Keyword = strings.ToLower(opts.Keyword)

	searchQuery := "%" + opts.Keyword + "%"
	images = make([]*Image, 0, opts.PageSize)

	// Append conditions
	sess := x.NewSession()

	if len(opts.Keyword) != 0 {
		sess.Where("LOWER(caption) LIKE ?", searchQuery)
	}

	if len(opts.Tag) != 0 {
		sess.Where("LOWER(tag_name) = ?", strings.ToLower(opts.Tag))
	}

	var countSess xorm.Session
	countSess = *sess
	count, err := countSess.Count(new(Image))
	if err != nil {
		return nil, 0, fmt.Errorf("Count: %v", err)
	}

	if len(opts.OrderBy) > 0 {
		sess.OrderBy(opts.OrderBy)
	} else {
		sess.OrderBy("instagram_id DESC")
	}

	return images, count, sess.Limit(opts.PageSize, (opts.Page-1)*opts.PageSize).Find(&images)
}
