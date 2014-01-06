touch
=====

like unix touch
change a file's timestamp, or create a file

Build

- Use go build tools.

$ go build touch.go
or
$ go install

Cross-compile

$ gox -output="touch-{{.OS}}-{{.Arch}}"

TODO

- implement all the unix flags
