package log

import "github.com/RichardKnop/logging"

var (
	logger = logging.New(nil, nil, new(logging.ColouredFormatter))

	// INFO ...
	INFO = logger[logging.INFO]
	// WARNING ...
	WARNING = logger[logging.WARNING]
	// ERROR ...
	ERROR = logger[logging.ERROR]
	// FATAL ...
	FATAL = logger[logging.FATAL]
)
