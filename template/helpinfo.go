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
		fmt.Println("usage: cosupload -url $BucketURL -sd $SecretID -sk $SecretKEY -fnm $FileNAME -fdir $FileDIR")
		fmt.Println("       cosupload -url $BucketURL -sds $SecretID -sks $SecretKEY -fnm $FileNAME -fdir $FileDIR")
		fmt.Printf("%5s  %s\n", "-url ", "bucket_url")
		fmt.Printf("%5s  %s\n", "-sd  ", "secretid")
		fmt.Printf("%5s  %s\n", "-sk  ", "secretkey")
		fmt.Printf("%5s  %s\n", "-sds ", "receive encrypt secretid")
		fmt.Printf("%5s  %s\n", "-sks ", "receive encrypt secretkey")
		fmt.Printf("%5s  %s\n", "-bfn ", "the file full dir and file name on bucket")
		fmt.Printf("%5s  %s\n", "-enc ", "encrypt the string will use as encrypt secretid or secretkey ")
		fmt.Printf("%5s  %s\n%s\n", "-fdir", "the upload file dir and file name in the system", "eg 1:")
		fmt.Println("       cosupload -url https://********* -sd ****** -sk ****** -fnm /APP_BACKUP/test/test.tar.gz -fdir /usr/local/test.tar.gz")
		fmt.Printf("%5s\n", "eg 2:")
		fmt.Println("       ssecretid=`cosupload -enc secretid`; ssecretkey=`cosupload-v1 -enc secretkey` ")
		fmt.Println("       cosupload -url https://********* -sds $ssecretid -sks $ssecretkey -fnm /APP_BACKUP/test/test.tar.gz -fdir /usr/local/test.tar.gz")
		os.Exit(0)
	}

}
