// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"gopkg.in/urfave/cli.v2"

	"github.com/rodkranz/photoins/router"
	"github.com/rodkranz/photoins/models"
	"github.com/rodkranz/photoins/modules/verify"
	"github.com/rodkranz/photoins/modules/instagram"
)

var Service = &cli.Command{
	Name:        "service",
	Usage:       "Run Service",
	Description: `Start Service that fetch data from instagram.`,
	Action:      runService,
	Flags:       []cli.Flag{

	},
}

func runService(ctx *cli.Context) error {
	routers.GlobalInit()
	verify.CheckVersion()

	tag := ctx.String("tag")
	b, err := instagram.FetchFromTag(tag)
	if err != nil {
		return err
	}

	b, err = instagram.ParseJsonData(b)
	if err != nil {
		return err
	}

	obj := new(instagram.Data)
	err = obj.Parser(b)
	if err != nil {
		return err
	}

	for _, tagsPage := range obj.Data.TagPages {
		tagModel, err := persistTagData(obj.CountryCode, tagsPage.Tags)
		if err != nil {
			return err
		}

		for _, node := range tagsPage.Tags.Media.Nodes {
			err := persistNodeData(tagModel, node)
			if err != nil {
				return err
			}
		}
	}

	ctx.App.Writer.Write([]byte("Imported!"))
	return nil
}

func persistTagData(countryCode string, tags instagram.Tag) (*models.Tag, error) {
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

func persistNodeData(tag *models.Tag, data instagram.Node) error {
	img, err := models.GetImageByInstagramID(data.Id)
	if err != nil && models.IsErrImageNotExist(err) {
		img = &models.Image{}
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

	if img.ID != 0 {
		// Update
		return models.UpdateImage(img)
	}

	// Create new
	return models.CreateImage(img)
}
