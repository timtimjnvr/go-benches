### Results 

Using decoder from the standard library (Uses at least 3 times the input size in memory )
```
Input size      File Name      Total Duration           Ns/operation          Total Bytes   InputSize / Bytes Used         Allocs/op
19B                0.json               66720            20793 ns/op             928 B/op                       48        7 allocs/op
1.0K            1000.json               27530            39369 ns/op            3488 B/op                      3.4        9 allocs/op
2.0K            2000.json               19594            52809 ns/op            8608 B/op                      4.3       10 allocs/op
2.9K            3000.json               18009            75105 ns/op            9632 B/op                      3.3       10 allocs/op
3.9K            4000.json               10000           167577 ns/op           18848 B/op                      4.8       11 allocs/op
4.9K            5000.json                9279           109695 ns/op           20128 B/op                      4.1       11 allocs/op
5.9K            6000.json                9045           164464 ns/op           20896 B/op                      3.5       11 allocs/op
6.9K            7000.json                6140           171029 ns/op           22944 B/op                      3.3       11 allocs/op
7.8K            8000.json                7688           160563 ns/op           39328 B/op                      5.0       12 allocs/op
8.8K            9000.json                8335           138152 ns/op           40608 B/op                      4.6       12 allocs/op
9.8K           10000.json                8241           154276 ns/op           41376 B/op                      4.2       12 allocs/op
```

Using io.ReadAll with Unmarshall (Uses at least 3 times the input size in memory )
```
Input size      File Name      Total Duration           Ns/operation          Total Bytes   InputSize / Bytes Used         Allocs/op
19B                0.json             69044             16699 ns/op             928 B/op                        48       7 allocs/op
1.0K            1000.json             37334             30200 ns/op            3488 B/op                       3.4       9 allocs/op
2.0K            2000.json             25239             45507 ns/op            8608 B/op                       4.3      10 allocs/op
2.9K            3000.json             20848             58730 ns/op            9632 B/op                       3.3      10 allocs/op
3.9K            4000.json             16563             70478 ns/op           18848 B/op                       4.8      11 allocs/op
4.9K            5000.json             14438             84464 ns/op           20128 B/op                       4.1      11 allocs/op
5.9K            6000.json             14416             74996 ns/op           20896 B/op                       3.5      11 allocs/op
6.9K            7000.json             14218             84582 ns/op           22944 B/op                       3.3      11 allocs/op
7.8K            8000.json             12044             97303 ns/op           39328 B/op                       5.0      12 allocs/op
8.8K            9000.json             10000            105972 ns/op           40608 B/op                       4.6      12 allocs/op
9.8K           10000.json             10000            115313 ns/op           41376 B/op                       4.2      12 allocs/op
```

Using Unmarshall and read only the needed size (Uses approximately 2 times the size in memory)
```
Input size      File Name      Total Duration           Ns/operation          Total Bytes   Bytes Used/InputSize         Allocs/op
19B                0.json               88873             11315 ns/op             264 B/op                  13.9       6 allocs/op
1.0K            1000.json               54381             21430 ns/op            2288 B/op                   2.3       7 allocs/op
2.0K            2000.json               39126             30931 ns/op            4336 B/op                   1.6       7 allocs/op
2.9K            3000.json               30237             38391 ns/op            6384 B/op                   2.2       7 allocs/op
3.9K            4000.json               23294             53725 ns/op            8432 B/op                   2.2       7 allocs/op
4.9K            5000.json               20494             97887 ns/op           10992 B/op                   2.2       7 allocs/op
5.9K            6000.json               17862             89295 ns/op           12528 B/op                   2.1       7 allocs/op
6.9K            7000.json               15214             78487 ns/op           16624 B/op                   2.4       7 allocs/op
7.8K            8000.json               13380             87365 ns/op           16624 B/op                   2.1       7 allocs/op
8.8K            9000.json               12396             96871 ns/op           19184 B/op                   2.2       7 allocs/op
9.8K           10000.json               10000            110796 ns/op           20720 B/op                   2.1       7 allocs/op
```

Read only the needed size and give a reader to the decoder with the given data (Same as giving the initial reader to the decoder)
```
Input size      File Name      Total Duration           Ns/operation          Total Bytes   Bytes Used/InputSize         Allocs/op
19B                0.json              111681             10848 ns/op             928 B/op                   48        7 allocs/op
1.0K            1000.json               50431             23478 ns/op            3488 B/op                   3.5       9 allocs/op
2.0K            2000.json               34380             34883 ns/op            8608 B/op                   4.3      10 allocs/op
2.9K            3000.json               27733             42722 ns/op            9632 B/op                   3.3      10 allocs/op
3.9K            4000.json               20848             57253 ns/op           18848 B/op                   4.8      11 allocs/op
4.9K            5000.json               18577             64198 ns/op           20128 B/op                   4.1      11 allocs/op
5.9K            6000.json               16399             73831 ns/op           20896 B/op                   3.5      11 allocs/op
6.9K            7000.json               14498             81732 ns/op           22944 B/op                   3.3      11 allocs/op
7.8K            8000.json               12253             97372 ns/op           39328 B/op                   5.0      12 allocs/op
8.8K            9000.json               10513            110578 ns/op           40608 B/op                   4.6      12 allocs/op
9.8K           10000.json                9726            126214 ns/op           41376 B/op                   4.2      12 allocs/op
```