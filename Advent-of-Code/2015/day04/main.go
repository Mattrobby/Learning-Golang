package main

import(
  "fmt"
  "crypto/md5"
  "encoding/hex"
)

func main() {
  secret := "iwrupvqb"
  num := 1

  for {
    hasher := md5.New()
    input := fmt.Sprintf("%s%d", secret, num)

    hasher.Write([]byte(input))
    sumBytes := hasher.Sum(nil)
    sum := hex.EncodeToString(sumBytes)

    if sum[:5] == "00000" {
      fmt.Println("Part 1 | MD5 Number:", num)
      break
    } else {
      num++
    }
  }

  for {
    hasher := md5.New()
    input := fmt.Sprintf("%s%d", secret, num)

    hasher.Write([]byte(input))
    sumBytes := hasher.Sum(nil)
    sum := hex.EncodeToString(sumBytes)

    if sum[:6] == "000000" {
      fmt.Println("Part 2 | MD5 Number:", num)
      break
    } else {
      num++
    }
  }
}
