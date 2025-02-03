package event

import (
	"sync"

	"github.com/IBM/sarama"
)

var (
	offsetManagerInstance *OffsetManager
	offsetOnce            sync.Once
)

// OffsetManager almacena el último offset de Kafka
type OffsetManager struct {
	offset int64
	mutex  sync.Mutex
}

// GetOffsetManager retorna un singleton del OffsetManager
func GetOffsetManager() *OffsetManager {
	offsetOnce.Do(func() {
		offsetManagerInstance = &OffsetManager{offset: sarama.OffsetNewest}
	})
	return offsetManagerInstance
}

// SetOffset actualiza el offset global
func (o *OffsetManager) SetOffset(offset int64) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	o.offset = offset
}

// GetOffset retorna el offset almacenado
func (o *OffsetManager) GetOffset() int64 {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	return o.offset
}
