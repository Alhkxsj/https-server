1. Certificate File Location

After running hcertgen, a CA certificate file will be generated
The certificate file is placed in the home directory by default
~/hserve-ca.crt


---

2. Copy Certificate to Phone Storage

You can copy it to:

/storage/emulated/0/Download/
You don't have to copy it here specifically, just make sure you can find it when selecting the certificate file

---

3. Android Installation Steps

1. Open Settings


2. Security → Encryption & Credentials


3. Install Certificate → CA Certificate


4. Select hserve-ca.crt


5. If you can't find it, search in the settings top search box (certificate) and find the relevant certificate installation search results, then install the certificate.


---

6. Notes

Android will warn "Certificate can monitor traffic" - this is normal

Certificate is only for your locally generated HTTPS service

No upload, no internet, no sharing