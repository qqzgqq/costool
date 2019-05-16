package main

import (
	"cos_tool/operation"
	"cos_tool/template"
	"flag"
	"fmt"
	"os"
)

//init first start
func init() {
	template.InitreceiveString()
}

// CHEckouT *string to string
func CHEckouT(s *string) string {
	return *s
}

func main() {
	flag.Parse()
	// pd use judge
	var pd string
	// use costool -enc strings for encrypt strings
	ENCryptstringS := CHEckouT(template.ENCryptstring)
	if ENCryptstringS != "" {
		fmt.Println("++++++++++++++encrypt string+++++++++:", template.StringUPSET(ENCryptstringS))
		os.Exit(0)
	}

	// output the tool of costool help info
	template.HelpInfo(template.Helpdoc, template.Helpdoc2, template.Upload, template.Download, template.Delete, template.Move)

	// receive the strings
	BucketURLS := CHEckouT(template.BucketURL)
	var SecretIDS string = CHEckouT(template.SecretID)
	SSecretIDS := CHEckouT(template.SSecretID)
	var SecretKEYS = CHEckouT(template.SecretKEY)
	SSecretKEYS := CHEckouT(template.SSecretKEY)
	BFILENameS := CHEckouT(template.BFILEName)
	SYSFILEDirS := CHEckouT(template.SYSFILEDir)
	SRCS := CHEckouT(template.SRC)
	DSTS := CHEckouT(template.DST)

	// checkout the secretid and secretkey exist
	if SecretIDS == "" && SSecretIDS == "" {
		fmt.Println("SecretID is not find,please check `costool -h`")
		os.Exit(0)
	}
	if SecretKEYS == "" && SSecretKEYS == "" {
		fmt.Println("SecretKEY is not find,please check `costool -h`")
		os.Exit(0)
	}
	if SecretIDS != "" && SSecretIDS != "" {
		fmt.Println("SecretID and encrypt SecretID cann't use as the same time ,please check `costool -h`")
		os.Exit(0)
	}
	if SecretKEYS != "" && SSecretKEYS != "" {
		fmt.Println("SecretKEY and encrypt SecretKEY cann't use as the same time ,please check `costool -h`")
		os.Exit(0)
	}
	if SecretIDS != "" && SecretKEYS == "" || SSecretIDS != "" && SSecretKEYS == "" || SecretIDS == "" && SecretKEYS != "" || SSecretIDS == "" && SSecretKEYS != "" {
		fmt.Printf("cosupload can't use secretid and encrypt secretid as the same time!\ncosupload can't use secretkey and encrypt secretkey as the same time!")
		os.Exit(0)
	}
	if SecretIDS == "" && SecretKEYS == "" && SSecretIDS != "" && SSecretKEYS != "" {
		SecretIDS = template.StringRESTORE(SSecretIDS)
		SecretKEYS = template.StringRESTORE(SSecretKEYS)
	}

	// if upload is true
	if *template.Upload {
		//checkout the file exist
		if _, err := os.Stat(SYSFILEDirS); os.IsNotExist(err) {
			fmt.Println("the file name and dir in system is not exist")
			os.Exit(0)
		}
		operation.COSUpload(BucketURLS, SecretIDS, SecretKEYS, BFILENameS, SYSFILEDirS)
		operation.CosGetList(BucketURLS, SecretIDS, SecretKEYS)
		os.Exit(0)
	}

	// if download is true
	if *template.Download {
		if *template.GetList {
			operation.CosGetList(BucketURLS, SecretIDS, SecretKEYS)
			fmt.Println("please cat the file name as bucketdirlist.log in this dir")
			os.Exit(0)
		}
		if BFILENameS == "" {
			fmt.Println("the filename will be download is not find,please check `cosdownload -h`")
			os.Exit(0)
		}
		operation.CosDownLoad(BucketURLS, SecretIDS, SecretKEYS, BFILENameS)
		operation.CosGetList(BucketURLS, SecretIDS, SecretKEYS)
		os.Exit(0)
	}

	// if delete is true
	if *template.Delete {
		//checkout the file exist
		if operation.COSCheckoutfile(BucketURLS, SecretIDS, SecretKEYS, BFILENameS) == false {
			fmt.Println(BFILENameS + " can't find on the bucket")
			os.Exit(0)
		}
		fmt.Printf("Are you sure you want to delete " + BFILENameS + " ? please input  Y/N:")
		fmt.Scanln(&pd)
		if pd == "y" || pd == "Y" {
			operation.COSDelete(BucketURLS, SecretIDS, SecretKEYS, BFILENameS)
			operation.CosGetList(BucketURLS, SecretIDS, SecretKEYS)
		} else {
			os.Exit(0)
		}

	}

	if *template.Move {
		if SRCS == "" || DSTS == "" {
			fmt.Printf("the source or dest can't be empty")
			os.Exit(0)
		}
		if operation.COSCheckoutfile(BucketURLS, SecretIDS, SecretKEYS, SRCS) == false {
			fmt.Println(SRCS + " can't find on the bucket")
			os.Exit(0)
		}
		if operation.COSCheckoutfile(BucketURLS, SecretIDS, SecretKEYS, DSTS) {
			fmt.Println(DSTS + " is exist on the bucket")
			fmt.Printf("Are you sure you want to cover " + DSTS + " ? please input  Y/N:")
			fmt.Scanln(&pd)
			if pd == "y" || pd == "Y" {
				operation.COSDelete(BucketURLS, SecretIDS, SecretKEYS, DSTS)
				operation.COSCopy(BucketURLS, SecretIDS, SecretKEYS, SRCS, DSTS)
				operation.CosGetList(BucketURLS, SecretIDS, SecretKEYS)
			} else {
				operation.CosGetList(BucketURLS, SecretIDS, SecretKEYS)
				os.Exit(0)
			}
		} else {
			operation.COSCopy(BucketURLS, SecretIDS, SecretKEYS, SRCS, DSTS)
			operation.CosGetList(BucketURLS, SecretIDS, SecretKEYS)
		}

	}

}
