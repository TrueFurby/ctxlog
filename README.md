# ctxlog

ctxlog is a logging library for Context which tries to avoid any external dependency lock in.

# Why

The point of ctxlog is to allow you to use your prefered log library and define when you want to log and what level, with out having to hard code it.

# Examples

## standard library log example
```
package main

import (
	"context"
	"log"
) 

func main() {
	ctx := ctxlog.WithLogger(context.Background(), log.Print)
	ctxlog.Log(ctx, "this is a test")
}
```

## logrus example
```
package main

import (
	"context"

	log "github.com/Sirupsen/logrus"
)

func main() {
	ctx := ctxlog.WithLoggerf(ctx.Background(), log.Warnf)
	ctxlog.Log(ctx, "%s %s", "value1", "value2")
}
```


## intercept log example
```
package main

import (
	"context"
	"fmt"
	"log"

	"github.rackspace.com/wilb5254/ctxlog"
)

func logUUID(uuid string) func(v ...interface{}) {
	return func(v ...interface{}) {
		log.Printf("UUID: %s, %s", uuid, fmt.Sprint(v...))
	}
}

func main() {
	ctx := ctxlog.WithLogger(context.Background(), logUUID("12345-6789"))
	ctxlog.Log(ctx, "value1", "value2")
}
```
