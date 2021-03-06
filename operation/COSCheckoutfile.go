package operation

import (
	"context"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	cos "github.com/tencentyun/cos-go-sdk-v5"
)

// COSCheckoutfile checkout the file on buket exist
func COSCheckoutfile(btURL, secID, secKEY, fileNAME string) bool {
	// Test false is not find the filename on buket
	var Test bool
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

	//对象键（Key）是对象在存储桶中的唯一标识。
	var name string
	if m, _ := regexp.MatchString("^\\/", fileNAME); m {
		FishuzU := strings.SplitAfterN(fileNAME, "", 2)
		name = FishuzU[1]
	} else {
		name = fileNAME
	}

	//检测存储桶中所有对象别表
	opt := &cos.BucketGetOptions{
		Prefix: "",
	}
	v, _, err := c.Bucket.Get(context.Background(), opt)
	if err != nil {
		panic(err)
	}
	//Test 用于下面判断

	//定义切片接收for循环中传递的c.key
	var FileStrIng []string
	//for循环将c.key的字符串apend至数组FileStrIng
	for _, c := range v.Contents {
		FileStrIng = append(FileStrIng, c.Key)
	}
	//判断刚上传的fileNAME是否在存在，存在Test为1，并退出循环
	for i := 0; i < len(FileStrIng); i++ {
		if FileStrIng[i] == name {
			Test = true
			break
		}
	}
	return Test
}
