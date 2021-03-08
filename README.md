# **imgcli**
A simple tool written in Go that prints images to the command line and more.

<hr />

## **Setup**
```
go get github.com/branogarbo/imgcli
```

## **CLI usage**
```
imgcli

A simple tool written in Go that prints images to the command line and more.

Usage:
   imgcli [--mode=<mode>] [--width=<number>] [--invert] [--save] <path-to-image>
   imgcli [--mode=<mode>] [--width=<number>] [--invert] [--save] [--web] "<image-url>"

Flags:
   --invert
      whether or not the the print will be inverted
   --mode string
      the mode the image will be printed in. (color, ascii, or box) (default "box")
   --save
      whether or not the the print will be written to a text file
   --web
      whether the image is in the filesystem or fetched from the web
   --width int
      the number of characters in each row of the printed image (default 100)


Use "imgcli --help" or "imgcli -h" to view flag usage.
```