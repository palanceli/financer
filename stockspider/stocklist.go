package stockspider

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang/glog"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type StockType int32

const (
	StockType_CNSZ = iota
	StockType_CNSH
	StockType_US
)

type StockInfo struct {
	Name   string
	Symbol string
	SType  StockType
}

func decodeToGBK(text string) (string, error) {

	dst := make([]byte, len(text)*2)
	tr := simplifiedchinese.GB18030.NewDecoder()
	nDst, _, err := tr.Transform(dst, []byte(text), true)
	if err != nil {
		return text, err
	}

	return string(dst[:nDst]), nil
}

func parseFromEastMoney(url string, fnStockType func(symbol string) StockType) []StockInfo {
	var result []StockInfo
	glog.Infof("Querying %s", url)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		glog.Fatalf("FAILED to query %s", url)
	}
	doc.Find("#quotesearch").Find("li").Each(func(i int, li *goquery.Selection) {
		liText := li.Text()
		stockName := liText[:strings.Index(liText, "(")]
		stockSymbol := liText[strings.Index(liText, "(")+1 : strings.Index(liText, ")")]
		stockType := fnStockType(stockSymbol)
		if strings.Contains(liText, "...") {
			a := li.Find("a")
			if a != nil {
				title, exists := a.Attr("title")
				if exists {
					stockName = title
				}
			}
		}
		stockNameGBK, _ := decodeToGBK(stockName)
		// glog.Infof("%s, %s", stockNameGBK, stockSymbol)
		result = append(result, StockInfo{stockNameGBK, stockSymbol, stockType})
	})
	return result
}

// ListAllStocks 列出所有股票名称、代号
func ListAllStocks() []StockInfo {
	cnStocks := parseFromEastMoney("http://quote.eastmoney.com/stocklist.html", func(symbol string) StockType {
		if symbol[0] == '6' {
			return StockType_CNSH
		}
		return StockType_CNSZ
	})
	usStock := parseFromEastMoney("http://quote.eastmoney.com/usstocklist.html", func(symbol string) StockType {
		return StockType_US
	})
	result := cnStocks
	result = append(result, usStock...)
	return result
}
