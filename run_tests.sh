#!/usr/bin/sh

# go test "github.com/DylanSp/fuzzing-for-binary-search-overflow/search"
go test -fuzz "FuzzBinarySearch" "github.com/DylanSp/fuzzing-for-binary-search-overflow/search"