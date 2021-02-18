package utils

/*
令牌桶限流
1. token 固定速率放入桶中(用当前时间 - 最后一次放入时间 * 速率)
2. 新请求来的时候移除一个 token
3. 如果移除的速率 <= 放入的，则正常
4. 否则，则拒绝
*/

type TokenBucket struct {
    rate         int64 //固定的token放入速率, r/s
    capacity     int64 //桶的容量
    tokens       int64 //桶中当前token数量
    lastTokenSec int64 //桶上次放token的时间戳 s

    lock sync.Mutex
}

func (l *TokenBucket) Allow() bool {
    l.lock.Lock()
    defer l.lock.Unlock()

    now := time.Now().Unix()
    l.tokens = l.tokens + (now-l.lastTokenSec)*l.rate // 先添加令牌
    if l.tokens > l.capacity {
        l.tokens = l.capacity
    }
    l.lastTokenSec = now
    if l.tokens > 0 {
        // 还有令牌，领取令牌
        l.tokens--
        return true
    } else {
        // 没有令牌,则拒绝
        return false
    }
}

func (l *TokenBucket) Set(r, c int64) {
    l.rate = r
    l.capacity = c
    l.tokens = 0
    l.lastTokenSec = time.Now().Unix()
}
