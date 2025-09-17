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

# See Also
* [Language Mechanics On Stacks And Pointers](https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-stacks-and-pointers.html)
* [Garbage Collection In Go : Part I - Semantics](https://www.ardanlabs.com/blog/2018/12/garbage-collection-in-go-part1-semantics.html)
* [A Guide to the Go Garbage Collector](https://tip.golang.org/doc/gc-guide)