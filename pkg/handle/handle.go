package handle

import "github.com/yanshuf0/cross/pkg/data"

// Env will be the receiver for our handlers. It can pass various
// environment variables we may need in our handlers, like the
// database pool as implement here, along with loggers, templates etc.
type Env struct {
	DB        data.Datastore
	AssetsDir *string
}
