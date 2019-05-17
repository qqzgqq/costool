package template

import (
	"fmt"
	"os"
)

// HelpInfo cosupload-v1 helpinfo
func HelpInfo(s, s2, s3, s4, s5, s6 *bool) {
	// var receive2 = flag.Arg(0)
	var oper = false
	if *s3 || *s4 || *s5 || *s6 {
		oper = true
	}
	if *s || *s2 || oper == false {
		fmt.Println("usage: costool [OPERATION] -url $BucketURL -sd $SecretID -sk $SecretKEY [PARAMETER]")
		fmt.Printf("       costool [OPERATION] -url $BucketURL -sds $SSecretID -sks $SSecretKEY [PARAMETER]\n\n")
		fmt.Println("OPERATION:")
		fmt.Printf("%10s  %s\n", "-upload   ", "Transfer the files to the cos bucket")
		fmt.Printf("%10s  %s\n", "-download ", "download the files from the cos bucket")
		fmt.Printf("%10s  %s\n", "-move     ", "Moves files around in buckets,can use rename")
		fmt.Printf("%10s  %s\n\n", "-delete   ", "delete the files from the cos bucket")
		fmt.Println("PARAMETER:")
		fmt.Printf("%5s  %s\n", "-url ", "bucket_url")
		fmt.Printf("%5s  %s\n", "-sd  ", "secretid")
		fmt.Printf("%5s  %s\n", "-sk  ", "secretkey")
		fmt.Printf("%5s  %s\n", "-sds ", "receive encrypt secretid")
		fmt.Printf("%5s  %s\n", "-sks ", "receive encrypt secretkey")
		fmt.Printf("%5s  %s\n", "-src ", "the source full dir and file name on bukect")
		fmt.Printf("%5s  %s\n", "-dst ", "the dest full dir and file name on bukect")
		fmt.Printf("%5s  %s\n", "-bfn ", "the file full dir and file name on bucket")
		fmt.Printf("%5s  %s\n", "-enc ", "encrypt the string will use as encrypt secretid or secretkey ")
		fmt.Printf("%5s  %s\n\n", "-fdir", "the upload file dir and file name in the system")
		fmt.Printf("%5s\n", "OPERATION EXAMPLE:")
		fmt.Printf("%5s\n", "costool upload operation eg1:")
		fmt.Println("        costool -upload -url https://********* -sd ****** -sk ****** -bfn APP_BACKUP/test/test.tar.gz -fdir /usr/local/test.tar.gz")
		fmt.Printf("%5s\n", "costool upload operationeg 2:")
		fmt.Println("        ssecretid=`costool -enc secretid`; ssecretkey=`costool -enc secretkey` ")
		fmt.Printf("        costool -upload -url https://********* -sds $ssecretid -sks $ssecretkey -bfn APP_BACKUP/test/test.tar.gz -fdir /usr/local/test.tar.gz\n\n")
		fmt.Printf("%5s\n", "costool download operation eg1:")
		fmt.Println("        costool -download -url https://********* -sd ****** -sk ****** -glt")
		fmt.Println("        please cat the file name as bucketdirlist.log in this dir , get you will pull the file dir and name")
		fmt.Println("        costool -download -url https://********* -sd ****** -sk ****** -bfn APP_BACKUP/test/test.tar.gz")
		fmt.Printf("%5s\n", "costool download operationeg 2:")
		fmt.Println("        ssecretid=`costool -enc secretid`; ssecretkey=`costool -enc secretkey` ")
		fmt.Println("        costool -download -url https://********* -sds $ssecretid -sks $ssecretkey -glt")
		fmt.Println("        please cat the file name as bucketdirlist.log in this dir , get you will pull the file dir and name")
		fmt.Printf("        costool -download -url https://********* -sds $ssecretid -sks $ssecretkey -bfn APP_BACKUP/test/test.tar.gz\n\n")
		fmt.Printf("%5s\n", "costool move operation eg1:")
		fmt.Println("        costool -move -url https://********* -sd ****** -sk ****** -src APP_BACKUP/test/test.tar.gz -dst APP_BACKUP/test2.tar.gz")
		fmt.Printf("%5s\n", "costool move operationeg 2:")
		fmt.Println("        ssecretid=`costool -enc secretid`; ssecretkey=`costool -enc secretkey` ")
		fmt.Printf("        costool -move -url https://********* -sds $ssecretid -sks $ssecretkey -src APP_BACKUP/test/test.tar.gz -dst APP_BACKUP/test2.tar.gz\n\n")
		fmt.Printf("%5s\n", "costool delete operation eg1:")
		fmt.Println("        costool -delete -url https://********* -sd ****** -sk ****** -bfn APP_BACKUP/test/test.tar.gz")
		fmt.Printf("%5s\n", "costool delete operationeg 2:")
		fmt.Println("        ssecretid=`costool -enc secretid`; ssecretkey=`costool -enc secretkey` ")
		fmt.Println("        costool -delete -url https://********* -sds $ssecretid -sks $ssecretkey -bfn APP_BACKUP/test/test.tar.gz")
		os.Exit(0)
	}

}
