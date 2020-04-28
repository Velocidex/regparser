package main

import (
	"fmt"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"www.velocidex.com/golang/regparser"
	"www.velocidex.com/golang/regparser/appcompatcache"
)

var (
	appcompatcache_command = app.Command(
		"appcompatcache", "List the application compatibility cache.")

	appcompatcache_command_file_arg = appcompatcache_command.Arg(
		"file", "Registry hive file",
	).Required().OpenFile(os.O_RDONLY, 0600)
)

const (
	appcompatcache_path = "/ControlSet001/Control/Session Manager/AppCompatCache"
)

func parseAppCompatibilityCache(buffer []byte) {
	for idx, entry := range appcompatcache.ParseValueData(buffer) {
		fmt.Printf("%d: %v  %v\n", idx, entry.Time, entry.Name)
	}
}

func doAppCompatCache() {
	registry, err := regparser.NewRegistry(*appcompatcache_command_file_arg)
	kingpin.FatalIfError(err, "Open hive")

	key := registry.OpenKey(appcompatcache_path)
	if key == nil {
		kingpin.Fatalf("Key path not found %v", appcompatcache_path)
	}

	for _, value := range key.Values() {
		if value.ValueName() != "AppCompatCache" {
			continue
		}

		parseAppCompatibilityCache(value.ValueData().Data)
	}

}
