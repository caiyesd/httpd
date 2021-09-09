package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var (
	addr string
	port int
	dir  string
)

func init() {
	flag.StringVar(&addr, "a", "0.0.0.0", "listen address")
	flag.IntVar(&port, "p", 8080, "http port")
	flag.StringVar(&dir, "d", ".", "http root dir")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `httpd - 0.0.1
Usage: httpd [-a ip-addr] [-p port] [-d dir]
Options:
`)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	http.ListenAndServe(fmt.Sprintf("%s:%d", addr, port), http.FileServer(http.Dir(dir)))
}
