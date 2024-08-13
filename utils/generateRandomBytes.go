package utils

import (
    "math/rand"
)

func GenerateRandomBytes(n int) []byte {
    bytesA := make([]byte, n)
    for i := 0; i < n; i++ {
        bytesA[i] = byte(rand.Intn(75) + 48)
        for {
            if bytesA[i] == ':' || bytesA[i] == '?' || bytesA[i] == '\\' || bytesA[i] == '`' ||
            bytesA[i] == '=' || bytesA[i] == ';' || bytesA[i] == '@' || bytesA[i] == '^' ||
            bytesA[i] == '<' || bytesA[i] == '>' || bytesA[i] == '_' || bytesA[i] == '[' || bytesA[i] == ']' {
                bytesA[i] = byte(rand.Intn(75) + 48)
            }else {
                break
            }
        }
    }

    return bytesA
}
