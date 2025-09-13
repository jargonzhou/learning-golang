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