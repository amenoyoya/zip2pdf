# zip2pdf

Convert zip file to pdf file

## Environment

- Shell:
    - PowerShell
    - bash
- Golang: `1.17.1`

***

## Usage

```bash
$ go run zip2pdf.go <ZipFile>

# => generated: <ZipFile>.pdf
```

***

## Development

### Setup
```bash
# initialize dependent modules management file
## => generated: go.mod
$ go mod init zip2pdf

# install golang.org/x/image/webp
$ go get golang.org/x/image/webp

# install github.com/signintech/gopdf
$ go get github.com/signintech/gopdf

# clean the go.mod file
$ go mod tidy
```

### Build an executable file
```bash
$ go build ./zip2pdf.go

# => generated: executable file
## - Linux: ./zip2pdf
## - macOS: ./zip2pdf.app
## - Windows: ./zip2pdf.exe

# test
$ ./zip2pdf ./test/test.zip
```
