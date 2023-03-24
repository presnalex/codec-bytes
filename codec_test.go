//go:build ignore
// +build ignore

package bytes_test

import (
	"testing"

	"go.unistack.org/micro-codec-json/v3"
	"go.unistack.org/micro-codec-jsonpb/v3"
	"go.unistack.org/micro-codec-proto/v3"
	"go.unistack.org/micro/v3/codec"

	bytes "github.com/presnalex/codec-bytes"
)

type testRWC struct{}

func (rwc *testRWC) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (rwc *testRWC) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (rwc *testRWC) Close() error {
	return nil
}

func getCodecs() map[string]codec.Codec {
	return map[string]codec.Codec{
		"bytes":  bytes.NewCodec(),
		"json":   json.NewCodec(),
		"jsonpb": jsonpb.NewCodec(),
		"proto":  proto.NewCodec(),
	}
}

func Test_WriteEmptyBody(t *testing.T) {
	for name, c := range getCodecs() {
		err := c.Write(&testRWC{}, &codec.Message{
			Type:   codec.Error,
			Header: map[string]string{},
		}, nil)
		if err != nil {
			t.Fatalf("codec %s - expected no error when writing empty/nil body: %s", name, err)
		}
	}
}
