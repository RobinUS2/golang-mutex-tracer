# golang-mutex-tracer
Helps you debug slow lock/unlocks for golang sync Mutex/RWMutex 

Example usage
====
Import with an alias for minimal impact. 

Enable per lock:
```
import sync "github.com/RobinUS2/golang-mutex-tracer"

l := sync.Mutex{}
l.EnableTracer()
l.Lock()
l.Unlock()
```

Enable with customer settings per lock:
```
import sync "github.com/RobinUS2/golang-mutex-tracer"

l := sync.Mutex{}
l.EnableTracerWithOpts(sync.Opts{
    Threshold: 10 * time.Millisecond,
})
l.Lock()
l.Unlock()
```

Enable for all locks (that use the import):
```
import sync "github.com/RobinUS2/golang-mutex-tracer"

l := sync.Mutex{}
sync.SetGlobalOpts(sync.Opts{
    Threshold: 100 * time.Millisecond,
    Enabled:   true,
})
l.Lock()
l.Unlock()
```