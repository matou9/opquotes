package sina

import (
	"bytes"
	"fmt"
	"opquotes/quote"
	"strings"

	//"opquotes/pub"
	"reflect"
	"testing"
	//"time"
)

func Test_sms(t *testing.T){

	quote.UnderlyingTick()
	//
	//time.Sleep(time.Second*100)
	////tick,_:=GetStockLastTick("510050")
	////_,errdb:=pub.DB.Table("optiontick").Insert(tick)
	////fmt.Println(errdb)
	//UnderlyingTick()
	//fmt.Println(pub.IsTradeTime())
}
func GenInsertSQL(src interface{},o string)string{
	var buffer bytes.Buffer
	primarykeys:=make([]string,0)
	names:=make([]string,0)
	types:=make([]string,0)
	srctype :=reflect.TypeOf(src)
	tablename:=srctype.Name()
	kind:=reflect.TypeOf(src).Kind()
	if kind==reflect.Struct{
		for i:=0;i<srctype.NumField();i++{
			fieldobj:=srctype.Field(i)
			if strings.Contains(fieldobj.Tag.Get("xorm"),"-")==false{
				names=append(names,fieldobj.Name)
				types=append(types,fieldobj.Type.String())
			}
			if strings.Contains(fieldobj.Tag.Get("xorm"),"pk"){
				primarykeys=append(primarykeys,fieldobj.Name)
			}
		}
	}
	t:=make([]string,0)
	for _,typename:=range types{
		switch typename{
		case "string":
			t=append(t,"'%s'")
		case "time.Time":
			t=append(t,"'%s'")
		case "float64":
			t=append(t,"%f")
		case "int64":
			t=append(t,"%d")
		case "int":
			t=append(t,"%d")
		}
	}
	buffer.WriteString("`insert into "+tablename+"(")
	buffer.WriteString(strings.Join(names,",")+") values (")
	buffer.WriteString(strings.Join(t,",")+") on conflict("+strings.Join(primarykeys,",")+") do update set ")
	upsql:=make([]string,0)
	objfields:=make([]string,0)
	for index,name:=range names{
		for _,key:=range primarykeys{
			if name==key{
				continue
			}else{
				upsql=append(upsql,name+"="+"EXCLUDED."+name)
				break
			}
		}
		if types[index]=="time.Time"{
			objfields=append(objfields,"utils.Time2Str("+o+"."+name+")")
		}else{
			objfields=append(objfields,o+"."+name)
		}

	}
	buffer.WriteString(strings.Join(upsql,","))
	buffer.WriteString("`,")
	buffer.WriteString(strings.Join(objfields,","))

	fmt.Println(buffer.String())
	return buffer.String()
}
