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

// This package is the service that deletes old exposure keys; it is intended to be invoked over HTTP by Cloud Scheduler.
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/exposure-notifications-server/internal/api/wipeout"
	"github.com/google/exposure-notifications-server/internal/database"
	"github.com/google/exposure-notifications-server/internal/logging"
	"github.com/google/exposure-notifications-server/internal/serverenv"
)

const (
	timeoutEnvVar  = "WIPEOUT_TIMEOUT"
	defaultTimeout = 10 * time.Minute
)

func main() {
	ctx := context.Background()
	logger := logging.FromContext(ctx)

	timeout := serverenv.ParseDuration(ctx, timeoutEnvVar, defaultTimeout)
	logger.Infof("Using timeout %v (override with $%s)", timeout, timeoutEnvVar)

	env, err := serverenv.New(ctx, serverenv.WithSecretManager)
	if err != nil {
		logger.Fatalf("unable to connect to secret manager: %v", err)
	}

	db, err := database.NewFromEnv(ctx, env)
	if err != nil {
		logger.Fatalf("unable to connect to database: %v", err)
	}
	defer db.Close(ctx)

	http.Handle("/", wipeout.NewExportHandler(db, timeout))
	logger.Info("starting export wipeout server")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", env.Port()), nil))
}
