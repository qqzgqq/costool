package template

import (
	"cos_tool/operation"
	"fmt"
	"os"
)

// OPerationJudge use receive operation and run
func OPerationJudge() {

	// pd use judge
	var pd string

	// ReceiveStringsCheckout use receive the *string to strings and judge

	// receive the *strings to strings
	BucketURLS := CHEckouT(BucketURL)
	SecretIDS, SecretKEYS := ReceiveStringsCheckout(SecretID, SSecretID, SecretKEY, SSecretKEY)
	BFILENameS := CHEckouT(BFILEName)
	SYSFILEDirS := CHEckouT(SYSFILEDir)
	SRCS := CHEckouT(SRC)
	DSTS := CHEckouT(DST)
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
