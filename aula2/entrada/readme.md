# 💖 Aula de Entrada de Dados em Go 💖

## 🌟 Código de exemplo

```go
package main

import "fmt"

func main() {
  var numero int
  fmt.Println("Olá! Digite um número:")
  if _, err := fmt.Scan(&numero); err != nil {
    fmt.Println("Erro ao ler o número:", err)
    return
  }
  fmt.Printf("O número digitado foi: %d\n", numero)
}
```plaintext

---

## ✨ O que está rolando aqui?

Esse código é um exemplo básico e poderoso de **entrada de dados** com tratamento de erro 🧠

1. A diva digita um número no terminal.
2. O programa tenta ler o número com `fmt.Scan`.
3. Se ocorrer erro, ele avisa com classe e encerra o programa.

### 💡 Detalhes técnicos

- `fmt.Scan(&numero)` faz a leitura do que o usuário digitar.
- `_, err :=` verifica se houve erro durante a leitura.
- Se o input for válido, imprime o número digitado.

---

## 🖥️ Exemplo de entrada/saída:

```plaintext
``
Olá! Digite um número:
123
O número digitado foi: 123
``

---

Feito com 💖 por Paola ✨

