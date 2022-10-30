# Public Shopify Monitor

Welcome to my public shopify monitor, a project which I used to sell until I decided to open source it. Although the code quality isn't the best, it's still a learning resource for beginners.

This was a project I made a while ago.

If you want to use it, just run main.go (or you can do `go build` in order to make an exe).

Put your proxies in `proxies.txt`. The shopify API bucket rate is around 1 request/4 seconds to `products.json`, keep in mind that the monitor will rotate proxy on every request. The more proxies, the better.
