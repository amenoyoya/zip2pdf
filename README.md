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

# clean the go.mod file
$ go mod tidy
```
