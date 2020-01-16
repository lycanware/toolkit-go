## Lycanware Toolkit For Go
Useful tools that are used all the time and not built-in to Go.

## Install
```sh
go get -u github.com/lycanware/toolkit-go/filesys/copy
```

## Example
First create a directory called `source_directory` and add some test files that will be copied. Running the following program will
copy the directory and all files to a new directory called `destination_directory`
```go
package main

import (
	"fmt"
	"log"
	"github.com/lycanware/toolkit-go/filesys/copy"
)

func main() {
	if err := copy.Copy("source_directory", "destination_directory"); err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Copy Complete")
}
```

## Author
Craig Sherlock

## License
lycanware/toolkit-go is licensed under the MIT license. See the LICENSE file for more info.