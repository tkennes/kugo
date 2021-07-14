package src

import (
	"strconv"
)

var (
	TolerationsHeaders = []string{"name", "namespace", "effect", "key", "value", "operator", "toleration-seconds"}
)

func GetAndShowTolerations(namespace string) (res [][]string) {
	tolerations := GetPods(namespace)

	for _, toleration := range tolerations {
		name := *toleration.Metadata.Name
		namespace := *toleration.Metadata.Namespace
		for _, tol := range toleration.Spec.Tolerations {
			var seconds string = ""
			var effect string = ""
			var value string = ""
			var operator string = ""
			if tol.TolerationSeconds != nil {
				seconds = strconv.Itoa(*tol.TolerationSeconds)
			}
			if tol.Effect != nil {
				effect = *tol.Effect
			}
			if tol.Value != nil {
				value = *tol.Value
			}
			if tol.Operator != nil {
				operator = *tol.Operator
			}
			res = append(res,
				[]string{name,
					namespace,
					effect,
					*tol.Key,
					value,
					operator,
					seconds})
		}
	}
	return
}
