package config_parser_test

import (
	"testing"

	"github.com/millirud/intermocker/internal/service/config_parser"
)

const config_1 = `
{
    "routes" : [
        {
            "req": {
                "path": "/test/api",
                "method": "POST"
            },
            "res": {
                "type": "application/json",
                "content": {
                    "success": "true"
                },
                "timeout": 15000
            }
        }
    ]
}
`

func TestParser(t *testing.T) {

	service := config_parser.New()

	service.ParseConfig(config_1)
}
