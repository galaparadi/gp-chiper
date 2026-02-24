package main

import (
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/galaparadi/gp-chiper/chiper"
)

func main() {
	plain := "GALA PARADI"
	max := 9999
	min := 0

	ko := rand.IntN(max+1-min) + min
	vo := rand.IntN(max+1-min) + min
	ks := rand.IntN(max+1-min) + min
	vs := rand.IntN(max+1-min) + min

	fmt.Printf("ko : %d vo : %d ks : %d vs : %d\n", ko, vo, ks, vs)

	chip := chiper.NewROT(chiper.NewROTConf(ko, vo, ks, vs))

	fmt.Println(string(chip.Encrypt([]byte(strings.ToUpper(plain)))))
}
