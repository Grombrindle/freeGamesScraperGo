package vpn

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func PID() string {
	cmd := exec.Command("tasklist", "/FI", "IMAGENAME eq ProtonVPN.Client.exe")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	var pid string
	// pid = ""

	lines := bytes.Split(output, []byte{'\n'})
	for _, line := range lines {
		text := string(line)
		if strings.HasPrefix(text, "ProtonVPN.Client.exe") {
			fields := strings.Fields(text)
			if len(fields) >= 2 {
				pid := fields[1]
				fmt.Println("PID:", fields[1])
				return pid
			}
		}
	}
	fmt.Println("Process not found.")
	return pid
}
