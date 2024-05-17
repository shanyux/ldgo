/*
 * Copyright (C) distroy
 */

package ldlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newWrapper(log *zap.Logger, sugar *zap.SugaredLogger) Wrapper {
	return Wrapper{
		log:   log,
		sugar: sugar,
	}
}

type Wrapper struct {
	log      *zap.Logger
	sugar    *zap.SugaredLogger
	sequence string
	rate     rateConfig
}

func (l *Wrapper) Sync() { l.log.Sync() }

func (l *Wrapper) Enabled(lvl zapcore.Level) bool { return l.enabled(lvl) }

func (l *Wrapper) Logger() *Logger           { return newLogger(*l) }
func (l *Wrapper) Core() *zap.Logger         { return l.log }
func (l *Wrapper) Sugar() *zap.SugaredLogger { return l.sugar }

func (l *Wrapper) zCore(skip int) *zap.Logger {
	if !l.checkRateOrInterval(skip + 1) {
		return Discard().Core()
	}
	return l.log
}
func (l *Wrapper) zSugar(skip int) *zap.SugaredLogger {
	if !l.checkRateOrInterval(skip + 1) {
		return Discard().Sugar()
	}
	return l.sugar
}

func (l *Wrapper) enabled(lvl zapcore.Level) bool { return l.log.Core().Enabled(lvl) }

func (l *Wrapper) Debugf(fmt string, args ...interface{}) { l.zSugar(1).Debugf(fmt, args...) }
func (l *Wrapper) Debug(args ...interface{})              { l.zSugar(1).Debug(pw(args)) }
func (l *Wrapper) Debugln(args ...interface{})            { l.zSugar(1).Debug(pw(args)) }
func (l *Wrapper) Debugz(fmt string, fields ...zap.Field) { l.zCore(1).Debug(fmt, fields...) }

func (l *Wrapper) Infof(fmt string, args ...interface{}) { l.zSugar(1).Infof(fmt, args...) }
func (l *Wrapper) Info(args ...interface{})              { l.zSugar(1).Info(pw(args)) }
func (l *Wrapper) Infoln(args ...interface{})            { l.zSugar(1).Info(pw(args)) }
func (l *Wrapper) Infoz(fmt string, fields ...zap.Field) { l.zCore(1).Info(fmt, fields...) }

func (l *Wrapper) Printf(fmt string, args ...interface{}) { l.zSugar(1).Infof(fmt, args...) }
func (l *Wrapper) Print(args ...interface{})              { l.zSugar(1).Info(pw(args)) }
func (l *Wrapper) Println(args ...interface{})            { l.zSugar(1).Info(pw(args)) }
func (l *Wrapper) Printz(fmt string, fields ...zap.Field) { l.zCore(1).Info(fmt, fields...) }

func (l *Wrapper) Logf(fmt string, args ...interface{}) { l.zSugar(1).Infof(fmt, args...) }
func (l *Wrapper) Log(args ...interface{})              { l.zSugar(1).Info(pw(args)) }
func (l *Wrapper) Logln(args ...interface{})            { l.zSugar(1).Info(pw(args)) }
func (l *Wrapper) Logz(fmt string, fields ...zap.Field) { l.zCore(1).Info(fmt, fields...) }

func (l *Wrapper) Warnf(fmt string, args ...interface{}) { l.zSugar(1).Warnf(fmt, args...) }
func (l *Wrapper) Warn(args ...interface{})              { l.zSugar(1).Warn(pw(args)) }
func (l *Wrapper) Warnln(args ...interface{})            { l.zSugar(1).Warn(pw(args)) }
func (l *Wrapper) Warnz(fmt string, fields ...zap.Field) { l.zCore(1).Warn(fmt, fields...) }

func (l *Wrapper) Warningf(fmt string, args ...interface{}) { l.zSugar(1).Warnf(fmt, args...) }
func (l *Wrapper) Warning(args ...interface{})              { l.zSugar(1).Warn(pw(args)) }
func (l *Wrapper) Warningln(args ...interface{})            { l.zSugar(1).Warn(pw(args)) }
func (l *Wrapper) Warningz(fmt string, fields ...zap.Field) { l.zCore(1).Warn(fmt, fields...) }

func (l *Wrapper) Errorf(fmt string, args ...interface{}) { l.Sugar().Errorf(fmt, args...) }
func (l *Wrapper) Error(args ...interface{})              { l.Sugar().Error(pw(args)) }
func (l *Wrapper) Errorln(args ...interface{})            { l.Sugar().Error(pw(args)) }
func (l *Wrapper) Errorz(fmt string, fields ...zap.Field) { l.Core().Error(fmt, fields...) }

func (l *Wrapper) Fatalf(fmt string, args ...interface{}) { l.Sugar().Fatalf(fmt, args...) }
func (l *Wrapper) Fatal(args ...interface{})              { l.Sugar().Fatal(pw(args)) }
func (l *Wrapper) Fatalln(args ...interface{})            { l.Sugar().Fatal(pw(args)) }
func (l *Wrapper) Fatalz(fmt string, fields ...zap.Field) { l.Core().Fatal(fmt, fields...) }

func (l *Wrapper) V(v int) bool {
	if v <= 0 {
		return !l.Enabled(zapcore.DebugLevel)
	}
	return true
}
