// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"errors"
	"fmt"

	"github.com/urfave/cli"

	"github.com/rodkranz/photoins/modules/instagram"
	"github.com/rodkranz/photoins/modules/verify"
	"github.com/rodkranz/photoins/router"
)

var Import = cli.Command{
	Name:        "service",
	Usage:       "Run Import",
	Description: `Start Import images from instagram.`,
	Action:      runImport,
	Flags: []cli.Flag{
		stringFlag("tag", "", "the tag name to import"),
	},
}

func runImport(ctx *cli.Context) error {
	if !ctx.IsSet("tag") {
		return errors.New("Tag is not specified.")
	}

	routers.GlobalInit()
	verify.CheckVersion()

	tag := ctx.String("tag")
	ir := instagram.ImportTag(tag)
	if ir.Err != nil {
		return ir.Err
	}

	layout := `
The tag #%s.

Total of new images         : %d
Total of images updated     : %d
Total of error new images   : %d
Total of error update images: %d

Error: %v

`

	fmt.Fprintf(ctx.App.Writer,
		layout,
		tag,
		len(ir.NewImages),
		len(ir.EditImages),
		len(ir.NewImagesErr),
		len(ir.EditImagesErr),
		ir.Err,
	)

	return nil
}
