package task

import "context"

type SyncJob func(ctx context.Context) error
type AsyncJob func(ctx context.Context) error
