package hw10programoptimization

import (
	"bufio"
	"bytes"
	"io"

	"github.com/valyala/fastjson"
)

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	var fParser fastjson.Parser
	result := make(DomainStat)
	searchBytes := []byte("." + domain)
	delimiter := []byte("@")

	bScanner := bufio.NewScanner(r)

	for bScanner.Scan() {
		val, err := fParser.ParseBytes(bScanner.Bytes())
		if err != nil {
			return nil, err
		}

		email := val.GetStringBytes("Email")
		email = bytes.ToLower(email)
		email = bytes.TrimSpace(email)

		// Если email заканчивается на нужный домен
		if bytes.HasSuffix(email, searchBytes) {
			result[string(bytes.SplitN(email, delimiter, 2)[1])]++
		}
	}

	return result, nil
}
