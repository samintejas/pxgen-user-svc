package log

import (
	"context"
	"log"
	"os"
	"runtime"
	"strings"

	"pxgen.io/user/internal/constants"
)

var logger *log.Logger

const (
	TraceFmt = "TRACE > "
	InfoFmt  = "INFO > "
	DebugFmt = "DEBUG > "
	WarnFmt  = "WARN > "
	ErrorFmt = "ERROR > "
)

type logWriter struct{}

func (lw logWriter) Write(p []byte) (n int, err error) {
	msg := string(p)
	msg = strings.TrimSuffix(msg, "\n")
	Error(msg)
	return len(p), nil
}

func init() {
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	log.SetOutput(logWriter{})
}

func Info(msg string) {
	logger.Printf("%s%s", InfoFmt, msg)
}

func Debug(msg string) {
	logger.Printf("%s%s", DebugFmt, msg)
}

func Warn(msg string) {
	logger.Printf("%s%s", WarnFmt, msg)
}

func Error(msg string) {
	logger.Printf("%s%s\n%s", ErrorFmt, msg, formatStackTrace())
}

func Infof(format string, v ...any) {
	logger.Printf(InfoFmt+format, v...)
}

func Debugf(format string, v ...any) {
	logger.Printf(DebugFmt+format, v...)
}

func Warnf(format string, v ...any) {
	logger.Printf(WarnFmt+format, v...)
}

func Errorf(format string, v ...any) {
	logger.Printf(ErrorFmt+format, v...)
}

func Infoc(ctx context.Context, msg string) {
	logger.Printf("%s%s%s", InfoFmt, addContextProps(ctx), msg)
}

func Debugc(ctx context.Context, msg string) {
	logger.Printf("%s%s%s", InfoFmt, addContextProps(ctx), msg)
}

func Warnc(ctx context.Context, msg string) {
	logger.Printf("%s%s%s", InfoFmt, addContextProps(ctx), msg)
}

func Errorc(ctx context.Context, msg string) {
	logger.Printf("%s%s%s", InfoFmt, addContextProps(ctx), msg)
}

func Infocf(ctx context.Context, format string, v ...any) {
	logger.Printf(InfoFmt+addContextProps(ctx)+format, v...)
}

func Debugcf(ctx context.Context, format string, v ...any) {
	logger.Printf(DebugFmt+addContextProps(ctx)+format, v...)
}

func Warncf(ctx context.Context, format string, v ...any) {
	logger.Printf(WarnFmt+addContextProps(ctx)+format, v...)
}

func Errorcf(ctx context.Context, format string, v ...any) {
	logger.Printf(ErrorFmt+addContextProps(ctx)+format, v...)
}

func addContextProps(ctx context.Context) string {
	prop := ""
	if traceId, ok := ctx.Value(constants.TRACE_ID_KEY).(string); ok {
		prop = prop + "[" + traceId + "] "
	}
	return prop
}

func Logger() *log.Logger {
	return logger
}

func getStackTrace() string {
	stackBuf := make([]byte, 4096)
	n := runtime.Stack(stackBuf, false)
	return string(stackBuf[:n])
}

func formatStackTrace() string {
	stackBuf := make([]byte, 4096)
	n := runtime.Stack(stackBuf, false)
	stackTrace := string(stackBuf[:n])

	lines := strings.Split(stackTrace, "\n")
	var formattedLines []string
	formattedLines = append(formattedLines, "Stack Trace:")
	for _, line := range lines {
		if strings.HasPrefix(line, "runtime/") || strings.HasPrefix(line, "go/") {
			continue
		}
		formattedLines = append(formattedLines, "    "+line)
	}
	return strings.Join(formattedLines, "\n")
}
