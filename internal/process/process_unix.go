//go:build !windows

package process

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// FindPids finds the PIDs of processes listening on the specified port.
func FindPids(port string) ([]int, error) {
	// lsof -i :<port> -t
	// -t: terse mode, output only PIDs
	cmd := exec.Command("lsof", "-i", ":"+port, "-t")
	var out bytes.Buffer
	cmd.Stdout = &out
    // lsof returns exit code 1 if no process is found, which results in an error in Run().
    // We should handle that as "no processes found" rather than a hard error.
	err := cmd.Run()
	if err != nil {
        // Check if exit code is 1 (no process found)
        if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 1 {
            return []int{}, nil
        }
		return nil, fmt.Errorf("failed to run lsof: %w", err)
	}

	var pids []int
	scanner := bufio.NewScanner(&out)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		pid, err := strconv.Atoi(line)
		if err == nil {
			pids = append(pids, pid)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to parse lsof output: %w", err)
	}

	return pids, nil
}

// KillProcess kills the process with the given PID.
func KillProcess(pid int) error {
	proc, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("failed to find process %d: %w", pid, err)
	}
	// sends SIGKILL
	return proc.Kill()
}
