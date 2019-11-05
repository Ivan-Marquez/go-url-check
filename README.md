# go-url-check
> Minimal Golang CLI to check if url(s) are online

## Usage
```bash
make build # bin/go-url-check
bin/go-url-check --url https://google.com --time 10s # check google.com for 10 seconds
```

## Command flags
`--url`: specify url to check  
`--urls`: specify comma separated list of urls  
`--time`: check time period (defaults to 60 seconds)  
