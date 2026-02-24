package chiper

import (
	"strings"
)

const kl = "BCDFGHJKLMNPQRSTVWXYZ"
const vl = "AEIOU"

type ROTConf struct {
	ko int
	vo int
	ks int
	vs int
}

func NewROTConf(ko, vo, ks, vs int) ROTConf {
	return ROTConf{ko, vo, ks, vs}
}

type ROT struct {
	conf ROTConf
}

func NewROT(conf ROTConf) *ROT {
	return &ROT{conf}
}

func (r *ROT) Encrypt(plain []byte) []byte {
	chiper := make([]byte, 0)
	tks := r.conf.ks
	tvs := r.conf.vs

	for _, v := range plain {
		if idx := strings.Index(string(kl), string(v)); idx > -1 {
			cv := kl[(idx+r.conf.ko+tks)%len(kl)]

			chiper = append(chiper, byte(cv))
			tks *= 2
			continue
		}

		if idx := strings.Index(string(vl), string(v)); idx > -1 {
			cv := vl[(idx+r.conf.vo+tvs)%len(vl)]

			chiper = append(chiper, byte(cv))
			tvs *= 2
			continue
		}

		chiper = append(chiper, v)
	}

	return chiper
}

func (r *ROT) Decrypt(chiper []byte) []byte {

	return []byte("")
}
