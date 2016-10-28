package millionth

// Millionth
// API
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	//"errors"
	"fmt"
	//"math/rand"
	//"runtime"
	"sync"
	"sync/atomic"
	//"time"
)

const POOL_SIZE int = 6000
const SECTION_SIZE uint64 = 6100
const SECTION_LIMIT uint64 = 6000
const TRIAL_LIMIT int = 20000000

// New - create a new Millionth-struct
func New() *Millionth {
	m := &Millionth{}
	m.shift = Cursor{uint64(POOL_SIZE), 0}
	m.base = append(m.base, Section{}) // dummy первая секция
	for i := 1; i <= POOL_SIZE; i++ {
		sct := Section{}
		sct.data = make([][]byte, 0, SECTION_SIZE)
		sct.length = 0
		sct.lock = 0
		m.base = append(m.base, sct)
		m.cursors = append(m.cursors, Cursor{cursor: uint64(i), lock: 0})
	}
	return m
}

// const CONF_FILE string = "config.ini"

// Millionth structure
type Millionth struct {
	mu      sync.Mutex
	shift   Cursor
	base    []Section
	cursors []Cursor
	swtch   uint64
}
type Section struct {
	data   [][]byte
	lock   uint64
	length uint64
}

type Cursor struct {
	cursor uint64
	lock   uint64
}

// Merge - создать новую запись присоединением и получить её ID
// Эта функция полученные данные ПРИСОЕДИНЯЕТ в новую запись, т.е.
// клиентский код МОЖЕТ её в дальнейшем изменить коственно.
func (m *Millionth) Merge(record []byte) uint64 {
	numSection, curSwitch := m.getNumSection()
	if !m.lock(&m.base[numSection].lock) {
		return 0
	}
	m.base[numSection].data = append(m.base[numSection].data, record)
	m.base[numSection].length++
	n := m.base[numSection].length
	m.base[numSection].lock = 0
	if n == SECTION_LIMIT {
		m.createNewSection(curSwitch)
	}
	return (n - 1) + numSection<<32
}

// Create - создать новую запись копированием и получить её ID
// Эта функция полученные данные КОПИРУЕТ в новую запись, т.е.
// клиентский код НЕ может её в дальнейшем изменить коственно.
func (m *Millionth) Create(record []byte) uint64 {
	numSection, curSwitch := m.getNumSection()
	if !m.lock(&m.base[numSection].lock) {
		return 0
	}
	m.base[numSection].data = append(m.base[numSection].data, []byte{})
	m.base[numSection].length++
	n := m.base[numSection].length
	m.base[numSection].data[n-1] = append(m.base[numSection].data[n-1], record...)
	m.base[numSection].lock = 0
	if n == SECTION_LIMIT {
		m.createNewSection(curSwitch)
	}
	return (n - 1) + numSection<<32
}

// Read - прочитать содержимое записи
func (m *Millionth) Read(id uint64) []byte {
	numSection := id >> 32
	numRecord := (id << 32) >> 32 // uint64(uint32(id))
	if m.shift.cursor < numSection ||
		m.base[numSection].length < numRecord { // || ns.data == nil || ns.data[numRecord] == nil
		return nil
	}
	if !m.lock(&m.base[numSection].lock) {
		return nil
	}
	out := m.base[numSection].data[numRecord]
	m.base[numSection].lock = 0
	return out
}

// Del - удаляет содержимое записи, номер записи не удаляется
func (m *Millionth) Delete(id uint64) bool {
	numSection := id >> 32
	numRecord := (id << 32) >> 32
	if m.shift.cursor < numSection ||
		m.base[numSection].length < numRecord {
		return false
	}
	if !m.lock(&m.base[numSection].lock) {
		return false
	}
	if m.base[numSection].data[numRecord] != nil {
		m.base[numSection].data[numRecord] = nil
		m.base[numSection].lock = 0
		return true
	}
	m.base[numSection].lock = 0
	return false
}

// Add - добавляет содержимое к записи
func (m *Millionth) Add(id uint64, record []byte) bool {
	numSection := id >> 32
	numRecord := (id << 32) >> 32
	if m.shift.cursor < numSection ||
		m.base[numSection].length < numRecord {
		return false
	}
	if !m.lock(&m.base[numSection].lock) {
		return false
	}
	m.base[numSection].data[numRecord] = append(m.base[numSection].data[numRecord], record...)
	m.base[numSection].lock = 0
	return true
}

// Write - заменяет содержимое записи на новое
func (m *Millionth) Write(id uint64, record []byte) bool {
	numSection := id >> 32
	numRecord := (id << 32) >> 32
	if m.shift.cursor < numSection ||
		m.base[numSection].length < numRecord {
		return false
	}
	if !m.lock(&m.base[numSection].lock) {
		return false
	}
	m.base[numSection].data[numRecord] = []byte{}
	m.base[numSection].data[numRecord] = append(m.base[numSection].data[numRecord], record...)
	m.base[numSection].lock = 0
	return true
}

func (m *Millionth) createNewSection(curSwitch uint64) {
	m.mu.Lock()
	sct := Section{}
	sct.data = make([][]byte, 0, SECTION_SIZE)
	sct.length = 0
	sct.lock = 0
	m.base = append(m.base, sct)
	m.shift.cursor++
	m.cursors[curSwitch].cursor = m.shift.cursor
	m.mu.Unlock()
}

func (m *Millionth) getNumSection() (uint64, uint64) {
	curSwitch := m.swtch //atomic.LoadUuint64(&m.swtch)
	var newSwitch uint64 = curSwitch + 1
	if newSwitch >= uint64(POOL_SIZE) {
		newSwitch = 0
	}
	m.swtch = newSwitch                           // atomic.CompareAndSwapUint64(&m.swtch, curSwitch, newSwitch)
	return m.cursors[curSwitch].cursor, curSwitch // numSection := atomic.LoadUint64(&m.cursors[curSwitch].cursor)
}

func (m *Millionth) lock(lock *uint64) bool {
	for i := TRIAL_LIMIT; i > 0; i-- {
		if *lock == 0 && atomic.CompareAndSwapUint64(lock, 0, 1) {
			break
		}
		if i == 5 {
			return false
		}
	}
	return true
}

func (m *Millionth) unlock(lock *uint64) bool {
	for i := TRIAL_LIMIT; i > 0; i-- {
		if *lock == 0 && atomic.CompareAndSwapUint64(lock, 1, 0) {
			break
		}
		if i == 5 {
			return false
		}
	}
	return true
}

func (m *Millionth) dummy() {
	fmt.Print("\n ````` ")
}
