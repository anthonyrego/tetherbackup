package main

import (
	"fmt"
	"os"
	"os/signal"
)

const coverText = `

████████╗███████╗████████╗██╗  ██╗███████╗██████╗    
╚══██╔══╝██╔════╝╚══██╔══╝██║  ██║██╔════╝██╔══██╗   
   ██║   █████╗     ██║   ███████║█████╗  ██████╔╝   
   ██║   ██╔══╝     ██║   ██╔══██║██╔══╝  ██╔══██╗   
   ██║   ███████╗   ██║   ██║  ██║███████╗██║  ██║   
   ╚═╝   ╚══════╝   ╚═╝   ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝   
                                                     
    ██████╗  █████╗  ██████╗██╗  ██╗██╗   ██╗██████╗ 
    ██╔══██╗██╔══██╗██╔════╝██║ ██╔╝██║   ██║██╔══██╗
    ██████╔╝███████║██║     █████╔╝ ██║   ██║██████╔╝
    ██╔══██╗██╔══██║██║     ██╔═██╗ ██║   ██║██╔═══╝ 
    ██████╔╝██║  ██║╚██████╗██║  ██╗╚██████╔╝██║     
    ╚═════╝ ╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝ ╚═════╝ ╚═╝     
                                            
`

func isDirectory(path string) bool {
	srcInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if !srcInfo.IsDir() {
		fmt.Println(fmt.Sprintf("%s is not a directory", srcInfo.Name()))
		return false
	}
	return true
}

func main() {
	args := os.Args[1:]
	if len(args) <= 1 {
		fmt.Println("Usage: {src} {...destinations}")
		return
	}
	if !isDirectory(args[0]) {
		return
	}
	for _, dest := range args[1:] {
		if !isDirectory(dest) {
			return
		}
	}
	go func() {
		sigchan := make(chan os.Signal, 10)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		fmt.Println("\nProgram closing...")
		os.Exit(0)
	}()
	clearScreen()
	fmt.Println(colorizeString(coverText, magenta))
	watchDirectory(args[0], args[1:])
}
