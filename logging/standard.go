package logging

import "log"

//Standard logs to a log.Logger
//
//EX: logging.Standard{Logger: log.Default()}
type Standard struct {
	Logger log.Logger
}

func (s *Standard) Log(level Level, data string) {
	s.Logger.Printf("%v: %v", level, data)
}
