// Package models
package models

import (
	"github.com/go-xorm/xorm"
	"strings"
	"time"
)

type Tag struct {
	ID        int64  `xorm:"pk autoincr"`
	LowerName string `xorm:"VARCHAR(255)"`
	Name      string `xorm:"VARCHAR(255)"`
	Country   string `xorm:"TEXT"`

	Count int64 `xorm:"INT"`

	Created     time.Time `xorm:"-"`
	CreatedUnix int64

	Updated     time.Time `xorm:"-"`
	UpdatedUnix int64
}

func (t *Tag) BeforeInsert() {
	t.CreatedUnix = time.Now().UTC().Unix()
}

func (t *Tag) BeforeUpdate() {
	t.UpdatedUnix = time.Now().UTC().Unix()
}

func (t *Tag) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "name":
		t.LowerName = strings.ToLower(t.Name)
	case "created_unix":
		t.Created = time.Unix(t.CreatedUnix, 0).Local()
	case "updated_unix":
		t.Updated = time.Unix(t.UpdatedUnix, 0).Local()
	}
}

func CreateTag(t *Tag) (err error) {
	sess := x.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return err
	}

	t.LowerName = strings.ToLower(t.Name)
	if _, err = sess.Insert(t); err != nil {
		sess.Rollback()
		return err
	}

	return sess.Commit()
}

func UpdateTag(t *Tag) error {
	t.LowerName = strings.ToLower(t.Name)

	_, err := x.Id(t.ID).AllCols().Update(t)
	return err
}

func GetTagByID(id int64) (*Tag, error) {
	t := new(Tag)
	has, err := x.Id(id).Get(t)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrTagNotExist{id, ""}
	}
	return t, nil
}

func GetTagByName(name string) (*Tag, error) {
	t := new(Tag)
	has, err := x.Where("name = ?", strings.ToLower(name)).Get(t)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, ErrTagNotExist{0, name}
	}
	return t, nil
}

func IsTagExist(id int64, name string) (bool, error) {
	if len(name) == 0 {
		return false, nil
	}

	return x.Where("id!=?", id).Get(&Tag{LowerName: strings.ToLower(name)})
}
