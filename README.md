godatastruct
============

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/8f99b21a6dcc4f4ab9adc4fd2a836210)](https://app.codacy.com/manual/egoroff/godatastruct?utm_source=github.com&utm_medium=referral&utm_content=aegoroff/godatastruct&utm_campaign=Badge_Grade_Dashboard)
[![CI](https://github.com/aegoroff/godatastruct/actions/workflows/ci.yml/badge.svg)](https://github.com/aegoroff/godatastruct/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/aegoroff/godatastruct/branch/master/graph/badge.svg?token=zJDEi5IIc3)](https://codecov.io/gh/aegoroff/godatastruct)
[![](https://tokei.rs/b1/github/aegoroff/godatastruct)](https://github.com/aegoroff/godatastruct)

Implementations of various fundamental data structures as tree, stacks, queues etc. in Go

## Packages

| Package | Description |
|:--|:--|
| rbtree | Red-black binary tree implementation that supports ordered statistic |
| rbtree/special | Contains specialized Red-black search binary tree implementations |
| countingsort | Counting sort is an algorithm for sorting a collection of objects according to keys that are small integers; that is, it is an integer sorting algorithm. |
| collections | Various containers. Now only int, int64 and string hashsets implemented |

## Benchmarks

Here are some benchmark that compares RbTree, Google BTree and HashSet in various OS's:

```commandline
goos: linux
goarch: amd64
pkg: github.com/aegoroff/godatastruct/rbtree
cpu: Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz
Benchmark_RbTree_Insert-2            	     146	   8065465 ns/op	 1437960 B/op	   39744 allocs/op
Benchmark_RbTree_ReplaceOrInsert-2   	      94	  12322666 ns/op	 1437958 B/op	   39744 allocs/op
Benchmark_BTree_ReplaceOrInsert-2    	     138	   8593302 ns/op	  946512 B/op	   19955 allocs/op
Benchmark_RbTree_Search-2            	    8384	    155340 ns/op	    2931 B/op	      47 allocs/op
Benchmark_BTree_Search-2             	    7102	    160549 ns/op	    3012 B/op	      42 allocs/op
Benchmark_StringHashSet_Search-2     	  194384	      5856 ns/op	     122 B/op	       1 allocs/op
Benchmark_RbTree_Ascend-2            	18868369	        59.53 ns/op	      16 B/op	       1 allocs/op
Benchmark_BTree_Ascend-2             	     990	   1035405 ns/op	    5676 B/op	     101 allocs/op
Benchmark_RbTree_Descend-2           	17982375	        62.05 ns/op	      16 B/op	       1 allocs/op
Benchmark_BTree_Descend-2            	    1006	   1035565 ns/op	    5582 B/op	      99 allocs/op
```
```commandline
goos: darwin
goarch: amd64
pkg: github.com/aegoroff/godatastruct/rbtree
cpu: Intel(R) Xeon(R) CPU E5-1650 v2 @ 3.50GHz
Benchmark_RbTree_Insert-3            	     205	   5889469 ns/op	 1437958 B/op	   39744 allocs/op
Benchmark_RbTree_ReplaceOrInsert-3   	     127	   9342955 ns/op	 1437961 B/op	   39744 allocs/op
Benchmark_BTree_ReplaceOrInsert-3    	     132	   8444145 ns/op	  949028 B/op	   19956 allocs/op
Benchmark_RbTree_Search-3            	    8524	    130391 ns/op	    2883 B/op	      46 allocs/op
Benchmark_BTree_Search-3             	    8367	    134109 ns/op	    2566 B/op	      35 allocs/op
Benchmark_StringHashSet_Search-3     	  208464	      4917 ns/op	     113 B/op	       1 allocs/op
Benchmark_RbTree_Ascend-3            	21137170	        56.48 ns/op	      16 B/op	       1 allocs/op
Benchmark_BTree_Ascend-3             	    1352	    773357 ns/op	    4129 B/op	      74 allocs/op
Benchmark_RbTree_Descend-3           	20798305	        51.98 ns/op	      16 B/op	       1 allocs/op
Benchmark_BTree_Descend-3            	    1664	    676728 ns/op	    3365 B/op	      60 allocs/op
```
```commandline
goos: windows
goarch: amd64
pkg: github.com/aegoroff/godatastruct/rbtree
cpu: Intel(R) Xeon(R) CPU E5-2673 v3 @ 2.40GHz
Benchmark_RbTree_Insert-2            	     146	   7957292 ns/op	 1437957 B/op	   39744 allocs/op
Benchmark_RbTree_ReplaceOrInsert-2   	      99	  12133983 ns/op	 1437957 B/op	   39744 allocs/op
Benchmark_BTree_ReplaceOrInsert-2    	     136	   8682621 ns/op	  946224 B/op	   19955 allocs/op
Benchmark_RbTree_Search-2            	    5938	    204376 ns/op	    4140 B/op	      67 allocs/op
Benchmark_BTree_Search-2             	    5336	    197659 ns/op	    3997 B/op	      56 allocs/op
Benchmark_StringHashSet_Search-2     	  186262	      6147 ns/op	     127 B/op	       1 allocs/op
Benchmark_RbTree_Ascend-2            	15858182	        68.24 ns/op	      16 B/op	       1 allocs/op
Benchmark_BTree_Ascend-2             	    1263	    955326 ns/op	    4420 B/op	      79 allocs/op
Benchmark_RbTree_Descend-2           	15209337	        70.29 ns/op	      16 B/op	       1 allocs/op
Benchmark_BTree_Descend-2            	    1226	    967683 ns/op	    4553 B/op	      82 allocs/op
```