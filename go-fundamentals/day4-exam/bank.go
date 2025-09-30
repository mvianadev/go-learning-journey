package main

import (
	"fmt"
)

/*
Tu misión: Crear un sistema bancario completo con validaciones.
Requisitos:
2. Methods a implementar:

3. En main(), simular:
go- Crear 2 cuentas
- Depositar en ambas
- Intentar retiro válido
- Intentar retiro inválido (sin fondos)
- Transferir entre cuentas
- Mostrar balances finales
*/

/*
Tu misión: Crear un sistema bancario completo con validaciones.
Requisitos:
1. Struct BankAccount: ✅
gotype BankAccount struct {
    AccountNumber string
    Owner         string
    Balance       float64
}*/

type BankAccount struct {
	AccountNumber string
	Owner         string
	Balance       float64
}

/*
Deposit(amount float64) error

Si amount <= 0, retornar error "monto inválido"
Agregar al balance
Usar pointer receiver
*/

func (b *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("monto inválido")
	}
	b.Balance += amount
	return nil
}

/*
Withdraw(amount float64) error

Si amount <= 0, retornar error "monto inválido"
Si amount > balance, retornar error "fondos insuficientes"
Restar del balance
Usar pointer receiver
*/

func (b *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("monto inválido")
	}
	if amount > b.Balance {
		return fmt.Errorf("fondos insuficientes")
	}

	b.Balance -= amount

	return nil
}

/*
GetBalance() float64

Retornar balance actual
Value receiver está bien
*/

func (b *BankAccount) GetBalance() float64 {
	return b.Balance
}

/*
Transfer(target *BankAccount, amount float64) error

Usar Withdraw y Deposit
Si Withdraw falla, no hacer nada
*/

func (b *BankAccount) Transfer(target *BankAccount, amount float64) error {
	err := b.Withdraw(amount)
	if err != nil {
		return err // Si falla withdraw, no hacer nada
	}
	return target.Deposit(amount) // Depositar en TARGET
}

/*
3. En main(), simular:
go- Crear 2 cuentas
- Depositar en ambas
- Intentar retiro válido
- Intentar retiro inválido (sin fondos)
- Transferir entre cuentas
- Mostrar balances finales
*/

func main() {

	var err error

	banUno := &BankAccount{AccountNumber: "1111", Owner: "Mauricio", Balance: 65.99}
	banDos := &BankAccount{AccountNumber: "2222", Owner: "Andres", Balance: 125.32}

	fmt.Printf("Balance: %.2f, numero de cuenta: %s\n", banUno.Balance, banUno.AccountNumber)
	fmt.Printf("Balance: %.2f, numero de cuenta: %s\n", banDos.Balance, banDos.AccountNumber)

	err = banUno.Deposit(10.99)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("Depositando: 15$ en la cuenta: %s\n", banUno.AccountNumber)
		fmt.Printf("Balance: %.2f \n", banUno.GetBalance())
	}

	err = banDos.Deposit(25)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("Depositando 25$ en la cuenta: %s\n", banDos.AccountNumber)
		fmt.Printf("Balance: %.2f \n", banDos.GetBalance())
	}

	fmt.Printf("Balance: %.2f, numero de cuenta: %s\n", banUno.Balance, banUno.AccountNumber)
	fmt.Printf("Balance: %.2f, numero de cuenta: %s\n", banDos.Balance, banDos.AccountNumber)

	err = banUno.Withdraw(15)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("Retirando 15$ de la cuenta: %s\n", banUno.AccountNumber)
		fmt.Printf("Balance: %.2f \n", banUno.GetBalance())
	}

	err = banDos.Withdraw(25)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("Retirando 25$ de la cuenta: %s\n", banDos.AccountNumber)
		fmt.Printf("Balance: %.2f \n", banDos.GetBalance())
	}

	err = banUno.Withdraw(500)
	if err != nil {
		fmt.Printf("Error: %v, en la cuenta: %s\n", err, banUno.AccountNumber)
	} else {
		fmt.Printf("Retirando 500$ de la cuenta: %s\n", banUno.AccountNumber)
		fmt.Printf("Balance: %.2f \n", banUno.GetBalance())
	}

	err = banDos.Withdraw(2000)
	if err != nil {
		fmt.Printf("Error: %v, en la cuenta: %s\n", err, banDos.AccountNumber)
	} else {
		fmt.Printf("Retirando 2000$ de la cuenta: %s\n", banDos.AccountNumber)
		fmt.Printf("Balance: %.2f \n", banDos.GetBalance())
	}

	err = banUno.Transfer(banDos, 15)
	if err != nil {
		fmt.Printf("Error: %v, en la cuenta: %s\n", err, banUno.AccountNumber)
	} else {
		fmt.Printf("transferiendo 15$ a la cuenta: %s\n", banDos.AccountNumber)
		fmt.Printf("Balance: %.2f \n", banUno.GetBalance())
	}

	err = banDos.Transfer(banUno, 5)
	if err != nil {
		fmt.Printf("Error: %v, en la cuenta: %s\n", err, banDos.AccountNumber)
	} else {
		fmt.Printf("transferiendo 5$ a la cuenta: %s\n", banUno.AccountNumber)
		fmt.Printf("Balance: %.2f \n", banDos.GetBalance())
	}

	fmt.Printf("Balance: %.2f, numero de cuenta: %s\n", banUno.Balance, banUno.AccountNumber)
	fmt.Printf("Balance: %.2f, numero de cuenta: %s\n", banDos.Balance, banDos.AccountNumber)
}
