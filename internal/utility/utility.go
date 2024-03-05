package utility

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
    rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomInt(min, max int64) int64 {
    return min + rand.Int63n(max - min + 1)
}

func RandomString(n int) string {
    var sb strings.Builder
    k := len(alphabet)

    for i := 0; i < n; i++ {
        c := alphabet[rand.Intn(k)]
        sb.WriteByte(c)
    }

    return sb.String()
}

func RandomUsername() string {
    return RandomString(6)
}

func RandomPassword() string {
    return RandomString(18)
}

func RandomFullName() string {
    return fmt.Sprintf("%s %s", RandomUsername(), RandomUsername())
}

func RandomBalance() int64 {
    return RandomInt(0, 100)
}

func RandomEmail(username string) string {
    return fmt.Sprintf("%s@%s.com", username, RandomString(4))
}

func RandomCurrency() string {
    currencies := []string{"EUR", "USD", "CAD"}
    n := len(currencies)
    return currencies[rand.Intn(n)]
}

