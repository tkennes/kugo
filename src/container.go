package kugo_src

import (
	"fmt"
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
		name := *pod.Metadata.Name
		namespace := *pod.Metadata.Namespace
		for _, container := range pod.Spec.Containers {
			res = append(res,
				[]string{name,
				namespace,
				cleanContainerImageName(*container.Image)})
		}
	}
	return
}

func GetAndShowContainerPorts(namespace string) (res [][]string) {
	pods := GetPods(namespace)

	for _, pod := range pods {
		name := *pod.Metadata.Name
		namespace := *pod.Metadata.Namespace
		for _, container := range pod.Spec.Containers {
			for _, port := range container.Ports {
				port_ := strconv.Itoa(*port.ContainerPort)
				res = append(res,
					[]string{name,
					namespace,
					cleanContainerImageName(*container.Image),
					port_,
					*port.Protocol})
			}
		}
	}
	return
}

func GetAndShowContainerEnvs(namespace string) (res [][]string) {
	pods := GetPods(namespace)

	for _, pod := range pods {
		name := *pod.Metadata.Name
		namespace := *pod.Metadata.Namespace
		for _, container := range pod.Spec.Containers {
			for _, env := range container.Env {
				var var_name string
				var var_value string
				var type_ string
				if env.Value == nil && env.ValueFrom == nil {
					fmt.Println("nill env")
					fmt.Println(env)
					fmt.Println(container)
					continue
				}
				if env.Value == nil {
					if env.ValueFrom.ConfigMapKeyRef != nil {
						var_name = *env.ValueFrom.ConfigMapKeyRef.Key
						var_value = *env.ValueFrom.ConfigMapKeyRef.Name
						type_ = "cfm"
					} else if env.ValueFrom.FieldRef != nil {
						var_name = *env.ValueFrom.FieldRef.ApiVersion
						var_value = *env.ValueFrom.FieldRef.FieldPath
						type_ = "field"
					} else if env.ValueFrom.ResourceFieldRef != nil {
						var_name = *env.ValueFrom.ResourceFieldRef.ContainerName
						var_value = *env.ValueFrom.ResourceFieldRef.Resource
						type_ = "rfield"
					} else {
						var_name = *env.ValueFrom.SecretKeyRef.Name
						var_value = *env.ValueFrom.SecretKeyRef.Key
						type_ = "skv"
					}
				} else {
					var_name = *env.Name
					var_value = *env.Value
				}
				res = append(res,
					[]string{name,
					namespace,
					cleanContainerImageName(*container.Image),
					type_,
					var_name,
					var_value})
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
