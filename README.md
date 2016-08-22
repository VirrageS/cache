# cache
Simple, in-memory, thread-safe TTL cache for Go

## Usage

```go
import (
    "fmt"
    "time"

    "github.com/VirrageS/cache"
)

func API(c *cache.Cache) []string {
    cached_items, exists := c.Get("api")
    if exists {
        return cached_items
    }

    items := []string{"a", "lot", "of", "data"}
    c.Set("api", items)
    return items
}

func main() {
    cache := NewCache(time.Minute)
    items := API(cache)

    fmt.Printf("items: %v", items)
}
```
