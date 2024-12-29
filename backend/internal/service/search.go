package service

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"goodvs/server"
	"io"
	"log"
	"os"
	"os/exec"
)

// SearchCallByFrontend search
func SearchCallByFrontend(reqStr string) (data []server.ProductByCraw, err error) {
	return Search(reqStr, 0)
}

// Search search
func Search(reqStr string, opt int) (data []server.ProductByCraw, err error) {
	// implement Search
	// step1: 处理请求参数，进行分词处理
	// step2: 调用python爬虫脚本，获取搜索结果
	// 上面两步可以合并为一个步骤，调用python爬虫脚本，获取搜索结果
	//fmt.Println(reqStr)
	var choice string
	if opt == 0 {
		choice = "f"
	} else {
		choice = "r"
	}
	result, err := WebCrawlingByPython(reqStr, choice)
	if err != nil {
		return nil, err
	}
	var category = result[0].Category // 暂时只取第一个商品的category
	for _, v := range result {
		r := server.ProductByCraw{
			Id:       v.Id,
			Url:      v.Url,
			Name:     v.Name,
			ImgUrl:   v.ImgUrl,
			Price:    v.Price,
			Title:    v.Title,
			Category: category,
			Platform: v.Platform,
		}
		data = append(data, r)
	}
	return data, nil
}

// WebCrawlingByPython 调用Python爬虫脚本
// 利用goroutine并发调用JD和SN的爬虫脚本
func WebCrawlingByPython(input string, choice string) (result []server.ProductByCraw, err error) {
	chJd := make(chan []server.ProductByCraw, 200)
	chSn := make(chan []server.ProductByCraw, 200)
	go JDScript(input, choice, chJd)
	go SNScript(input, choice, chSn)
	dataJd := <-chJd
	if len(dataJd) == 0 || dataJd == nil {
		err = fmt.Errorf("JD WebCrawler Python script executed failed")
	}
	dataSn := <-chSn
	if len(dataSn) == 0 || dataSn == nil {
		err = fmt.Errorf("SN WebCrawler Python script executed failed")
	}
	result = append(dataJd, dataSn...)
	return result, err
}

// JDScript 调用JD爬虫脚本
func JDScript(input string, choice string, c chan []server.ProductByCraw) {
	result, err := ExecutePyCrawlerScript("jd", input, choice)
	if err == nil {
		fmt.Println("JD WebCrawler Python script executed successfully")
	} else {
		fmt.Println("JD WebCrawler Python script executed failed")
		logrus.Fatal(err)
	}
	c <- result
	close(c)
}

// SNScript 调用SN爬虫脚本
func SNScript(input string, choice string, c chan []server.ProductByCraw) {
	result, err := ExecutePyCrawlerScript("sn", input, choice)
	if err == nil {
		fmt.Println("SN WebCrawler Python script executed successfully")
	} else {
		fmt.Println("SN WebCrawler Python script executed failed")
		logrus.Fatal(err)
	}
	c <- result
	close(c)
}

// ExecutePyCrawlerScript 执行Python爬虫脚本
// platform: 爬虫脚本平台
// input: 搜索关键字
// choice: 选择搜索模式
func ExecutePyCrawlerScript(platform string, input string, choice string) ([]server.ProductByCraw, error) {
	cmd := exec.Command("python", "./scripts/"+platform+"_crawler.py", choice, input)
	fmt.Println(cmd.String())
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing "+platform+" WebCrawler Python script:", err)
		return nil, err
	}
	file, err := os.Open("./tmp/" + platform + "_result.json")
	if err != nil {
		fmt.Println("Error opening "+platform+"_result.json:", err)
		return nil, err
	}
	fileContents, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err) // 如果读取失败，则退出程序
		return nil, err
	}
	//fmt.Println(string(output)) // 这里最好输出用JSON转成对象
	var data []server.ProductByCraw
	err = json.Unmarshal(fileContents, &data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return data, nil
}
