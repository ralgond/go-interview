# go-interview
go的面试题

## 协程
- 启动3个协程，按顺序打印"cat"、"dog"和"fish", 重复这个过程100次。[答案](https://github.com/ralgond/go-interview/blob/main/cmd/goroutine_run_sequentially.go)

## 锁
- sync.Mutex是协程级别的还是线程级别的锁？答案：线程级别
- sync.Mutex是悲观锁还是乐观锁？答案：悲观锁
- sync.RWMutex是悲观锁还是乐观锁？答案：悲观锁
