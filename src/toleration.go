package kugo_src

import (
	"strconv"
)

var (
	TolerationsHeaders = []string{"name", "namespace", "effect", "key", "value", "operator", "toleration-seconds"}
)

func GetAndShowTolerations(namespace string) (res [][]string) {
	tolerations := GetPods(namespace)

	for _, toleration := range tolerations {
		name := toleration.Metadata.Name
		namespace := toleration.Metadata.Namespace
		for _, tol := range toleration.Spec.Tolerations {
			seconds := strconv.Itoa(tol.TolerationSeconds)
			res = append(res,
				[]string{name,
					namespace,
					tol.Effect,
					tol.Key,
					tol.Value,
					tol.Operator,
					seconds})
		}
	}
	return
}
