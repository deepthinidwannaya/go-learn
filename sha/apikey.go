package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
)

func main() {
	input := "MzU1MDI0NjAwZmM5YjkwOGY2N2RmN2QzNzljNWEzYzM5YTgzMzdkZGViYzUyMzU2MTExNWZiZDczYmVhOGJiMw"
	sha_512 := sha512.New()
	sha_512.Write([]byte("saltnpepper"))
	sha_512.Write([]byte(input))
	result := sha_512.Sum(nil)
	shaString := hex.EncodeToString(result)
	fmt.Printf("%x", result)
	fmt.Println(reflect.TypeOf(result))
	filter := bson.M{"key_hash": shaString}
	fmt.Printf("keyHash is %v", shaString)
	fmt.Printf("filter is %v", filter)
}
