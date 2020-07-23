package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/go-vgo/robotgo"
)

type windowsProcessObj struct {
	pid  int32
	name string
}

// WindowsProcess represents the process we find
var WindowsProcess *windowsProcessObj = &windowsProcessObj{}

func checkFileExists() (exists bool) {
	pslist := DataHome() + "pslist64.exe"
	pssuspend := DataHome() + "pssuspend64.exe"
	// Check if pslist exists
	var pslistExists bool
	if _, err := os.Stat(pslist); err == nil {
		pslistExists = true
	}
	// Check if pssuspend exists
	var pssuspendExists bool
	if _, err := os.Stat(pssuspend); err == nil {
		pssuspendExists = true
	}
	if pslistExists && pssuspendExists {
		return true
	}
	log.Println("Need tools, downloading..")
	return false
}

func getTools() {
	toolsExists := checkFileExists()
	if toolsExists == false {
		// Download pslist
		pslist := "( New-Item -Path " + DataHome() + "pslist64.exe" + " -Force )"
		cmd, err := exec.Command("cmd", "/C",
			"powershell.exe",
			"Invoke-WebRequest",
			"http://live.sysinternals.com/tools/pslist64.exe",
			"-OutFile",
			pslist).CombinedOutput()
		log.Printf("%s\n", cmd)
		if err != nil {
			log.Println("Error downloading pslist: ", err)
		}
		// Download pssuspend
		pssuspend := "( New-Item -Path " + DataHome() + "pssuspend64.exe" + " -Force )"
		cmd, err = exec.Command("cmd", "/C",
			"powershell.exe",
			"Invoke-WebRequest",
			"http://live.sysinternals.com/tools/pssuspend64.exe",
			"-OutFile",
			pssuspend).CombinedOutput()
		log.Printf("%s\n", cmd)
		if err != nil {
			log.Println("Error downloading pssuspend: ", err)
		}
	}
}

func getActiveWindowWindows() {
	// Get name
	activeWindowName := robotgo.GetTitle()
	log.Println("activeWindowName: ", activeWindowName)
	WindowsProcess.name = activeWindowName
	// Get PID
	activeWindowPID := robotgo.GetPID()
	log.Println("activeWindowPID: ", activeWindowPID)
	WindowsProcess.pid = activeWindowPID
}

// ToggleSuspendWindows will.. toggle suspend on Windows
func ToggleSuspendWindows() {
	getTools()
	getActiveWindowWindows()
	/* // Check process status
	stringPID := strconv.Itoa(int(WindowsProcess.pid))
	log.Println("stringPID: ", stringPID)
	cmd, err := exec.Command("cmd", "/C",
		"C:\\Users\\Merritt\\Downloads\\pslist64.exe",
		"-accepteula",
		"-x",
		stringPID).CombinedOutput()
	log.Printf("%s\n", cmd)
	if err != nil {
		log.Println("Error running pslist: ", err)
	} */

	// Save suspended process details to file
	SaveProcessFile(WindowsProcess.name, WindowsProcess.pid)
	/* -------------------------------------------------------------------------- */
	// Check if a saved process file exists
	/* name, pid, err := LoadProcessFile()
	switch {
	case err == nil:
		log.Println("Found saved process details - name:", name, "PID:", pid)
		process, err := process.NewProcess(pid)
		Check(err)
		status, err := process.Status()
		Check(err)
		// If a saved file exists, try to resume that process
		if status == "T" {
			log.Println(name, "is stopped - resuming.")
			NotifyResume(name)
			err := os.Remove(SavedProcessFile)
			Check(err)
		} else {
			err := os.Remove(SavedProcessFile)
			Check(err)
			log.Println(name, "is not suspended, removed invalid cache.")
		}
	// If no saved process file is found, suspend the active window
	case err != nil:
		log.Println("No saved process details found")
		name, pid := findProcess()
		process, err := process.NewProcess(pid)
		Check(err)
		status, err := process.Status()
		Check(err)
		if status == "T" {
			status = "Suspended"
		} else {
			status = "Running"
		}
		log.Println("Checking process - name:", name, "PID:", pid, "status:", status)
		switch status {
		case "Running":
			log.Println("Suspending", name)
			NotifySuspend(name)
			process.Suspend()
		case "Suspended":
			log.Println("Resuming", name)
			NotifyResume(name)
			process.Resume()
		}
		// Save suspended process details to file
		SaveProcessFile(name, pid)
	} */
}
