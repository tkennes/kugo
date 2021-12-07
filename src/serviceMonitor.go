package kugo_src

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"log"
	"time"
)

type ServiceMonitorResponse struct {
	APIVersion string   `json:"apiVersion"`
	Items      []ServicemonitorItems  `json:"items"`
	Kind       string   `json:"kind"`
	Metadata   ServicemonitorMetadata `json:"metadata"`
}

type ServicemonitorMetadata struct {
	Labels            ServicemonitorLabels          `json:"labels"`
	Name              string          `json:"name"`
	Namespace         string          `json:"namespace"`
	ResourceVersion   string          `json:"resourceVersion"`
	UID               string          `json:"uid"`
	SelfLink          string 		  `json:"selfLink"`
	CreationTimestamp time.Time       `json:"creationTimestamp"`
	Generation        int             `json:"generation"`
}

type ServicemonitorParams struct {
	Module []string `json:"module"`
	Target []string `json:"target"`
}

type ServicemonitorRelabelings struct {
	Replacement string `json:"replacement"`
	TargetLabel string `json:"targetLabel"`
}

type ServicemonitorNamespaceSelector struct {
	MatchNames []string `json:"matchNames"`
}

type ServicemonitorMatchLabels struct {
	AppKubernetesIoInstance string `json:"app.kubernetes.io/instance"`
	AppKubernetesIoName     string `json:"app.kubernetes.io/name"`
}

type ServicemonitorSelector struct {
	MatchLabels ServicemonitorMatchLabels `json:"matchLabels"`
}

type ServicemonitorMetricRelabelings struct {
	SourceLabels []string `json:"sourceLabels,omitempty"`
	TargetLabel  string   `json:"targetLabel"`
	Replacement  string   `json:"replacement,omitempty"`
}

type ServicemonitorEndpoints struct {
	Interval          string              				`json:"interval"`
	MetricRelabelings []ServicemonitorMetricRelabelings `json:"metricRelabelings"`
	Params            ServicemonitorParams              `json:"params"`
	Path              string              				`json:"path"`
	Port              string              				`json:"port"`
	Scheme            string              				`json:"scheme"`
	ScrapeTimeout     string              				`json:"scrapeTimeout"`
	HonorTimestamps   bool          		 			`json:"honorTimestamps"`
	Relabelings       []ServicemonitorRelabelings 		`json:"relabelings"`
}

type ServicemonitorSpec struct {
	Endpoints         []ServicemonitorEndpoints       `json:"endpoints"`
	JobLabel          string            `json:"jobLabel"`
	NamespaceSelector ServicemonitorNamespaceSelector `json:"namespaceSelector"`
	Selector          ServicemonitorSelector          `json:"selector"`
}

type ServicemonitorLabels struct {
	App                string `json:"app"`
	Chart              string `json:"chart"`
	Heritage           string `json:"heritage"`
	IoCattleFieldAppID string `json:"io.cattle.field/appId"`
	Release            string `json:"release"`
	Source             string `json:"source"`
	ArgocdArgoprojIoInstance string `json:"argocd.argoproj.io/instance"`
	Prometheus               string `json:"prometheus"`
}

type ServicemonitorItems struct {
	APIVersion string                 `json:"apiVersion"`
	Kind       string                 `json:"kind"`
	Metadata   ServicemonitorMetadata `json:"metadata,omitempty"`
	Spec       ServicemonitorSpec     `json:"spec,omitempty"`
}

var (
	ServiceMonitorHeaders = []string{"Name", "Namespace", "Target"}
)

var (
	ServiceMonitorRelabelingsHeaders = []string{"Name", "Namespace", "Replacement", "TargetLabel"}
)

func GetServiceMonitors() []ServicemonitorItems {
	command_canonical := "kubectl get servicemonitors -A -o json"
	cmd := exec.Command("bash", "-c", command_canonical)
	cmd.Env = os.Environ()
	out, err := cmd.Output()
    if err != nil {
		fmt.Println("ERROR IN EXECUTION: " + command_canonical)
        log.Printf("verbose error info: %#v", err)
	}
	return parseServiceMonitors(string(out))
}

func parseServiceMonitors(responseData string) []ServicemonitorItems {
	var obj ServiceMonitorResponse
	if err := json.Unmarshal([]byte(responseData), &obj); err != nil {
		log.Fatal(err)
	}
	return obj.Items
}

func GetAndShowServiceMonitors() (res [][]string){
	serviceMonitors := GetServiceMonitors()
	var target string = ""

	for _, serviceMonitor := range serviceMonitors {
		for _, endpoint := range serviceMonitor.Spec.Endpoints {
			if endpoint.Params.Target != nil {
				target = endpoint.Params.Target[0]
			}
		}
		res = append(res,
			[]string{serviceMonitor.Metadata.Name,
				serviceMonitor.Metadata.Namespace,
				target,
		})
	}
	return
}

func GetAndShowServiceMonitorsRelabelings() (res [][]string){
	serviceMonitors := GetServiceMonitors()

	for _, serviceMonitor := range serviceMonitors {
		for _, endpoint := range serviceMonitor.Spec.Endpoints {
			for _, relabelings := range endpoint.Relabelings {
				res = append(res,
					[]string{serviceMonitor.Metadata.Name,
						serviceMonitor.Metadata.Namespace,
						relabelings.Replacement,
						relabelings.TargetLabel,
				})
			}
		}
	}
	return
}