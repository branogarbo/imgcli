# 📷 **imgcli**
**A simple tool written in Go that prints images to the command line and more.**

---

## 🔧 **Setup**
Download and compile from sources:
```
go get github.com/branogarbo/imgcli
```
Install just the binary with Go:
```
go install github.com/branogarbo/imgcli@latest
```

Or get the pre-compiled binaries for your platform on the [releases page](https://github.com/branogarbo/imgcli/releases)


## 💻 **CLI usage**
```
imgcli

A simple tool written in Go that prints images to the command line and more.

Usage:
  imgcli [command]

Available Commands:
  help        Help about any command
  print       Prints images to the command line   
  save        Saves converted image to a text file

Flags:
  -h, --help   help for imgcli

Use "imgcli [command] --help" for more information about a command.
```

### **Command Usage:**
```
Usage:
  imgcli [command] [flags]

Examples:
imgcli print --invert ./images/pic.jpg
imgcli save -w 200 -W "https://url-to-some/image.jpg"

Flags:
  -h, --help           Help for command
  -i, --invert         Whether or not the the print will be inverted
  -m, --mode string    The mode the image will be printed in (default "ascii")
  -W, --web            Whether the source image is in the filesystem or fetched from the web
  -w, --width int      The number of characters in each row of the output (default 100)
  -q, --quiet          Whether or not the save output is quiet (Does not matter for print command)
  -p, --ascii string   The pattern of ascii characters from least to greatest 
                       visibility (Does not matter for non-ascii modes) (default " .:-=+*#%@")

Use "imgcli [command] --help" for more information about a command.
```

---

## 📁 **Package Usage**
Install the package:
```
go get github.com/branogarbo/imgcli/util
```

### **Example:**
``` go
package main

import (
	"fmt"
	"log"

	ic "github.com/branogarbo/imgcli/util"
)

func main() {
	printConfig := ic.OutputConfig{
		Src:          "https://github.com/branogarbo/imgcli/blob/master/examples/images/portrait.jpg?raw=true",
		OutputMode:   "ascii",
		AsciiPattern: " .:-=+*#%@",
		OutputWidth:  100,
		IsUseWeb:     true,
	}

	printString, err := ic.OutputImage(printConfig)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(printString)
}
```