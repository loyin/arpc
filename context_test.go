// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package arpc

import (
	"reflect"
	"testing"
)

func TestContext_Body(t *testing.T) {
	ctx := &Context{
		Client:  &Client{Codec: DefaultCodec},
		Message: Message([]byte{0, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8}),
	}
	if got := ctx.Body(); !reflect.DeepEqual(got, []byte{1, 2, 3, 4, 5, 6, 7, 8}) {
		t.Errorf("Context.Body() = %v, want %v", got, []byte{1, 2, 3, 4, 5, 6, 7, 8})
	}
}

func TestContext_Bind(t *testing.T) {
	ctx := &Context{
		Client:  &Client{Codec: DefaultCodec},
		Message: Message([]byte{1, 0, 1, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 'a', 'b', 'c', 'd'}),
	}
	if err := ctx.Bind(nil); err == nil {
		t.Errorf("Context.Bind() error = nil, want %v", err)
	}
}

func TestContext_Write(t *testing.T) {
	ctx := &Context{
		Client:  &Client{Codec: DefaultCodec},
		Message: Message([]byte{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}),
	}
	if err := ctx.Write(nil); err != ErrShouldOnlyResponseToRequestMessage {
		t.Errorf("Context.Write() error = %v, wantErr %v", err, ErrShouldOnlyResponseToRequestMessage)
	}
}

func TestContext_WriteWithTimeout(t *testing.T) {
}

func TestContext_Error(t *testing.T) {
}

func TestContext_newRspMessage(t *testing.T) {
}

func TestContext_write(t *testing.T) {
}

func Test_newContext(t *testing.T) {
}
