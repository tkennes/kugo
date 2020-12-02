package kugo_src

import (
	"encoding/json"
	"os"
	"os/exec"
	"log"
	"strconv"
	"time"
)

type IngressResponse struct {
	APIVersion string   			  `json:"apiVersion"`
	IngressItems      []IngressItem   `json:"items"`
	Kind       string   			   `json:"kind"`
	Metadata   MetadataIngress `json:"metadata"`
}
type Annotations struct {
	IngressKubernetesIoServerAlias              string `json:"ingress.kubernetes.io/server-alias"`
	KubectlKubernetesIoLastAppliedConfiguration string `json:"kubectl.kubernetes.io/last-applied-configuration"`
}
type MetadataIngress struct {
	Annotations       Annotations `json:"annotations"`
	CreationTimestamp time.Time   `json:"creationTimestamp"`
	Generation        int         `json:"generation"`
	Name              string      `json:"name"`
	Namespace         string      `json:"namespace"`
	ResourceVersion   string      `json:"resourceVersion"`
	SelfLink          string      `json:"selfLink"`
	UID               string      `json:"uid"`
}
type Backend struct {
	ServiceName string `json:"serviceName"`
	ServicePort int    `json:"servicePort"`
}
type Paths struct {
	Backend Backend `json:"backend"`
	Path    string  `json:"path"`
}
type HTTP struct {
	Paths []Paths `json:"paths"`
}
type Rules struct {
	Host string `json:"host"`
	HTTP HTTP   `json:"http"`
}
type SpecIngress struct {
	Rules []Rules `json:"rules"`
}
type IngressItem struct {
	APIVersion string   `json:"apiVersion"`
	Kind       string   `json:"kind"`
	MetadataIngress   MetadataIngress `json:"metadata"`
	SpecIngress       SpecIngress     `json:"spec"`
}

var IngressTableHeaders = []string{"name", "namespace", "host", "path", "service-name", "service-port"}

func GetIngress(namespace string) []IngressItem {
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
        log.Fatal(err)
    }
	return parseIngress(string(out))
}

func GetAndShowIngress(namespace string) (res [][]string) {
	ingresses := GetIngress(namespace)
	
	for _, ingress := range ingresses {
		metadata_name := firstN(ingress.MetadataIngress.Name, 32)
		metadata_namespace := firstN(ingress.MetadataIngress.Namespace, 32)
		for _, rule := range ingress.SpecIngress.Rules {
			host := rule.Host
			for _, path := range rule.HTTP.Paths {
				port := strconv.Itoa(path.Backend.ServicePort)
				svc_name := firstN(path.Backend.ServiceName, 32)
				res = append(res, 
					[]string{metadata_name,
						metadata_namespace,
						host, 
						path.Path,
						svc_name,
						port})
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

func parseIngress(responseData string) []IngressItem {
	var obj IngressResponse
	if err := json.Unmarshal([]byte(responseData), &obj); err != nil {
		log.Fatal(err)
	}
	return obj.IngressItems
}
