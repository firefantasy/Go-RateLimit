package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// 发送 get 请求接收 json 响应，返回解码结果
func Get(url string) (res map[string]interface{}, err error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	res = make(map[string]interface{})
	err = json.Unmarshal(body, &res)
	return
}

func main() {
	for {
		for i := 0; i < 10; i++ {
			go Get("http://127.0.0.1:8082/user/list?org=org2")
		}
		time.Sleep(time.Second)
	}
}
