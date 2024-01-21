package main

import (
	"fmt"
)

// Trong ví dụ này:

// BankAccount là một struct chứa thông tin về tài khoản, bao gồm owner (chủ sở hữu) và balance (số dư). Cả hai trường này đều được đặt là private,
//  tức là không thể truy cập trực tiếp từ bên ngoài.
// Các phương thức Deposit, Withdraw, và GetBalance cho phép tương tác với dữ liệu đóng gói trong BankAccount.
//  Chúng là cách duy nhất để thay đổi hoặc truy xuất thông tin từ BankAccount, đảm bảo rằng các quy tắc về việc xử lý tài khoản được tuân thủ
//  (ví dụ: không thể rút nhiều hơn số dư hiện có).
// Constructor NewBankAccount được sử dụng để tạo một đối tượng BankAccount mới, cho phép đặt giá trị ban đầu cho owner và balance.
// Đây là một ví dụ điển hình về cách đóng gói dữ liệu và hành vi liên quan đến dữ liệu đó trong một struct trong Golang.

// BankAccount là một struct đóng gói thông tin tài khoản
type BankAccount struct {
	owner   string // chủ sở hữu, không thể truy cập từ bên ngoài
	balance int    // số dư, không thể truy cập từ bên ngoài
}

// NewBankAccount là constructor tạo một BankAccount mới
func NewBankAccount(owner string, initialBalance int) *BankAccount {
	return &BankAccount{owner: owner, balance: initialBalance}
}

// Depossit cho phép gửi tiền vào tài khoản
func (ba *BankAccount) Deposit(amount int) {
	ba.balance += amount
}

// Withdraw cho phép rút tiền từ tài khoản
func (ba *BankAccount) Withdraw(amount int) {
	if amount <= ba.balance {
		ba.balance -= amount
	}
}

// GetBalance trả về số dư hiện tại của tài khoản
func (ba *BankAccount) GetBalance() int {
	return ba.balance
}

func main() {
	// tạo một tài khoản mới
	account := NewBankAccount("Thang Tran", 1000)

	// Gửi tiền
	account.Deposit(5000)
	fmt.Println("Balance after deposit: ", account.GetBalance())

	// Rút tiền
	account.Withdraw(200)
	fmt.Println("Balance after withdrawal: ", account.GetBalance())
}
