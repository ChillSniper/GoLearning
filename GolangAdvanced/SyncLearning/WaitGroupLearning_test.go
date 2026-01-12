package SyncLearning

import "testing"

func TestUsingChannelsToSync(t *testing.T) {
	UsingChannelsToSync()
}

func TestMyGoroutine(t *testing.T) {
	MyGoroutine()
}

func TestDeadLock(t *testing.T) {
	DeadLock()
}

func TestCircularDependency(t *testing.T) {
	CircularDependency()
}

func TestLearnMap(T *testing.T) {
	LearnMap()
}

func TestLearnSyncMap(t *testing.T) {
	LearnSyncMap()
}

func TestLearningAtomicUsage(t *testing.T) {
	LearnAtomicUsage()
}

func TestLearningSyncPool(t *testing.T) {
	LearnSyncPool()
}
