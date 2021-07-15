package main

import (
	"fmt"
	"os"

	"github.com/miihael/goconf"
)

type testOptions struct {
	goconf.AutoOptions
	HTTPAddress string   `default:"0.0.0.0:0000"`
	Hosts       []string `flag:"hosts" cfg:"hosts" default:"127.0.0.0,127.0.0.1"`
	LogLevel    int      `default:"3"`
	BoolVar     bool     `default:"false"`
	IntSlice    []int64  `cfg:"int_slice"`
	FromEnv     string   `default:"undefined"`
}

func main() {
	ops := &testOptions{}
	os.Setenv("FROM_ENV", "zzz")
	// conf_3 inherit from conf_1 and conf_2
	goconf.MustResolve(ops, "conf_1.toml", "conf_1.json", "conf_3.toml")
	fmt.Println(ops)
}
