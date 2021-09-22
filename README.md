# zip2pdf

Convert zip file to pdf file

## Environment

- Shell:
    - PowerShell
    - bash
- Node.js: `12.21.0`
    - npm: `7.21.1`
    - yarn: `1.22.11`

***

## Usage

```bash
$ node zip2pdf.js <ZipFile>

# => generated: <ZipFile>.pdf
```

***

## Development

### Setup
```bash
# using npm
$ npm i

# or using yarn
$ yarn

# => installed: node_modules
## <dependencies>
### * node-zip
### * pdfkit
### * sharp
## <devDependencies>
### * nexe
```

### Build executable file
```bash
# prepare `bin` directory
$ mkdir bin
$ cp node_modules -r bin/node_modules

# linux
$ npx nexe zip2pdf.js -o bin/zip2pdf --target linux-x64-14.15.3

# mac
$ npx nexe zip2pdf.js -o bin/zip2pdf.app --target mac-x64-14.15.3

# windows
$ npx nexe zip2pdf.js -o bin/zip2pdf.exe --target windows-x64-14.15.3
```
