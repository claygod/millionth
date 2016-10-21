package millionth

// Millionth
// Test
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
	//"sync/atomic"
	"testing"
	//"unsafe"
	//"runtime"
	//"sync"
	//"math/rand"
	//"time"
)

const ADD_COUNT = 700000

// Test add new, ns/op

func BenchmarkWrite(b *testing.B) {
	limit := uint64(65000)
	b.StopTimer()
	m := New()
	var arr []uint64
	for i := uint64(0); i <= limit; i++ {
		mas := []byte{byte(i + 5)}
		key := m.Create(mas)
		arr = append(arr, key)
	}
	counter := uint64(1)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		//rez := m.Write(arr[counter], []byte{199})
		m.Write(arr[counter], []byte{199})
		//if rez != true {
		//	fmt.Print("\nПолучили ошибку на запрос Write ключ ", counter)
		//}
		if counter >= limit-100 {
			counter = 1
		}
		counter++
	}
}

func BenchmarkWriteParallel(b *testing.B) {
	limit := uint64(65000)
	b.StopTimer()
	m := New()
	var arr []uint64
	for i := uint64(0); i <= limit; i++ {
		mas := []byte{byte(i + 5)}
		key := m.Create(mas)
		arr = append(arr, key)
	}
	counter := uint64(1)
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//rez := m.Add(arr[counter], []byte{199})
			m.Write(arr[counter], []byte{199})
			//if rez != true {
			//	fmt.Print("\nПолучили ошибку на запрос Write ключ ", counter)
			//}
			if counter >= limit-100 {
				counter = 1
			}
			counter++
		}
	})
}

func BenchmarkAdd(b *testing.B) {
	limit := uint64(65000)
	b.StopTimer()
	m := New()
	var arr []uint64
	for i := uint64(0); i <= limit; i++ {
		mas := []byte{byte(i + 5)}
		key := m.Create(mas)
		arr = append(arr, key)
	}
	counter := uint64(1)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		//rez := m.Add(arr[counter], []byte{199})
		m.Add(arr[counter], []byte{199})
		//if rez != true {
		//	fmt.Print("\nПолучили ошибку на запрос Add ключ ", counter)
		//}
		if counter >= limit-100 {
			counter = 1
		}
		counter++
	}
}

func BenchmarkAddParallel(b *testing.B) {
	limit := uint64(65000)
	b.StopTimer()
	m := New()
	var arr []uint64
	for i := uint64(0); i <= limit; i++ {
		mas := []byte{byte(i + 5)}
		key := m.Create(mas)
		arr = append(arr, key)
	}
	counter := uint64(1)
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//rez := m.Add(arr[counter], []byte{199})
			m.Add(arr[counter], []byte{199})
			//if rez != true {
			//	fmt.Print("\nПолучили ошибку на запрос Add ключ ", counter)
			//}
			if counter >= limit-100 {
				counter = 1
			}
			counter++
		}
	})
}

func BenchmarkDelete(b *testing.B) {
	limit := 65000
	b.StopTimer()
	m := New()
	var arr []uint64
	for i := 0; i <= limit; i++ {
		mas := []byte{byte(i + 5)}
		key := m.Create(mas)
		arr = append(arr, key)
	}
	counter := 0
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		//rez := m.Delete(arr[i])
		m.Delete(arr[counter])
		//if rez == false {
		//	fmt.Print("\nПолучили ошибку на запрос DEL ключ ", arr[i])
		//	counter++
		//}
		if counter >= limit-100 {
			counter = 1
		}
		counter++
	}
}

func BenchmarkDeleteParallel(b *testing.B) {
	limit := 65000
	b.StopTimer()
	m := New()
	var arr []uint64
	for i := 0; i <= limit; i++ {
		mas := []byte{byte(i + 5)}
		key := m.Create(mas)
		arr = append(arr, key)
	}
	counter := 0
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Delete(arr[counter])
			if counter >= limit-100 {
				counter = 1
			}
			counter++
		}
	})
}

func BenchmarkRead(b *testing.B) {
	limit := 65000
	b.StopTimer()
	m := New()
	var arr []uint64
	for i := 0; i <= limit; i++ {
		mas := []byte{byte(i + 5)}
		key := m.Create(mas)
		arr = append(arr, key)
	}
	counter := 0
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		//key := uint64(uint16(i))
		m.Read(arr[counter])
		//rez := m.Read(arr[key])
		//if rez == nil {
		//	fmt.Print("\nПолучили ошибку на запрос READ ключ ", arr[i])
		//counter++
		//}
		if counter >= limit-100 {
			counter = 1
		}
		counter++
	}
	//fmt.Print("\nВсего ошибок чтения - ", counter)
}

func BenchmarkReadParallel(b *testing.B) {
	limit := 65000
	b.StopTimer()
	m := New()
	var arr []uint64
	for i := 0; i <= limit; i++ {
		mas := []byte{byte(i + 5)}
		key := m.Create(mas)
		arr = append(arr, key)
	}
	counter := 0
	//i := 0
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			//key := uint64(uint16(i))
			//rez := m.Read(arr[key])
			m.Read(arr[counter])
			//if rez == nil {
			//	fmt.Print("\nПолучили ошибку на запрос READ ключ ", key)
			//counter++
			//}
			//if i >= limit {
			//	i = 0
			//}

			if counter >= limit-100 {
				counter = 1
			}
			counter++
		}
	})
}

func BenchmarkCreate(b *testing.B) {
	b.StopTimer()
	m := New()
	b.StartTimer()
	m.Create([]byte{200, 201, 202})
	for i := 0; i < b.N; i++ {
		m.Create([]byte{200})
	}
}

func BenchmarkCreateParallel(b *testing.B) {
	b.StopTimer()
	m := New()
	m.Create([]byte{200, 111, 112})
	//b.SetParallelism(1000)
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Create([]byte{200})
		}
	})
}

func BenchmarkMerge(b *testing.B) {
	b.StopTimer()
	m := New()
	b.StartTimer()
	m.Merge([]byte{200, 201, 202})
	for i := 0; i < b.N; i++ {
		m.Merge([]byte{200})
	}
}

func BenchmarkMergeParallel(b *testing.B) {
	b.StopTimer()
	m := New()
	m.Merge([]byte{200, 111, 112})
	//b.SetParallelism(1000)
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Merge([]byte{200})
		}
	})
}

/*
func BenchmarkAddMap(b *testing.B) {
	b.StopTimer()
	//type Millionth struct {}
	d := make(map[int]int)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d[i] = i
	}
}

func BenchmarkAddMapParallel(b *testing.B) {
	b.StopTimer()
	type db struct {
		mu   sync.Mutex
		data map[int]int
	}
	d := db{}
	d.data = make(map[int]int)
	b.StartTimer()
	i := 0
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			d.mu.Lock()
			d.data[i] = i
			d.mu.Unlock()
			i++
		}
	})
}

func BenchmarkAddSlice(b *testing.B) {
	b.StopTimer()
	type db struct {
		mu   sync.Mutex
		data []int
	}
	d := db{}
	d.data = make([]int, 10000)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d.data = append(d.data, i)
	}
}

func BenchmarkAddSliceParallelNoMutex(b *testing.B) {
	limit := 256
	b.StopTimer()
	type section struct {
		lock int64
		db   []int
	}
	data := make([]section, limit)
	for i := 0; i < limit; i++ {
		data[i] = section{} //0, make([]int, 100), sync.Mutex
		data[i].lock = 0
		data[i].db = make([]int, 1000)
	}
	b.StartTimer()
	i := 0
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := int(byte(i))

			data[key].db = append(data[key].db, i)
			i++
			if i == limit {
				i = 0
			}
		}
	})
}

func BenchmarkAddSliceParallelMutex(b *testing.B) {
	b.StopTimer()
	type db struct {
		mu   sync.Mutex
		data []int
	}
	d := db{}
	d.data = make([]int, 10000)
	b.StartTimer()
	i := 0
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			d.mu.Lock()
			d.data = append(d.data, i)
			d.mu.Unlock()
			i++
		}
	})
}




func BenchmarkAdd64a(b *testing.B) {
	//runtime.GOMAXPROCS(2)
	//runtime.GOMAXPROCS(runtime.NumCPU())
	//runtime.GC()
	//b.SetParallelism(10)
	b.StopTimer()
	m := New()
	m.Create64a([]byte{200, 222})
	b.StartTimer()
	for i := 0; i < b.N; i++ { // b.N
		m.Create64a([]byte{200})
	}
}

func BenchmarkAdd64Parallela(b *testing.B) {
	b.StopTimer()
	m := New()
	m.Create64a([]byte{200, 111, 112})
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Create64a([]byte{200})
		}
	})
}

func BenchmarkAdd64b(b *testing.B) {
	//runtime.GOMAXPROCS(2)
	//runtime.GOMAXPROCS(runtime.NumCPU())
	//runtime.GC()
	//b.SetParallelism(10)
	b.StopTimer()
	m := New()
	m.Create64b([]byte{200, 222})
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		m.Create64b([]byte{200})
	}
}

func BenchmarkAdd64Parallelb(b *testing.B) {
	b.StopTimer()
	m := New()
	m.Create64b([]byte{200, 111, 112})
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Create64b([]byte{200})
		}
	})
}

func BenchmarkAdd32(b *testing.B) {
	m := New()
	m.Create32([]byte{200, 201, 202})
	for i := 0; i < b.N; i++ { //b.N
		m.Create32([]byte{200})
	}
}
*/
func BenchmarkDummy(b *testing.B) {
	fmt.Print("END")
}
