package main

import (
    "fmt"
    "os"
  
    "github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
    Use:    "ipmitool"
    Short:  ""
}

funct Execte() {
    if err := Cmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}

func main() {
    Execute()
}
