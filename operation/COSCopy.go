package operation

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	cos "github.com/tencentyun/cos-go-sdk-v5"
)

// COSCopy use copy the source to dest on buket
func COSCopy(btURL, secID, secKEY, sourCe, deSt string) {
	u, _ := url.Parse(btURL)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			//如实填写账号和密钥，也可以设置为环境变量
			SecretID:  secID,
			SecretKey: secKEY,
		},
	})
	var deStS string
	soruceURL := fmt.Sprintf("%s/%s", u.Host, sourCe)
	if strings.IndexAny(deSt, "/") != 0 {
		deStS = "/" + deSt
	}
	_, _, err := c.Object.Copy(context.Background(), deStS, soruceURL, nil)
	if err != nil {
		panic(err)
	}

	_, err2 := c.Object.Get(context.Background(), deSt, nil)
	if err2 != nil {
		fmt.Println(deSt + " move failure")
		panic(err2)
	} else {
		fmt.Println(deSt + " move sucess")
		_, err1 := c.Object.Delete(context.Background(), sourCe)
		if err1 != nil {
			panic(err1)
		}
	}

}
