package operation

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	cos "github.com/tencentyun/cos-go-sdk-v5"
)

// CosGetList get the file list from cos bucket
func CosGetList(btURL, secID, secKEY string) {

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

	//检测存储桶中APP_BACKUP内的对象别表
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
		FileStrIng = append(FileStrIng, c.Key+"     \n")
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
}

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
	num2 := strings.Count(arr[num], ".")
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

// COSUpload put the backup file to cos bucket
func COSUpload(btURL, secID, secKEY, fileNAME, fileDIR string) {
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

	//Local file
	// f := strings.NewReader(fileDIR)
	f, err := os.Open(fileDIR)
	if err != nil {
		panic(err)
	}
	// s, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(s.Size())
	_, err = c.Object.Put(context.Background(), name, f, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(fileDIR, " UPLOAD SECUSS")
	//检测存储桶中所有对象别表
	opt := &cos.BucketGetOptions{
		Prefix: "",
	}
	v, _, err := c.Bucket.Get(context.Background(), opt)
	if err != nil {
		panic(err)
	}
	//Test 用于下面判断
	var Test int
	//定义切片接收for循环中传递的c.key
	var FileStrIng []string
	//for循环将c.key的字符串apend至数组FileStrIng
	for _, c := range v.Contents {
		FileStrIng = append(FileStrIng, c.Key)
	}
	//判断刚上传的fileNAME是否在存在，存在Test为1，并退出循环
	for i := 0; i < len(FileStrIng); i++ {
		if FileStrIng[i] == name {
			Test = 1
			break
		}
	}
	if Test == 1 {
		fmt.Println(fileNAME, " is find on the bucket")
	} else {
		fmt.Println(fileNAME, " can't find on the bucket")
	}
}

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
		os.Exit(0)
	} else {
		fmt.Println(fileNAME + " delete failure")
	}

}
