package main

import "testing"

type EvalItf interface {
	Eval()
}
type EvalPtrItf interface {
	EvalPtr()
}

type Array1K struct {
	arr [1024]int
}
type Array10K struct {
	arr [10 * 1024]int
}
type Array16K struct {
	arr [16 * 1024]int
}
type Array17K struct {
	arr [17 * 1024]int
}
type Array20K struct {
	arr [20 * 1024]int
}

func (arr Array1K) Eval() {
	_ = arr.arr[0]
}
func (arr *Array1K) EvalPtr() {
	_ = arr.arr[0]
}
func (arr Array10K) Eval() {
	_ = arr.arr[0]
}
func (arr *Array10K) EvalPtr() {
	_ = arr.arr[0]
}
func (arr Array16K) Eval() {
	_ = arr.arr[0]
}
func (arr *Array16K) EvalPtr() {
	_ = arr.arr[0]
}
func (arr Array17K) Eval() {
	_ = arr.arr[0]
}
func (arr *Array17K) EvalPtr() {
	_ = arr.arr[0]
}
func (arr Array20K) Eval() {
	_ = arr.arr[0]
}
func (arr *Array20K) EvalPtr() {
	_ = arr.arr[0]
}

// cpu: Intel(R) Core(TM) i3 CPU       M 380  @ 2.53GHz
// BenchmarkValueReceiver/1K-4         	335045390	         3.773 ns/op	       0 B/op	       0 allocs/op
// BenchmarkValueReceiver/10K-4        	335033006	         3.749 ns/op	       0 B/op	       0 allocs/op
// BenchmarkValueReceiver/16K-4        	293596389	         3.768 ns/op	       0 B/op	       0 allocs/op
// BenchmarkValueReceiver/17K-4        	   12381	     98653 ns/op	  139264 B/op	       1 allocs/op
// BenchmarkValueReceiver/20K-4        	    9061	    117815 ns/op	  163840 B/op	       1 allocs/op
// 总结：应该与cpu缓存大小有关
func BenchmarkValueReceiver(b *testing.B) {
	ss := []struct {
		name string
		itf  EvalItf
	}{
		{"1K", Array1K{}},
		{"10K", Array10K{}},
		{"16K", Array16K{}},
		{"17K", Array17K{}},
		{"20K", Array20K{}},
	}
	for _, s := range ss {
		b.Run(s.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s.itf.Eval()
			}
		})
	}
}

// cpu: Intel(R) Core(TM) i3 CPU       M 380  @ 2.53GHz
// BenchmarkPointerReceiver/1K-4         	360013178	         3.352 ns/op	       0 B/op	       0 allocs/op
// BenchmarkPointerReceiver/10K-4        	321220201	         3.354 ns/op	       0 B/op	       0 allocs/op
// BenchmarkPointerReceiver/16K-4        	376611500	         3.358 ns/op	       0 B/op	       0 allocs/op
// BenchmarkPointerReceiver/17K-4        	376555268	         3.303 ns/op	       0 B/op	       0 allocs/op
// BenchmarkPointerReceiver/20K-4        	320996722	         3.358 ns/op	       0 B/op	       0 allocs/op
func BenchmarkPointerReceiver(b *testing.B) {
	ss := []struct {
		name string
		itf  EvalPtrItf
	}{
		{"1K", &Array1K{}},
		{"10K", &Array10K{}},
		{"16K", &Array16K{}},
		{"17K", &Array17K{}},
		{"20K", &Array20K{}},
	}
	for _, s := range ss {
		b.Run(s.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s.itf.EvalPtr()
			}
		})
	}

}
