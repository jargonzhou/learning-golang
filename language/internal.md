# Go Internal

# Go Scheduler
* [Scheduling In Go : Part II - Go Scheduler](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html)

![](https://www.ardanlabs.com/images/goinggo/94_figure2.png)
Here:
- GRQ: Global Run Queue
- LRQ: Local Run Queue
- M: Machine, OS Thread
- P: Logical Processor for every virtual core that is identified on the host machine
- G: Goroutine
	- [Coroutine - wikipedia](https://en.wikipedia.org/wiki/Coroutine)

## See Also
* [Go's work-stealing scheduler](https://rakyll.org/scheduler/) - 2017-07-16
* [Work stealing](https://en.wikipedia.org/wiki/Work_stealing) - wikipedia
  * **Child stealing** vs. **continuation stealing**: see [A Primer on Scheduling Fork–Join Parallelism with Work Stealin](https://www.open-std.org/jtc1/sc22/wg21/docs/papers/2014/n3872.pdf) - 2014-01-15

# See Also
* [golang/proposal](https://github.com/golang/proposal): Proposing Changes to Go. This document describes the process for proposing, documenting, and implementing changes to the Go project.
* [Language Mechanics On Stacks And Pointers](https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-stacks-and-pointers.html)
* [Garbage Collection In Go : Part I - Semantics](https://www.ardanlabs.com/blog/2018/12/garbage-collection-in-go-part1-semantics.html)
* [A Guide to the Go Garbage Collector](https://tip.golang.org/doc/gc-guide)