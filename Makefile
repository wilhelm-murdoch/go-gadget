BIN=go1.18

build:
	${BIN} build -v .

test:
	${BIN} test -race -v .

bench:
	${BIN} test -benchmem -count 3 -bench .

coverage:
	${BIN} test -v -coverprofile cover.out .

watch-template:
	${BIN} run . --path stub --format template; echo "---boundary---"
	fswatch main.go README.tpl | while read file; do ${BIN} run . --path stub --format template; echo "---boundary---"; done

watch-json:
	${BIN} run . --path stub --format json | jq -r
	fswatch main.go | while read file; do ${BIN} run . --path stub --format json | jq -r; echo "---boundary---"; done

watch-debug:
	${BIN} run . --path stub --format debug
	fswatch main.go | while read file; do ${BIN} run . --path stub --format debug; echo "---boundary---"; done