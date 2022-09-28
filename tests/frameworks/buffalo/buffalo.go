package buffalo

import (
	"github.com/GoAdminGroup/themes/adminlte"
	// add buffalo adapter
	_ "github.com/vafinvr/go-admin/adapter/buffalo"
	"github.com/vafinvr/go-admin/modules/config"
	"github.com/vafinvr/go-admin/modules/language"
	"github.com/vafinvr/go-admin/plugins/admin/modules/table"

	// add mysql driver
	_ "github.com/vafinvr/go-admin/modules/db/drivers/mysql"
	// add postgresql driver
	_ "github.com/vafinvr/go-admin/modules/db/drivers/postgres"
	// add sqlite driver
	_ "github.com/vafinvr/go-admin/modules/db/drivers/sqlite"
	// add mssql driver
	_ "github.com/vafinvr/go-admin/modules/db/drivers/mssql"
	// add adminlte ui theme
	_ "github.com/GoAdminGroup/themes/adminlte"

	"github.com/vafinvr/go-admin/template"
	"github.com/vafinvr/go-admin/template/chartjs"

	"net/http"
	"os"

	"github.com/gobuffalo/buffalo"
	"github.com/vafinvr/go-admin/engine"
	"github.com/vafinvr/go-admin/plugins/admin"
	"github.com/vafinvr/go-admin/plugins/example"
	"github.com/vafinvr/go-admin/tests/tables"
)

func newHandler() http.Handler {
	bu := buffalo.New(buffalo.Options{
		Env:  "test",
		Addr: "127.0.0.1:9033",
	})

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(tables.Generators)
	adminPlugin.AddGenerator("user", tables.GetUserTable)

	examplePlugin := example.NewExample()

	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin, examplePlugin).Use(bu); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	bu.ServeFiles("/uploads", http.Dir("./uploads"))

	return bu
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) http.Handler {
	bu := buffalo.New(buffalo.Options{
		Env:  "test",
		Addr: "127.0.0.1:9033",
	})

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(gens)

	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfig(&config.Config{
		Databases: dbs,
		UrlPrefix: "admin",
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language:    language.EN,
		IndexUrl:    "/",
		Debug:       true,
		ColorScheme: adminlte.ColorschemeSkinBlack,
	}).
		AddPlugins(adminPlugin).Use(bu); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	bu.ServeFiles("/uploads", http.Dir("./uploads"))

	return bu
}
