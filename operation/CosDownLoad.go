package operation

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	cos "github.com/tencentyun/cos-go-sdk-v5"
)

// CosDownLoad download the file from cos bucket
func CosDownLoad(btURL, secID, secKEY, fileName string) {

	//将<bucketname>、<appid>和<region>修改为真实的信息
	//例如：http://test-1253846586.cos.ap-guangzhou.myqcloud.com
	u, _ := url.Parse(btURL)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			//如实填写账号和密钥，也可以设置为环境变量
			SecretID:  secID,
			SecretKey: secKEY,
		},
	})

	//检测存储桶中APP_BACKUP内的对象列表
	opt := &cos.BucketGetOptions{
		Prefix: "",
	}
	v, _, err := c.Bucket.Get(context.Background(), opt)
	if err != nil {
		panic(err)
	}
	//定义切片接收for循环中传递的c.key
	var FileStrIng []string
	//for循环将c.key的字符串apend至数组FileStrIng
	for _, c := range v.Contents {
		FileStrIng = append(FileStrIng, c.Key+"\n")
	}
	//数组转字符串
	FileStrIngS := strings.Join(FileStrIng, "")

	//生成本地文件
	ff, err := os.Create("bucketdirlist.log")
	defer ff.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = ff.Write([]byte(FileStrIngS))
		if err != nil {
			panic(err)
		}
	}
	resp, err := c.Object.Get(context.Background(), fileName, nil)
	if err != nil {
		panic(err)
	}
	//字符串截取下载文件名
	num := strings.Count(fileName, "/")
	arr := strings.Split(fileName, "/")
	//新增创建下载文件夹
	var num2 = strings.Count(arr[num], ".")
	if num2 == 0 {
		num2 = 1
	}
	arr2 := strings.Split(arr[num], ".")
	_ = os.Mkdir(arr2[num2-1], 777)
	file, err := os.Create(arr2[num2-1] + "/" + arr[num])
	if err != nil {
		panic(err)
	}
	io.Copy(file, resp.Body)
	resp.Body.Close()
	fmt.Println(arr[num], " download succes")
}
