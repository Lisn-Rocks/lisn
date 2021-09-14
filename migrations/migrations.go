package migrations

import "embed"

//go:embed *.up.sql
var Up embed.FS

// //go:embed *.down.sql
var Down embed.FS
