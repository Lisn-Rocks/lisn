package configs

import "github.com/pborman/getopt/v2"

type Flags struct {
	ConfigPath *string
}

func parseFlags() (flags Flags) {
	defer getopt.ParseV2()

	flags.ConfigPath = getopt.StringLong(
		"conf", 'c', "", "Specify custom config path")

	return
}
