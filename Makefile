#!/bin/sh

GO = main.go

run:
	go run $(GO)

run_watch:
	watchexec -r "go run $(GO)"
