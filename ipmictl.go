package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "os"
    "os/exec"
    "strings"
)

var (
    host        string
    username    string
    password    string
    command     string
    file        string
)

func init() {
    flag.StringVar(&host, "H", "", "主机地址")
    flag.StringVar(&username, "U", "ADMIN", "主机用户")
    flag.StringVar(&password, "P", "ADMIN", "主机密码")
    flag.StringVar(&command, "p", "status", "操作命令")
    flag.StringVar(&file, "i", "", "主机清单 hosts.txt")

    flag.Usage = usage
}

func usage() {
    fmt.Fprintf(os.Stderr, `
Usage:
    ipmictl -H 127.0.0.1 -U username -P password -p [ status | on | off ]
    ipmictl -i hosts.txt -p [ status | on | off ]

Inventoy:
    vi hosts.txt

    Example 1:
    127.0.0.1 

    Example 2:
    127.0.0.1 username password

    Example 3:
    127.0.0.1 username password [ status | on | off ]

Options:
`)

    flag.PrintDefaults()
}

func Cmd(h string, u string, p string, c string) {
    args := fmt.Sprintf("ipmitool -I lan -H %s -U %s -P %s power %s", h, u, p, c)
    res  := exec.Command("/bin/bash", "-c", args)
    output, err := res.Output()

    if err != nil {
        fmt.Printf("Error: %s %s\n", h, err.Error() )
        return
    }

    fmt.Printf("Output: %s %s\n", h, string(output))
}

func InventoryCmd() {
    text, err := ioutil.ReadFile(file)

    if err != nil {
        panic(err)
        os.Exit(1)
    }

    list := strings.Split(string(text), "\n")
    for _, v := range list {
        if len(v) == 0 {
            continue
        }
        t := strings.Split(v, " ")
        if len(t) == 1 {
            host = t[0]
        }
        if len(t) == 3 {
            host     = t[0]
            username = t[1]
            password = t[2]
        }
        if len(t) == 4 {
            host     = t[0]
            username = t[1]
            password = t[2]
            command  = t[3]
        }

        Cmd(host, username, password, command)
    }
}

func main() {
    flag.Parse()

    if (len(host) != 0) &&  (len(file) ==0) {
        Cmd(host, username, password, command)

    } else if (len(host) == 0) && (len(file) !=0) {
        InventoryCmd()
    } else {
        flag.Usage()
        os.Exit(1)
    }
}
