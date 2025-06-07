package vpn

import (
	"log"
	"os/exec"
	// "time"
)

func OpenVpn() {
	cmd := exec.Command("C:/Program Files/Proton/VPN/ProtonVPN.Launcher.exe")

	err := cmd.Start()
	if err != nil {
		log.Fatalf("Failed to start app: %v", err)
	}
	log.Println("Application started")
	// time.Sleep(20 * time.Second)
	// err = cmd.Process.Kill()
	// if err != nil {
	// 	log.Fatalf("Failed to kill app: %v", err)
	// }

	// log.Println("Application closed")

}
