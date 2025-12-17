# hserve

A tool for quickly setting up a local HTTPS server.

## Features

- [x] Single executable file with multi-subcommand design
- [x] Smart certificate management (auto-generate, install)
- [x] Access path whitelist control
- [x] Support for external TLS certificates
- [x] Interactive language selection during installation
- [x] Dynamic Chinese/English interface switching
- [x] Multi-architecture cross-compilation support
- [x] Termux environment optimization

## Usage

> [WARNING] Note: After each new version installation, it's recommended to regenerate certificates and install them into the system.

For detailed usage instructions, please refer to [Usage Documentation](./docs/usage.md).

## Subcommand Details

hserve supports the following subcommands:

### gen-cert - Generate Certificate
```bash
hserve gen-cert           # Generate certificate
hserve gen-cert --force   # Force regenerate certificate
```

### serve - Start Server
```bash
hserve serve                                # Start server (default port 8443, current directory)
hserve serve --port 9443 --dir /sdcard     # Specify port and directory
hserve serve --quiet                       # Quiet mode (no access logs)
hserve serve --allow /sdcard --allow /home # Set access whitelist
hserve serve --auto-gen                    # Auto-generate certificates for first run
hserve serve --tls-cert-file cert.pem --tls-key-file key.pem # Use external certificates
```

### language - Switch language
```bash
hserve language en    # Switch to English
hserve language zh    # Switch to Chinese
```

### install-ca - Install CA certificate to Termux trust store
```bash
hserve install-ca    # Install CA certificate to Termux trust store
```

### export-ca - Export CA certificate
```bash
hserve export-ca     # Export CA certificate to download directory for manual installation to Android system
```

## Build and Installation

### Requirements
- Go 1.21 or higher
- Termux environment (for installation and usage)

### Build Commands

```bash
# Build binary
make build

# Build multi-architecture versions
make multiarch

# Build deb package
make deb

# Format code
make fmt

# Check code
make vet

# Run tests
make test

# Clean build files
make clean
```

### Installation Methods

**Method 1: Direct binary installation**
```bash
make install
```

**Method 2: Install deb package**
```bash
# Build and install
make install-deb

# Or manual installation
dpkg -i dist/*.deb
```

## Architecture

### Project Structure
```
hserve/
├── cmd/                 # Command line entry
│   ├── hserve/          # Main program entry
│   └── root.go          # Cobra root command
├── internal/
│   ├── certmanager/     # Certificate management module
│   │   ├── generate.go  # Certificate generation logic
│   │   ├── check.go     # Certificate check logic
│   │   └── install.go   # Certificate installation logic
│   ├── server/          # HTTP server module
│   ├── tls/             # TLS configuration policy
│   └── i18n/            # Internationalization support
├── scripts/             # Build scripts
│   ├── build-deb.sh     # deb package build script
│   └── build-multiarch.sh # Multi-architecture build script
└── docs/                # Documentation
```

## License

This project is licensed under [LICENSE](./LICENSE).

## Usage Notes
Attention: After each new version installation, you need to regenerate the certificate and remove the old certificate from the system. Then reinstall the new certificate.
For detailed usage instructions, please refer to [Usage Documentation](./docs/usage.md).

...

For information about the security model, please refer to [Security Model Documentation](./docs/security-model.md).

...

For detailed steps on installing CA certificates on Android devices, please refer to [Android CA Installation Documentation](./docs/android-ca-install.md).

## Build and Installation

### Build

```bash
make build
```

### Build deb package

```bash
make deb
```

### Install

```bash
make install
```

Install deb package:

```bash
dpkg -i dist/*.deb
```

## License

This project is licensed under [LICENSE](./LICENSE).