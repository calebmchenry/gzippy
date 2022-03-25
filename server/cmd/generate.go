package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
)

type Thing struct {
	Rules []Rule
}

type Rule struct {
	IDs []string
}

func main() {
	thing := Thing{}
	for i := 0; i < 100; i++ {
		rule := Rule{}
		for j := 0; j < 5_000; j++ {
			rule.IDs = append(rule.IDs, id())
		}
		thing.Rules = append(thing.Rules, rule)
	}
	file, _ := json.MarshalIndent(thing, "", " ")

	_ = ioutil.WriteFile("static/10MB.json", file, 0644)
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func id() string {
	const length = 20
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
