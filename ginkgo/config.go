package ginkgo

import (
	"github.com/onsi/ginkgo/config"
	"os"
	"runtime"
)

func Configure() {
	config.DefaultReporterConfig.SlowSpecThreshold = 70
	config.DefaultReporterConfig.NoColor = runtime.GOOS == "windows" && len(os.Getenv("ANSICON")) < 1
}
