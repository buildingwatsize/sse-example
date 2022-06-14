# SSE Example

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)

SSE Example is an example of SSE (Server Sent Events) in Golang (Fiber) with an easy client html

## Don't waste your time, clone and play it

### Step 1/3

```bash
  git clone https://github.com/buildingwatsize/sse-example.git
  cd sse-example
  go run main.go
```

### Step 2/3

```text
Open browser and go to `http://localhost:7777` and you will see a simple webpage
```

### Step 3/3

```text
Try to call api `http://localhost:7777/computed` which a sleep function for 5-second and then it will return some message to SSE Client (a webpage in step 2)

or just click "Send Request" via using REST Client Extension in VSCode (file location: `/tools/computed.http`)
```

> Made with ❤️ Powered by Watsize