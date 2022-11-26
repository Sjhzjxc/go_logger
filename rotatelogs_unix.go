//go:build !windows

package go_logger

import (
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

func GetWriteSyncer(filePath, fileName, linkName string, withConsole bool) (zapcore.WriteSyncer, error) {
	fileWriter, err := rotatelogs.New(
		path.Join(filePath, fileName+"_%Y-%m-%d.log"),
		rotatelogs.WithLinkName(linkName),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		return nil, err
	}
	if withConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), nil
	}
	return zapcore.AddSync(fileWriter), nil
}
