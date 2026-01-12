package SyncLearning

import "sync"

type Config struct{}

var instance *Config
var once sync.Once

func InitConfig() *Config {
	once.Do(func() {
		// 注意！这个once保证了传入的那个匿名函数只会被执行一次，这样的话，实际上起到了单例模式的作用。
		instance = &Config{}
	})

	return instance
}
