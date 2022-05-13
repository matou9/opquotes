package quote

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/page"
	"io/ioutil"

	//"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
	"time"
)

type Scrawl struct{
	JsPath string
}
func NewScrawl()*Scrawl{
	return &Scrawl{JsPath:"d:\\s.js"}
}
func (s *Scrawl)GetDomFromRequest(url string)*goquery.Document{
	if err,content,_:=s.Request(url,true);err==nil {
		dom, err := goquery.NewDocumentFromReader(strings.NewReader(content))
		if err != nil {
			log.Fatal(err.Error())
		} else {
			return dom
		}
	}
	return nil
}
func (s *Scrawl)Request(url string,isexecjs bool)(err error,content string,rurl string){
	var jsstr string
	var newURL string
	if isexecjs==true{
		b, err := ioutil.ReadFile("d:\\s.js") // just pass the file name
		if err != nil {
			fmt.Print(err)
		}
		jsstr = string(b)
	}

	opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", true),chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.Flag("no-default-browser-check", true),chromedp.Flag("no-sandbox", false),chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 20*time.Second)
	defer cancel()
	var res string
	if isexecjs==true{
		if err := chromedp.Run(ctx,
			chromedp.ActionFunc(func(ctx context.Context) error {
				_, err := page.AddScriptToEvaluateOnNewDocument(jsstr).Do(ctx)
				if err != nil {
					return err
				}
				return nil
			}),
			chromedp.Navigate(url),
			chromedp.Evaluate(`window.location.href`, &newURL),
			chromedp.Sleep(1 * time.Second),
			chromedp.OuterHTML("html",&res),
		); err != nil {
			fmt.Println("出现异常错误" + err.Error())
			return err,"",url

		}
	}else{
		if err := chromedp.Run(ctx,
			chromedp.Navigate(url),
			chromedp.Sleep(2 * time.Second),
			chromedp.OuterHTML("html",&res),
		); err != nil {
			fmt.Println("出现异常错误" + err.Error())
			return err,"",url

		}
	}

	return nil ,res,newURL
}
