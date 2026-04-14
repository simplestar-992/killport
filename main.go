package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	port := flag.Arg(0)

	if port == "" {
		fmt.Println("KillPort - Kill process on port")
		fmt.Println("Usage: killport <port>")
		return
	}

	pid := getPID(port)
	if pid == 0 {
		fmt.Printf("No process found on port %s\n", port)
		return
	}

	fmt.Printf("Killing PID %d on port %s\n", pid, port)
	exec.Command("kill", strconv.Itoa(pid)).Run()
	fmt.Println("Done")
}

func getPID(port string) int {
	out, _ := exec.Command("ss", "-tlnp").Output()
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if strings.Contains(line, ":"+port+" ") {
			// Extract PID from ss output
			parts := strings.Fields(line)
			for _, p := range parts {
				if strings.Contains(p, "pid=") {
					pidStr := strings.TrimPrefix(p, "pid=")
					pidStr = strings.Split(pidStr, ",")[0]
					pid, _ := strconv.Atoi(pidStr)
					return pid
				}
			}
		}
	}
	return 0
}
