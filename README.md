# FPGACryptoToolkitSDKX

## Description



## Features

- Implements hardware acceleration for SHA-3 and Keccak hash functions using pipelined architectures.
- Provides a high-level synthesis (HLS) interface for generating custom cryptographic cores from C++ code.
- Integrates a parameterized AES core with configurable key sizes and modes of operation.
- Supports side-channel attack resistance through masking and shuffling techniques implemented in VHDL.
- Offers a secure key management system with FPGA-based true random number generation (TRNG).
- Includes a pre-verified IP core for post-quantum cryptography (PQC) algorithms like CRYSTALS-Kyber.
- Generates optimized bitstreams for Xilinx and Intel FPGA families, targeting specific resource utilization and performance.
- Verifies cryptographic implementations against NIST test vectors using a built-in self-test framework.
## Installation

```bash
go get github.com/uhsr/FPGACryptoToolkitSDKX
```

## Usage

```go
package main

import (
    "fpgacryptotoolkitsdkx/internal/fpgacryptotoolkitsdkx"
)

func main() {
    app := fpgacryptotoolkitsdkx.NewApp(false)
    app.Run()
}
```

## Contributing

We welcome contributions! Here's how to get started:

1. Fork this repository
2. Create a new branch for your feature (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -am 'Add some awesome feature'`)
4. Push to the branch (`git push origin feature/your-feature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.
