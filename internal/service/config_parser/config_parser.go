package config_parser

import (
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog"
)

func New(logger *zerolog.Logger) *ConfigParser {
	return &ConfigParser{
		logger: logger,
	}
}

type configRequest struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

type configResponse struct {
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
	Timeout int64       `json:"timeout"`
}

type configRoute struct {
	Req configRequest  `json:"req"`
	Res configResponse `json:"res"`
}

type configModel struct {
	Routes []configRoute `json:"routes"`
}

type ConfigParser struct {
	logger *zerolog.Logger
}

func (c *ConfigParser) parse(config string) (configModel, error) {
	var parsed configModel

	err := json.Unmarshal([]byte(config), &parsed)

	if err != nil {
		c.logger.Err(err).Msg("ConfigParser.parse.Error")
		return parsed, err
	}

	return parsed, nil
}

func (c *ConfigParser) ParseConfig(config string) error {

	parsed, err := c.parse(config)

	if err != nil {
		return err
	}

	for _, route := range parsed.Routes {
		fmt.Println(route.Req.Method)
	}

	return nil
}
