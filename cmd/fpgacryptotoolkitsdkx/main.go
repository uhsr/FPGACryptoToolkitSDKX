// cmd/fpgacryptotoolkitsdkx/main.go
package main

import (
"flag"
"log"
"os"

"fpgacryptotoolkitsdkx/internal/fpgacryptotoolkitsdkx"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := fpgacryptotoolkitsdkx.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
