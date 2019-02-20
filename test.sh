#!/usr/bin/env bash
go fmt .
go vet -v .
go test -v --race .
go test -run=XXX -bench=.