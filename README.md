# gobaseline

This is a simple _hello world_ HTTP Go server to find a baseline number for performance comparisons.

## Run Test

You'll need `Go` and `Node` installed.

    go run main.go

Apache Bench didn't work so I used [loadtest](https://www.npmjs.com/package/loadtest).

    loadtest -n 100000 -c 8 -k http://127.0.0.1:8080/

## Performance Results

The test was run on a MacBook Pro (Retina, 13-inch, Mid 2014), 3 GHz Intel Core i7 with 16 GB 1600 MHz DDR3 memory.

    Concurrency level:   8
    Agent:               keepalive

    Completed requests:  100000
    Total errors:        0
    Total time:          19.489376383 s
    Requests per second: 5131
    Total time:          19.489376383 s

    Percentage of the requests served within a certain time
      50%      1 ms
      90%      1 ms
      95%      1 ms
      99%      2 ms
     100%      24 ms (longest request)

## Problems Found

It's slow. Need to work out what's going wrong.
