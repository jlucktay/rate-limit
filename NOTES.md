# Notepad for random odds and sods

## Inspiration

- The go repo wiki page on [Rate Limiting](https://github.com/golang/go/wiki/RateLimiting)
- Blog posts:
  - [How to Rate Limit HTTP Requests](https://www.alexedwards.net/blog/how-to-rate-limit-http-requests)
  - [An alternative approach to rate limiting][figma]
    - I like this article a lot, because it dives into the theory and technical design and at the same time does not
    give away any implementation details, which strikes a good balance for my purposes
    - This article references [another][classdojo] which does detail some Node.js implementation, but nothing in Go
  - [Scaling your API with rate limiters](https://stripe.com/au/blog/rate-limiters)
- [This discussion][discussion] on the [golang-nuts group][group]

## Known issues

1. Due to the nature of the implementation around sliding windows and how they are at some points converted from
nanoseconds to full seconds, using a sliding window factor of 60 (the current hard-coded value) and setting a duration
on the Limiter below one minute will result in sub-par behaviour.

    One potential way to work around this would be to drop the conversions into full seconds and keep everything in
nanoseconds.

    Maybe save this for a v2 rewrite? :)

1. As described in the [linked blog posts][figma] that served as an inspiration, under the **Practical considerations**
section, there is some small amount of leeway in terms of the number of requests allowed through.

    Within my implementation, this stems from additional visits no longer being counted *after* the limit is hit, such
that the visit count total for a restricted visitor ends up one higher than the configured limit.

    If subsequent visits are attempted squarely within each consecutive time window, then the visitor may be allowed
back through a small time later than they might otherwise expect.

    Again, this might be mitigated by keeping all time measurements at the nanosecond level, and is only noticeably
restrictive when a very low request limit is configured.

## TODO

- setup a shutdown/cleanup channel, to close down when SIGHUP/KILL/etc signal is received
- configurables:
  - cleanup interval
    - currently derived from the duration and sliding window factor
  - ~~duration~~ an argument of `New()`
  - ~~number of requests~~ an argument of `New()`
- identify users/clients based on..? IP? HTTP header(s)?
  - the `jrl` demo binary is currently using an arbitrary `JRL-ID` HTTP header
- throw Redis into the mix
  - running purely in-memory at the moment, while I nail down the algorithm behaviour
- ~~set up pruning at the leaf level also~~ visitor structs have their own mutex locks for this
- ~~start a pruning goroutine in the constructor~~ `New()` kicks this off
- consolidate visit counts from time windows prior to the current one, in order to minimise memory footprint

[figma]: https://medium.com/figma-design/an-alternative-approach-to-rate-limiting-f8a06cf7c94c
[classdojo]: https://engineering.classdojo.com/blog/2015/02/06/rolling-rate-limiter/
[discussion]: https://groups.google.com/d/topic/golang-nuts/LBzvaXkH3QE
[group]: https://groups.google.com/forum/#!forum/golang-nuts
