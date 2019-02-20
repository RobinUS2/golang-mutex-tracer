package muxtracer

import (
	"log"
	"sync/atomic"
	"time"
)

func (m *Mutex) traceBeginAwaitLock() {
	atomic.StoreUint64(&m.beginAwaitLock, now())
}

func (m *Mutex) traceEndAwaitLock(threshold uint64) {
	ts := now() // first obtain the current time
	start := atomic.LoadUint64(&m.beginAwaitLock)
	atomic.StoreUint64(&m.lockObtained, uint64(ts))
	var took uint64
	if start < ts {
		// check for no overflow
		took = ts - start
	}
	if took >= threshold {
		logViolation(m, Threshold(threshold), Actual(took), Now(ts), ViolationLock)
	}
}

func (m *Mutex) traceBeginAwaitUnlock() {
	atomic.StoreUint64(&m.beginAwaitUnlock, now())
}

func (m *Mutex) traceEndAwaitUnlock(threshold uint64) {
	ts := now() // first obtain the current time

	// lock obtained time (critical section)
	lockObtained := atomic.LoadUint64(&m.lockObtained)
	var took uint64
	if lockObtained < ts {
		// check for no overflow
		took = ts - lockObtained
	}

	if took >= threshold && lockObtained > 0 {
		// lockObtained = 0 when the tracer is enabled half way
		logViolation(m, Threshold(threshold), Actual(took), Now(ts), ViolationCritical)
	}
}

func logViolation(m *Mutex, threshold Threshold, actual Actual, now Now, violationType ViolationType) {
	beginAwaitLock := atomic.LoadUint64(&m.beginAwaitLock)
	lockObtained := atomic.LoadUint64(&m.lockObtained)
	beginAwaitUnlock := atomic.LoadUint64(&m.beginAwaitUnlock)
	log.Printf("violation %s section took %s %d (threshold %s, beginAwaitLock %s, lockObtained %s, beginAwaitUnlock %s, now %s)", violationType.String(), time.Duration(actual).String(), actual, time.Duration(threshold).String(), time.Unix(0, int64(beginAwaitLock)).String(), time.Unix(0, int64(lockObtained)).String(), time.Unix(0, int64(beginAwaitUnlock)).String(), time.Unix(0, int64(now)).String())
}

type Threshold uint64
type Actual uint64
type Now uint64
