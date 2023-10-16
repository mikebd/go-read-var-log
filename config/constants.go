package config

const (
	HttpPort             = 8080
	LargeFileBytes       = 10_000_000
	LargeFileSeekBytes   = 1_000_000
	LogDirectory         = "/var/log"
	MaxResultLines       = 2_000
	NumberOfLogLines     = 25
	UnsupportedFileTypes = "gz|zip|bz2|tar|tgz|7z|rar"
)
