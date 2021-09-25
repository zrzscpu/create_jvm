package main

import "fmt"

func main() {

	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println(cmd.versionFlag)

	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startjvm(cmd)
	}

}
func startjvm(cmd *Cmd) {

	fmt.Printf("classpath: %s args: %s", cmd.class, cmd.args)
}
