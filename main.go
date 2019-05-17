package main

import (
	"cos_tool/template"
	"flag"
	"fmt"
	"os"
)

//init first start use recive strings
func init() {
	template.InitreceiveString()
}

func main() {

	// parse strings
	flag.Parse()

	// use costool -enc strings for encrypt strings
	ENCryptstringS := template.CHEckouT(template.ENCryptstring)
	if ENCryptstringS != "" {
		fmt.Println("++++++++++++++encrypt string+++++++++:", template.StringUPSET(ENCryptstringS))
		os.Exit(0)
	}

	// output the tool of costool help info
	template.HelpInfo(template.Helpdoc, template.Helpdoc2, template.Upload, template.Download, template.Delete, template.Move)

	// OPerationJudge use receive operation and run
	template.OPerationJudge()

}
