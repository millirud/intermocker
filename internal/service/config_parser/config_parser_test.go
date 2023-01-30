package config_parser_test

import (
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/millirud/intermocker/internal/service/config_parser"
	"github.com/millirud/intermocker/pkg/logger"
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
                "json": {
                    "success": "true"
                },
                "timeout": 5000
            }
        },
        {
            "req": {
                "path": "/view/image",
                "method": "GET"
            },
            "res": {
                "type": "image/png",
                "file": "./params/demo.png",
                "timeout": 3000
            }
        },
        {
            "req": {
                "path": "/view/image",
                "method": "GET"
            },
            "res": {
                "type": "text/html",
                "file": "./params/index.html",
                "timeout": 3000
            }
        }
    ]
}
`

func TestParser(t *testing.T) {

	service := config_parser.New(
		logger.New("debug", os.Stdout),
	)

	service.ParseConfig(config_1)

	assert.Equal(t, 1, 1)

}
