1. Design Goals

The security model of this project follows these principles:

Local-first

Full user control

No third-party CA dependency

No complex PKI system



---

2. Trust Model

User
 └─ Install local CA (user actively trusts)
      └─ hserve (only valid locally)

CA is only generated on this device

Private key never leaves the device



---

3. Why Use Self-Signed CA

Reasons:

Let's Encrypt is not suitable for local / IP usage

Android local development certificate needs are clear

Self-signed CA = User actively trusts


This is a developer tool, not a public network service.


---

4. TLS Policy

TLS minimum version: TLS 1.2

Disable insecure protocols

Certificate validity is longer to reduce repeated operations


Specific parameters defined in:

internal/tls/policy.go


---

5. Things Not Done (Intentionally)

[X] Automatically install system certificates

[X] Bypass Android security prompts

[X] Background resident service


Users must clearly know what they are doing.


---

4. Applicable Scenarios Summary

Local HTTPS development testing

Android ↔ PC file sharing

LAN device access


Not suitable for:

Public network deployment

Commercial HTTPS services



---

5. Conclusion

This is a tool designed for clear-minded people.

No magic

No hidden behavior

All certificates and trusts are in your hands


If you understand HTTPS, you will like it.