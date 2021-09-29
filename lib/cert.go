package lib

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"strings"
)

/**
 * 获取域名证书链
 */
func GetCertList(host string) ([]*x509.Certificate, error) {
	cfg := tls.Config{}
	conn, err := tls.Dial("tcp", host+":443", &cfg)
	if err != nil {
		return nil, err
	}
	// Grab the last certificate in the chain
	certChain := conn.ConnectionState().PeerCertificates

	/*
		// 打印证书信息
		cert := certChain[len(certChain)-1]
		// Print the certificate
		result, err := certinfo.CertificateText(cert)
		if err != nil {
			log.Fatal(err)
		}
		log.Print(result)
	*/

	return certChain, nil
}

/**
 * 从线上获取域名证书文件，用于客户端证书绑定
 */
func GetCert(host string) ([]string, error) {

	certChain, err := GetCertList(host)
	if err != nil {
		return nil, err
	}

	if len(certChain) == 0 {
		return nil, errors.New("certChain is empty")
	}

	rows := []string{}
	for _, val := range certChain {
		c := base64.StdEncoding.EncodeToString(val.Raw)
		mix := []string{"-----BEGIN CERTIFICATE-----"}
		for {
			if len(c) >= 64 {
				mix = append(mix, c[:64])
				c = c[64:]
			} else {
				mix = append(mix, c)
				break
			}
		}
		mix = append(mix, "-----END CERTIFICATE-----")
		rows = append(rows, strings.Join(mix, "\n"))
	}

	return rows, nil
}
