// Package convert
package convert

import (
	"github.com/rodkranz/photoins/entity"
	"github.com/rodkranz/photoins/models"
)

func ToImage(model *models.Image) *entity.Image {
	myEntity := new(entity.Image)

	myEntity.ID = model.ID
	myEntity.InstagramID = model.InstagramID
	myEntity.DisplaySrc = model.DisplaySrc
	myEntity.ThumbnailSrc = model.ThumbnailSrc
	myEntity.IsVideo = model.IsVideo
	myEntity.Code = model.Code
	myEntity.Date = model.Date
	myEntity.Caption = model.Caption
	myEntity.TagID = model.TagID
	myEntity.TagName = model.TagName
	myEntity.Height = model.Height
	myEntity.Width = model.Width
	myEntity.Owner = model.Owner
	myEntity.Comments = model.Comments
	myEntity.Likes = model.Likes
	myEntity.Created = model.Created
	myEntity.Updated = model.Updated
	myEntity.IsNew = model.IsNew

	return myEntity
}

func ToImages(models []*models.Image) []*entity.Image {
	list := make([]*entity.Image, len(models))

	for i, model := range models {
		list[i] = ToImage(model)
	}

	return list
}
