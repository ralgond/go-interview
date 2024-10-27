# go-interview
go的面试题

## 协程
- 启动3个协程，按顺序打印"cat"、"dog"和"fish", 重复这个过程100次。[答案](https://github.com/ralgond/go-interview/blob/main/cmd/goroutine_run_sequentially.go)

## 锁
- sync.Mutex是协程级别的还是线程级别的锁？答案：线程级别
- sync.Mutex是悲观锁还是乐观锁？答案：悲观锁
- sync.RWMutex是悲观锁还是乐观锁？答案：悲观锁
- 你用sync.Mutex做过什么？回答：锁住LRU Cache，因为每次访问都会将被访问的item移动到队列头，每次插入都可能要删除队列尾的元素。

## CPU Cache和内存序
- go语言cpu改变了变量a，但还没有刷到主存中，这时协程被切换出去，请问a会刷到主存中吗？猜测：会的，我猜是在切换协程前利用内存序机制保证a刷到主存中。
