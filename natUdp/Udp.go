package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
)

func main() {

	_ = os.WriteFile("./config.ini",
		[]byte(`[default]
logto=stdout
authtoken=xxx
loglevel=INFO`), 0777)
	if runtime.GOOS == "windows" {
		natAppWindows()
	} else {
		natAppLinux()
	}

}

func natAppLinux() {
	for {
		out, _ := exec.Command("ps", "-C", "natappUdp").Output()
		out, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(out)
		if strings.Contains(string(out), "natappUdp") {
			//正则
			reg := regexp.MustCompile(`\d{3,}`)
			pid := reg.FindString(string(out))
			//结束进程
			_ = exec.Command("kill", pid).Run()
			// fmt.Println("kill natapp")
		}
		time.Sleep(time.Second * 1)
		//启动进程
		exec.Command("chmod", "777", "natappUdp").Run()
		cmd := exec.Command("./natappUdp")
		stdout, _ := cmd.StdoutPipe()
		cmd.Start()
		go func() {
			buf := make([]byte, 1024)
			for {
				n, _ := stdout.Read(buf)

				//正则匹配域名
				re := regexp.MustCompile(`cc:\d{4,5}`)
				port := re.FindString(string(buf[:n]))
				if port != "" {
					fmt.Println("域名为:" + strings.Replace(port, "cc:", "", -1))
					//发送推送消息
					// PushMI(domain)
					sendUdp(strings.Replace(port, "cc:", "", -1))
				}
			}
		}()
		time.Sleep(20 * time.Minute)
	}
}

func sendUdp(port string) {
	var url = "http://api.liusui.xyz/1/classes/Url/TgA8666j"
	client := &http.Client{}
	type UdpPort struct {
		UdpPort string `json:"udpPort"`
	}
	type RespData struct {
		UpdatedAt string `json:"updatedAt"`
	}
	var updatedAt = RespData{}
	var udpPort = UdpPort{
		UdpPort: port,
	}
	bytes, _ := json.Marshal(udpPort)
	req, e := http.NewRequest("PUT", url, strings.NewReader(string(bytes))) // 建立一个请求
	if e != nil {
		fmt.Println(e.Error())
	}
	req.Header.Add("X-Bmob-Application-Id", "0bdcb075b0490f37d55593302851da5d")
	req.Header.Add("X-Bmob-REST-API-Key", "eed1a0509dfafce55a1e7b35ba3e10ca")
	req.Header.Add("Content-Type", "application/json")
	res, e := client.Do(req)
	if e != nil {
		fmt.Println(e.Error())
	}
	by, _ := io.ReadAll(res.Body)
	_ = json.Unmarshal(by, &updatedAt)
	fmt.Println(updatedAt)
	// 代码执行完毕后，关闭Body
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)
}

func natAppWindows() {
	//判断程序是否已经运行
	for {
		cmd := exec.Command("cmd", "/c", "tasklist")
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
			return
		}
		out, err = simplifiedchinese.GB18030.NewDecoder().Bytes(out)
		if err != nil {
			panic(err)
		}
		if strings.Contains(string(out), "natappUdp.exe") {
			//读取进程pid
			cmd = exec.Command("cmd", "/c", "tasklist | findstr natappUdp.exe")
			out, err = cmd.Output()
			if err != nil {
				fmt.Println(err)
				return
			}
			out, err = simplifiedchinese.GB18030.NewDecoder().Bytes(out)
			if err != nil {
				panic(err)
			}
			//获取进程pid
			//正则
			reg := regexp.MustCompile(`\d+ C`)
			pid := reg.FindAllString(string(out), -1)
			for i := 0; i < len(pid); i++ {
				pid[i] = strings.Replace(pid[i], " C", "", -1)
				//结束进程
				cmd = exec.Command("cmd", "/c", "taskkill /f /pid "+pid[i])
				cmd.Run()
			}
			fmt.Println("结束所有NatApp进程")
		}

		//等待100ms
		time.Sleep(100 * time.Millisecond)
		//启动程序
		cmd = exec.Command("cmd", "/c", "cd", "/d", ".\\", "&", "natappUdp.exe")
		//输出程序的输出信息
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Println(err)
			return
		}

		//启动程序
		err = cmd.Start()
		if err != nil {
			fmt.Println(err)
			return
		}
		//读取程序的输出信息
		go func() {
			buf := make([]byte, 1024)
			for {
				n, err := stdout.Read(buf)
				if err != nil {
					fmt.Println(err)
					return
				}
				//正则匹配域名
				re := regexp.MustCompile(`cc:\d{4,5}`)
				port := re.FindString(string(buf[:n]))
				if port != "" {
					fmt.Println("域名为:" + strings.Replace(port, "cc:", "", -1))
					//发送推送消息
					sendUdp(strings.Replace(port, "cc:", "", -1))
				}
			}
		}()
		time.Sleep(20 * time.Minute)
	}
}
