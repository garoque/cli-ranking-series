/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/garoque/cli-ranking-series/cmd"
	"github.com/garoque/cli-ranking-series/store"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	store.OpenDatabase()
	cmd.Execute()
}
