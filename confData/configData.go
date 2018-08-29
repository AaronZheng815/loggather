package confData

type Config struct {
	LogLevel   string
	LogPath    string
	CConf      []CollectConf
	ChanSize   int
	KfkSerIp   string
	KfkSerPort int
}

type CollectConf struct {
	LogPath string
	Topic   string
}

// 定义全局变量
var (
	AppConfig *Config
)
