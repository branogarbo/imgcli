# 📷 **imgcli**
**A simple tool written in Go that prints images to the command line and more.**

---

## 🔧 **Setup**
Install the binary with go:
```
go install github.com/branogarbo/imgcli@latest
```

Or get the pre-compiled binaries for your platform on the [releases page](https://github.com/branogarbo/imgcli/releases)


## 💻 **CLI usage**
```
imgcli

A simple tool written in Go that prints images to the command line and more.

Usage:
   imgcli [--mode=<mode>] [--width=<number>] [--invert] [--save] [--ascii] <path-to-image>
   imgcli [--mode=<mode>] [--width=<number>] [--invert] [--save] [--ascii] [--web] "<image-url>"

Flags:
   -ascii string
         the pattern of ascii characters from least to greatest visibility. pattern of 
         over 8 characters is not recommended (default " .-+*#%@")
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