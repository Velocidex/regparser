package main

import (
	"fmt"
	"os"
	"path"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"www.velocidex.com/golang/regparser"
)

var (
	app = kingpin.New("regparser",
		"A tool for parsing registry hives.")

	ls_command = app.Command(
		"ls", "List a path in the registry hive")

	ls_command_file_arg = ls_command.Arg(
		"file", "Registry hive file",
	).Required().OpenFile(os.O_RDONLY, 0600)

	ls_command_path = ls_command.Arg(
		"path", "Path to list").Default("").String()
)

func doLs() {
	registry, err := regparser.NewRegistry(*ls_command_file_arg)
	kingpin.FatalIfError(err, "Open hive")

	key := registry.OpenKey(*ls_command_path)
	if key == nil {
		kingpin.Fatalf("Key path not found %v", *ls_command_path)
	}

	fmt.Printf("Listing key %s (%v)\n\n", path.Join(*ls_command_path, key.Name()),
		key.LastWriteTime())

	fmt.Printf("Subkeys:\n")
	for _, subkey := range key.Subkeys() {
		fmt.Printf(" %s - %v\n", subkey.Name(), subkey.LastWriteTime())
	}

	fmt.Printf("\nValues:\n")
	for _, value := range key.Values() {
		fmt.Printf(" %s : %#v\n", value.ValueName(), value.ValueData())
	}
}

func main() {
	app.HelpFlag.Short('h')
	app.UsageTemplate(kingpin.CompactUsageTemplate)
	command := kingpin.MustParse(app.Parse(os.Args[1:]))
	switch command {

	case ls_command.FullCommand():
		doLs()

	case appcompatcache_command.FullCommand():
		doAppCompatCache()

	}
}
