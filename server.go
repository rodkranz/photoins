// Copyright 2016 Kranz. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package main

import (
	"os"
	"runtime"

	"gopkg.in/urfave/cli.v2"

	"github.com/rodkranz/photoins/cmd"
	"github.com/rodkranz/photoins/modules/setting"
)

const VER = "1.0.0"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	setting.AppVer = VER
}

func main() {
	app := cli.App{
		Name:    "PhotoIns",
		Usage:   "Instagram photo by tag",
		Version: VER,
		Commands: []*cli.Command{
			cmd.Server,
			cmd.Import,
		},
	}

	app.Flags = append(app.Flags, []cli.Flag{}...)
	if len(os.Args) == 1 {
		os.Args = append(os.Args, cmd.Server.Name)
	}
	app.Run(os.Args)
}
