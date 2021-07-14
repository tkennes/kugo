package kugo_src

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"log"

	model "github.com/tkennes/kugo/model"
)

type NodeResponse struct {
	APIVersion string   `json:"apiVersion"`
	Items      []model.Io_k8s_api_core_v1_Node  `json:"items"`
	Kind       string   `json:"kind"`
}

var (
	NodeHeaders = []string{"Name", "IP", "Architecture",
		"ContainerRuntimeVersion", "KernelVersion",
		"KubeProxyVersion", "KubeletVersion",
		"OperatingSystem", "OsImage",
	}
)

func GetNodes() []model.Io_k8s_api_core_v1_Node {
	command_canonical := "kubectl get nodes -o json"
	cmd := exec.Command("bash", "-c", command_canonical)
	cmd.Env = os.Environ()
	out, err := cmd.Output()
    if err != nil {
		fmt.Println("ERROR IN EXECUTION: " + command_canonical)
        log.Printf("verbose error info: %#v", err)
	}
	return parseNodes(string(out))
}

func parseNodes(responseData string) []model.Io_k8s_api_core_v1_Node{
	var obj NodeResponse
	if err := json.Unmarshal([]byte(responseData), &obj); err != nil {
		log.Fatal(err)
	}
	return obj.Items
}

func GetAndShowNodes() (res [][]string){
	nodes := GetNodes()

	for _, node := range nodes {
		ip_address := ""
		for _, address := range node.Status.Addresses {
			if *address.Type == "InternalIP" {
				ip_address = *address.Address
			}
		}
		res = append(res,
			[]string{*node.Metadata.Name,
				ip_address,
				*node.Status.NodeInfo.Architecture,
				*node.Status.NodeInfo.ContainerRuntimeVersion,
				*node.Status.NodeInfo.KernelVersion,
				*node.Status.NodeInfo.KubeProxyVersion,
				*node.Status.NodeInfo.KubeletVersion,
				*node.Status.NodeInfo.OperatingSystem,
				*node.Status.NodeInfo.OsImage,
		})
	}
	return
}