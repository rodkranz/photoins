// Package instagram
package instagram

import (
	"github.com/rodkranz/photoins/models"
)

type ImportResult struct {
	NewImages     []*models.Image `json:"new_images"`
	NewImagesErr  []error         `json:"new_images_err"`
	EditImages    []*models.Image `json:"edit_images"`
	EditImagesErr []error         `json:"edit_images_err"`
	Err           error           `json:"err"`
}

func (ir *ImportResult) NewImage(image *models.Image) {
	ir.NewImages = append(ir.NewImages, image)
}

func (ir *ImportResult) NewImageError(err error) {
	ir.NewImagesErr = append(ir.NewImagesErr, err)
}

func (ir *ImportResult) EditImage(image *models.Image) {
	ir.EditImages = append(ir.EditImages, image)
}

func (ir *ImportResult) EditImageError(err error) {
	ir.EditImagesErr = append(ir.EditImagesErr, err)
}

func ImportTag(tag string) (ir *ImportResult) {
	ir = new(ImportResult)
	b, err := FetchFromTag(tag)
	if err != nil {
		ir.Err = err
		return ir
	}

	b, err = ParseJsonData(b)
	if err != nil {
		ir.Err = err
		return ir
	}

	obj := new(Data)
	err = obj.Parser(b)
	if err != nil {
		ir.Err = err
		return ir
	}

	Persist := func(tag *models.Tag, data Node) {
		img, IsNew, err := persistNodeData(tag, data)

		if IsNew && err != nil {
			ir.NewImageError(err)
			return
		}

		if IsNew && err == nil {
			ir.NewImage(img)
			return
		}

		if err != nil {
			ir.EditImageError(err)
			return
		}

		ir.EditImage(img)
	}

	for _, tagsPage := range obj.Data.TagPages {
		tagModel, err := persistTagData(obj.CountryCode, tagsPage.Tags)
		if err != nil {
			ir.Err = err
			return ir
		}

		for _, node := range tagsPage.Tags.Media.Nodes {
			Persist(tagModel, node)
		}

		for _, node := range tagsPage.Tags.TopPost.Nodes {
			Persist(tagModel, node)
		}
	}

	return ir
}

func persistTagData(countryCode string, tags Tag) (*models.Tag, error) {
	tagModel, err := models.GetTagByName(tags.Name)
	if models.IsErrTagNotExist(err) {
		tagModel = &models.Tag{Name: tags.Name, Count: tags.Media.Count, Country: countryCode}
		if err := models.CreateTag(tagModel); err != nil {
			return nil, err
		}

		return tagModel, nil
	}

	tagModel.Count = tags.Media.Count
	if err := models.UpdateTag(tagModel); err != nil {
		return nil, err
	}

	return tagModel, nil
}

func persistNodeData(tag *models.Tag, data Node) (img *models.Image, IsNew bool, err error) {
	IsNew = false
	img, err = models.GetImageByInstagramID(data.Id)
	if err != nil && models.IsErrImageNotExist(err) {
		img = &models.Image{}
		IsNew = true
	}

	img.InstagramID = data.Id
	img.DisplaySrc = data.DisplaySrc
	img.ThumbnailSrc = data.ThumbnailSrc
	img.IsVideo = data.IsVideo
	img.Code = data.Code
	//img.Date = data.Date
	img.Caption = data.Caption
	img.TagID = tag.ID
	img.TagName = tag.Name
	img.Height = data.Dimensions.Height
	img.Width = data.Dimensions.Width
	img.Owner = data.Owner.Id
	img.Comments = data.Comments.Count
	img.Likes = data.Likes.Count
	img.IsNew = IsNew

	if img.ID != 0 {
		// Update
		err = models.UpdateImage(img)
		return
	}

	// Create new
	err = models.CreateImage(img)
	return
}
