package main

import (
	"context"

	_ "github.com/redpanda-data/benthos/v4/public/components/io"
	_ "github.com/redpanda-data/benthos/v4/public/components/pure"
	"github.com/redpanda-data/benthos/v4/public/service"
	_ "github.com/sagikazarmark/benthos-qrcode"
)

func main() {
	service.RunCLI(context.Background())
}
