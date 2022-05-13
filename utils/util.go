package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"math"
	"opquotes/pub"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// StringFromGBK 转换GBK
func StringFromGBK(src string) (dst string) {
	data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewDecoder()))
	if err == nil {
		dst = string(data)
	}
	return
}
func Time2Str(t time.Time) string {
	const shortForm = "2006-01-02 15:04:05"
	temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	str := temp.Format(shortForm)
	return str
}
// StringFromGBK2 StringFromGBK2
func StringFromGBK2(src []byte) (dst string) {
	data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader(src), simplifiedchinese.GBK.NewDecoder()))
	if err == nil {
		dst = string(data)
	}
	return
}

// ParseBeijingTime 解析北京时间
func ParseBeijingTime(layout, value string) time.Time {
	loc, err := time.LoadLocation("Asia/Chongqing") // 北京时间
	if err == nil {
		tx, err := time.ParseInLocation(layout, value, loc)
		if err == nil {
			return tx
		}
		return time.Time{}
	}

	return time.Time{}
}

func ToString(value interface{}) string {
	switch value.(type) {
	case string:
		return value.(string)
	case int8:
		return strconv.FormatInt(int64(value.(int8)), 10)
	case int16:
		return strconv.FormatInt(int64(value.(int16)), 10)
	case int32:
		return strconv.FormatInt(int64(value.(int32)), 10)
	case int64:
		return strconv.FormatInt(int64(value.(int64)), 10)
	case uint8:
		return strconv.FormatUint(uint64(value.(uint8)), 10)
	case uint16:
		return strconv.FormatUint(uint64(value.(uint16)), 10)
	case uint32:
		return strconv.FormatUint(uint64(value.(uint32)), 10)
	case uint64:
		return strconv.FormatUint(uint64(value.(uint64)), 10)
	case float32:
		return strconv.FormatFloat(float64(value.(float32)), 'g', -1, 64)
	case float64:
		return strconv.FormatFloat(float64(value.(float64)), 'g', -1, 64)
	case bool:
		return strconv.FormatBool(value.(bool))
	default:
		return fmt.Sprintf("%+v", value)
	}
}
func Encode(src string) (dst string) {
	data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewEncoder()))
	if err == nil {
		dst = string(data)
	}
	return
}

// Decode 转UTF8
func Decode(src string) (dst string) {
	data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewDecoder()))
	if err == nil {
		dst = string(data)
	}
	return
}

// ParseInt ParseInt
func ParseInt(src string) int {
	i, err := strconv.Atoi(src)
	if err != nil {
		return 0
	}
	return i
}

// ParseInt32 ParseInt
func ParseInt32(src string) int32 {
	i, err := strconv.Atoi(src)
	if err != nil {
		return 0
	}
	return int32(i)
}

// ParseFloat ParseFloat
func ParseFloat(src string) float64 {
	f, err := strconv.ParseFloat(src, 10)
	if err != nil {
		return 0
	}
	return f
}

// GetMD5 转md5
func GetMD5(str string) string {
	h := md5.New()
	h.Write([]byte(str)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}




// ParseInt32 ParseInt


func Round(src float64) float64 {
	return math.Round(src*100)/100
}
func Round4(src float64) float64 {
	return math.Round(src*10000)/10000
}
func ParseInt64(src string) int64 {
	if price,err:= strconv.ParseInt(src, 10,64);err==nil{
		return price
	}else{
		return 0
	}
}
func UnixToTime(src int64)time.Time{
	return time.Unix(src, 0)
}
func ParseFloat64(src string) float64 {
	f, err := strconv.ParseFloat(src, 10)
	if err != nil {
		return 0
	}
	return f
}
func Strtotime2(timestr string)time.Time{
	todayZero, err := time.ParseInLocation("20060102",timestr ,time.Local)
	if err==nil{
		return todayZero
	}
	return time.Time{}


}
func Strtotime(timestr string)time.Time{
	todayZero, err := time.ParseInLocation("2006-01-02 15:04:05",timestr ,time.Local)
	if err==nil{
		return todayZero
	}
	return time.Time{}


}
/*时间戳->字符串*/
func Stamp2Str(stamp int64) string{
	timeLayout := "2006-01-02 15:04:05"
	str:=time.Unix(stamp,0).Format(timeLayout)
	return str
}
/*时间戳->时间对象*/
func Stamp2Time(stamp int64)time.Time{
	stampStr:=Stamp2Str(stamp)
	timer:=Str2Time(stampStr)
	return timer
}
func Time2StrLast(t time.Time)string{
	const shortForm = "2006-01-02 15:04:05"
	temp := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 59, time.Local)
	str := temp.Format(shortForm)
	return str
}
func Time2StrDate(t time.Time)string{
	const shortForm = "2006-01-02"
	temp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	str := temp.Format(shortForm)
	return str
}
func Str2Time(formatTimeStr string) time.Time{
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, formatTimeStr, loc) //使用模板在对应时区转化为time.time类型

	return theTime

}

func IsTradeTime()bool{
	morning_trade_start:=Str2Time(time.Now().Format("2006-01-02")+" 09:30:00")
	morning_trade_end:=Str2Time(time.Now().Format("2006-01-02")+" 11:30:00")
	after_trade_start:=Str2Time(time.Now().Format("2006-01-02")+" 13:00:00")
	after_trade_end:=Str2Time(time.Now().Format("2006-01-02")+" 15:00:01")
	now:=time.Now()
	if (now.After(morning_trade_start)&&now.Before(morning_trade_end))||(now.After(after_trade_start)&&now.Before(after_trade_end)){
		return true
	}
	return false
}
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func Int64ToString(src int64)string{
	s := strconv.FormatInt(src, 10)
	return s
}
func GenInsertSQL(src interface{},o string)string{
	var buffer bytes.Buffer
	primarykeys:=make([]string,0)
	names:=make([]string,0)
	fnames:=make([]string,0)
	types:=make([]string,0)
	srctype :=reflect.TypeOf(src)
	tablename:=srctype.Name()
	kind:=reflect.TypeOf(src).Kind()
	if kind==reflect.Struct{
		for i:=0;i<srctype.NumField();i++{
			fieldobj:=srctype.Field(i)
			if strings.Contains(fieldobj.Tag.Get("xorm"),"-")==false{
				if fieldobj.Tag.Get("json")!=""{
					names=append(names,fieldobj.Tag.Get("json"))
				}else{
					names=append(names,fieldobj.Name)
				}
				fnames=append(fnames,fieldobj.Name)
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
			objfields=append(objfields,"utils.Time2Str("+o+"."+fnames[index]+")")
		}else{
			objfields=append(objfields,o+"."+fnames[index])
		}

	}
	buffer.WriteString(strings.Join(upsql,","))
	buffer.WriteString("`,")
	buffer.WriteString(strings.Join(objfields,","))

	fmt.Println(buffer.String())
	return buffer.String()
}

func GetFromDBToJson(sql string)(data []interface{},err error){
	var jsondata interface{}
	var querydata interface{}
	retdata:=make([]interface{},0)
	//sql := `select row_to_json(contract) from contract limit 10` //select row_to_json(t) from (select orderid, volume from suborders)t
	//select array_to_json(array_agg(row_to_json(t))) from
	//(select price,vol,amount,to_char(datetime, 'HH12MI'),case pre_close when 0 then 0 else (price-pre_close)/pre_close end risefall from mindata)t
	if results, e := pub.DB.QueryString(sql);e!=nil{
		return nil,e
	}else{
		for _,v:=range results{
			if _data,ok:=v["row_to_json"];ok{
				querydata=_data
			}else if _data,ok:=v["array_to_json"];ok{
				querydata=_data
			}
			if querydata!=nil{
				if _querydata,ok:=querydata.(string);ok{
					if err=json.Unmarshal([]byte(_querydata), &jsondata);err!=nil{
						return nil,errors.New("GetFromDBToJson格式化字符串失败")
					}
					retdata=append(retdata,jsondata)
				}

			}
		}
		return retdata,nil
	}
	return
}
func RemainDay(day time.Time)int64{
	current:=time.Now().Unix()
	timestamp := day.Unix()    //转化为时间戳 类型是int64
	res:=int64((timestamp-current)/86400) //相差值
	return res

}
func SplitStrQuote(src string)string{
	var buffer bytes.Buffer
	for _,code :=range strings.Split(src,","){
		buffer.WriteString("'")
		buffer.WriteString(code)
		buffer.WriteString("'")
		buffer.WriteString(",")
	}

	return strings.TrimSuffix(buffer.String(),",")
}