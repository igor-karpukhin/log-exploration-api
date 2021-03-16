#!/bin/bash

populate_es() {
	curl -v -X POST -H 'Content-type: application/json' http://localhost:9200/app-000001/1 -d '{"msg":"test-message-app", "tag": "logs"}'
	curl -v -X POST -H 'Content-type: application/json' http://localhost:9200/infra-000001/1 -d '{"msg":"test-message-infra", "tag": "logs"}'
	curl -v -X POST -H 'Content-type: application/json' http://localhost:9200/audit-000001/1 -d '{"msg":"test-message-audit", "tag": "logs"}'
}


populate_es
