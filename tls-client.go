package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	certFile = flag.String("cert", "/etc/puppetlabs/puppet/ssl/certs/klika-dev.cgca.uwm.edu.pem", "A PEM eoncoded certificate file.")
	keyFile  = flag.String("key", "/etc/puppetlabs/puppet/ssl/private_keys/klika-dev.cgca.uwm.edu.pem", "A PEM encoded private key file.")
	caFile   = flag.String("CA", "/etc/puppetlabs/puppet/ssl/certs/ca.pem", "A PEM eoncoded CA's certificate file.")
)

func letstls(stringin string) string{
	flag.Parse()

	// Load client cert
	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Fatal(err)
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile(*caFile)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	// Do GET something
	resp, err := client.Get(stringin)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Dump response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

//func main() {
//	fmt.Println(letstls("klika-dev.cgca.uwm.edu"))
//}
