# millionth
Millionth in-memory database
# Usage
```Go
m := millionth.New()
newId :=	m.Create([]byte{200})
```

# Bench

- BenchmarkCreate-4           	10000000	       192 ns/op
- BenchmarkCreateParallel-4   	10000000	       154 ns/op
- BenchmarkWrite-4            	 5000000	       289 ns/op
- BenchmarkWriteParallel-4    	10000000	       131 ns/op
- BenchmarkAdd-4              	10000000	       183 ns/op
- BenchmarkAddParallel-4      	10000000	       171 ns/op
- BenchmarkDelete-4           	20000000	       112 ns/op
- BenchmarkDeleteParallel-4   	20000000	        72.9 ns/op
- BenchmarkRead-4             	20000000	        70.1 ns/op
- BenchmarkReadParallel-4     	20000000	       100 ns/op
- BenchmarkMerge-4            	20000000	       115 ns/op
- BenchmarkMergeParallel-4    	20000000	        86.5 ns/op
