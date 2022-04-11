// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2022 Datadog, Inc.

// Package logrus provides a log/span correlation hook for the sirupsen/logrus package (https://github.com/sirupsen/logrus).
package logrus

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// DDContextLogHook logs any span in the log context by implementing the logrus.Hook interface
type DDContextLogHook struct{}

func (d *DDContextLogHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel, logrus.WarnLevel, logrus.InfoLevel, logrus.DebugLevel, logrus.TraceLevel}
}

func (d *DDContextLogHook) Fire(e *logrus.Entry) error {
	span, found := tracer.SpanFromContext(e.Context)
	if !found {
		return nil
	}
	e.Data["dd.trace_id"] = span.Context().TraceID()
	e.Data["dd.span_id"] = span.Context().SpanID()
	return nil
}
