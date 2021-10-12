package basic

import "embed"

//go:embed tpl
var tpl embed.FS

func GenerateBasic(appname string) (err error) {
	generateGitignore(appname)
	return
}
