package base

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"opquotes/basedata"
	"opquotes/log"
	"opquotes/quote"
	"opquotes/utils"
	"sort"
	"strings"
)
type ops []*basedata.Option
func (s ops) Len() int {
	return len(s)
}

//Less():成绩将有低到高排序
func (s ops) Less(i, j int) bool {
	return s[i].Strikeprice<s[j].Strikeprice
}

//Swap()
func (s ops) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func GetComm(c *gin.Context){
	uid :=utils.ParseInt64(c.PostForm("uid"))
	sql:=fmt.Sprintf(`select row_to_json(t) from (select accountid,openfee,closefee,marginfee from subaccount where uid=%d)t `,uid)
	if data,err:=utils.GetFromDBToJson(sql);err!=nil{
		c.JSON(http.StatusOK, gin.H{"status": 204, "data": gin.H{}, "message": "获取费率数据失败，失败的错误="+err.Error()})
	}else if data==nil||len(data)==0{
		c.JSON(http.StatusOK, gin.H{"status": 204, "data": gin.H{}, "message": "没有查到该用户的费率数据,代码="+utils.ToString(uid)})
	}else if len(data)==1{
		c.JSON(http.StatusOK,gin.H{"status":200,"message":"返回用户的费率数率成功","data":data[0]})
	}

}
func GetOptions(c *gin.Context) {
	Type:=strings.ToUpper(strings.TrimSpace(c.PostForm("type")))
	OptionByMonth:=make(map[string]map[string][]*basedata.Option)
	if err,options:= quote.OptionMgr.GetOptions();err!=nil{
		log.Logger.Info("获取期权标的出现错误,错误代码="+err.Error())
		c.JSON(http.StatusOK,gin.H{"status":200,"data":gin.H{},"message":"获取options出现错误"+err.Error()})
	}else{
		opss:=ops(options)
		sort.Sort(opss)
		if Type=="LIST"{
			c.JSON(http.StatusOK,gin.H{"status":200,"message":"返回成功","data":gin.H{"options":opss}})
		}else{
			for _,op:=range opss{
				if _,ok:=OptionByMonth[op.Underlying];!ok{
					OptionByMonth[op.Underlying]=make(map[string][]*basedata.Option)
				}
				OptionByMonth[op.Underlying][op.Month]=append(OptionByMonth[op.Underlying][op.Month],op)
			}
			c.JSON(http.StatusOK,gin.H{"status":200,"message":"返回成功","data":gin.H{"options":OptionByMonth}})
		}

	}
}

func GetOptionsByMap(c *gin.Context) {
	if err,options:= quote.OptionMgr.GetOptionsMap();err!=nil{
		log.Logger.Info("获取期权标的出现错误,错误代码="+err.Error())
		c.JSON(http.StatusOK,gin.H{"status":200,"data":gin.H{},"message":"获取options出现错误"+err.Error()})
	}else{
		c.JSON(http.StatusOK,gin.H{"status":200,"message":"返回成功","data":gin.H{"options":options}})

	}
}
func Notice(c *gin.Context) {

	sql:=fmt.Sprintf(`select row_to_json(notice) from notice order by datetime desc limit 1 `)
	if data,err:=utils.GetFromDBToJson(sql);err!=nil{
		c.JSON(http.StatusOK, gin.H{"status": 204, "data": gin.H{}, "message": "获取最新公告数据失败，失败的错误="+err.Error()})
	}else if data==nil||len(data)==0{
		c.JSON(http.StatusOK, gin.H{"status": 204, "data": gin.H{}, "message": "没有查到最新公告"})
	}else if len(data)==1{
		c.JSON(http.StatusOK,gin.H{"status":200,"message":"返回成功","data":gin.H{"notice":data[0]}})
	}
}