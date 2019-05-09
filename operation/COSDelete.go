package operation

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	cos "github.com/tencentyun/cos-go-sdk-v5"
)

// COSDelete delete the filename on bucket
func COSDelete(btURL, secID, secKEY, fileNAME string) {
	u, _ := url.Parse(btURL)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			//如实填写账号和密钥，也可以设置为环境变量
			SecretID:  secID,
			SecretKey: secKEY,
		},
	})

	_, err1 := c.Object.Delete(context.Background(), fileNAME)
	if err1 != nil {
		panic(err1)
	}

	_, err2 := c.Object.Get(context.Background(), fileNAME, nil)
	if err2 != nil {
		fmt.Println(fileNAME + " delete sucess")
	} else {
		fmt.Println(fileNAME + " delete failure")
	}

}
