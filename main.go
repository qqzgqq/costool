package main

import (
	"cos_tool/operation"
	"cos_tool/template"
	"flag"
	"fmt"
	"os"
)

// receive is usege scanln user input info
var receive string

// Upload if Upload is true then upload to cos
var Upload *bool

// Download if  Download is true then Download from cos
var Download *bool

// Delete if  Delete is true then Delete from cos
var Delete *bool

// Move if Move is true then move to cos dir
var Move *bool

// GetList get file list from cos bucket
var GetList *bool

// Helpdoc if input -h or --h  then print helpdoc
var Helpdoc *bool

// Helpdoc2 if input -help or --help then print helpdoc
var Helpdoc2 *bool

// BucketURL use receive bucket_url
var BucketURL *string

// SecretID use receive secretid
var SecretID *string

// SecretKEY use receive secretkey
var SecretKEY *string

// SSecretID use encrypt secretid
var SSecretID *string

// SSecretKEY use encrypt secrekey
var SSecretKEY *string

// ENCryptstring use encrypt string
var ENCryptstring *string

// SRC use the bucket file name as source
var SRC *string

// DST use the bucket file name as dest
var DST *string

// BFILEName the file name in bucket
var BFILEName *string

// SYSFILEDir the file name and dir in system
var SYSFILEDir *string

//init fast start
func init() {
	BucketURL = flag.String("url", receive, "bucket_url")
	Helpdoc = flag.Bool("h", false, "-h for helpdoc")
	Helpdoc2 = flag.Bool("help", false, "--help for helpdoc")
	Upload = flag.Bool("upload", false, "upload to cos")
	Download = flag.Bool("download", false, "download from cos")
	Delete = flag.Bool("delete", false, "delete from cos")
	GetList = flag.Bool("glt", false, "get file list from cos bucket")
	Move = flag.Bool("move", false, "move to cos dir")
	SecretID = flag.String("sd", receive, "secretid")
	SecretKEY = flag.String("sk", receive, "secretkey")
	SSecretID = flag.String("sds", receive, "encrypt secretid")
	SSecretKEY = flag.String("sks", receive, "encrypt secretkey")
	BFILEName = flag.String("bfn", receive, "the file full dir and file name on bucket")
	SYSFILEDir = flag.String("fdir", receive, "the file dir and name in system")
	ENCryptstring = flag.String("enc", receive, "encrypt string")
	SRC = flag.String("src", receive, "the bucket file name as source")
	DST = flag.String("dst", receive, "the bucket file name as dest")
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
	ENCryptstringS := CHEckouT(ENCryptstring)
	if ENCryptstringS != "" {
		fmt.Println("++++++++++++++encrypt string+++++++++:", template.StringUPSET(ENCryptstringS))
		os.Exit(0)
	}

	// output the tool of costool help info
	template.HelpInfo(Helpdoc, Helpdoc2, Upload, Download, Delete, Move)

	// receive the strings
	BucketURLS := CHEckouT(BucketURL)
	var SecretIDS string = CHEckouT(SecretID)
	SSecretIDS := CHEckouT(SSecretID)
	var SecretKEYS string = CHEckouT(SecretKEY)
	SSecretKEYS := CHEckouT(SSecretKEY)
	BFILENameS := CHEckouT(BFILEName)
	SYSFILEDirS := CHEckouT(SYSFILEDir)
	SRCS := CHEckouT(SRC)
	DSTS := CHEckouT(DST)

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
	if *Upload {
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
	if *Download {
		if *GetList {
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
	if *Delete {
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

	if *Move {
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
