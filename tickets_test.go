package gogroovehq_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/roadmap-space/gogroovehq"
)

func Test_Tickets_Get(t *testing.T) {
	t.Parallel()

	gapi := gogroovehq.New(os.Getenv("GROOVE_TOKEN"))
	ticket, err := gapi.Tickets.Get("8")
	if err != nil {
		t.Fatal(err)
	} else if ticket == nil {
		t.Fatal(fmt.Errorf("no ticket returned"))
	}
	fmt.Println(ticket.Ticket.Links.Customer.Href)
}
