package util

import (
	"log/slog"
	"math/rand"
	"time"
)

func RandomWithSeed() {
	rand.Seed(time.Now().Unix())
	a := rand.Intn(2000)
	seed := rand.New(rand.NewSource(time.Now().Unix()))
	b := seed.Intn(2000)
	if a == b {
		slog.Info("生成的随机数", slog.Int("a", a), slog.Int("b", b))
	} else {
		slog.Info("不相等")
	}
}
