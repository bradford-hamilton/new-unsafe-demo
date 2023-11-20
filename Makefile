.PHONY: run

run: 
	go run -gcflags=all=-d=checkptr main.go
