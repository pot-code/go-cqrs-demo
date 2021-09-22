package mq

import "github.com/Shopify/sarama"

type KafkaHeaders struct {
	headers []*sarama.RecordHeader
	reg     map[string]*sarama.RecordHeader
}

func NewKafkaHeaders() *KafkaHeaders {
	return &KafkaHeaders{headers: make([]*sarama.RecordHeader, 0), reg: make(map[string]*sarama.RecordHeader)}
}

func (kh *KafkaHeaders) SetType(typeName string) *KafkaHeaders {
	if h, ok := kh.reg["type"]; !ok {
		h = &sarama.RecordHeader{
			Key:   []byte("type"),
			Value: []byte(typeName),
		}
		kh.reg["type"] = h
		kh.headers = append(kh.headers, h)
	} else {
		h.Value = []byte(typeName)
	}
	return kh
}

func (kh *KafkaHeaders) SetTraceID(traceId string) *KafkaHeaders {
	if h, ok := kh.reg["trace_id"]; !ok {
		h = &sarama.RecordHeader{
			Key:   []byte("trace_id"),
			Value: []byte(traceId),
		}
		kh.reg["trace_id"] = h
		kh.headers = append(kh.headers, h)
	} else {
		h.Value = []byte(traceId)
	}

	return kh
}

func (kh *KafkaHeaders) getHeader(key string) string {
	if v, ok := kh.reg[key]; ok {
		return string(v.Value)
	}

	return ""
}

func (kh *KafkaHeaders) GetType() string {
	return kh.getHeader("type")
}

func (kh *KafkaHeaders) GetTraceID() string {
	return kh.getHeader("trace_id")
}

func (kh *KafkaHeaders) ExportHeaders() []sarama.RecordHeader {
	exp := make([]sarama.RecordHeader, len(kh.headers))
	for i, h := range kh.headers {
		exp[i] = *h
	}
	return exp
}

func FromHeaders(headers []sarama.RecordHeader) *KafkaHeaders {
	hs := make([]*sarama.RecordHeader, len(headers))
	reg := make(map[string]*sarama.RecordHeader)
	for i := len(hs) - 1; i > -1; i-- {
		header := &headers[i]
		hs[i] = header
		reg[string(header.Key)] = header
	}
	return &KafkaHeaders{
		headers: hs,
		reg:     reg,
	}
}

func FromPointerHeaders(headers []*sarama.RecordHeader) *KafkaHeaders {
	reg := make(map[string]*sarama.RecordHeader)
	for _, h := range headers {
		reg[string(h.Key)] = h
	}
	return &KafkaHeaders{
		headers: headers,
		reg:     reg,
	}
}
