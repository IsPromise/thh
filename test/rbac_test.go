package test

import (
	"fmt"
	"testing"
)

func TestRbac(t *testing.T) {
	admin := map[string]any{
		"admin": []string{
			"admin1",
			"admin2",
		},
		"adminRole": []string{
			"id",
		},
		"adminPermission": []string{
			"id",
		},
		"adminRoleRelation": []any{
			struct {
				adminId int64
				roleId  int64
			}{},
		},
		"rolePermissionRelation": []any{
			struct {
				permission int64
				roleId     int64
			}{},
		},
	}

	fmt.Println(admin)
}
