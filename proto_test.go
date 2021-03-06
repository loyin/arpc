// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package arpc

import (
	"reflect"
	"testing"
)

func TestHeader_BodyLen(t *testing.T) {
	msg := NewMessage(CmdRequest, "hello", "hello", DefaultCodec)
	if msg.BodyLen() != 10 {
		t.Errorf("Header.BodyLen() = %v, want %v", msg.BodyLen(), 10)
	}
}

func TestHeader_message(t *testing.T) {
	msg := NewMessage(CmdRequest, "hello", "hello", DefaultCodec)
	if msg.BodyLen() != 10 {
		t.Errorf("Header.BodyLen() = %v, want %v", msg.BodyLen(), 10)
	}
	head := Header(msg[:HeadLen])
	msg2, err := head.message()
	if err != nil {
		t.Errorf("Header.message() error = %v", err)
	}
	if len(msg) != len(msg2) {
		t.Errorf("len(Header.message()) = %v, want %v", len(msg2), len(msg))
	}

	head[4], head[5], head[6], head[7] = 0xFF, 0xFF, 0xFF, 0xFF
	_, err = head.message()
	if err == nil {
		t.Errorf("Header.message() error = nil")
	}
}

func TestMessage_Cmd(t *testing.T) {
	msg := NewMessage(CmdRequest, "hello", "hello", DefaultCodec)
	if got := msg.Cmd(); got != CmdRequest {
		t.Errorf("Message.Cmd() = %v, want %v", got, CmdRequest)
	}

	msg = NewMessage(CmdResponse, "hello", "hello", DefaultCodec)
	if got := msg.Cmd(); got != CmdResponse {
		t.Errorf("Message.Cmd() = %v, want %v", got, CmdResponse)
	}

	msg = NewMessage(CmdNotify, "hello", "hello", DefaultCodec)
	if got := msg.Cmd(); got != CmdNotify {
		t.Errorf("Message.Cmd() = %v, want %v", got, CmdNotify)
	}
}

func TestMessage_Async(t *testing.T) {
	msg := NewMessage(CmdRequest, "hello", "hello", DefaultCodec)
	if got := msg.Async(); got != 0 {
		t.Errorf("Message.Async() = %v, want %v", got, 0)
	}
	msg[headerIndexAsync] = 1
	if got := msg.Async(); got != 1 {
		t.Errorf("Message.Async() = %v, want %v", got, 1)
	}
}

func TestMessage_IsAsync(t *testing.T) {
	msg := NewMessage(CmdRequest, "hello", "hello", DefaultCodec)
	if got := msg.IsAsync(); got != false {
		t.Errorf("Message.IsAsync() = %v, want %v", got, false)
	}
	msg[headerIndexAsync] = 1
	if got := msg.IsAsync(); got != true {
		t.Errorf("Message.IsAsync() = %v, want %v", got, true)
	}
}

func TestMessage_IsError(t *testing.T) {
	msg := NewMessage(CmdRequest, "hello", "hello", DefaultCodec)
	if got := msg.IsError(); got != false {
		t.Errorf("Message.IsError() = %v, want %v", got, false)
	}
	msg[headerIndexError] = 1
	if got := msg.IsError(); got != true {
		t.Errorf("Message.IsError() = %v, want %v", got, true)
	}
}

func TestMessage_Error(t *testing.T) {
	msg := NewMessage(CmdRequest, "hello", "hello", DefaultCodec)
	if got := msg.Error(); got != nil {
		t.Errorf("Message.Error() = %v, want %v", got, nil)
	}
}

func TestMessage_MethodLen(t *testing.T) {
	msg := NewMessage(CmdRequest, "hello", "hello", DefaultCodec)
	if got := msg.MethodLen(); got != 5 {
		t.Errorf("Message.MethodLen() = %v, want %v", got, 5)
	}
}

func TestMessage_Method(t *testing.T) {
	msg := NewMessage(CmdRequest, "hello", "hello", DefaultCodec)
	if got := msg.Method(); got != "hello" {
		t.Errorf("Message.Method() = %v, want %v", got, "hello")
	}
}

func TestMessage_BodyLen(t *testing.T) {
	msg := NewMessage(CmdRequest, "hello", "hello", DefaultCodec)
	if got := msg.BodyLen(); got != 10 {
		t.Errorf("Message.BodyLen() = %v, want %v", got, 10)
	}
}

func TestMessage_Seq(t *testing.T) {
	msg := NewMessage(CmdRequest, "hello", "hello", DefaultCodec)
	if got := msg.Seq(); got != 0 {
		t.Errorf("Message.Seq() = %v, want %v", got, 0)
	}
}

func TestMessage_Data(t *testing.T) {
	msg := NewMessage(CmdRequest, "hello", "hello", DefaultCodec)
	if got := msg.Data(); !reflect.DeepEqual(got, []byte("hello")) {
		t.Errorf("Message.Data() = %v, want %v", got, []byte("hello"))
	}
}

func TestNewMessage(t *testing.T) {
	msg := NewMessage(CmdRequest, "hello", "hello", DefaultCodec)
	if msg == nil {
		t.Errorf("NewMessage() = nil")
	}
}
