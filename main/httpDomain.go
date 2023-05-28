package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
)

// func startClient() {
// 	watcher, e := fsnotify.NewWatcher()
// 	if e != nil {
// 		return
// 	}
// 	defer func(watcher *fsnotify.Watcher) {
// 		err := watcher.Close()
// 		if err != nil {
// 			return
// 		}
// 	}(watcher)
// 	done := make(chan bool)
// 	go func() {
// 		for {
// 			select {
// 			case event := <-watcher.Events:
// 				{
// 					if event.Op&fsnotify.Write == fsnotify.Write ||
// 						event.Op&fsnotify.Create == fsnotify.Create {
// 						//创建文件
// 						// 发现文件更改，上传文件
// 						//go func() {
// 						//	//// 发送文件
// 						//	//conn, err := net.Dial("tcp", serverAddr)
// 						//	//if err != nil {
// 						//	//	fmt.Println("net.Dial err = ", err)
// 						//	//	return
// 						//	//}
// 						//	//defer conn.Close()
// 						//	//// 发送文件名
// 						//	//println("event.Name = ", event.Name)
// 						//	//_, err = conn.Write([]byte(strings.Split(event.Name, "\\")[strings.Count(event.Name, "\\")]))
// 						//	//if err != nil {
// 						//	//	fmt.Println("conn.Write err = ", err)
// 						//	//	return
// 						//	//}
// 						//	//// 发送文件内容
// 						//	//file, err := os.Open(event.Name)
// 						//	//if err != nil {
// 						//	//	fmt.Println("os.Open err = ", err)
// 						//	//	return
// 						//	//}
// 						//	//defer file.Close()
// 						//	//buf := make([]byte, 1024*4)
// 						//	//for {
// 						//	//	n, err := file.Read(buf)
// 						//	//	if err != nil {
// 						//	//		if err == io.EOF {
// 						//	//			fmt.Println("文件发送完毕")
// 						//	//		} else {
// 						//	//			fmt.Println("file.Read err = ", err)
// 						//	//		}
// 						//	//		return
// 						//	//	}
// 						//	//	if n == 0 {
// 						//	//		fmt.Println("n == 0 文件发送完毕")
// 						//	//		break
// 						//	//	}
// 						//	//	_, err = conn.Write(buf[:n])
// 						//	//	if err != nil {
// 						//	//		fmt.Println("conn.Write err = ", err)
// 						//	//		return
// 						//	//	}
// 						//	//}
// 						//
// 						//}()
// 						//截取文件内容，将其发送给MiPush平台
// 						file, r := ioutil.ReadFile("./a.txt")
// 						if r != nil {
// 							fmt.Println(r.Error() + "86")
// 						}
// 						regex, e := regexp.Compile(":\\d{5}")
// 						if e != nil {
// 							fmt.Println(e.Error() + "90")
// 						}
// 						m := regex.FindString(string(file))
// 						//将匹配到的字符串发送到MiPush
// 						if m != "" {
// 							str := strings.Replace(m, ":", "", 1)
// 							_, err := os.Stat("./port")
// 							//os.Stat获取文件信息
// 							//port文件用于保存端口
// 							// 作用是为了防止多次更改导致的多次发送端口
// 							if err != nil { //判断文件是否存在
// 								if os.IsExist(err) {
// 									fmt.Println(err.Error() + "100")
// 									return
// 								}
// 							}
// 							//读取port文件
// 							s, _ := ioutil.ReadFile("./port")
// 							//判断文件端口与natapp获取到的端口是否一致
// 							if string(s) != str {
// 								//不一致，则写入新的端口
// 								_ = os.WriteFile("./port", []byte(str), 0777)
// 								//并发送端口到Mipush服务器
// 								PushMI(str)
// 								//打印端口
// 								fmt.Println(str)
// 							}
// 							continue
// 						}
// 					}

// 				}
// 			case err := <-watcher.Errors:
// 				fmt.Println(err)
// 			}

// 		}
// 	}()
// 	err := watcher.Add("./")
// 	checkError(err)
// 	<-done
// }

// // PushMI 推送函数
// type PushData struct {
// 	Source     string     `json:"source"`
// 	Appkey     string     `json:"appkey"`
// 	PushTarget PushTarget `json:"pushTarget"`
// 	PushNotify PushNotify `json:"pushNotify"`
// }
// type PushTarget struct {
// 	Target      int      `json:"target"`
// 	AppPackages []string `json:"appPackages"`
// }
// type PushNotify struct {
// 	Plats         []int         `json:"plats"`
// 	Content       string        `json:"content"`
// 	Type          int           `json:"type"`
// 	AndroidNotify AndroidNotify `json:"androidNotify"`
// }
// type AndroidNotify struct {
// 	Content []string `json:"content"`
// 	Warn    string   `json:"warn"`
// 	Style   int      `json:"style"`
// }

//	func PushMI(m string) {
//		// 设置Push推送URL
//		var urlPush = "http://api.push.mob.com/v3/push/createPush"
//		// 初始化http客户端
//		client := &http.Client{}
//		// 初始化url Body值
//		data := new(PushData)
//		data.Appkey = "367a3b20e615a"
//		data.PushNotify = PushNotify{
//			Plats:   []int{1},
//			Content: "这是Maser的域名更新通知",
//			Type:    1,
//			AndroidNotify: AndroidNotify{
//				Content: []string{m},
//				Warn:    "123",
//				Style:   1,
//			},
//		}
//		data.PushTarget = PushTarget{
//			Target:      1,
//			AppPackages: []string{"com.wusui.server"},
//		}
//		data.Source = "webapi"
//		bytes, e := json.Marshal(data)
//		checkError(e)
//		// 发送Request请求 请求方式为POST
//		req, e := http.NewRequest("POST", urlPush, strings.NewReader(string(bytes))) // 建立一个请求
//		// 检查错误函数
//		checkError(e)
//		// request请求添加Headers
//		//更改Mipush服务器的认证信息
//		req.Header.Add("key", "367a3b20e615a") //40BWA3vzT/o7xHXZiO1wWw==
//		req.Header.Add("Content-Type", "application/json")
//		req.Header.Add("sign", str2Md5(string(bytes)+"013830b877671b48cce021a48871c8b0"))
//		// 获取http响应
//		res, e := client.Do(req)
//		if e != nil {
//			fmt.Println(e.Error())
//		}
//		// 代码执行完毕后，关闭Body
//		defer func(Body io.ReadCloser) {
//			err := Body.Close()
//			checkError(err)
//		}(res.Body)
//	}
//
//	func str2Md5(str string) string {
//		return fmt.Sprintf("%x", md5.Sum([]byte(str)))
//	}
func checkError(e error) {
	if e != nil {
		log.Fatalln(e)
		return
	}

}
func main() {
	//写入文件
	//在authtoken=后填入natapp所提供的token
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
		out, _ := exec.Command("ps", "-C", "natappHttp").Output()
		out, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(out)
		if strings.Contains(string(out), "natappHttp") {
			//正则
			reg := regexp.MustCompile(`\d{3,}`)
			pid := reg.FindString(string(out))
			//结束进程
			_ = exec.Command("kill", pid).Run()
			// fmt.Println("kill natapp")
		}
		time.Sleep(time.Second * 1)
		//启动进程
		exec.Command("chmod", "777", "natappHttp").Run()
		cmd := exec.Command("./natappHttp")
		stdout, _ := cmd.StdoutPipe()
		cmd.Start()
		go func() {
			buf := make([]byte, 1024)
			for {
				n, _ := stdout.Read(buf)

				//正则匹配域名
				re := regexp.MustCompile(`http://[a-z0-9].*?.natappfree.cc`)
				domain := re.FindString(string(buf[:n]))
				if domain != "" {
					fmt.Println("域名为:" + domain)
					//发送推送消息
					// PushMI(domain)
					sendDomain(domain)

				}
			}
		}()
		time.Sleep(20 * time.Minute)
	}
}

func sendDomain(domain string) {

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
		if strings.Contains(string(out), "natappHttp.exe") {
			//读取进程pid
			cmd = exec.Command("cmd", "/c", "tasklist | findstr natappHttp.exe")
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
		cmd = exec.Command("cmd", "/c", "cd", "/d", ".\\", "&", "natappHttp.exe")
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
				re := regexp.MustCompile(`http://[a-z0-9].*?.natappfree.cc`)
				domain := re.FindString(string(buf[:n]))
				if domain != "" {
					fmt.Println("域名为:" + domain)
					//发送推送消息
					sendDomain(domain)

				}
			}
		}()
		time.Sleep(20 * time.Minute)
	}
}
