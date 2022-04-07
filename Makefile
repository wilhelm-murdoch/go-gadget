BIN=go

build:
	${BIN} build -v ./cmd/gadget

test:
	${BIN} test -race -v .

lint:
	staticcheck -f stylish

bench:
	${BIN} test -run . -bench . -benchtime 5s -count 10 -benchmem -cpuprofile cpu.out -memprofile mem.out -trace trace.out

pprof-cpu:
	${BIN} tool pprof -http :8800 cpu.out

pprof-mem:
	${BIN} tool pprof -http :8900 mem.out

trace:
	${BIN} tool trace trace.out

coverage:
	${BIN} test -v -coverprofile cover.out .

watch-template:
	${BIN} run ./cmd/gadget -format template; echo "---boundary---"
	fswatch ./cmd/gadget/main.go README.tpl | while read file; do ${BIN} run ./cmd/gadget --format template; echo "---boundary---"; done

watch-json:
	${BIN} run ./cmd/gadget | jq -r
	fswatch ./cmd/gadget/main.go | while read file; do ${BIN} run ./cmd/gadget | jq -r; echo "---boundary---"; done

watch-debug:
	${BIN} run ./cmd/gadget --format debug
	fswatch ./cmd/gadget/main.go | while read file; do ${BIN} run ./cmd/gadget --format debug; echo "---boundary---"; done