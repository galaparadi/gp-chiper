package main

import (
	"fmt"
	"strings"

	"github.com/galaparadi/gp-chiper/chiper"
)

func main() {
	plain := "GALA PARADI"

	chip := chiper.NewROT(chiper.NewROTConf(377469, 6829, 141, 15))

	fmt.Println(string(chip.Encrypt([]byte(strings.ToUpper(plain)))))
}
