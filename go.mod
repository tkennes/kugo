module kugo

go 1.15

replace (
	github.com/tkennes/kugo/cmd => ./cmd
	github.com/tkennes/kugo/src => ./src
	github.com/tkennes/kugo/model => ./model
)

require (
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/mapstructure v1.3.3
	github.com/olekukonko/tablewriter v0.0.4
	github.com/spf13/cobra v1.1.1
	gopkg.in/yaml.v2 v2.3.0
)