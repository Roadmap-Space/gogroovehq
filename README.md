# Groove API Go library

This is a **partial** implementation of [Groove](https://groovehq.com) that we've built at [Roadmap](https://roadmap.space) 
when we implemented our Groove integration.

## Install

```shell
go get -u github.com/roadmap-space/gogroovehq
```

## Usage

```go
package main

import (
	"fmt"
	"log"
	"github.com/roadmap-space/gogroovehq
)

func main() {
	gapi := gogroovehq.New("YOUR_GROOVE_API_TOKEN")
	tiket, err := gapi.Tickets.Get("ticket_id")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ticket)
}
```

## A tiny implementation

We were only needing `Webhooks` create, `Tickets` get by id and `Customers` get by id.

But this is a good basis to continue implementing the rest of Groove's API.

## Unit tests

The unit tests uses hard-coded resources ID.