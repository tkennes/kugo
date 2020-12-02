module kugo

go 1.15

replace (
	github.com/tkennes/kugo/cmd => ./cmd
	github.com/tkennes/kugo/src => ./src
	github.com/tkennes/kugo/model => ./model
)

require (
	github.com/tkennes/kugo/cmd v0.0.0-00010101000000-000000000000
	github.com/tkennes/kugo/src v0.0.0-00010101000000-000000000000
)
