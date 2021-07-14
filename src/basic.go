package src

import (
	"fmt"
	"os"
	"os/exec"
	"log"
)

func GetAllResources(namespace string) string {
	command_canonical := "kubectl api-resources --verbs=list --namespaced -o name | xargs -n 1 kubectl get --show-kind --ignore-not-found"
	if namespace != "" {
		command_canonical = command_canonical + " -n " + namespace
	} else {
		command_canonical = command_canonical + " --all-namespaces"
	}
	cmd := exec.Command("bash", "-c", command_canonical)
	cmd.Env = os.Environ()
	out, err := cmd.Output()
    if err != nil {
		fmt.Println("ERROR IN EXECUTION: " + command_canonical)
        log.Fatal(err)
    }
	return string(out)
}
