package millionth

// Millionth
// Test
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	//"fmt"
	//"sync/atomic"
	"testing"
	//"unsafe"
)

func TestCreate1000000sum(t *testing.T) {
	m := New()
	for i := 0; i < 1000000; i++ {
		m.Create([]byte{200})
	}
	c := 0
	for i := len(m.base) - 1; i >= 0; i-- {
		c += len(m.base[i].data)
	}
	if c != 1000000 {
		t.Error("Ожидаю размер базы данных == 1000000, а получил: ", c)
	}
}

func TestRead1000000(t *testing.T) {
	m := New()
	addrs := make(map[uint64]byte)
	for i := 0; i < 1000000; i++ {
		id := m.Create([]byte{byte(i + 5)})
		//fmt.Print("\n", id, "\n")
		addrs[id] = byte(i + 5)
	}
	for key, value := range addrs {
		v := m.Read(key)
		if v[0] != value {
			t.Error("Ключ `", key, " получено:", v[0], " вместо:", value)
		}
	}
}

/*
func TestCreate64a(t *testing.T) {
	m := New()
	for i := 0; i < 65256; i++ {
		m.Create64a([]byte{200, 201, 202})
	}
}
func TestCreate64b(t *testing.T) {
	m := New()
	for i := 0; i < 65256; i++ {
		m.Create64b([]byte{200, 201, 202})
	}
}

func TestCreate32(t *testing.T) {
	m := New()
	for i := 0; i < 65256; i++ {
		m.Create32([]byte{200, 201, 202})
	}
}
*/
func TestDummy(t *testing.T) {
}
