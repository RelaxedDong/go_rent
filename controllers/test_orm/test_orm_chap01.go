package test_orm

import (
	"github.com/astaxie/beego"
)

type TestORMController struct {
	beego.Controller
}

func (r *TestORMController) Get() {
	// 插入
	//users := []models.User{{Name: "donghao", Age: 24, Addr: "天通苑"}, {Name: "tanyajuan", Age: 24, Addr: "天通苑"}}
	// 最多插入20条
	//created_cnt, err := models.OrmManager.InsertMulti(20, users)
	//fmt.Println(created_cnt, err) // 0 Error 1062 (23000): Duplicate entry 'donghao' for key 'auth_user.name'

	//id, err := models.OrmManager.Insert(&models.User{Name: "donghao", Age: 24, Addr: "天通苑"})
	//if err != nil {
	//	fmt.Println("插入失败..", err)
	//}

	// 查询
	//user := models.User{Name: "donghao1", Age: 24}
	//err := models.OrmManager.Read(&user, "Name", "Age") // 默认是根据id查询，这里换成Name字段查询
	//if err == orm.ErrNoRows {
	//	fmt.Println("查询失败..", err) // <QuerySeter> no row found
	//}
	//r.TplName = "test_orm/test_orm_chap01.html"
	//r.Data["err"] = err
	//r.Data["user"] = user
}
