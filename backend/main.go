package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/anyinone/go-zero/core/logx"
	"github.com/anyinone/go-zero/core/service"
	"github.com/anyinone/go-zero/rest"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"smartpro/core"
	"strings"
)

type User struct {
	DataName             string `json:"data_name"`
	Identifier           string `json:"identifier"`
	Author               string `json:"author"`
	ProjectName          string `json:"project_name"`
	DataSize             string `json:"data_size"`
	Version              string `json:"version"`
	DownloadURL          string `json:"download_url"`
	DataSource           string `json:"data_source"`
	CreationOrganization string `json:"creation_organization"`
}

type Experiment struct {
	ExperimentId string `json:"experiment_id"`
}

func main() {
	srv, err := rest.NewServer(rest.RestConf{
		Port: 8080, // 侦听端口
		ServiceConf: service.ServiceConf{
			Log: logx.LogConf{Path: "./logs"}, // 日志路径
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer srv.Stop()
	// 注册路由
	srv.AddRoute(rest.Route{
		Method:  http.MethodPost,
		Path:    "/api/user",
		Handler: submit,
	})
	srv.AddRoute(rest.Route{
		Method:  http.MethodPost,
		Path:    "/api/get",
		Handler: get,
	})

	srv.Start() // 启动服务
}

func submit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// 添加 CORS 头
		headers := w.Header()
		headers.Set("Access-Control-Allow-Origin", "*")
		headers.Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		headers.Set("Access-Control-Allow-Headers", "Content-Type")
		err := r.ParseMultipartForm(10 << 20) // 最大文件大小设置为 10MB
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// 解析其他表单字段
		userJson := r.FormValue("data")
		var user User
		err = json.Unmarshal([]byte(userJson), &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// 解析上传的文件
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		h := sha256.New()

		// 将文件内容写入 hash 对象
		if _, err := io.Copy(h, file); err != nil {
			fmt.Println(err)
		}

		// 计算 hash 值并转换为字符串返回
		hash := fmt.Sprintf("%x", h.Sum(nil))
		fmt.Println("哈希", hash)
		defer file.Close()
		// 处理文件，这里直接输出文件名
		fmt.Printf("Received file: %v\n", handler.Filename)
		// 在这里处理您的逻辑
		fmt.Printf("Received data: %+v\n", user)
		// 返回响应

		// 调用智能合约函数
		logsStr := core.Smartcontract(user.DataName, user.Identifier, user.Author, user.ProjectName, user.DataSize, user.Version, user.DownloadURL, user.CreationOrganization, hash)
		fmt.Printf(logsStr)
		logs := strings.Split(logsStr, " ")
		experimentId := logs[2]
		//transactionId := logs[3]
		//txHash := logs[6]

		response := map[string]interface{}{
			//	"message": "success\n" + "experimentid=" + experimentId + "\n" + "transactionid=" + transactionId + "\n" + txHash + "\n",
			"message": "success\n" + "experimentid=" + experimentId + "\n",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else if r.Method == "OPTIONS" {
		// 添加 CORS 头
		headers := w.Header()
		headers.Set("Access-Control-Allow-Origin", "*")
		headers.Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		headers.Set("Access-Control-Allow-Headers", "Content-Type")
	}

	data, err := ioutil.ReadFile("path/to/file")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	hash := sha256.Sum256(data)
	hashString := hex.EncodeToString(hash[:])

	fmt.Println("File hash:", hashString)

}
func get(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// 添加 CORS 头
		headers := w.Header()
		headers.Set("Access-Control-Allow-Origin", "*")
		headers.Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		headers.Set("Access-Control-Allow-Headers", "Content-Type")
		// 解析其他表单字段
		userJson := r.FormValue("data")
		var experiment Experiment
		err := json.Unmarshal([]byte(userJson), &experiment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println("数据数据数据", experiment.ExperimentId)
		// 调用智能合约函数
		data := core.GetExperiment(experiment.ExperimentId)

		response := map[string]interface{}{
			//	"message": "success\n" + "experimentid=" + experimentId + "\n" + "transactionid=" + transactionId + "\n" + txHash + "\n",

			"message": "数据名称：" + data.DataName + "\n" + "    标识代码：" + data.IdentificationCode + "\n" + "   实验人：" + data.Author + "\n" + "   项目名称：" + data.ProjectName + "\n" + "   数据量：" + data.DataVolume + "\n" + "   版本：" + data.Version + "\n" + "   下载地址：" + data.DataTraceability + "\n" + "   实验日期：" + data.Timestamp + "\n" + "   创建机构：" + data.CreationOrganization + "\n" + "   文件hash值：" + data.DataFileHash + "\n",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else if r.Method == "OPTIONS" {
		// 添加 CORS 头
		headers := w.Header()
		headers.Set("Access-Control-Allow-Origin", "*")
		headers.Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		headers.Set("Access-Control-Allow-Headers", "Content-Type")
	}
}
