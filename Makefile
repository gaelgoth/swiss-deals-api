SWAG=${HOME}/go/bin/swag
# DOCS=docs/docs.go

build:
	go build -o out/server main.go

run: build
	out/server

watch:
#	${SWAG} init -g ${DOCS}
	${SWAG} init
	reflex -s -r '\.go$$' make run