package kugo_src

import (
	"errors"
	"fmt"

	model "github.com/tkennes/kugo/model"
)


type VolumeResponse struct {
	APIVersion string   `json:"apiVersion"`
	Items      []model.Io_k8s_api_core_v1_Pod  `json:"items"`
	Kind       string   `json:"kind"`
}

var (
	VolumesHeaders = []string{"name", "namespace", "type", "object-name", "alias", "Reference"}
)

func GetAndShowVolumes(namespace string) (res [][]string) {
	pods := GetPods(namespace)
   
	for _, pod := range pods {
		name := *pod.Metadata.Name
		namespace := *pod.Metadata.Namespace
		for _, vol := range pod.Spec.Volumes {
			type_ := ""
			alias := *vol.Name
			object_name := ""
			reference := ""
			var err error
			if vol.PersistentVolumeClaim != nil {
				type_ = "pvc"
				object_name = *vol.PersistentVolumeClaim.ClaimName
			} else if vol.ConfigMap != nil {
				type_ = "cfm"
				object_name = *vol.ConfigMap.Name
			} else if vol.Secret != nil {
				type_ = "sec"
				object_name = *vol.Secret.SecretName
			} else if vol.EmptyDir != nil {
				type_ = "emptyDir"
			} else if *vol.HostPath.Path != "" {
				type_ = "HostPath"
				object_name = *vol.HostPath.Path
			} else {
				fmt.Println("Found Unknown Storage Class")
				fmt.Println("Volume:")
				fmt.Println(vol)
				fmt.Println("Pod:")
				fmt.Println(pod)
				break
			}

			// Find the reference in the container
			if inSlice([]string{"pvc", "cfm", "sec"}, type_) {
				reference, err = findVolumeMount(pod, alias)
				if err != nil {
					reference, err = findSecretReference(pod, alias)
					if err != nil {
						fmt.Println(fmt.Sprintf("Could not find name: %s in pod %s, namespace %s", alias, name, namespace))
						fmt.Println("Marking as non-used. Report if this is a bug.")
						reference = "NOT USED"
					}
				}
			}
			res = append(res,
				[]string{name,
					namespace,
					type_,
					object_name,
					alias,
					reference})
		}
	}
	return
}

func findVolumeMount(pod model.Io_k8s_api_core_v1_Pod, name string) (string, error) {
	for _, container := range pod.Spec.Containers {
		for _, volume_mount := range container.VolumeMounts {
			if *volume_mount.Name == name {
				return *volume_mount.MountPath, nil
			}
		}
	}
	return "", errors.New(fmt.Sprintf("Could not find volume with name: %s", name))
}

func findSecretReference(pod model.Io_k8s_api_core_v1_Pod, name string) (string, error) {
	for _, container := range pod.Spec.Containers {
		for _, env := range container.Env {
			if env.ValueFrom != nil {
				if env.ValueFrom.SecretKeyRef != nil {
					if *env.ValueFrom.SecretKeyRef.Name == name {
						return "ENV: " + *env.Name, nil
					}
				}
			}
		}
	}
	return "", errors.New(fmt.Sprintf("Cound not find secret key reference with name: %s", name))
}

func inSlice(slice []string, val string) bool {
    for _, item := range slice {
        if item == val {
            return true
        }
    }
    return false
}
