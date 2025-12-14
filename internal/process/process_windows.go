//go:build windows

package process

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// FindPids finds the PIDs of processes listening on the specified port.
func FindPids(port string) ([]int, error) {
	// netstat -ano
	cmd := exec.Command("netstat", "-ano")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to run netstat: %w", err)
	}

	var pids []int
    seen := make(map[int]bool)
	scanner := bufio.NewScanner(&out)
	for scanner.Scan() {
		line := scanner.Text()
		// Line format example: "  TCP    0.0.0.0:8080           0.0.0.0:0              LISTENING       1234"
        // We want lines containing our port and "LISTENING"
        if !strings.Contains(line, "LISTENING") {
            continue
        }
        
        // Simple check for port presence. 
        // Robustness: check if it actually matches :<port> or <ip>:<port> boundaries to avoid matching 80800 for 8080
        // netstat output usually uses colon separator. 
        if !strings.Contains(line, ":"+port+" ") && !strings.HasSuffix(line, ":"+port) {
             // Basic substring check might fail for port 80 if 8080 exists. 
             // Let's rely on fields parsing.
             // We'll parse fields and check the local address.
        }

        fields := strings.Fields(line)
        if len(fields) < 5 {
            continue
        }
        
        // Field 0: Proto (TCP)
        // Field 1: Local Address (0.0.0.0:8080)
        // Field 2: Foreign Address
        // Field 3: State (LISTENING)
        // Field 4: PID (1234)

        localAddr := fields[1]
        // Check if localAddr ends with :port
        if !strings.HasSuffix(localAddr, ":"+port) {
            continue
        }

        pidStr := fields[len(fields)-1] // PID is the last column
        pid, err := strconv.Atoi(pidStr)
		if err == nil {
            if !seen[pid] {
                pids = append(pids, pid)
                seen[pid] = true
            }
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to parse netstat output: %w", err)
	}

	return pids, nil
}

// KillProcess kills the process with the given PID.
func KillProcess(pid int) error {
	// taskkill /F /PID <pid>
	cmd := exec.Command("taskkill", "/F", "/PID", strconv.Itoa(pid))
	// taskkill output goes to stdout/stderr, we might want to suppress it or capture it for debug
    // "Process <PID> is listening..." user just wants it gone.
	return cmd.Run()
}
