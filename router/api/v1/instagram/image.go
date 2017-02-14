// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package instagram

import (
	"fmt"

	"github.com/rodkranz/photoins/models"
	"github.com/rodkranz/photoins/modules/context"
)


func GetInfoImageByInstagramId(ctx *context.APIContext) {
	image, err := models.GetImageByInstagramID(ctx.Params(":instagramId"))
	if err != nil {
		if models.IsErrImageNotExist(err) {
			ctx.Error(404, "Image not found", nil)
			return
		}

		ctx.Error(500, "GetImageByInstagramID", err)
		return
	}

	ctx.Render(200, fmt.Sprintf("Image found [instagram_id: %v]", image.ID), image)
}

func Search(ctx *context.APIContext) {
	opts := &models.SearchImageOptions{
		Keyword: ctx.Query("q"),
		Tag:  ctx.Query("tag"),
		Page: ctx.QueryInt("page"),
		PageSize: ctx.QueryInt("limit"),
	}

	if opts.Page <= 1 {
		opts.Page = 1
	}

	if opts.PageSize <= 0 {
		opts.PageSize = 10
	}

	images, total, err := models.SearchImage(opts)
	if err != nil {
		ctx.Error(400, "Error to search images", err)
		return
	}

	ctx.JSON(200, map[string]interface{}{
		"message":  "List of images found",
		"status":   200,
		"total":    total,
		"resource": images,
	})
}
