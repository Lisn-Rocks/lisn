package database

import (
	"testing"

	"github.com/lisn-rocks/lisn/configs"
	"github.com/lisn-rocks/lisn/lisntest"
)

func init() {
	configs.Default()
}

func TestInit(t *testing.T) {
	lisntest.SkipIfNotIntegration(t)
	db := Init()
	db.Down()
}
