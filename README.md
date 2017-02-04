# gss

go simple server for serving a website directory 

Add the executable to a directory in your path, renaming it to gss, then to use, just navigate to the directory
you want to serve and run

```
gss
```

## Build 

on windows (powershell)

```
$GOOS="windows"; $GOARCH="amd64"; go build -o gss.exe .\main.go
$GOOS="darwin"; $GOARCH="amd64"; go build -o gss .\main.go
```

on mac/linux

```
GOOS=windows GOARCH=amd64 go build -o gss.exe .\main.go
GOOS=darwin GOARCH=amd64 go build -o gss .\main.go
```
