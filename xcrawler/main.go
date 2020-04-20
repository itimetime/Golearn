package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var (
	reEmail = `[0-9]+?@[0-9A-z]+\.com`
	reHref  = `href="(http[s]?://[\S]+?)"`
	reImage = `src="(http[s]?://[\S]+.(jpg|jpeg))"`
)

var chanImageUrls chan string
var wg sync.WaitGroup
var chanTask chan string

func myTest(url string) {
	//pageStr := getPage(url)
}

func downloadFile(url, filename string) (ok bool) {
	resp, err := http.Get(url)
	handError(err, "open image url failed!")
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	handError(err, "read image url failed!")
	filename = "D:\\go\\src\\golearnProject\\xcrawler\\image\\" + filename
	err = ioutil.WriteFile(filename, bytes, 0644)
	if err != nil {
		return false
	}
	return true
}

func getPage(url string) (pageStr string) {
	resp, err := http.Get(url)
	handError(err, "http get 错误")
	defer resp.Body.Close()
	pageBytes, err := ioutil.ReadAll(resp.Body)
	handError(err, "open file err!")
	pageStr = string(pageBytes)
	return
}

func reMatch(url string, rule string) [][]string {
	pageStr := getPage(url)
	re := regexp.MustCompile(rule)
	results := re.FindAllStringSubmatch(pageStr, -1) //写-1 取全部
	return results
}

func getFileName(url string) string {
	lastIndex := strings.LastIndex(url, "/")
	return url[lastIndex+1:]
}

func crawlers(url string, rule string) {
	pageStr := getPage(url)
	re := regexp.MustCompile(rule)
	results := re.FindAllStringSubmatch(pageStr, -1) //写-1 取全部
	for _, url := range results {
		fileName := getFileName(url[1])
		downloadFile(url[1], fileName)
	}
	wg.Done()
}

func handError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}

}

func main() {
	chanImageUrls = make(chan string, 100)
	chanTask = make(chan string, 30)

	for i := 1; i < 27; i++ {
		wg.Add(1)
		go crawlers("https://www.umei.cc/p/gaoqing/cn/"+strconv.Itoa(i)+".htm", reImage)
	}
	wg.Wait()

	//matchQQEmail("https://tieba.baidu.com/p/5535170914?red_tag=2341465931")

}
