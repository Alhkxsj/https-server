1. Project Introduction

hserve is a simple and easy-to-use HTTPS file server:

Auto-generate CA and server certificates

Suitable for local development / LAN file sharing

Specifically adapted for Termux (Android) environment

No external CA dependency, no internet connection required



---

2. Installation

Termux

make termux-install

After installation, you will get:

hserve - HTTPS file server (with gen-cert subcommand for certificate generation)



---

3. Generate Certificate (Required)

Before first use, you must generate certificates:

hcertgen

Generated content:

CA root certificate (for installation to Android system)

Server certificate + private key (for server use)



---

4. Install CA Certificate to Android

See documentation: android-ca-install.md

[WARNING] Without installing CA, browsers will show "Not Secure" warning.


---

5. Start Server

hserve

Common parameters:

-port   Listening port (default 8443)
-dir    Shared directory (default current directory)
-quiet  Quiet mode (no access logs)

Example:

hserve -dir=/sdcard -port=9443