# go-interview
go的面试题

## 协程
- 启动3个协程，按顺序打印"cat"、"dog"和"fish", 重复这个过程100次。[答案](https://github.com/ralgond/go-interview/blob/main/cmd/goroutine_run_sequentially.go)
- 已经有线程了，为什么还要有协程？[答案](https://github.com/ralgond/go-interview/blob/main/%E5%B7%B2%E7%BB%8F%E6%9C%89%E7%BA%BF%E7%A8%8B%E4%BA%86%EF%BC%8C%E4%B8%BA%E4%BB%80%E4%B9%88%E8%BF%98%E8%A6%81%E6%9C%89%E5%8D%8F%E7%A8%8B%EF%BC%9F.md)
  
## 锁
- sync.Mutex是协程级别的还是线程级别的锁？答案：线程级别
- sync.Mutex是悲观锁还是乐观锁？答案：悲观锁
- sync.RWMutex是悲观锁还是乐观锁？答案：悲观锁
- 你用sync.Mutex做过什么？回答：锁住LRU Cache，因为每次访问都会将被访问的item移动到队列头，每次插入都可能要删除队列尾的元素。

## CPU Cache和内存序
- go语言cpu改变了变量a，但还没有刷到主存中，这时协程被切换出去，请问a会刷到主存中吗？猜测：会的，我猜是在切换协程前利用内存序机制保证a刷到主存中。
