module main

go 1.19

replace (
	melato.org/command => ../../command
	melato.org/trace => ../
)

require (
	melato.org/command v0.0.0-00010101000000-000000000000
	melato.org/trace v0.0.0-00010101000000-000000000000
)
