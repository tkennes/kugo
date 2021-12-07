package kugo_src

import (
	"os"
	"fmt"
	"log"
	"reflect"

	tablewriter "github.com/olekukonko/tablewriter"
	yaml "gopkg.in/yaml.v2"
)


func Table(jobs [][]string, header []string) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader(header)
	// ALIGN LEFT
	table.SetAlignment(3)
	table.SetCenterSeparator("*")
	table.SetRowSeparator("=")

	for _, v := range jobs {
		table.Append(v)
	}
	table.Render()
}

func YAML(obj interface{}) {
	d, err := yaml.Marshal(&obj)
	if err != nil {
		ErrorLog(err)
	}
	// Log without datetime prefix
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	log.Printf(string(d))
}

func ReflectOnStruct(s interface{}) {
	v := reflect.ValueOf(s)
    typeOfS := v.Type()

    for i := 0; i< v.NumField(); i++ {
        fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
    }
}