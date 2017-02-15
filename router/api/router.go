// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package api

import (
	"gopkg.in/macaron.v1"

	"github.com/rodkranz/photoins/modules/context"
	"github.com/rodkranz/photoins/router/api/v1/instagram"
)

// Contexter middleware already checks token for user sign in process.
func reqToken() macaron.Handler {
	return func(ctx *context.Context) {
		if !ctx.IsSigned {
			ctx.Error(401)
			return
		}
	}
}

func reqBasicAuth() macaron.Handler {
	return func(ctx *context.Context) {
		if !ctx.IsBasicAuth {
			ctx.Error(401)
			return
		}
	}
}

func RegisterRoutes(m *macaron.Macaron) {
	//bind := binding.Bind

	m.Group("/v1", func() {

		m.Get("/image/search", instagram.Search)
		m.Get("/image/:instagramId", instagram.GetInfoImageByInstagramId)
		m.Get("/image/import/tag/:tag", instagram.ImportByTag)

	}, context.APIContexter())
}
