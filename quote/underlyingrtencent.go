package quote

import (
	"errors"
	"io/ioutil"
	"net/http"
	base "opquotes/basedata"
	"opquotes/log"
	util "opquotes/utils"
	"strings"
)
type Tencent struct{

}
func (t *Tencent)getRawTickString(market string, symbol string) []string {
	resp, err := http.Get("http://web.sqt.gtimg.cn/q=" + market + symbol)
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			tickArr := strings.Split(string(body), "~")
			if len(tickArr) > 1 {
				tickArr[1] = util.StringFromGBK(tickArr[1])
			} else {
				log.Logger.Info("getRawTickString"+market+"-"+symbol)
			}
			return tickArr
		}
	}
	return nil
}

func (t *Tencent)GetMarketFromCode(code string)string{
	Shprefix:=[]string{"50","51","60","90","110","113","132","204"}
	Szprefix:=[]string{"00","13","18","15","16","20","30","39","115","1318"}
	for _,prefix:=range Shprefix{
		if strings.HasPrefix(code,prefix){
			return "sh"
		}
	}
	for _,prefix:=range Szprefix{
		if strings.HasPrefix(code,prefix){
			return "sz"
		}
	}
	return ""
}
func (t *Tencent)UnderlyingTick(){
	Underlying:=[]string{"510050","510300"}
	for _,code:=range Underlying{
		_,err:=t.GetStockLastTick(code)
		if err!=nil{
			log.Logger.Info("spidertencent插入tick数据报错"+err.Error())
		}
	}
}
func (t *Tencent)GetStockLastTick(symbol string) (*base.Optiontick, error) {
	ret := &base.Optiontick{}
	market:=t.GetMarketFromCode(symbol)
	tickArr := t.getRawTickString(market, symbol)
	if tickArr == nil || len(tickArr) < 38 {
		return nil, errors.New("ErrGetStockTick")
	}
	if tickArr != nil && len(tickArr) >= 38 {
		timeStr := tickArr[30]
		ret.Datetime = util.ParseBeijingTime("20060102150405", timeStr)
		ret.Exchange=market
		ret.Code = symbol
		ret.Name=tickArr[1]
		ret.Lastprice = util.ParseFloat(tickArr[3])
		//ret.clode= ret.Price
		ret.Preclose = util.ParseFloat(tickArr[4])
		if ret.Preclose!=0{
			ret.Chg =(ret.Lastprice-ret.Preclose)*100/ret.Preclose
		}
		ret.Open = util.ParseFloat(tickArr[5])
		ret.High = util.ParseFloat(tickArr[33])
		ret.Low = util.ParseFloat(tickArr[34])
		ret.Vol = util.ParseInt64(tickArr[6])
		ret.Amount = float64(util.ParseInt(tickArr[37]) * 10000)
		//ret.UpperLimit = utils.ParseFloat(tickArr[47])
		//ret.LowerLimit = utils.ParseFloat(tickArr[48])

		ret.Bidvolume5 = util.ParseFloat(tickArr[18])
		ret.Bidprice5 = util.ParseFloat(tickArr[17])

		ret.Askvolume5 = util.ParseFloat(tickArr[28])
		ret.Askprice5 = util.ParseFloat(tickArr[27])

		ret.Bidvolume4 = util.ParseFloat(tickArr[16])
		ret.Bidprice4 = util.ParseFloat(tickArr[15])
		ret.Askvolume4 = util.ParseFloat(tickArr[26])
		ret.Askprice4 = util.ParseFloat(tickArr[25])

		ret.Bidvolume3 = util.ParseFloat(tickArr[14])
		ret.Bidprice3= util.ParseFloat(tickArr[13])
		ret.Askvolume3 = util.ParseFloat(tickArr[24])
		ret.Askprice3 = util.ParseFloat(tickArr[23])


		ret.Bidvolume2 = util.ParseFloat(tickArr[12])
		ret.Bidprice2= util.ParseFloat(tickArr[11])
		ret.Askvolume2 = util.ParseFloat(tickArr[22])
		ret.Askprice2 = util.ParseFloat(tickArr[21])

		ret.Bidvolume1 = util.ParseFloat(tickArr[10])
		ret.Bidprice1= util.ParseFloat(tickArr[9])
		ret.Askvolume1 = util.ParseFloat(tickArr[20])
		ret.Askprice1 = util.ParseFloat(tickArr[19])
		ret.Save()

	}
	return ret, nil
}

