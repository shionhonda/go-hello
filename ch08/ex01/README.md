# Wall Clock

```
$ TZ=US/Eastern go run clock/clock.go  -port 8010 &
$ TZ=Asia/Tokyo go run clock/clock.go  -port 8020 &
$ TZ=Europe/London go run clock/clock.go -port 8030 &
$ go run main.go NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030
         NewYork: 03:18:30
           Tokyo: 17:18:30
          London: 08:18:30
           Tokyo: 17:18:31
          London: 08:18:31
         NewYork: 03:18:31
          London: 08:18:32
           Tokyo: 17:18:32
         NewYork: 03:18:32
           Tokyo: 17:18:33
         NewYork: 03:18:33
          London: 08:18:33
         NewYork: 03:18:34
           Tokyo: 17:18:34
          London: 08:18:34
           Tokyo: 17:18:35
          London: 08:18:35
         NewYork: 03:18:35
           Tokyo: 17:18:36
         NewYork: 03:18:36
          London: 08:18:36
         NewYork: 03:18:37
           Tokyo: 17:18:37
          London: 08:18:37
           Tokyo: 17:18:38
          London: 08:18:38
         NewYork: 03:18:38
```