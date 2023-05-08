package algorithm

import (
	"fmt"
	"testing"
)

func TestEncryptPassword(t *testing.T) {
	password := "testpassword"

	encodedHash, encodedSalt, err := EncryptPassword(password)
	if err != nil {
		t.Errorf("Error encrypting password: %v", err)
	}

	// 检查返回值是否有效
	if encodedHash == "" || encodedSalt == "" {
		t.Error("Empty return value")
	}
	fmt.Println(encodedHash, encodedSalt)
}

func TestVerifyPassword(t *testing.T) {
	password := "testpassword"

	encodedHash, encodedSalt, err := EncryptPassword(password)
	if err != nil {
		t.Errorf("Error encrypting password: %v", err)
	}

	// 验证正确的密码
	err = VerifyPassword(encodedHash, encodedSalt, password)
	if err != nil {
		t.Errorf("Incorrect password: %v", err)
	}

	// 验证错误的密码
	err = VerifyPassword(encodedHash, encodedSalt, "wrongpassword")
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func BenchmarkEncryptPassword(b *testing.B) {
	password := "testpassword"

	for i := 0; i < b.N; i++ {
		_, _, _ = EncryptPassword(password)
	}
}

func BenchmarkVerifyPassword(b *testing.B) {
	password := "testpassword"
	encodedHash, encodedSalt, _ := EncryptPassword(password)

	for i := 0; i < b.N; i++ {
		_ = VerifyPassword(encodedHash, encodedSalt, password)
	}
}

func BenchmarkEncryptPasswordOld(b *testing.B) {
	password := "testpassword"

	for i := 0; i < b.N; i++ {
		_, _, _ = EncryptPassword(password)
	}
}

func TestPassword(t *testing.T) {
	tmpPassWord := "asdasdas"
	password, err := MakePassword(tmpPassWord)
	if err != nil {
		t.Error("密码生成异常", err)
	}
	err = VerifyEncryptPassword(password, tmpPassWord)
	if err != nil {
		fmt.Println(err)
	}
	err = VerifyEncryptPassword(password, "tmpPassWord")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(password)
	// mw/jF8gnoHyxo/a9da9gewoNJBWKe/BUO//HANYqkkk=:+owmh2syRgHrpUngFnzTOQmP+UtAA+aZ5j0O1WyDO6Q=
	// IoYUxGCZ29Cm0QM1qFSCZj3grzaIsdD+TXRmOBQBTgQ=:qGu3D5bjoLWAk2RmhiFWRFjfikc/p6cqET85+uVhYlA=
}
