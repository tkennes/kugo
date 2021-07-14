package kugo_src

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"log"

	model "github.com/tkennes/kugo/model"
)

type PodResponse struct {
	APIVersion string   `json:"apiVersion"`
	Items      []model.Io_k8s_api_core_v1_Pod  `json:"items"`
	Kind       string   `json:"kind"`
}

func GetPods(namespace string) []model.Io_k8s_api_core_v1_Pod {
	command_canonical := "kubectl get pods -o json"
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
        log.Printf("verbose error info: %#v", err)
	}
	return parsePods(string(out))
}

func parsePods(responseData string) []model.Io_k8s_api_core_v1_Pod{ 
	var obj PodResponse
	if err := json.Unmarshal([]byte(responseData), &obj); err != nil {
		log.Fatal(err)
	}
	return obj.Items
}
