package pipelines

import (
	"testing"

	"github.com/fengyfei/definitions/pipelines/bits"
	"github.com/fengyfei/definitions/pipelines/operation"
	"github.com/go-playground/assert/v2"
)

func TestDeltaPipeline(t *testing.T) {
	floatVal := float32(1.0)

	pipeline := operation.NewDeltaPipeline(100)
	if err := pipeline.Apply(&floatVal); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, floatVal, float32(101))
}

func TestScalePipeline(t *testing.T) {
	floatVal := float32(1.0)

	pipeline := operation.NewScalePipeline(2)
	if err := pipeline.Apply(&floatVal); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, floatVal, float32(2))
}

func TestBitsPipeline(t *testing.T) {
	intVal := int64(0x4)

	pipeline := bits.NewBitsPipeline(2, 1)
	if err := pipeline.Apply(&intVal); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, intVal, int64(0x1))
}