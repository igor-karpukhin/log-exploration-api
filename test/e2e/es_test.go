package e2e

import (
	"testing"

	"github.com/ViaQ/log-exploration-api/pkg/configuration"
	"github.com/ViaQ/log-exploration-api/pkg/elastic"
	"go.uber.org/zap"
)

func Test_E2EGetAllLogs(t *testing.T) {
	esconfig := &configuration.ElasticsearchConfig{
		UseTLS:    false,
		EsAddress: "http://localhost:9200",
	}

	es, err := elastic.NewElasticRepository(zap.L(), esconfig)
	if err != nil {
		t.Fatal(err)
	}
	logs, err := es.GetAllLogs()
	if err != nil {
		t.Fatal()
	}
	t.Log("got logs from ES:", logs)
	//TODO: verify logs
}
