package brush

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	. "github.com/fishedee/util"
	. "github.com/fishedee/web"
	"strings"
	"sync"
)

var (
	proxyMutex     = &sync.Mutex{}
	proxyData      = []string{}
	proxyNextIndex = 1
)

type BrushProxyAoModel struct {
	Model
}

func (this *BrushProxyAoModel) refreshXiciProxy() error {
	var data []byte
	err := DefaultAjaxPool.Get(&Ajax{
		Url: fmt.Sprintf("http://www.xicidaili.com/nn/%v", proxyNextIndex),
		Header: map[string]string{
			"User-Agent": "User-Agent:Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.84 Safari/537.36",
		},
		ResponseData: &data,
	})
	if err != nil {
		return err
	}
	proxyNextIndex++
	if proxyNextIndex >= 1000 {
		proxyNextIndex = 1
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err != nil {
		return err
	}

	proxyData = []string{}
	doc.Find("#ip_list tr").Each(func(index int, s *goquery.Selection) {
		if index == 0 {
			return
		}
		ip := s.Find("td").Eq(1).Text()
		port := s.Find("td").Eq(2).Text()
		protocol := strings.ToLower(s.Find("td").Eq(5).Text())
		proxyData = append(proxyData, protocol+"://"+ip+":"+port)
	})
	return nil
}

func (this *BrushProxyAoModel) GetXiciProxy() string {
	proxyMutex.Lock()
	defer proxyMutex.Unlock()

	if len(proxyData) == 0 {
		err := this.refreshXiciProxy()
		if err != nil {
			proxyNextIndex = 1
			panic(err)
		}
	}
	result := proxyData[0]
	proxyData = proxyData[1:]
	return result
}
