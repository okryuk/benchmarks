# This repo contains various benchmarks results


JSON-PATCHING-RFC6902 benchmark was created to measure the efficacy of json changes within the specification for RFC6902. Main criterias to assess were:
- minimum or no use of the json stdlib Marshal/Unmarshalling to increase the throughput of the system
- speed and minimum memory allocation in the priority
- add, replace and remove operations were of a focus

JSONPATCH represents https://github.com/evanphx/json-patch repo. Standard functionality used.

SJSON represents https://github.com/tidwall/sjson repo. Since it only allows to change values, following benchmark variations were used:
- raw mode with the json patching operations hardcoded to asses the native speed
- UniJSON function that imitated the json operation instructions layer with the version, that used RawMessage with stdlib json Unmarshalling, and another version build around complimentary https://github.com/tidwall/gjson repo.

# Performance results

```
Benchmark_Add_JSONPATCH-10                                 93003             12356 ns/op            4492 B/op         95 allocs/op
Benchmark_Add_SJSON_rawOperation-10                       988942              1063 ns/op            2528 B/op          9 allocs/op
Benchmark_Add_UniJSONPATCH_EQ_withRawMessage-10           563065              2024 ns/op            3080 B/op         20 allocs/op
Benchmark_Add_UniJSONPATCH_EQ_Optimized-10                770680              1494 ns/op            2641 B/op         12 allocs/op
Benchmark_Replace_JSONPATCH-10                             98708             11940 ns/op            4340 B/op         93 allocs/op
Benchmark_Replace_SJSON_rawOperation-10                  1230950             976.3 ns/op            2464 B/op          9 allocs/op
Benchmark_Replace_UniJSONPATCH_EQ_withRawMessage-10       579582              1974 ns/op            3016 B/op         20 allocs/op
Benchmark_Replace_UniJSONPATCH_EQ_Optim-10                816120              1393 ns/op            2577 B/op         12 allocs/op
Benchmark_Remove_JSONPATCH-10                             107253             10975 ns/op            3988 B/op         83 allocs/op
Benchmark_Remove_SJSON_rawOperation-10                   1433566             819.7 ns/op            2384 B/op          7 allocs/op
Benchmark_Remove_UniJSONPATCH_EQ_withRawMessage-10        673614              1707 ns/op            2888 B/op         18 allocs/op
Benchmark_Remove_UniJSONPATCH_EQ_Optim-10                 914898              1266 ns/op            2529 B/op         11 allocs/op

```
