package mq

import (
	"testing"

	"github.com/Shopify/sarama"
	"github.com/stretchr/testify/assert"
)

func TestKafkaHeaders_AddType(t *testing.T) {
	kh := NewKafkaHeaders()

	kh.SetType("test")
	assert.Equal(t, "test", kh.GetType())

	kh.SetType("twice")
	assert.Equal(t, "twice", kh.GetType())

	kh.SetType("")
	assert.Equal(t, "", kh.GetType())
}

func TestKafkaHeaders_AddTraceID(t *testing.T) {
	kh := NewKafkaHeaders()

	kh.SetTraceID("7f808d7f-00be-4154-8a8f-9df7d798113d")
	assert.Equal(t, "7f808d7f-00be-4154-8a8f-9df7d798113d", kh.GetTraceID())

	kh.SetTraceID("03a832b9-b5c2-40f5-ab19-bce0b5006b68")
	assert.Equal(t, "03a832b9-b5c2-40f5-ab19-bce0b5006b68", kh.GetTraceID())

	kh.SetTraceID("")
	assert.Equal(t, "", kh.GetTraceID())
}

func TestKafkaHeaders_ExportHeaders(t *testing.T) {
	kh := NewKafkaHeaders()

	assert.Len(t, kh.ExportHeaders(), 0)

	kh.SetType("test")
	assert.Len(t, kh.ExportHeaders(), 1)
	assert.Contains(t, kh.ExportHeaders(), sarama.RecordHeader{
		Key:   []byte("type"),
		Value: []byte("test"),
	})

	kh.SetTraceID("test")
	assert.Len(t, kh.ExportHeaders(), 2)
	assert.Contains(t, kh.ExportHeaders(), sarama.RecordHeader{
		Key:   []byte("trace_id"),
		Value: []byte("test"),
	})
}

func TestFromHeaders(t *testing.T) {
	headers := []sarama.RecordHeader{
		{
			Key:   []byte("header0"),
			Value: []byte("value0"),
		},
		{
			Key:   []byte("type"),
			Value: []byte("type0"),
		},
		{
			Key:   []byte("trace_id"),
			Value: []byte("trace"),
		},
	}

	kh := FromHeaders(headers)
	assert.Len(t, kh.ExportHeaders(), 3)
	assert.Equal(t, "type0", kh.GetType())
	assert.Equal(t, "trace", kh.GetTraceID())

	kh.SetType("type1")
	assert.Equal(t, "type1", kh.GetType())

	kh.SetTraceID("trace1")
	assert.Equal(t, "trace1", kh.GetTraceID())
}

func TestFromPointerHeaders(t *testing.T) {
	headers := []*sarama.RecordHeader{
		{
			Key:   []byte("header0"),
			Value: []byte("value0"),
		},
		{
			Key:   []byte("type"),
			Value: []byte("type0"),
		},
		{
			Key:   []byte("trace_id"),
			Value: []byte("trace"),
		},
	}

	kh := FromPointerHeaders(headers)
	assert.Len(t, kh.ExportHeaders(), 3)
	assert.Equal(t, "type0", kh.GetType())
	assert.Equal(t, "trace", kh.GetTraceID())

	kh.SetType("type1")
	assert.Equal(t, "type1", kh.GetType())

	kh.SetTraceID("trace1")
	assert.Equal(t, "trace1", kh.GetTraceID())
}
