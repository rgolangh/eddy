package main

import (
	"fmt"
	"github.com/rgolangh/eddy/pkg/eddy"
	flag "github.com/spf13/pflag"
	"os"
)


func main() {
	fmt.Println(os.Args)

	var unitFile eddy.UnitFile
	initFlags(&unitFile)

	iniAsString, err := eddy.Write(unitFile)

	if err != nil {
		exit(1, err.Error())
	}

	// print to out
	exit(0, iniAsString)
}


func initFlags(unitFile *eddy.UnitFile) {
	defer flag.Parse()

	flag.StringVar(&unitFile.Unit.Description, "unit-description","",  "a description of the unit")
	flag.StringVar(&unitFile.Service.ExecStart,"service-exec-start", "", "the ExecStart command")
	flag.StringVar(&unitFile.Service.PidFile,"service-pid-file", "", "the pid file of the service, standard is /var/run/name.pid")
	flag.StringVar(&unitFile.Service.Type,"service-type", "", "the of service e.g forking")
	flag.StringVar(&unitFile.Install.WantedBy,"install-wanted-by", "", "a WantedBy specification, e.g multi-user.target")
	flag.StringVar(&unitFile.Install.RequiredBy,"install-required-by", "", "a RequiredBy specification, e.g multi-user.target")
}

func exit(code int, message string) {
	fmt.Println(message)
	os.Exit(code)
}