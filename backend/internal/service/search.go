package service

import (
	"fmt"
	"goodvs/server"
	"os/exec"
)

func Search(req server.SearchReq) (server.SearchRes, error) {
	// todo: implement Search
	// step1: 处理请求参数，进行分词处理
	// step2: 调用python爬虫脚本，获取搜索结果
	// 上面两步可以合并为一个步骤，调用python爬虫脚本，获取搜索结果
	WebCrawlingByPython(req.productStr)
	// step3: 如果是第一次搜索，将搜索结果(product)存入数据库

	// step4: 如果距离上次搜索时间超过一定时间，将搜索结果(product price list)存入数据库

	// step5: 返回搜索结果
	return server.SearchRes{}, nil
}

func WebCrawlingByPython(input string) {
	cmd := exec.Command("python", "crawler.py", input)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing Python script:", err)
		return
	}
	fmt.Println(string(output)) // 这里最好输出用JSON转成对象
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
