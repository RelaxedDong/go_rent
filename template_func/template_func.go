package template_func

import (
	"github.com/astaxie/beego"
)

func Indexaddone(index int) (index1 int) {
	index1 = index + 1
	return
}

func init() {
	beego.AddFuncMap("indexaddone", Indexaddone)
}
