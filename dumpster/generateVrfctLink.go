package utils

import (
    "fmt"
)

func GenerateVrfctLink(baseStr string, userId int, length int) (string, []byte){
    randomBytes:= GenerateRandomBytes(length) 
    vrfctLinkStr := baseStr + "token=" + string(randomBytes) + "&id=" + fmt.Sprintf("%d", userId) 
    return vrfctLinkStr, randomBytes
}
