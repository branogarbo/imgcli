# 📷 **imgcli-cobra**
**A simple tool written in Go that prints images to the command line and more.**

---

## 🔧 **Setup**
Download and compile from sources:
```
go get github.com/branogarbo/imgcli-cobra
```
Install just the binary with Go:
```
go install github.com/branogarbo/imgcli-cobra@latest
```

Or get the pre-compiled binaries for your platform on the [releases page](https://github.com/branogarbo/imgcli-cobra/releases)


## 💻 **CLI usage**
```
imgcli-cobra

A rough copy of imgcli written with cobra

Usage:
  imgcli-cobra [command]

Available Commands:
  help        Help about any command
  print       Prints images to the command line
  save        Saves converted image to a text file

Flags:
  -h, --help           help for imgcli-cobra
  -i, --invert         Whether the the print will be inverted or not
  -m, --mode string    he mode the image will be printed in (default "ascii")
  -W, --web            Whether the source image is in the filesystem or fetched from the web
  -w, --width int      The number of characters in each row of the output (default 100)
  -p, --ascii string   The pattern of ascii characters from least to greatest 
                       visibility. Patterns of over 8 characters are not recommended (default " .-+*#%@")

Use "imgcli-cobra [command] --help" for more information about a command.
```