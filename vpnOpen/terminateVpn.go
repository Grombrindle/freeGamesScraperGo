package vpn

import (
	// "syscall"
	"log"

	"golang.org/x/sys/windows"
)

func KillProcess(pid int) error {
	h, err := windows.OpenProcess(windows.PROCESS_TERMINATE, false, uint32(pid))
	if err != nil {
		log.Fatal(err)
	}
	defer windows.CloseHandle(h)
	return windows.TerminateProcess(h, 1)
}
