/*
Package walker handles walking through the given directory and extracting file data.
*/
package walker

import (
	"fmt"
	// "path/filepath"
)

func WalkDir(fp string) {
	fmt.Printf("Walking given directory %s", fp)
}
