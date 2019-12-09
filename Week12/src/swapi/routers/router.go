package routers

import (
	"swapi/controllers"
	"github.com/astaxie/beego"
	"github.com/boltdb/bolt"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//注册 namespace
	beego.AutoRouter(&controllers.APIController{})
	db, err := bolt.Open("swapi.db", 0600, nil)
    if err != nil {
        beego.Notice(err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		//判断要创建的表是否存在
		b := tx.Bucket([]byte("Film"))
			if b == nil {
				//创建叫"Film"的表
				_, err := tx.CreateBucket([]byte("Film"))
				if err != nil {
					beego.Notice(err)
				}
				_, err = tx.CreateBucket([]byte("People"))
				if err != nil {
					beego.Notice(err)
				}
				_, err = tx.CreateBucket([]byte("Planet"))
				if err != nil {
					beego.Notice(err)
				}
				_, err = tx.CreateBucket([]byte("Specie"))
				if err != nil {
					beego.Notice(err)
				}
				_, err = tx.CreateBucket([]byte("Starship"))
				if err != nil {
					beego.Notice(err)
				}
				_, err = tx.CreateBucket([]byte("Vehicle"))
				if err != nil {
					beego.Notice(err)
				}
			}
		return nil
	})
	//更新数据库失败
	if err != nil {
		beego.Notice(err)
	}
}
