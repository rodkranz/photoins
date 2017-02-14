package routers

import (
	"github.com/rodkranz/photoins/models"
	"github.com/rodkranz/photoins/modules/log"
	"github.com/rodkranz/photoins/modules/setting"
	"github.com/rodkranz/photoins/modules/verify"
)

func GlobalInit() {
	setting.NewContext()

	log.Trace("Custom path: %s", setting.CustomPath)
	log.Trace("Log path: %s", setting.LogRootPath)

	models.LoadConfigs()
	setting.NewServices()

	if err := models.NewEngine(); err != nil {
		log.Fatal(4, "Fail to initialize ORM engine: %v", err)
	}
	models.HasEngine = true

	verify.CheckRunMode()
}
