package template

import (
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

// InitreceiveString use receive string
func InitreceiveString() {
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

// ReceiveStringsCheckout use receive the *string to strings and judge
func ReceiveStringsCheckout(SecretID, SSecretID, SecretKEY, SSecretKEY *string) (string, string) {
	// receive the *strings to strings

	SecretIDS := CHEckouT(SecretID)
	SSecretIDS := CHEckouT(SSecretID)
	SecretKEYS := CHEckouT(SecretKEY)
	SSecretKEYS := CHEckouT(SSecretKEY)

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
		SecretIDS = StringRESTORE(SSecretIDS)
		SecretKEYS = StringRESTORE(SSecretKEYS)
	}
	return SecretIDS, SecretKEYS
}
