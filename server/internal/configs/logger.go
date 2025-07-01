package configs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

// InitZap 初始化完备的 zap 日志系统
func InitZap(z Zap) *zap.Logger {
	// 1. 解析日志级别
	level, err := zap.ParseAtomicLevel(z.Level)
	if err != nil {
		level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	// 2. 创建日志目录（如果不存在）
	if _, err := os.Stat(z.Director); os.IsNotExist(err) {
		if err := os.MkdirAll(z.Director, 0755); err != nil {
			panic("make logger file err!")
			return nil
		}
	}

	// 3. 配置编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 4. 创建多个写入器（控制台 + 文件）
	consoleDebugging := zapcore.Lock(os.Stdout)
	fileWriter := getLogWriter(z.Director,z.RetentionDay)

	// 5. 创建多个核心（不同级别不同处理）
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel && lvl >= level.Level()
	})

	// 6. 组合核心
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), consoleDebugging, highPriority),
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), fileWriter, highPriority),
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), consoleDebugging, lowPriority),
		zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), fileWriter, lowPriority),
	)

	// 7. 构建 logger
	logger := zap.New(core)

	if z.Prefix != "" {
		logger = logger.With(zap.String("prefix",z.Prefix))
	}

	if z.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	logger = logger.WithOptions(zap.AddStacktrace(zap.ErrorLevel))

	return logger
}

// 自定义时间编码器
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// 获取日志写入器（使用 lumberjack 实现日志分割）
func getLogWriter(logDir string,maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logDir + "/app.log",
		MaxSize:    10,   // 单个文件最大 10MB
		MaxBackups: 30,   // 最多保留 30 个备份
		MaxAge:     7,    // 最多保留 7 天
		Compress:   true, // 压缩旧日志
	}
	return zapcore.AddSync(lumberJackLogger)
}