# Installing Go Development Environment on Mac OS X

Go is pretty opinionated about the folder structure and needs a bit of figuring out.

## Getting Started

```
brew install go
```
This will install GO binaries in your /usr/local/Cellar/go/1.11.1/libexec

### IDE's

I use GOland for my development. It seems to be more convenient as I am used to IntelliJ .
Other choices are :
Visual Code Studio with GO Plugin
Atom



### Installing 3rd Party Packages directly from Github

Go is pretty stringent about finding packages from a specific location. To install any packages you have to run ```go get github.com/gorilla/mux``` from the `src` folder so that it creates `github.com` on the your path of `~/go/src` .

```
cd ~/go/src
go get github.com/gorilla/mux - URL router and dispatcher
go get github.com/subosito/gotenv - Load environment variables dynamically
go get github.com/lib/pq - postgres driver for Go

```