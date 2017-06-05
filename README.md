## What am I trying to do
I want to use be able to import the vendor files from code in my modules subfolder.

## What is happening
I get the following errors:

> $ go run main.go
modules\mystuff\mycode.go:6:2: cannot find package "github.com/deprepo" in any of:
        C:\Go\src\github.com\deprepo (from $GOROOT)
        C:\Users\dave.johnson\go\src\github.com\deprepo (from $GOPATH)

## My environment
* I working on a Windows 10 box with Go version 1.8.3
* I set GOPATH=C:/example/product1/src/myapp/vendor/
* I am working out of the folder C:/example/product1/src/myapp
* I am executing ```go run main.go``` for my test.

## Is this a valid work around
I move the dependency repos under `vendor/src` and change my GOPATH to `.../vendor`.   I am then able to `go run main.go` and get the following output:
```$ go run main.go
start main
inside vendor mymod, from:  main
from inside modules/mystuff
inside vendor mymod, from:  from mystuff
end main```