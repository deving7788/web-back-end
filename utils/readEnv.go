package utils

import (
    "fmt"
    "os"
    "strings"
)

func ReadEnv(name string) (string, error) {
    var value = ""
    content, err := os.ReadFile("utils/../.env")
    if err != nil {
        return "", fmt.Errorf("error reading .env: %v\n", err) 
    }
    pairs := strings.Split(string(content), string(10))
    for _, v := range pairs {
        kvp := strings.Split(v, "=")
        if kvp[0] == name {
            value = kvp[1]
            break
        }
    }
    return value, nil
}
