# 📷 **imgcli**
**A simple tool written in Go that prints images to the command line and more.**

---

🚨 **Warning: imgcli may be replaced by [imgcli-cobra](https://github.com/branogarbo/imgcli-cobra) in the near future under the imgcli name** 🚨

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
   imgcli [--mode=<mode>] [--width=<number>] [--ascii=<pattern>] [--invert] [--save] <path-to-image>
   imgcli [--mode=<mode>] [--width=<number>] [--ascii=<pattern>] [--invert] [--save] [--web] "<image-url>"

Flags:
   -ascii string
         The pattern of ascii characters from least to greatest visibility. pattern of 
         over 8 characters is not recommended (default " .-+*#%@")
   -invert
         Whether or not the the print will be inverted
   -mode string
         The mode the image will be printed in. (color, ascii, or box) (default "ascii")
   -save
         Whether or not the the print will be written to a text file
   -web
         Whether the image is in the filesystem or fetched from the web
   -width int
         The number of characters in each row of the printed image (default 100)


Use "imgcli --help" or "imgcli -h" to view flag usage.
```