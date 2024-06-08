package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

func main() {

	fmt.Println("Shell is running\nFor exit enter \\quit")

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		txt := sc.Text()
		if strings.EqualFold(txt, "\\quit") {
			break
		}
		cmdArr := strings.Split(txt, "|")
		for _, el := range cmdArr {
			txtCmd := strings.Split(el, " ")

			if len(txtCmd) == 0 {
				continue
			}
			if txtCmd[0] == "exec" {
				pid := os.Getpid()

			}

			switch txtCmd[0] {
			case "cd":
				var dirname string
				if len(txtCmd) < 2 {
					var err error
					dirname, err = os.UserHomeDir()
					if err != nil {
						log.Fatal(err)
					}

				} else {
					dirname = txtCmd[1]
				}
				os.Chdir(dirname)
				newDir, err := os.Getwd()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Current Working Direcotry: %s\n", newDir)
			case "pwd":
				newDir, err := os.Getwd()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Current Working Direcotry: %s\n", newDir)
			case "echo":
				fmt.Println(strings.Join(txtCmd[1:], " "))
			case "kill":
				processes, err := process.Processes()
				if err != nil {
					log.Fatal(err)
				}
				pid, err := strconv.ParseInt(txtCmd[1], 10, 32)
				if err != nil {
					log.Fatal(err)
				}
				var procFound bool
				for _, p := range processes {
					if int32(pid) == p.Pid {
						p.Kill()
						procFound = true
						break
					}
				}
				if !procFound {
					fmt.Println("process not found")
				}
			case "ps":
				processes, err := process.Processes()
				if err != nil {
					log.Fatal(err)
				}
				w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
				for _, p := range processes {
					username, err := p.Username()
					if err != nil {
						log.Fatal(err)
					}
					createTime, err := p.CreateTime()
					if err != nil {
						log.Fatal(err)
					}
					//startDate := time.Now().Local().Add((time.Duration(createTime*(-1)) * time.Millisecond))
					cmd, err := p.Cmdline()
					if err != nil {
						log.Fatal(err)
					}
					fmt.Fprintf(w, "%s\t%d\t%s\t%s\t\n", username, p.Pid, time.Unix(0, createTime*int64(time.Millisecond)), cmd[:min(len(cmd), 10)])
				}
				w.Flush()
			default:
				fmt.Println("Command not found")
			}
		}
	}
}
