package yolobeer

type Configuration struct {
	ServerConf serverInstanceConfiguration `toml:"server"`
	YOLOBeers  yoloConfiguration           `toml:"yolo_beers"`
}

type serverInstanceConfiguration struct {
	Host       string `toml:"host"`
	Port       int32  `toml:"port"`
	QueueLimit int    `toml:"queue_limit"`
}

type yoloConfiguration struct {
	Cfg       string  `toml:"cfg"`
	Weights   string  `toml:"weights"`
	Threshold float32 `tom:"thresold"`
}
