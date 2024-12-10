# README

I have created the server certificates and key by following this blog, feel free to create your own !

[How to Implement TLS ft. Golang](https://medium.com/@harsha.senarath/how-to-implement-tls-ft-golang-40b380aae288)

If you are lazy like me, here you go :

```
openssl req -new -newkey rsa:2048 -keyout ca.key -x509 -sha256 -days 365 -out ca.crt
```

create server.cnf and don't forget to replace CHANGEME

```
[req]
default_md = sha256
prompt = no
req_extensions = v3_ext
distinguished_name = req_distinguished_name

[req_distinguished_name]
CN = #CHANGEME

[v3_ext]
keyUsage = critical,digitalSignature,keyEncipherment
extendedKeyUsage = critical,serverAuth,clientAuth
subjectAltName = DNS:#CHANGEME
```

After this, run the following commands to generate the certificate and keys

```
openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr -config server.cnf
openssl req -noout -text -in server.csr
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key \
  -CAcreateserial -out server.crt -days 365 -sha256 -extfile server.cnf -extensions v3_ext
```

Now just keep the server.key and server.crt and throw the rest if you want to. 

Throwing Certificate Authority certificate and key might comeback biting you later if you are deploying this to prod !!


