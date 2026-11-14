package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "3.8" )

func h5JTFkAUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lu5aJQM9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cWDCAzSxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FfA7w9uLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oFg1UmlNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 93Ub8n82Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Za3ESXPQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8oaWLF7IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mOfupLJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FlPDjlbkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qsvr1PNKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YrN3B8hFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TgcqacrhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6NJuomUXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2fjGScf8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C4mnoKyCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ACmmDf8zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cq9XAwkYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1dWG8kioWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nYxEikKUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VckdDHRPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qigvQkzYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h7ftd89uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m7mCw04HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zuASCUBAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MteWvBXdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wLHeHJUhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ob9DoGQIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fsfbsvI7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H4uu6XukWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AUEHf9biWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3kHQPOiHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wrqXSrUhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WyFVE3POWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XrGqBCbGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8qymwJhdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SimeD6RXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K3cVrx7HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F407CLmyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0cWrxcPXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rPgQU6W8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TckvZgywWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NWDN0DoHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wz8msV3qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v75dYUZKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hWZ5NrmwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RFP5hwFZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DmTUvJB8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DxcJNTndWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GmqaScWJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 92D52OcfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fm1Hr0PLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func soeJqSqBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ADiI9SBeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ueQodc9wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IikIGtIsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RKEvGRk3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8uVGCEpEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IJopI0QvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iI5ePa9lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t4bHCIhLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func to2JCnJBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RCMLel4TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JKC9iqR4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mDjiMiRZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SfZbMioYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ediW6NLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nFaT7XRfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8fAT3pCuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E3Jly144Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mwVNQiSOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

