# The underlying principles of Coroutine

## Channel

Go语言常见的一种设计模式，是通过**通信**的方式去**共享内存**。这和许多主流的编程语言不同，比如 Java主要通过共享内存的方式进行线程之间的通信，结果会造成很多并发安全问题。

Go语言提供的**并发模型**，称为**通信顺序进程**，即 CSP (Communicating sequential processes)。在Go编程中，Goroutine作为CSP中的实体，而channel作为传递信息的媒介。Goroutine之间通过channel来传递数据
