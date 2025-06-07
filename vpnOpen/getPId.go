package vpn

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func PID() int {
	cmd := exec.Command("tasklist", "/FI", "IMAGENAME eq ProtonVPN.Client.exe")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	var pid int
	// pid = ""

	lines := bytes.Split(output, []byte{'\n'})
	for _, line := range lines {
		text := string(line)
		if strings.HasPrefix(text, "ProtonVPN.Client.exe") {
			fields := strings.Fields(text)
			if len(fields) >= 2 {
				pid := fields[1]
				fmt.Println("PID:", fields[1])
				pidNumber, err := strconv.Atoi(pid)
				if err != nil {
					log.Fatal(err)
				}
				return pidNumber
			}
		}
	}
	fmt.Println("Process not found.")
	return pid
}
