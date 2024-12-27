package service

import (
	"encoding/json"
	"fmt"
	"goodvs/server"
	"io"
	"log"
	"os"
	"os/exec"
)

type CrawlerData struct {
	Name     string  `json:"name"`
	ImgUrl   string  `json:"img_url"`
	Price    float64 `json:"price"`
	Title    string  `json:"title"`
	Category string  `json:"category"`
	Platform string  `json:"platform"`
}

// Search search
func Search(reqStr string) (resps server.SearchRes, err error) {
	// todo: implement Search
	// step1: 处理请求参数，进行分词处理
	// step2: 调用python爬虫脚本，获取搜索结果
	// 上面两步可以合并为一个步骤，调用python爬虫脚本，获取搜索结果
	fmt.Println(reqStr)
	data, err := WebCrawlingByPython(reqStr)
	if err != nil {
		return server.SearchRes{}, err
	}
	for i, v := range data {
		resp := server.ProductByCraw{
			Id:       fmt.Sprintf("%d", i),
			Name:     v.Name,
			ImgUrl:   v.ImgUrl,
			Price:    v.Price,
			Category: v.Category,
			Platform: v.Platform,
		}
		resps.Results = append(resps.Results, resp)
	}
	//fmt.Println(data)

	// step3: 如果是第一次搜索，将搜索结果(product)存入数据库

	// step4: 如果距离上次搜索时间超过一定时间，将搜索结果(product price list)存入数据库

	// step5: 返回搜索结果

	return resps, nil
}

// WebCrawlingByPython 调用Python爬虫脚本
func WebCrawlingByPython(input string) ([]CrawlerData, error) {
	cmd := exec.Command("python", "./scripts/crawler.py", input)
	fmt.Println(cmd.String())
	_, err := cmd.CombinedOutput()
	//fmt.Println(stdOutStdErr)
	if err != nil {
		fmt.Println("Error executing Python script:", err)
		return []CrawlerData{}, err
	}
	file, err := os.Open("result.json")
	if err != nil {
		fmt.Println("Error opening result.json:", err)
		return []CrawlerData{}, err
	}
	fileContents, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err) // 如果读取失败，则退出程序
	}
	//fmt.Println(string(output)) // 这里最好输出用JSON转成对象
	var data []CrawlerData
	err = json.Unmarshal(fileContents, &data)
	if err != nil {
		log.Fatal(err)
		return []CrawlerData{}, err
	}
	fmt.Println("Python script executed successfully")
	return data, nil
}

// 已弃用
//func SegmentText(text string) ([]string, error) {
//	jieba := gojieba.NewJieba()
//	defer jieba.Free()
//	// 精确模式分词
//	words := jieba.Cut(text, true)
//	fmt.Println("精确模式:", words)
//	if words == nil {
//		return nil, fmt.Errorf("jieba cut failed")
//	}
//	return words, nil
//}
