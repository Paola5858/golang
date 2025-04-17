
# 💖 Aula de Operações Bancárias em Go 💖

```md

## 🌟 Código de exemplo

```go
package main

import "fmt"

func main() {
  var saldo float64
  var valor float64
  var escolha string

  fmt.Println("🎀 Oii, diva! Informe seu saldo inicial:")
  fmt.Scan(&saldo)

  fmt.Println("💰 Agora você pode escolher: 'sacar' ou 'depositar'")
  fmt.Scan(&escolha)

  if escolha == "sacar" {
    fmt.Println("👜 Quanto deseja sacar?")
    fmt.Scan(&valor)

    if valor <= saldo {
      saldo -= valor
      fmt.Println("Saque realizado com sucesso!")
    } else {
      fmt.Println("Ops! Você não tem saldo suficiente, princesa.")
    }
  } else if escolha == "depositar" {
    fmt.Println("💸 Quanto deseja depositar?")
    fmt.Scan(&valor)
    saldo += valor
    fmt.Println("✨ Depósito realizado com sucesso, gata!")
  } else {
    fmt.Println("Eita, esse comando é inválido, tente 'sacar' ou 'depositar'.")
  }

  fmt.Println("💳 Saldo atualizado é:", saldo)
  fmt.Println("💋 Arrasou, linda!")
}
```plaintext

---

## ✨ O que está rolando aqui?

Esse código simula um mini sistema bancário com toda a energia de diva 💅

1. A diva informa o **saldo inicial**.
2. Escolhe entre **sacar** ou **depositar**.
3. Dependendo da escolha, o saldo é atualizado ou ela recebe um aviso caso algo esteja errado.

### 🧠 Lógica por trás:

- Se escolher `sacar`, verifica se tem saldo suficiente.
- Se escolher `depositar`, só adiciona o valor ao saldo.
- Caso contrário, mensagem de comando inválido.

---

## 🖥️ Exemplo de entrada/saída

```plaintext
🎀 Oii, diva! Informe seu saldo inicial:
1000
💰 Agora você pode escolher: 'sacar' ou 'depositar'
sacar
👜 Quanto deseja sacar?
300
Saque realizado com sucesso!
💳 Saldo atualizado é: 700
💋 Arrasou, linda!
```plaintext
```plaintext

---

Feito com 💖 por Paola ✨
```
