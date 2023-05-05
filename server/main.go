package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func printHeader(r *http.Request) {
	log.Print("=============== Header ===============")
	for name, values := range r.Header {
		for _, value := range values {
			log.Printf("%v:%v", name, value)
		}
	}
}

func printConnState(state *tls.ConnectionState) {
	log.Print("=============== State ===============")
	log.Printf("Version: %x", state.Version)
	log.Printf("HandshakeComplete: %t", state.HandshakeComplete)
	log.Printf("DidResume: %t", state.DidResume)
	log.Printf("NegotiatedProtocol: %s", state.NegotiatedProtocol)
	log.Printf("NegotiatedProtocolIsMutual: %t", state.NegotiatedProtocolIsMutual)

	log.Printf("Certificate chain:")
	for i, cert := range state.PeerCertificates {
		subject := cert.Subject
		issuer := cert.Issuer
		log.Printf(" %d subject:/C=%v/ST=%v/L=%v/O=%v/OU=%v/CN=%s",
			i,
			subject.Country,
			subject.Province,
			subject.Locality,
			subject.Organization,
			subject.OrganizationalUnit,
			subject.CommonName,
		)
		log.Printf(" issuer:/C=%v/ST=%v/L=%v/O=%v/OU=%v/CN=%s",
			issuer.Country,
			issuer.Province,
			issuer.Locality,
			issuer.Organization,
			issuer.OrganizationalUnit,
			issuer.CommonName,
		)

	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	printHeader(r)
	if r.TLS != nil {
		printConnState(r.TLS)
	}
	log.Print("=============== End ===============")
	fmt.Println()
	io.WriteString(w, "Hello, world!\n")
}

func main() {
	port := 8080
	sslPort := 8443

	handler := http.NewServeMux()
	handler.HandleFunc("/hello", helloHandler)

	// Listen to port 8080
	go func() {
		server := http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: handler,
		}
		fmt.Printf("(HTTP) Listen on: %d\n", port)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("(HTTP) error listening to port: %v", err)
		}
	}()

	// load CA certificate file and add it to list of client CAs
	caCertFile, err := os.ReadFile("../external/ca.crt")
	if err != nil {
		log.Fatalf("error reading CA Certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCertFile)

	tlsConfig := &tls.Config{
		ClientCAs:                caCertPool,
		ClientAuth:               tls.RequireAndVerifyClientCert,
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		},
	}
	tlsConfig.BuildNameToCertificate()

	// listen on port 8443
	server := http.Server{
		Addr:      fmt.Sprintf(":%d", sslPort),
		Handler:   handler,
		TLSConfig: tlsConfig,
	}

	fmt.Printf("(HTTPS) Listen on: %d\n", sslPort)
	if err := server.ListenAndServeTLS("../external/server.crt", "../external/server.key"); err != nil {
		log.Fatalf("(HTTPS) error listening to port: %v", err)
	}
}
