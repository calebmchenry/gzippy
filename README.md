# GZip Comparison of \~20MB and \~10MB payloadeds

## Get started
`go run main.go`

## Results
Request order
* /gzip/20MB
* /gzip/10MB
* /raw/20MB
* /raw/10MB

![localhost results](./localhost_20MB.png)

☝️ At such fast speeds the gzipping takes longer than the request does.

![4G results](./4G_20MB.png)

☝️ At slow speeds the gzipped variations are significantly faster.
