// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This package contains utilities for exporting metrics.
package metrics

import (
	"context"

	"github.com/google/exposure-notifications-server/internal/logging"
	"go.uber.org/zap"
)

const logString = "!METRIC! Type = %v cumulative = %v value = %v"

// Exporter defines a generic metric exporter inferface used in this application.
type Exporter interface {
	WriteBool(name string, value bool)
	WriteInt(name string, cumulative bool, value int)
	WriteIntDistribution(name string, cumulative bool, values []int)
	WriteFloat64(name string, cumulative bool, value float64)
	WriteFloat64Distribution(name string, cumulative bool, values []float64)
}

type exporterImpl struct {
	logger *zap.SugaredLogger
}

func NewLogsBasedFromContext(ctx context.Context) Exporter {
	logger := logging.FromContext(ctx)
	return &exporterImpl{
		logger: logger,
	}
}

func NewLogsBasedExporter(log *zap.SugaredLogger) Exporter {
	return &exporterImpl{
		logger: log,
	}
}

func (e *exporterImpl) WriteBool(name string, value bool) {
	e.logger.Infof(logString, name, false, value)
}

func (e *exporterImpl) WriteInt(name string, cumulative bool, value int) {
	e.logger.Infof(logString, name, cumulative, value)
}

func (e *exporterImpl) WriteIntDistribution(name string, cumulative bool, values []int) {
	e.logger.Infof(logString, name, cumulative, values)
}

func (e *exporterImpl) WriteFloat64(name string, cumulative bool, value float64) {
	e.logger.Infof(logString, name, cumulative, value)
}

func (e *exporterImpl) WriteFloat64Distribution(name string, cumulative bool, values []float64) {
	e.logger.Infof(logString, name, cumulative, values)
}
