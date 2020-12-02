package kugo_src

import (
	"strings"
	"strconv"
)

var (
	ContainerHeaders = []string{"name", "namespace", "container"}
	ContainerPortHeaders = []string{"name", "namespace", "container", "port", "protocol"}
	ContainerEnvHeaders = []string{"name", "namespace", "container", "type", "key", "value"}
)

func GetAndShowContainers(namespace string) (res [][]string){
	pods := GetPods(namespace)

	for _, pod := range pods {
		name := pod.Metadata.Name
		namespace := pod.Metadata.Namespace
		for _, container := range pod.Spec.Containers {
			res = append(res,
				[]string{name,
				namespace,
				cleanContainerImageName(container.Image)})
		} 
	}
	return
}

func GetAndShowContainerPorts(namespace string) (res [][]string) {
	pods := GetPods(namespace)

	for _, pod := range pods {
		name := pod.Metadata.Name
		namespace := pod.Metadata.Namespace
		for _, container := range pod.Spec.Containers {
			for _, port := range container.Ports {
				port_ := strconv.Itoa(port.ContainerPort)
				res = append(res,
					[]string{name,
					namespace,
					cleanContainerImageName(container.Image),
					port_,
					port.Protocol})
			}
		} 
	}
	return
}

func GetAndShowContainerEnvs(namespace string) (res [][]string) {
	pods := GetPods(namespace)

	for _, pod := range pods {
		name := pod.Metadata.Name
		namespace := pod.Metadata.Namespace
		for _, container := range pod.Spec.Containers {
			for _, env := range container.Env {
				if env.Value == "" {
					res = append(res,
						[]string{name,
							namespace,
							cleanContainerImageName(container.Image),
							"sec",
							env.ValueFrom.SecretKeyRef.Name,
							env.ValueFrom.SecretKeyRef.Key})
				} else {
					res = append(res,
						[]string{name,
						namespace,
						cleanContainerImageName(container.Image),
						"kv",
						env.Name,
						env.Value})
				}
			}
		} 
	}
	return
}

func cleanContainerImageName(name string) string {
	if len(strings.Split(name, "/")) > 2 {
		return strings.Join(strings.Split(name, "/")[1:], "|")
	}
	return name
}
