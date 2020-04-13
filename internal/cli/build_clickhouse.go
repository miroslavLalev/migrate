// +build clickhouse

package cli

import (
	_ "github.com/ClickHouse/clickhouse-go"
	_ "github.com/miroslavLalev/migrate/v1/database/clickhouse"
)
