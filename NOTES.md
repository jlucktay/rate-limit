# Notepad for random odds and sods

## Inspiration

- The go repo wiki page on [Rate Limiting](https://github.com/golang/go/wiki/RateLimiting)
- Blog posts:
  - [An alternative approach to rate limiting][figma]
    - I like this article a lot, because it dives into the theory and technical design and at the same time does not
    give away any implementation details, which strikes a good balance for my purposes
    - This article references [another][classdojo] which does detail some Node.js implementation, but nothing in Go
- [This discussion][discussion] on the [golang-nuts group][group]

## TODO

- setup a shutdown/cleanup channel, to close down when SIGHUP/KILL/etc signal is received
- configurables:
  - cleanup interval
  - duration
  - number of requests
- identify users/clients based on..? IP? HTTP header(s)?
  - the `jrl` demo binary is currently using an arbitrary `JRL-ID` HTTP header
- throw Redis into the mix
  - running purely in-memory at the moment, while I nail down the algorithm behaviour

[figma]: https://medium.com/figma-design/an-alternative-approach-to-rate-limiting-f8a06cf7c94c
[classdojo]: https://engineering.classdojo.com/blog/2015/02/06/rolling-rate-limiter/
[discussion]: https://groups.google.com/d/topic/golang-nuts/LBzvaXkH3QE
[group]: https://groups.google.com/forum/#!forum/golang-nuts
