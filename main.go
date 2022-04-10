// This file is a tooling built to help on the overlay design process

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func main() {
	fmt.Println("💧 Water")

	watcher, err := fsnotify.NewWatcher();

	if err != nil {
		log.Fatal(err);
	}

	done := make(chan bool);

	go func() {
		defer close(done)

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Printf("%s: %s\n", event.Name, event.Op)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error:", err)
			}
		}
	}()

	dir, err := os.Getwd();
	inputPath := os.Args[1];

	if err != nil {
		log.Fatal(err);
	}

	finalPath := filepath.Join(dir, inputPath)

	log.Println("Watching: ", finalPath);

	err = watcher.Add(finalPath)
	if err != nil {
		log.Fatal("Add failed: ", err)
	}
	<-done
}
