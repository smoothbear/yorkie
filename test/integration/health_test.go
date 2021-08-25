//go:build integration

/*
 * Copyright 2021 The Yorkie Authors. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package integration

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func TestHealthCheck(t *testing.T) {
	conn, err := createConn()
	assert.NoError(t, err)
	defer func() {
		assert.NoError(t, conn.Close())
	}()

	cli := healthpb.NewHealthClient(conn)
	resp, err := cli.Check(context.Background(), &healthpb.HealthCheckRequest{})
	assert.NoError(t, err)
	assert.Equal(t, resp.Status, healthpb.HealthCheckResponse_SERVING)
}
