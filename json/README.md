### Json inputs size

Using decoder from the standard library (Uses at least 3 times the input size in memory )
Input size      File Name      Total Duration           Ns/operation          Total Bytes       InputSize / Bytes Used         Allocs/op
19B                0.json               66720            20793 ns/op             928 B/op       **(*48)**                    7 allocs/op
1.0K            1000.json               27530            39369 ns/op            3488 B/op       **(*3.4)**                   9 allocs/op
2.0K            2000.json               19594            52809 ns/op            8608 B/op       **(*4.3)**                  10 allocs/op
2.9K            3000.json               18009            75105 ns/op            9632 B/op       **(*3.3)**                  10 allocs/op
3.9K            4000.json               10000           167577 ns/op           18848 B/op       **(*4.8)**                  11 allocs/op
4.9K            5000.json                9279           109695 ns/op           20128 B/op       **(*4.1)**                  11 allocs/op
5.9K            6000.json                9045           164464 ns/op           20896 B/op       **(*3.5)**                  11 allocs/op
6.9K            7000.json                6140           171029 ns/op           22944 B/op       **(*3.3)**                  11 allocs/op
7.8K            8000.json                7688           160563 ns/op           39328 B/op       **(*5.0)**                  12 allocs/op
8.8K            9000.json                8335           138152 ns/op           40608 B/op       **(*4.6)**                  12 allocs/op
9.8K            10000.json               8241           154276 ns/op           41376 B/op       **(*4.2)**                  12 allocs/op