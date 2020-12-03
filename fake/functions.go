package fake 

import (
	"flag"
	"os"
)

func GetTerraformStdout() *os.File {
	var terraformStdout *os.File
	if v := flag.Lookup("test.v"); v == nil || v.Value.String() != "true" {
		terraformStdout = os.NewFile(uintptr(4), "stdout")
	} else {
		terraformStdout = os.Stdout
	}
	return terraformStdout
}