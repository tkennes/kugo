package kugo_src

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"log"
	"strings"
	"time"
)

type NamespaceResponse struct {
	APIVersion string            `json:"apiVersion"`
	Items      []NamespaceItem   `json:"items"`
	Kind       string            `json:"kind"`
}
type NamespaceSpec struct {
	Finalizers []string `json:"finalizers"`
}
type NamespaceStatus struct {
	Phase string `json:"phase"`
}
type NamespaceMetadata struct {
	CreationTimestamp time.Time   `json:"creationTimestamp"`
	Name              string      `json:"name"`
	ResourceVersion   string      `json:"resourceVersion"`
	SelfLink          string      `json:"selfLink"`
	UID               string      `json:"uid"`
}
type NamespaceItem struct {
	APIVersion string                    `json:"apiVersion"`
	Kind       string                    `json:"kind"`
	Metadata   NamespaceMetadata         `json:"metadata,omitempty"`
	Spec       NamespaceSpec             `json:"spec"`
	Status     NamespaceStatus           `json:"status"`
}

var (
	NamespacesHeaders = []string{"current", "name", "link"}
)

func GetNamespaces() []NamespaceItem {
	command_canonical := "kubectl get ns -o json"

	cmd := exec.Command("bash", "-c", command_canonical)
	cmd.Env = os.Environ()
	out, err := cmd.Output()
    if err != nil {
		fmt.Println("ERROR IN EXECUTION: " + command_canonical)
        log.Printf("verbose error info: %#v", err)
    }
	return parseNamespaces(string(out))
}

func GetAndShowNamespaces() (res [][]string){
	namespaces := GetNamespaces()
	current_namespace := GetCurrentNamespace()
	current := ""

	for _, namespace := range namespaces {
		if namespace.Metadata.Name == current_namespace {
			current = "CURRENT"
		} else {
			current = ""
		}
		res = append(res,[]string{current,
			namespace.Metadata.Name,
			namespace.Metadata.SelfLink})
	}
	return
}

func GetCurrentNamespace() string {
	command_canonical := "kubectl config view --minify | grep namespace"

	cmd := exec.Command("bash", "-c", command_canonical)
	cmd.Env = os.Environ()
	out, err := cmd.Output()
    if err != nil {
		fmt.Println("command failed")
		fmt.Println(out)
        log.Printf("verbose error info: %#v", err)
	}

	return cleanNamespace(string(out))
}

func SetCurrentnamespace(namespace string) {
	command_canonical := "kubectl config set-context --current --namespace=" + namespace

	cmd := exec.Command("bash", "-c", command_canonical)
	cmd.Env = os.Environ()
	out, err := cmd.Output()
    if err != nil {
		fmt.Println("command failed")
		fmt.Println(out)
        log.Printf("verbose error info: %#v", err)
	}
}

func parseNamespaces(responseData string) []NamespaceItem {
	var obj NamespaceResponse
	if err := json.Unmarshal([]byte(responseData), &obj); err != nil {
		log.Fatal(err)
	}
	return obj.Items
}

func cleanNamespace(namespace string) string {
	return strings.Split(strings.Split(namespace, ": ")[1], "\n")[0]
}
