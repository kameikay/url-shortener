package db

import (
	"os"
	"testing"

	"github.com/kameikay/url-shortener/tests/integration"
)

func TestMain(m *testing.M) {
	os.Exit(integration.TestMainIntegrationDB(m, "../../../sql/migrations"))
}
