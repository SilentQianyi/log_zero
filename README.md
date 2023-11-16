### zerolog + lumberjack + rotatelogs 实现zerolog 的根据时间和大小双重切割

```
type Config struct {
	Filename     string `json:"Filename"`           // 地址+文件名
	MaxSize      int    `json:"MaxSize"`            // 大小切割配置 单位b
	MaxAge       int    `json:"MaxAge"`             // 文件保存天数
	MaxBackups   int    `json:"MaxBackups"`         // 保留的旧日志文件的最大数量
	Compress     bool   `json:"Compress"`           // 是否压缩旧日志文件
	RotationTime int    `json:"RotationTime"`       // 日志切割时间间隔，单位：小时
}
```
