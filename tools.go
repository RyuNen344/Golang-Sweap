// +build tools

package tools

import (
	_ "github.com/GoogleCloudPlatform/cloud-sql-proxy/v2"
	_ "github.com/GoogleCloudPlatform/berglas/v2/pkg/auto"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
)
