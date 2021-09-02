# PublicShopifyMonitor

Welcome to my public shopify monitor!

This was a project I made a while ago.

If you want to use it, just run main.go (or you can do `go build` in order to make an exe).

Put your proxies in `proxies.txt`. The shopify API bucket rate is around 1 request/4 seconds to `products.json`, keep in mind that the monitor will rotate proxy on every request. The more proxies, the better.
