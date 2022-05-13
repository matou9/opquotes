package quote
import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
	"strings"
	"testing"
	"time"
)
func Test_util1(t *testing.T) {
	a:= NewScrawl()            //https://dealer.yiche.com/100030649/news/202103/566027248.html
	if err,data,rurl:=a.Request("https://dealer.yiche.com/100030649/news/202103/56602555557248.html",true);err==nil{
		if strings.Contains(data,"页面不存在，请打开腾讯新闻客户端")==true{
			fmt.Println("没反问道")
		}
		fmt.Println(rurl)
	}
	time.Sleep(time.Minute)


}

func Test_util2(t *testing.T) {
	//var i int
	var checkflag bool
	var url, check string
	a := NewScrawl()
	if f, err := excelize.OpenFile("e:\\excel\\last.xlsx"); err != nil {
		fmt.Println("打开EXCEL错误=", err)
	} else {
		rows, _ := f.GetRows("1")
		for index, row := range rows {
			if len(row) >= 10 {
				url = row[5]
			} else {
				continue
			}

			if strings.Contains(url, "http") {
				if len(row) >= 26 {
					check = row[26]
				} else {
					continue
				}

				if strings.Contains(url, "toutiao") {
					checkflag = false
				} else {
					checkflag = true
				}
				col := "AA" + strconv.Itoa(index+1)
				if strings.Contains(check, "未删除") == true {
					if err, data, rurl := a.Request(url, checkflag); err == nil {
						if strings.Contains(data, "您访问的页面找不回来了") == true ||
							strings.Contains(data, " 页面不存在") == true ||
							strings.Contains(data, " 内容被删除") == true ||
							strings.Contains(data, "出错了！文章没有找到哦") ||
							strings.Contains(data, "您要访问的页面弄丢了") ||
							strings.Contains(data, "页面不存在，请下载腾讯新闻客户端") ||
							strings.Contains(data, "文章暂时找不到了") ||
							strings.Contains(data, "404错误页") ||
							strings.Contains(data, "这篇文章从地球上消失了") ||
							strings.Contains(data, "文章不存在或已被删除") ||
							strings.Contains(data, "您访问的文章走失了") ||
							strings.Contains(data, "网页跑丢了") ||
							strings.Contains(data, "文章不存在") ||
							strings.Contains(data, "参数错误") ||
							strings.Contains(data, "您访问的帖子不存在") {
							fmt.Println("已经删除了该链接=", url, "所在列=", col)
							f.SetCellValue("1", col, "已经删除")

						} else if strings.Contains(url, "yiche") || strings.Contains(url, "bitauto") {
							if strings.Contains(rurl, "from=GlobalEx") || strings.Replace(url, "bitauto", "yiche", 1) != rurl {
								fmt.Println("易车网已经删除的文章", url, col, rurl)
								f.SetCellValue("1", col, "已经删除")
							}

						} else {
							//fmt.Println("未删除的文章",url,rurl)
						}
					} else {
						fmt.Println("验证网络请求连接url=", url, "发生错误")
					}
				}
			}

			fmt.Println("合计检测记录数=", index)

		}
		if err := f.SaveAs("20211011.xlsx"); err != nil {
			fmt.Println(err)
		}
	}
}

