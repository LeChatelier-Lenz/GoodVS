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

// ExecutePyCrawlerScript 执行Python爬虫脚本
// platform: 爬虫脚本平台
// input: 搜索关键字
// choice: 选择搜索模式
func ExecutePyCrawlerScript(platform string, input string, choice string) (interface{}, error) {
	var fileType string
	if choice == "f" {
		fileType = "result"
	} else if choice == "r" {
		fileType = "request"
	} else {
		return nil, fmt.Errorf("pyhon script choice error")
	}
	cmd := exec.Command("python", "./scripts/"+platform+"_crawler.py", choice, input)
	fmt.Println(cmd.String())
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing "+platform+" WebCrawler Python script:", err)
		return nil, err
	}
	file, err := os.Open("./tmp/" + platform + "_" + fileType + ".json")
	if err != nil {
		fmt.Println("Error opening "+platform+"_"+fileType+".json:", err)
		return nil, err
	}
	fileContents, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err) // 如果读取失败，则退出程序
		return nil, err
	}
	//fmt.Println(string(output)) // 这里最好输出用JSON转成对象
	if fileType == "result" {
		var data []server.ProductByCraw
		err = json.Unmarshal(fileContents, &data)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		return data, nil
	} else {
		var data []server.TimelyQueryReq
		err = json.Unmarshal(fileContents, &data)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		return data, nil
	}
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
	value, ok := result.([]server.ProductByCraw)
	if !ok {
		fmt.Println("Type conversion to []server.ProductByCraw failed")
		c <- nil
		close(c)
	}
	c <- value
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
	value, ok := result.([]server.ProductByCraw)
	if !ok {
		fmt.Println("Type conversion to []server.ProductByCraw failed")
		c <- nil
		close(c)
	}
	c <- value
	close(c)
}

func PlatformLogin(platform string) (err error) {
	fmt.Println("Process: platform login")
	cmd := exec.Command("python", "./scripts/"+platform+"_crawler.py", "login")
	fmt.Println(cmd.String())
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing "+platform+" WebCrawler Python script:", err)
		return err
	}
	return nil
}

// CheckByCrawler 调用爬虫脚本检查价格
// 返回值为检查后的商品列表
// 传入参数为需要检查的商品列表

func JDScriptForCheck(queries []server.TimelyQueryReq, c chan []server.TimelyQueryReq) {
	if len(queries) == 0 || queries == nil {
		c <- nil
		close(c)
		return
	}
	file, err := os.OpenFile("./tmp/jd_request.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Error opening jd_request.json:", err)
		logrus.Fatal(err)
	}
	// 将查询列表写入jd_request.json
	_, err = file.WriteString("")
	err = json.NewEncoder(file).Encode(queries)
	if err != nil {
		fmt.Println("Error writing queries to jd_request.json:", err)
		logrus.Fatal(err)
	}
	err = file.Close()
	if err != nil {
		fmt.Println("Error closing jd_request.json:", err)
		logrus.Fatal(err)
	}
	result, err := ExecutePyCrawlerScript("jd", "", "r")
	if err == nil {
		fmt.Println("JD WebCrawler Python script executed successfully")
	} else {
		fmt.Println("JD WebCrawler Python script executed failed")
		logrus.Fatal(err)
	}
	value, ok := result.([]server.TimelyQueryReq)
	if !ok {
		fmt.Println("Type conversion to []server.TimelyQueryReq failed")
		c <- nil
		close(c)
	}
	c <- value
	close(c)
}

func SNScriptForCheck(queries []server.TimelyQueryReq, c chan []server.TimelyQueryReq) {
	if len(queries) == 0 || queries == nil {
		c <- nil
		close(c)
		return
	}
	file, err := os.OpenFile("./tmp/sn_request.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Error opening sn_request.json:", err)
		logrus.Fatal(err)
	}
	// 清空sn_request.json
	_, err = file.WriteString("")
	err = json.NewEncoder(file).Encode(queries)
	if err != nil {
		fmt.Println("Error writing queries to sn_request.json:", err)
		logrus.Fatal(err)
	}
	err = file.Close()
	if err != nil {
		fmt.Println("Error closing sn_request.json:", err)
		logrus.Fatal(err)
	}
	result, err := ExecutePyCrawlerScript("sn", "", "r")
	if err == nil {
		fmt.Println("SN WebCrawler Python script executed successfully")
	} else {
		fmt.Println("SN WebCrawler Python script executed failed")
		logrus.Fatal(err)
	}
	value, ok := result.([]server.TimelyQueryReq)
	if !ok {
		fmt.Println("Type conversion to []server.TimelyQueryReq failed")
		c <- nil
		close(c)
	}
	c <- value
	close(c)
}

func CheckByCrawler(queries []server.TimelyQueryReq) (result []server.TimelyQueryReq, err error) {
	var jdData []server.TimelyQueryReq
	var snData []server.TimelyQueryReq
	for _, query := range queries {
		if query.Platform == "京东" {
			jdData = append(jdData, query)
		} else if query.Platform == "苏宁" {
			snData = append(snData, query)
		}
	}
	chJdn := make(chan []server.TimelyQueryReq, 200)
	chSnn := make(chan []server.TimelyQueryReq, 200)
	go JDScriptForCheck(jdData, chJdn)
	go SNScriptForCheck(snData, chSnn)
	dataJd := <-chJdn
	if len(dataJd) == 0 || dataJd == nil {
		if len(jdData) != 0 {
			err = fmt.Errorf("JD WebCrawler Python script executed failed\n")
		}
	}
	dataSn := <-chSnn
	if len(dataSn) == 0 || dataSn == nil {
		if len(snData) != 0 {
			err = fmt.Errorf("SN WebCrawler Python script executed failed\n")
		}
	}
	result = append(dataJd, dataSn...)
	return result, err
}
