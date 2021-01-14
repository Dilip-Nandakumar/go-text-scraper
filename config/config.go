package config

import "flag"

type Config struct {
	URL   string
	Depth int
}

func NewConfig() Config {
	url := flag.String("url", "", "URL to scrap")
	depth := flag.Int("depth", 1, "maximum number of levels you have to expand for urls within a url")

	flag.Parse()

	config := Config{
		URL:   *url,
		Depth: *depth,
	}

	return config
}
