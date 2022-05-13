package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"strings"
)
var Conf = &Config{}
type Config struct{
	v *viper.Viper
}
func Getstring(section string)string{
	return Conf.v.GetString(section)
}
func Getint64(section string)int64{
	return Conf.v.GetInt64(section)
}
func Getint(section string)int{
	return Conf.v.GetInt(section)
}
func Getfloat(section string)float64{
	return Conf.v.GetFloat64(section)
}
func GetStringList(section string)[]string{
	return  Conf.v.GetStringSlice(section)
}
func Get(section string)map[string]interface{}{
	ret :=Conf.v.Get(section)
	if retdata,ok := ret.(map[string]interface{});ok{
		return retdata
	}else{
		return nil
	}
}
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
func init(){
	filepath:=GetCurrentDirectory()
	fmt.Println(filepath)
	Conf.LoadConfigFromToml(filepath)
}
func (c *Config)LoadConfigFromToml(filepath string)error{
	c.v = viper.New()
	c.v.SetConfigName("app")
	c.v.AddConfigPath(filepath)
	c.v.SetConfigType("toml")
	if err:=c.v.ReadInConfig();err!=nil{
		return err
	}
	//section:=c.v.Get("app")
	//
	//
	//fmt.Println("字符类型=",c.v.GetString("jwt.token"))
	//fmt.Println(section.(map[string]interface{})["runmode"])
	//log.Printf("toml配置文件加载成功",reflect.ValueOf(c.v.Get("app")).Type())
	return nil
}
var (
	/* 0 - 500 请求错误*/
	 SecretKey = "02joepwkmldsa02jpoqkwmdlas"

)



