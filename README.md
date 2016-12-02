# gobaseline

This is a simple _hello world_ HTTP Go server to find a baseline number for performance comparisons.

## Run Test

You'll need `Go` and `Node` installed.

    go run main.go

Now run apache bench;

    ab -n 1000000 -c 8 -k http://127.0.0.1:8080/

## Performance Results

The test was run on a MacBook Pro (Retina, 13-inch, Mid 2014), 3 GHz Intel Core i7 with 16 GB 1600 MHz DDR3 memory.

    Concurrency Level:      8
    Time taken for tests:   20.064 seconds
    Complete requests:      1000000
    Failed requests:        0
    Keep-Alive requests:    998005
    Total transferred:      158990025 bytes
    HTML transferred:       20000000 bytes
    Requests per second:    49839.65 [#/sec] (mean)
    Time per request:       0.161 [ms] (mean)
    Time per request:       0.020 [ms] (mean, across all concurrent requests)
    Transfer rate:          7738.29 [Kbytes/sec] received

## Problems Found

None. But I thought it would be faster.
