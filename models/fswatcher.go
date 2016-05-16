package models

import (
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Watch(path string) chan error {
	stop := make(chan error)
	go func() {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			stop <- err
		}
		defer watcher.Close()
		// TODO handle error here
		filepath.Walk(path, func(p string, info os.FileInfo, _ error) error {
			if info.IsDir() {
				return watcher.Add(p)
			}
			return nil
		})
		stop := make(chan bool)
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Create == fsnotify.Create {
					s, e := os.Stat(event.Name)
					if e != nil {
						continue
					}
					// add a new path
					if s.IsDir() {
						watcher.Add(event.Name)
					}
				}
				if strings.HasSuffix(event.Name, ".md") {
					log.Println("md file changed:", event)
				}
			case err := <-watcher.Errors:
				log.Println("err:", err)
			case <-stop:
				break
			}
		}
	}()
	return stop
}
