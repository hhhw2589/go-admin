package beego

import (
	_ "github.com/GoAdminGroup/go-admin/adapter/beego"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	_ "github.com/GoAdminGroup/themes/adminlte"

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/GoAdminGroup/go-admin/plugins/example"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/http"
	"os"
)

func NewBeegoHandler() http.Handler {

	app := beego.NewApp()

	eng := engine.Default()
	adminPlugin := admin.NewAdmin(datamodel.Generators)
	adminPlugin.AddGenerator("user", datamodel.GetUserTable)

	examplePlugin := example.NewExample()

	if err := eng.AddConfigFromJson(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin, examplePlugin).Use(app); err != nil {
		panic(err)
	}

	template.AddComp(chartjs.NewChart())

	app.Handlers.Get("/admin", func(ctx *context.Context) {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return datamodel.GetContent()
		})
	})

	beego.BConfig.Listen.HTTPAddr = "127.0.0.1"
	beego.BConfig.Listen.HTTPPort = 9087

	return app.Handlers
}
