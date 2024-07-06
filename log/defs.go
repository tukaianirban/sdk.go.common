package log

//
// enumerated level definitions
//
type LOG_LEVEL int

const (
	LEVEL_DEBUG LOG_LEVEL = iota
	LEVEL_INFO
	LEVEL_WARNING
	LEVEL_ERROR
	LEVEL_ALERT
	LEVEL_FATAL
)

//
// level-based tag definitions
//
const (
	TAG_DEBUG   = "D"
	TAG_INFO    = "I"
	TAG_WARNING = "W"
	TAG_ERROR   = "E"
	TAG_ALERT   = "A"
	TAG_FATAL   = "F"
)
