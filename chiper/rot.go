package chiper

import (
	"strings"
)

const kl = "BCDFGHJKLMNPQRSTVWXYZ"
const vl = "AEIOU"

type KVROTConf struct {
	ko int
	vo int
	ks int
	vs int
}

func NewROTConf(ko, vo, ks, vs int) KVROTConf {
	return KVROTConf{ko, vo, ks, vs}
}

type KVROT struct {
	conf KVROTConf
}

func NewROT(conf KVROTConf) *KVROT {
	return &KVROT{conf}
}

func (r *KVROT) Encrypt(plain []byte) []byte {
	chiper := make([]byte, 0)
	tks := r.conf.ks
	tvs := r.conf.vs

	for _, v := range plain {
		if idx := strings.IndexRune(string(kl), rune(v)); idx > -1 {
			cv := kl[(idx+r.conf.ko+tks)%len(kl)]

			chiper = append(chiper, byte(cv))
			tks += tks % len(kl)
			continue
		}

		if idx := strings.IndexRune(string(vl), rune(v)); idx > -1 {
			cv := vl[(idx+r.conf.vo+tvs)%len(vl)]

			chiper = append(chiper, byte(cv))
			tvs += tks % len(vl)
			continue
		}

		chiper = append(chiper, v)
	}

	return chiper
}

func (r *KVROT) Decrypt(chiper []byte) []byte {
	plain := make([]byte, 0)
	tks := r.conf.ks
	tvs := r.conf.vs

	for _, v := range chiper {
		if idx := strings.IndexRune(string(kl), rune(v)); idx > -1 {
			cv := kl[(idx-(r.conf.ko+tks)%len(kl)+len(kl))%len(kl)]

			plain = append(plain, byte(cv))
			tks += tks % len(kl)
			continue
		}

		if idx := strings.IndexRune(string(vl), rune(v)); idx > -1 {
			cv := vl[(idx-(r.conf.vo+tvs)%len(vl)+len(vl))%len(vl)]

			plain = append(plain, byte(cv))
			tvs += tks % len(vl)
			continue
		}

		plain = append(plain, v)
	}

	return plain
}
