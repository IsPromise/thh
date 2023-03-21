package demo

import (
	"fmt"
	"thh/app/models/Users"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	appendCommand(&cobra.Command{Use: "demo:dbCreate", Short: "dbCreate", Run: dbCreate})
}
func dbCreate(_ *cobra.Command, _ []string) {
	user := Users.Users{Username: "nihao", Email: time.Now().String() + "@test.com"}
	user.SetPassword("niahao")
	err := Users.Create(&user)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("dbCreate")
}
