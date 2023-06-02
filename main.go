package main

import (
	"fmt"
	"github.com/percona/go-mysql/query"
	"os"
	"runtime"
	"strings"
)

const (
	GoOSWindows = "windows"
	GoOSLinux   = "linux"
)

func main() {
	if len(os.Args) <= 1 {
		if runtime.GOOS == GoOSLinux {
			fmt.Println("请输入 sql ,多个 sql 用空格分开\n" +
				"例：\n" +
				"./sql-fingerprint 'select * from table where id=1' 'select * from table where id=2'" +
				"\n" +
				"\n" +
				"Please enter sql, multiple sql separated by spaces\n" +
				"For example:\n" +
				"./sql-fingerprint 'select * from table where id=1' 'select * from table where id=2'")
		}

		if runtime.GOOS == GoOSWindows {
			fmt.Println("请输入 sql ,多个 sql 用空格分开\n" +
				"例：\n" +
				"sql-fingerprint.exe 'select * from table where id=1' 'select * from table where id=2'" +
				"\n" +
				"\n" +
				"Please enter sql, multiple sql separated by spaces\n" +
				"For example:\n" +
				"sql-fingerprint.exe 'select * from table where id=1' 'select * from table where id=2'")
		}
		return
	}

	for k, v := range os.Args {
		if k == 0 {
			continue
		}
		fmt.Println(strings.TrimSpace(query.Fingerprint(v)))
	}
}
