# hcpb

Simple console progress bar for Golang which shows the following info
- Progress bar
- Position
- Maximum value
- Rate
- Elapsed time
- ETA time

```
[>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>] 79.6k/79.6k @ 267/s in 04:58 ETA 00:00
```

It's buggy and incomplete and I don't recommend that you only use it for educational purposes

# Example

```
package main

import (
	"github.com/hypernova-za/hcpb"
	"time"
)

func main() {
	b := hcpb.New(79563)
	for i := 1; i <= 79563; i++ {
		b.Inc()
		time.Sleep(time.Millisecond * 3)
	}
}

```
 