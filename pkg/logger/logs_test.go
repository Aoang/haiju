package logger

import (
	"os"
	"testing"
)

// Logger.
var logger = NewLogger(os.Stdout)

func TestSetLevel(t *testing.T) {
	SetLevel("trace")
}

func TestTrace(t *testing.T) {
	SetLevel("trace")
	logger.Trace("trace")
	SetLevel("off")
	logger.Trace("trace")
}

func TestTracef(t *testing.T) {
	SetLevel("trace")
	logger.Tracef("tracef")
	SetLevel("off")
	logger.Tracef("tracef")
}

func TestDebug(t *testing.T) {
	SetLevel("debug")
	logger.Debug("debug")
	SetLevel("off")
	logger.Debug("debug")
}

func TestDebugf(t *testing.T) {
	SetLevel("debug")
	logger.Debugf("debugf")
	SetLevel("off")
	logger.Debug("debug")
}

func TestInfo(t *testing.T) {
	SetLevel("info")
	logger.Info("info")
	SetLevel("off")
	logger.Info("info")
}

func TestInfof(t *testing.T) {
	SetLevel("info")
	logger.Infof("infof")
	SetLevel("off")
	logger.Infof("infof")
}

func TestWarn(t *testing.T) {
	SetLevel("warn")
	logger.Warn("warn")
	SetLevel("off")
	logger.Warn("warn")
}

func TestWarnf(t *testing.T) {
	SetLevel("warn")
	logger.Warnf("warnf")
	SetLevel("off")
	logger.Warnf("warnf")
}

func TestError(t *testing.T) {
	SetLevel("error")
	logger.Error("error")
	SetLevel("off")
	logger.Error("error")
}

func TestErrorf(t *testing.T) {
	SetLevel("error")
	logger.Errorf("errorf")
	SetLevel("off")
	logger.Errorf("errorf")
}

func TestGetLevel(t *testing.T) {
	if getLevel("trace") != Trace {
		t.FailNow()

		return
	}

	if getLevel("debug") != Debug {
		t.FailNow()

		return
	}

	if getLevel("info") != Info {
		t.FailNow()

		return
	}

	if getLevel("warn") != Warn {
		t.FailNow()

		return
	}

	if getLevel("error") != Error {
		t.FailNow()

		return
	}

	if getLevel("fatal") != Fatal {
		t.FailNow()

		return
	}
}

func TestLoggerSetLevel(t *testing.T) {
	SetLevel("trace")

	if logger.level != Trace {
		t.FailNow()

		return
	}
}

func TestIsTraceEnabled(t *testing.T) {
	SetLevel("trace")

	if !logger.IsTraceEnabled() {
		t.FailNow()

		return
	}
}

func TestIsDebugEnabled(t *testing.T) {
	SetLevel("debug")

	if !logger.IsDebugEnabled() {
		t.FailNow()

		return
	}
}

func TestIsWarnEnabled(t *testing.T) {
	SetLevel("warn")

	if !logger.IsWarnEnabled() {
		t.FailNow()

		return
	}
}
