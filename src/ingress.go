package kugo_src

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"log"
	"strconv"

	model "github.com/tkennes/kugo/model"
)

var IngressTableHeaders = []string{"name", "namespace", "host", "path", "service-name", "service-port"}

func GetIngress(namespace string) []model.Io_k8s_api_networking_v1_Ingress {
	command_canonical := "kubectl get ingress -o json"
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
	return parseIngress(string(out))
}

func GetAndShowIngress(namespace string) (res [][]string) {
	ingresses := GetIngress(namespace)
	
	for _, ingress := range ingresses {
		metadata_name := firstN(*ingress.Metadata.Name, 32)
		metadata_namespace := firstN(*ingress.Metadata.Namespace, 32)
		for _, rule := range ingress.Spec.Rules {
			host := *rule.Host
			fmt.Println(host)
			for _, path := range rule.Http.Paths {
				if path.Backend.Service != nil {
					if path.Backend.Service.Port != nil && path.Backend.Service.Name != nil {
						port := strconv.Itoa(*path.Backend.Service.Port.Number)
						svc_name := firstN(*path.Backend.Service.Name, 32)
						res = append(res, 
							[]string{metadata_name,
								metadata_namespace,
								host, 
								*path.Path,
								svc_name,
								port})
					}
				}
			}
		}
	}
	return
}

func firstN(str string, n int) string {
    if len(str) > n {
		return str[:n] + "..."
	} else {
		return str
	}
}

func parseIngress(responseData string) []model.Io_k8s_api_networking_v1_Ingress {
	var obj model.Io_k8s_api_networking_v1_IngressList
	if err := json.Unmarshal([]byte(responseData), &obj); err != nil {
		log.Fatal(err)
	}
	return obj.Items
}
