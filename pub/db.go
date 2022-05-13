package pub

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"log"
	"opquotes/config"
	_ "opquotes/config"
	"sync"
)
var(
	host =config.Getstring("postgres.host")
	port =config.Getint("postgres.port")
	user = config.Getstring("postgres.user")
	password = config.Getstring("postgres.password")
	dbname = config.Getstring("postgres.dbname")
	constr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host,port,user,password,dbname)
	DB *xorm.Engine
	once sync.Once
)
func Init(){
	var err error
	once.Do(func(){
		if DB,err = xorm.NewEngine("postgres",constr);err!=nil{
			log.Fatal("数据库连接失败:", err)
		}
	})


	//f, err := os.Create("xormsql.log")
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
	//DB.SetLogger(xorm.NewSimpleLogger(f))
	//DB.ShowSQL()

}

func BeginTrans()*xorm.Session{
	ss:=DB.NewSession()
	defer ss.Close()
	if err:=ss.Begin();err!=nil {
		ss.Rollback()
		panic("开始事务发生错误"+err.Error())
	}else{
		return ss
	}
}