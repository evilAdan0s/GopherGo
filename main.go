package main

import (
	"GopherGo/utils"
	"fmt"
	"os"
)

func main() {
	echoLogo()
	echoHelp()
	//fmt.Println(len(os.Args))
	if len(os.Args) > 3 || len(os.Args) < 2 {
		fmt.Println("[!] Error parsing arguments")
		os.Exit(1)
	}
	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		echoHelp()
		os.Exit(0)
	}
	if os.Args[1] == "-e" || os.Args[1] == "--exploit" {
		utils.Exploit(os.Args[2])
		os.Exit(0)
	}
}

func echoLogo() {
	logoStr := `
   ______            __              ______    
  / ____/___  ____  / /_  ___  _____/ ____/___ 
 / / __/ __ \/ __ \/ __ \/ _ \/ ___/ / __/ __ \
/ /_/ / /_/ / /_/ / / / /  __/ /  / /_/ / /_/ /
\____/\____/ .___/_/ /_/\___/_/   \____/\____/ 
          /_/                                  
`
	fmt.Printf("\u001B[1;34;m%s\u001B[0m\n", logoStr)
	fmt.Printf("\u001B[1;32;m%s\u001B[0m\n", "author: @evilAdan0s  |  version: 0.0.1")
}

func echoHelp() {
	helpStr := `
Usage: gophergo [-h] [--exploit EXPLOIT]

Options:
  -h, --help    Show this help message and exit
  -e, --exploit mysql, redis, zabbix
`
	fmt.Println(helpStr)
}
