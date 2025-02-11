package kafka

import (
	"sync"

	"github.com/IBM/sarama"
)

var (
	offsetManagerInstance *OffsetManager
	offsetOnce            sync.Once
)

type OffsetManager struct {
	offset int64
	mutex  sync.Mutex
}

func GetOffsetManager() *OffsetManager {
	offsetOnce.Do(func() {
		offsetManagerInstance = &OffsetManager{offset: sarama.OffsetNewest}
	})
	return offsetManagerInstance
}

func (o *OffsetManager) SetOffset(offset int64) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	o.offset = offset
}

func (o *OffsetManager) GetOffset() int64 {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	return o.offset
}
