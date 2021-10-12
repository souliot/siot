package basic

import (
	"github.com/flosch/pongo2/v4"
	logs "github.com/souliot/siot-log"
)

func generateGitignore(appname string) (err error) {
	logs.Info(appname)
	gi, err := tpl.ReadFile("tpl/gitignore.sm")
	if err != nil {
		return
	}
	tpl_gi, err := pongo2.FromString(string(gi))
	if err != nil {
		return
	}
	res, err := tpl_gi.Execute(pongo2.Context{
		"appname": appname,
	})
	if err != nil {
		return
	}
	logs.Info(res)
	return
}
