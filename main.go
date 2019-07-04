package main

import (
	"fmt"
	//"github.com/spf13/viper"
	"github.com/svazist/sometests/cmd"
)

var (
	Version   = "dev"
	Build     = "none"
	BuildDate = "unknown"
)

func main() {
	cmd.Version = fmt.Sprintf("%v, commit %v, built at %v", Version, Build, BuildDate)
	fmt.Printf("Development version:\n Version: %v, commit: %v, built at: %v\n\n", Version, Build, BuildDate)
	cmd.Execute()
}
