module example

go 1.21.3

replace (
	example/util => ./util
	melato.org/trace => ../
)

require (
	example/util v0.0.0-00010101000000-000000000000
	melato.org/command v1.0.0
	melato.org/trace v0.0.0-00010101000000-000000000000
)

require gopkg.in/yaml.v2 v2.4.0 // indirect
