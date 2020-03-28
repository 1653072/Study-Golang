package LogConf

// Log configs
var (
	// Megabytes after which new file is created
	MaxSize 		= 50
	// Number of backups
	MaxBackups 		= 100
	// Days
	MaxAge			= 14
	// Enable gzip compress
	Compress		= true
)

// Log statuses
const (
	Debug	uint8     	= 0
	Info	uint8		= 1
	Warn	uint8		= 2
	Error	uint8		= 3
	Fatal	uint8		= 4
	Panic	uint8		= 5
)