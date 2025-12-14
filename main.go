package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"nuke-port/internal/process"
)

func main() {
	force := flag.Bool("force", false, "Force kill without confirmation")
	flag.BoolVar(force, "f", false, "Force kill without confirmation (shorthand)")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: nuke [--force] <port>")
		os.Exit(1)
	}

	port := args[0]

	pids, err := process.FindPids(port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error finding processes: %v\n", err)
		os.Exit(1)
	}

	if len(pids) == 0 {
		fmt.Printf("No process found listening on port %s\n", port)
		return
	}

	if !*force {
		fmt.Printf("Found process(es) %v listening on port %s. Kill them? (y/N): ", pids, port)
		reader := bufio.NewReader(os.Stdin)
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(strings.ToLower(response))

		if response != "y" && response != "yes" {
			fmt.Println("Aborted.")
			return
		}
	}

	for _, pid := range pids {
		err := process.KillProcess(pid)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to kill process %d: %v\n", pid, err)
		} else {
			fmt.Printf("Process %d killed.\n", pid)
		}
	}
}
