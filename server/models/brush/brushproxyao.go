package brush

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	. "github.com/fishedee/util"
	. "github.com/fishedee/web"
	"strings"
	"sync"
)

var (
	proxyXiciMutex     = &sync.Mutex{}
	proxyXiciData      = []string{}
	proxyXiciNextIndex = 1
	proxyMimvpMutex    = &sync.Mutex{}
	proxyMimvpData     = []string{}
)

type BrushProxyAoModel struct {
	Model
}

func (this *BrushProxyAoModel) refreshMimvpProxy() error {
	var data map[string]interface{}
	err := DefaultAjaxPool.Get(&Ajax{
		Url: "http://proxy.mimvp.com/api/fetch.php",
		Data: map[string]string{
			"orderid":       "billqiang@qq.com",
			"http_type":     "3",
			"result_fields": "1,2",
			"anonymous":     "5",
			"ping_time":     "1",
			"transfer_time": "5",
			"result_format": "json",
		},
		DataType:         "url",
		ResponseData:     &data,
		ResponseDataType: "json",
	})
	if err != nil {
		return nil
	}
	if data["code"] != nil && data["code"] != 0 {
		return errors.New("拉取代理失败," + data["code_msg"].(string))
	}

	proxyMimvpData = []string{}
	dataArray := data["result"].([]interface{})
	for _, singleData := range dataArray {
		singleDataObject := singleData.(map[string]interface{})
		ip := singleDataObject["ip:port"]
		protocol := singleDataObject["http_type"]
		if strings.Index(protocol.(string), "HTTPS") != -1 {
			protocol = "https"
		} else {
			protocol = "http"
		}
		proxyMimvpData = append(proxyMimvpData, fmt.Sprintf("%v://%v", protocol, ip))
	}

	return nil
}

func (this *BrushProxyAoModel) GetMimvpProxy() string {
	proxyMimvpMutex.Lock()
	defer proxyMimvpMutex.Unlock()

	if len(proxyMimvpData) == 0 {
		err := this.refreshMimvpProxy()
		if err != nil {
			panic(err)
		}
	}
	result := proxyMimvpData[0]
	proxyMimvpData = proxyMimvpData[1:]
	return result
}

func (this *BrushProxyAoModel) refreshXiciProxy() error {
	var data []byte
	err := DefaultAjaxPool.Get(&Ajax{
		Url: fmt.Sprintf("http://www.xicidaili.com/nn/%v", proxyXiciNextIndex),
		Header: map[string]string{
			"User-Agent": "User-Agent:Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.84 Safari/537.36",
		},
		ResponseData: &data,
	})
	if err != nil {
		return err
	}
	proxyXiciNextIndex++
	if proxyXiciNextIndex >= 1000 {
		proxyXiciNextIndex = 1
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err != nil {
		return err
	}

	proxyXiciData = []string{}
	doc.Find("#ip_list tr").Each(func(index int, s *goquery.Selection) {
		if index == 0 {
			return
		}
		ip := s.Find("td").Eq(1).Text()
		port := s.Find("td").Eq(2).Text()
		protocol := strings.ToLower(s.Find("td").Eq(5).Text())
		proxyXiciData = append(proxyXiciData, protocol+"://"+ip+":"+port)
	})
	return nil
}

func (this *BrushProxyAoModel) GetXiciProxy() string {
	proxyXiciMutex.Lock()
	defer proxyXiciMutex.Unlock()

	if len(proxyXiciData) == 0 {
		err := this.refreshXiciProxy()
		if err != nil {
			proxyXiciNextIndex = 1
			panic(err)
		}
	}
	result := proxyXiciData[0]
	proxyXiciData = proxyXiciData[1:]
	return result
}
