# Specification

We'd like you to design and implement a distributed rate limiting library. We expect it to cover the following
functionality:

- Rate limiting should be configurable with a duration and number of requests.
- Multiple instances of a service must be able to use this library with no adverse effects.
- There must be no race conditions or similar concurrency issues.
- It must be performant. The library must be usable in the hot path of a high traffic application.

Choose whatever remote storage you'd like. Third party libraries are allowed, but the rate limiting logic must be your
own and the library itself must be written in Go.

Make sure your code is available on a git repository (private is fine). Pay attention to proper git hygiene: Commits
should represent a unit of change and have a proper commit message that explains the **why**, not the **what**. It
doesn't have to be perfect, but do your best. A commit with the message `"bug fixes"` is never acceptable. If you have
doubts, read [this blog post][blog].

Finally, a piece of advice: Don't be afraid to experiment. A lot of the details are intentionally left vague, so you
can approach the problem however you want. When in doubt, instead of asking a question to us, make your own
assumptions. As long as you explain your assumptions and choices, we are open to pretty much anything.

Good luck!

[blog]: https://chris.beams.io/posts/git-commit/
