package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
)

func main2() {
	fmt.Print("请输入要访问的网址:")
	reader := bufio.NewReader(os.Stdin)

	// 定义请求的URL
	url, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("输入错误:", err)
		return
	}

	url = strings.TrimSpace(url)

	// 自定义DNS解析器，使用本地6653端口
	dialer := &net.Dialer{
		Resolver: &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				return net.Dial("udp", "119.29.29.29:53")
			},
		},
	}

	// 创建HTTP客户端，并使用自定义的Dialer
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: dialer.DialContext,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	// 发送GET请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("创建请求时出错:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求时出错:", err)
		return
	}
	defer resp.Body.Close()

	// 检查HTTP响应状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("请求失败，状态码: %d\n", resp.StatusCode)
		return
	}

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应体时出错:", err)
		return
	}

	// 打印响应信息
	fmt.Println("响应状态:", resp.Status)
	fmt.Println("响应头:", resp.Header)
	fmt.Println("响应体:", string(body))
}
