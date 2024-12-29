package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os/exec"
)

func PlatformLogin(c *gin.Context, platform string) (err error) {
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
