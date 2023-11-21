.PHONY: run bench

run: 
	go run -gcflags=all=-d=checkptr main.go

bench:
	go test -bench=. -benchmem
