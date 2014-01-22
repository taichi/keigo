package core

import (
	"bytes"
	"fmt"
	"github.com/taichi/keigo/log"
	"time"
)

type Seconds time.Duration // 1～32767 秒, 0は指定なしと同等

func (sec Seconds) write(buffer *bytes.Buffer, opt string) {
	if 0 < sec && sec < 32768 {
		fmt.Fprintf(buffer, " -%s %d", opt, sec)
	} else if 0 == sec {
		//do nothing.
	} else {
		log.Debugf("%d is suppressed. Seconds supports 1-32767", sec)
	}
}
