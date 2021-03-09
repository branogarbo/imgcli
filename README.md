# 📷 **imgcli**
A simple tool written in Go that prints images to the command line and more.

<hr />

## 🔧 **Setup**
...
<!-- Clone this repo and build the binary,
```

```
or run go get
```
go get github.com/branogarbo/imgcli
go build or go install
``` -->

## 💻 **CLI usage**
```
imgcli

A simple tool written in Go that prints images to the command line and more.

Usage:
   imgcli [--mode=<mode>] [--width=<number>] [--invert] [--save] [--ascii] <path-to-image>
   imgcli [--mode=<mode>] [--width=<number>] [--invert] [--save] [--ascii] [--web] "<image-url>"

Flags:
   -ascii string
         the pattern of ascii characters from least to greatest visibility. pattern of over 8 characters is not recommended (default " .-+*#%@")
   -invert
         whether or not the the print will be inverted
   -mode string
         the mode the image will be printed in. (color, ascii, or box) (default "ascii")
   -save
         whether or not the the print will be written to a text file
   -web
         whether the image is in the filesystem or fetched from the web
   -width int
         the number of characters in each row of the printed image (default 100)


Use "imgcli --help" or "imgcli -h" to view flag usage.
```