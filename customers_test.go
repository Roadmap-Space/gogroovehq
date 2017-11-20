package gogroovehq_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/roadmap-space/gogroovehq"
)

func Test_Customers_Get(t *testing.T) {
	t.Parallel()

	gapi := gogroovehq.New(os.Getenv("GROOVE_TOKEN"))
	cust, err := gapi.Customers.Get("support@roadmap.space")
	if err != nil {
		t.Fatal(err)
	} else if cust == nil {
		t.Fatal(fmt.Errorf("no customer returned"))
	}
	fmt.Println(cust.Customer.Email, cust.Customer.Name)
}
