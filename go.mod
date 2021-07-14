module kugo

go 1.15

replace github.com/tkennes/kugo/cmd => ./cmd

replace github.com/tkennes/kugo/model => ./model

replace github.com/tkennes/kugo/src => ./src

require (
	github.com/olekukonko/tablewriter v0.0.4
	github.com/spf13/cobra v1.1.1
	github.com/tkennes/kugo v1.0.15
	gopkg.in/yaml.v2 v2.3.0
)
