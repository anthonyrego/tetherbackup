package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

func watchDirectory(path string, destinations []string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()
	fmt.Println(colorizeString(fmt.Sprintf("Watching %s...", path), blue))
	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Create == fsnotify.Create {
					parts := strings.Split(event.Name, "/")
					filename := parts[len(parts)-1]
					fmt.Println(colorizeString("new file: ", green), filename)
					for _, dest := range destinations {
						copyFile(event.Name, dest+filename)
					}
				}
			case err = <-watcher.Errors:
				fmt.Println(colorizeString("error: ", red), err)
			}
			time.Sleep(1000 * time.Millisecond)
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func copyFile(srcName, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Println(colorizeString("error: ", red), err)
		return
	}

	dst, err := os.Create(dstName)
	if err != nil {
		fmt.Println(colorizeString("error: ", red), err)
		return
	}

	written, err = io.Copy(dst, src)
	dst.Close()
	src.Close()
	return
}
