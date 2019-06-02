/*
Package walker handles walking through the given directory and extracting file data.
*/
package walker

import (
	"fmt"
	"os"
	"path/filepath"
)

//WalkDir takes the given directory and walks through all the files, handling keying directory names
func WalkDir(fp string) chan string {
	// Gather absolute path name just in case
	abspath, _ := filepath.Abs(fp)
	fmt.Printf("Walking given directory %s\n", abspath)

	// Gather channel
	channel := GoWalk(abspath)
	return channel
}

//GoWalk found on internet, seems promising
func GoWalk(location string) chan string {
	// create channel of strings
	channel := make(chan string)
	// concurrently place the path in the channel
	go func() {
		filepath.Walk(location, func(path string, info os.FileInfo, _ error) (err error) {
			if !info.IsDir() {
				channel <- path
			}
			return
		})
		defer close(channel)
	}()
	return channel
}
