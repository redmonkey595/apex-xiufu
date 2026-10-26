package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "1.5" )

func gojuDLLeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tCZPUpRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vFjT9AS5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lOv4spQQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kx5yBXkVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func onWf99dmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wkS19qD9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lL2HoKoEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IG7jkBwjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sqMWZAsYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6hjAB81GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1HTyCuyrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MzpI5HI0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aWZLIsBDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pRzAJnUKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EXJVVcUmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KdPEOliSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6VHsE1RGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PNfqq0lwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pqaP5JnTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RKVltvtEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pF2zSzdGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oJAWXyr0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6LmdwpvOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3f5L0Gd6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xKiT06MoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dOOsEBKKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LlSHOkkeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a7CFSlRvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ulpPEJCKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YN6PmWhPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BniG3eE4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VSpaQzPpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4qEpAAI0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pfTzAby1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Moas43aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QNWAdT4WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nmX1rWpCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9SCZK40kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qmjjc4tvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func giamJjePWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o94h5dvbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func si1jleUUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AUUVqwtlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zu2MKCnXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kXA1fSPAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MT74zpW1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qOk2uHk4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ZYFqCTaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VIO31osJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5jSKNECnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wv472OsOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vpty1Lx5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ymBzOaAqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZrfgEKLJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aC637Je4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2bKDo8UCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KW0FTwW3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DxNc0IYZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RMhUAXx2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5AyGf9R9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sHIxJoA5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EXvbFj7GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func thiK3rORWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C7ZusBRUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IH7KE0euWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 65TiWhQcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JEmqXQaGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RGoHj4vpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NlYBB1zdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I0SRF2s4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KB3c29aLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ohQCIKlKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RIzaPfz0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J0tgolEIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Ce88aRvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uz4Oh9lIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sHYH98vWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LytHNlSLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I4d3mW9BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S4wmOcb5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sBFRZXBCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M8zdnZnJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ozc4rsgDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hgC5Vv7wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JPu8l6bsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5jyHuaFxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T32I9Fw3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BHQ99CySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pd2tsjzSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4XzOEfkwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MWFOXPzSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LFw1x1ElWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0GBGUp5TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V5mimeFGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZLeOUB8JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func saQHuHseWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MDijjpR9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pG6lHvptWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X8rPc7Z7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aOW5Z4yYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pvQNtOBLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xekEs6WMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r8ApRbteWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zvzlYLDTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FictIl6BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z5AFNaITWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BL6Oa1vCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XeVpUUO6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XuSJeD5QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NHjPZob9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m07MwRPUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4KaR3DznWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p1WwkW6HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UwAzcIM3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TkCk7FnqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mrir7L52Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iNbUIu6xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YFB5MdhkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w4vp8ZW9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jBh78gQCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xSOMwgRIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GLvPWeFRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IfRZ2xlhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G0FsZ1FAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DE6sjL0IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func afAQWaeRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2XJtFlAOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B04udUB3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MSg4Q2JcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ozFySg7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0I4Qh11VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zXQUuXJPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sSAHE7RDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TTaZSQZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DuxrMwkKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w48vetkIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H400u5IzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zQG4hV78Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hovm7fyoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f5eLwBBaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dfvRyaoOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HNPVYUZOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 04qZOdhgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QjI6j1RyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nobBHDHCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ORfYMexWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dzfiXrCOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OLkgDn4ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VdZFB926Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TKFKMYeKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C17ZxpRvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SYXI1ieEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8FhCA4W8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2jFUw1e6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z8zc4OhOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wdRSyndeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lgppdIafWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xYIryGVjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ldbjVkAQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AiOQQ40ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fz85GiQmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aArEfOxBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pqK4HUFEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tpb5oSicWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rhsSke90Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PUGplKeRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CKEFaG1qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2iBbwfh2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3v4ofzUrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8XH76KR5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 50n0DlJFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bB3QAy9bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yWXxtJAqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o2MOJXehWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ry3ofLMLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lAsQYFUcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 019kUpWXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7rfelMJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CqqgkO2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qgmpKR6yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sYaXvJlNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func onObVVUZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4x1iXDUzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DUuaafXvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rgFxjnC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 80Bfw9hGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4e51P2zIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VWWyoNynWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oibcfwBSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UvRHLWdNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D5SktQtpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GnTRZsKSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TpTAEp1MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QPFv4AhrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fsWYydrXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xgT0fDJPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VO8JC16UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7YWEEiHuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8wypPsvpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EhdC2dTkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7UH5YC7wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uSqoJsp3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YEC1fB4aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CnEdn8SsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9pobj5AuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cPqUUE3BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UEebUlhXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ErWd1HNSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qub0OOJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JRpfRbWZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vrjb4XNpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QfGFzNebWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LQ5GE2xnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8zAQCyHpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xoC97pR8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R24DatXhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ftjii90BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ohuBmXeeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b8RgeECVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func svu2c1RJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6K8HF5q4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JAq8WJn2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MZkhKTddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JAT4WnIzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QsmCBYB5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U8tjFnugWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vMfqLYBEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hqs02IKGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ste8jQk0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lZiQbYVGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8sA68ns4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o8Pl6JqOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qMthuBFPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rtsabxZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SmNOGfCCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kKJVsy3OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oShbG0beWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CbiL6rW9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6y3tfrTZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6x9U5JDXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kMBbJi6ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j6J16Yh7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z3im128KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5BuWar0GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DkmpAlclWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 41fiiBpEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0wWPkBKoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E1zcLq3sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xo30NmswWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3CEyjIddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QJak7QvAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q4tqJCD0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hkx0Gn2CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ax5QRq4WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CZIQ82ZaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aHoH11PbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RsKAKEPDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QyfdFypCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OHOnDw3LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gHAVPts9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T3l5TlPJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wEpan5OUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJrbgQARWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ThuLR8LuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ohs1UK0DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o0hvbJBYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mqqotC7ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C03tSPsKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JFPzIGRdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bm9UE9dTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V2Gp3AHLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9AfJSKNXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LRhz7PfZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EUQt5cNfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iTG91xeFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P598hSLfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fTQ82bZGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xwVErfaOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GqkiMtlGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func icGoftr3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KLwsIDUPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NlNuDvN7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ayWWa2GKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g7DU009dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L31Cra3wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zbtItVSSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oW3x2uS4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ku1mSilWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TQJ0dTTJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5J0AZTz5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WfHiuRxRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s02Npn4EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func twhNdiCCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8HC1jgGqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sFMev0YVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uVlG2LjUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FUd8SyGcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g4znFt1dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EmudVRnoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w56aSOCNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gaahYdtbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lm19h4CxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Ln47HWdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RRMwdIbgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d0IFDpYeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lb6U6wATWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MH2BzBSXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oPCTHF4XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gpfAQDnxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rKvkqHzrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kfzag7ciWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1oDTBOxrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0JEC7rweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EwqdGR0OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PfxlWr5BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1x8IKZ2oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xWS4xwWIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u8D36mD5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JiZ9aO4vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fa1XAjY9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CB2ZoPhFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CLoijKQVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AT1gcKHPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YwGaRTLlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NelkpYLpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q3fTWno1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OG9IwxDfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eWB3FC54Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F5I5hCp9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P3CGSBfkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nMJviZSdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 64EyZfDyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VvCQvlLkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hP3jumLzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VxRc724UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bpMn2Jg2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IzLBnrDyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ECRn1q85Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qntG4wvzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YN6UgzhIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ipc7hGCUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CBoLgR1kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dhajO4lhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QlnQ3eHgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jPgUwrzOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JX1zZS8iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ekwS75gIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BBzYisdsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RyQMF8RxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r6jRPrQSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4UzuU9ZtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fhePUx4oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zJbrI1IxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Gfz8UAfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 40QJct8NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mjHYEqoxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3rE3TDevWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MX3oEf33Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X44MsQ19Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 62ipqjdXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l4q4CnyyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pka6IsaRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qKrT1M25Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CkFLzKzbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gk3yHzxtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1dsm9rFoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5VO4vEGVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HkBkxlebWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func miHdiXgiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ii5Mq3gtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yZdOwUrNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nV2RmEyLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9DqCT9vaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dbw7KsYeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VdgZ32ZcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YUYDHP5xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WST0XGORWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PlPxtHIJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zk1klcB3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j0ivzU3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cugwhBZtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HsN47vnXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fQ2oiExkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C5wPFHNxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NDR39CauWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NAH5hZAuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YUszlmb2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cxMIrYrTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bTANN6T6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yVWADfeBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RYaAcNTIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AUoGb4ayWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0qs8X3OZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F0xolrX5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0bCyoxAZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JBkhMs0IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TBTrA28gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0cNP63jMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PEb3hMGzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wqpSXzxEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oKWQ7afgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sLjkevYXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AULs2RrpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HX3Ydp7iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h9zUA8lyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GiI4PONvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dQZhSPn6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NdcsVJ1bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 12BJJU81Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MBPeYMcxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mk168aDjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wEVDsrSXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E5qV6O3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MZmcT42WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3TaT4bgyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qBeiUSL6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fly9OvGlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SFkF6AyaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func St3p5ZLaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MpdekKiCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 898iOe5bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vTs1B2lZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wxlmLQJcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a1z2lM6mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U1OOtaJDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IjFtropgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IfI0HfMgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cQnkl1v4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 52d44qZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D3FS6gk9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qArtpMHJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QU3JFHgoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RQ72S78KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SYwRQyiFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ebh7cwn7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UBv0U60uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9iZ6CXzXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rUnokeaWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yhK9aH3vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mMENBCc3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZMyCqMLeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ppdk9CCBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VgYFylt2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iqLJLvRCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bUWnZYIEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qZalPmkcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wKDoNUdVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 57iiRmQaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ze9wVIWzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hMy7zhWBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3hvNY9WqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u61VgzPXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i4rNZ6EmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8h5RUPLvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3esJxLVGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zRL31A6TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wDWrBphSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cBrLNZpNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s5djni67Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k8gYjjRhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mau28D3zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BuUsWpzyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bwECZy1mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func suC7sA2AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q2gEuC7OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6COc0icGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ymHuLrp2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5StFRcgFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l5ZcYNdGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func izPIAjmJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w613Eq5dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Azo90vSiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tTIB0VO3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hIPMBQZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JZ6PmZ7XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jTqWD3jIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c5wwVsm0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pbqy3qDOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Vzm9N77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dIxDrIfdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G7e72SteWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WgJ2TqyLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ahYyhh2yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xKtjEndIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oP2YTCWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uqTob2yLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i2Vp1zq2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FNkQbbISWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V7ydktOpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5QPZXscUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7hRyLvsAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rICjG47FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LCPmlDoOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rbyvbiRuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uKkMDlieWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2A2MtuWaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xIrZLthAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JdZsND25Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func emK975A7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rOmEMiSeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HkbujUtxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9tndLluFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EGzoOe6aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func US7vHa05Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EQdpbRzmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kAthAA3UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oXNiyZenWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W1wNIlQKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6dZiIhRCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func puO8tpMTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q38QMEEbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uzjRaDPuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g5nIqMVTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5lyAkanxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nSphZmwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TYAaOV9sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fL6Er89uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vO7KFr8OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LqvtQ9RTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GGtQFaFUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p9Q4QlXYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func prm8rFZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V3Dhw5u6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6oiiVHl0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3nHYOJyGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zN3eVXzMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6nYDPnlyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pAMw7hkZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RdpoZ6GPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cvIQk63jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H9gofRTLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OsfXNhc1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LU6IJhlsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1miibkYWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func baTZUhcwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T9WUCNDOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ADuf4MzIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MDpxdUJ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t2auwg0ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wmYuKJq0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2A7SLgu1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CodwCojTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RzWSPiXUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kZI0cBS7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func godA2LeHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xVaadBsUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4gCUruvXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BKeET0bEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6xzzxPgMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vFGbPYnTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Nd88jsWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func esgvwvvNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JuIYalblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7NSnxjH4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kh2gFHGjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Erg8uSn9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bsuLOniLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ljjRHsbmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ftNKx83iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3mH3O0nKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2tMksTLkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NfjNQ9skWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m8Zp3WEcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZSOhnfUUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a0wJPSRTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OXkxEfFqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eebx3JyvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pkbKqlE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4tiY62zxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BXBFwQdBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p9HIYaNmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8TSo6f9aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CClHjRgcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1krRkH85Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3tm7aliRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CienaaNYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f2MlhX7IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c1Gy45oMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P2mxtnzuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EU3mgygEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RoE8MIoAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sGa1ek2MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4RxJAoCUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ic4aN7d4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x7GHbUI3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func owmQ8mL4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fg0W0gtjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v3wCUiyBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kIopmPjSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bp3AiaqnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MwuyZoj1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SKKXUrPLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cujR1HMKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H8sDl0vWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xRl9qrBwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A0hL9kfgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 61CDqgOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zmkg5icDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iYqxhiGJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zkuH9ufsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EZdnIDqjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QKu1lklOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oMbrdlP9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bRYI02zcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1FzRmBVmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5QxqOWVzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G7RXJsdvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jhk1jiKzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1HRgHcK0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e1NA10vMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LOCAgsXlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bws7N6xlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uFq4qAKDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FxayiAjfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3IDPlnjlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zBOP469CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 639Pib1kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7F6t0soEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WoXtqbusWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GXq5e6HZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ygkNZB0JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PMUFye4gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MT6hPd2tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vECcDmDXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dVTVKc33Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LUsYfBUeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KGa81DebWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fke0iBUhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func leZELDQYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PAM3oThBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pM6Uf2BVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pwUQ4KPpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I74qAlWwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 35vc0BrUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bJsmQRcxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B8qCe6UkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func InTu8lXEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SukgPFCxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jteNPncwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MNy2wSaFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J6TALzBcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PY5j8m6cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OEgP8Ei0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wB3NWcfeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MsEm7brYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AU2vAGWlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qXDmC19wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k78JA16qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v8qCJWk8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VGaDy2EaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Be8611deWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xy9RUHaeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jEJdme0kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7mRPuP0mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kB8tAbtSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OeYikPQRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6sp4bw6FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d9EGaUbvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HGH8nPKKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QglINf2dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VbTM1pIpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h95D1L2oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i3uYEBk5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vBPl39AUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3pM7Yz3vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ed6oq181Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ZGhRkjMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 72z4m8GnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WTturE8bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DCHD0YMKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CzpNEFGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4NePXgoeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zGmGo8KXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 68rwJtFYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B3ocLtD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3oEuedl5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4bQkHWVWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VzTUvJxyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qpABL8CoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R4xzmcpIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func etRAFDEiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XFfcHVA1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uhTB1wCTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 61Yh4VPfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cDHyBgtGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 34xgnzK2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SzKieqVqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wBoIXaUtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G7Xjc5JGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IlSs0Ml1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qnqIGVt6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1KR9CcjBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ko9hBz2RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k3I2BBRUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ecor6GKhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d1bM9Nz3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j7VA97XQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 71t8684JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k6QjHyCSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OhowKmWEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ORfckC0fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VhQ4sRs1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AHwg9YvTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uT6MFowkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JLCnDFsKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8pEa7YqnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G1vReCJIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7pME7nIfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RQNIdh01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func axbhHvsZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PiXZpBBdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func orqcfsTSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lpwbeMzrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dQjwB4zMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rr4t2QARWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e7RTGiWiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uLfI7CpkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SUl9JcvuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hCXVFFkBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T0vYGpURWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6irs2piDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TFvU32UuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lxupyQLeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i81wbToTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FdYmpr48Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qCB0V9tQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X4ZQJflfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o1SsreTZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vB8K2oCTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MAY9BGBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oxQU1yXzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iHOFxR01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YEt9nMMsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EJE7xcUGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V71laISHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rHchMhfYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Y7cUGlGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CAiluq2cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PKBKSMuhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MVvGqXOqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rjYM9v0FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QMXCVVgfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b1LznHIrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EkXRIsRyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FE7ZCJ22Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f73vFYMjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B2abaHPhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ozVQRmxzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MvnEnPeoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qgQv8qXEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eLK1egisWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 75Ng3zZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 86R3D2C6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 97ZrHk4bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eJQgF47lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gFTz9w2hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pVW3d8meWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PLAmtVH1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CIIoeSdhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xNmtyrmVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j7MJJ8kyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PdoNAEY3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mCxd80ylWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 732vLK9zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XFmsbrWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9zCquu9rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z2fwKhw2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d1AcJDFfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func thIY8DxZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6fBQkS0FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DCQZXDUrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func afFavgzmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4neePWf7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ef70scCwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tLdrOe1vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 77RgUzxsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a66Fz3BjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pMfY7BTPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LLgf4cVxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RjmgM4kGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6TkB7iNCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jIWVDhkgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U7EZAo1bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pqad3Zj1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YQdkaS2GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BzMpuojoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nsUeRWQrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hvsv8aTxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pZLPaIyNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n0PtLrJcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HLpZRvDmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iIeci79NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M1yfDAuAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cLK84ZNwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J4mmYJr9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eCj5rvJvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kOUrwakRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e9mzkQIzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1CSpX6hUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xygf2r1mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xdd1oKryWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SI6wT9VcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IY5JKujZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8u7RR9w3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BWism2ZwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nudthd1vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OKEgDmcdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uU9rBkG9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dpx9fqw3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func szvT2eqPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ht1pPloDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZUeAwMAmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jR7k57SkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func orWnuLxqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d7KD9IiyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IQEqKdWFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CtRsxKHOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kfug1oUaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F6kW0VN6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4sOYacRYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dTXrW1CBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MYUe3RIkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NBRdvpB0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dIhtHp85Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oMedYZ98Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GIpso90wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HAAvh7qkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f9zuDwd7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qM8oRu1ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IZ7mNF3eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lCoSJCQkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LsWQ3DYfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UmrOkXRfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TulMON58Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 50jsCpl1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z2Zx5P6WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QIKMDnsvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ZyXFW5EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aiz99mJ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4XEGuDNlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7bkuY0fNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aW3ol1iXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hZNYE3CWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2JBmiWirWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oKEEmumMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func favFtmCEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p7IHj5EEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C6x8GyXWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bbNBKjTGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O5K44DluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dADy3wtUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Heh698P3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EStkjSh6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i2q7uLcIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func snPnaLmGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FPyoIR6TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZSlG63hrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xtvPk5rSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C4WnKVLzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D0U70zQmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NzaYnV1cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eAiKhvmxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LWQqGZCcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d6BYy0ZHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pseIfSzHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X3oiseRPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XG4dNRWcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VvolyoTwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SztLb2jnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MpaUChENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ty2HnaO0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V8x28xlwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WSzfNIBnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kPkTZscBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VFU2ZzzoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6nk07P84Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8iR49hnsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GIRpgXtNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lHRlWZhIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k8C7zRatWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BJbPafhCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vgbGlEQeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qxnGWML4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

