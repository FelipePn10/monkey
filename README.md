# 🐵 Monkey Programming Language & Interpreter

Implementação de uma **linguagem de programação** e seu **interpretador**, desenvolvida em Go do absoluto zero sem dependencias, como parte do estudo prático do livro [*Writing an Interpreter in Go*](https://interpreterbook.com/), de Thorsten Ball.

## 📘 Sobre o projeto

Neste projeto desenvolvi do zero uma linguagem de programação chamada Monkey — incluindo sua sintaxe, semântica e estruturas de dados — e implementei junto a ela um interpretador completo para executá-la. Tudo é construído passo a passo, começando pela análise léxica e indo até a avaliação de expressões complexas.

Ao longo deste código:

- **Criei a linguagem Monkey**, definindo seus tokens, palavras-chave, sintaxe, estruturas de controle e sistema de tipos
- **Implementei o interpretador**, capaz de ler código-fonte Monkey e executá-lo
- Construi um **REPL** interativo para testar a linguagem em tempo real

## 🧠 Características da linguagem Monkey

Monkey é uma linguagem de programação simples mas poderosa, com:

- **Sintaxe inspirada em C/JavaScript**
- **Tipos de dados nativos:** inteiros, booleanos, strings, arrays e hashes (dicionários)
- **Expressões:** operadores aritméticos (`+`, `-`, `*`, `/`), relacionais (`>`, `<`, `==`, `!=`) e lógicos
- **Variáveis** declaradas com `let`
- **Funções de primeira classe** (`fn`), incluindo suporte a **closures**
- **Estruturas de controle:** `if` / `else`
- **Retorno explícito** com `return`
- **Funções embutidas (built-in):** `len`, `first`, `last`, `rest`, `push`, `puts`, entre outras

### Exemplo de código na linguagem Monkey

```js
// Mapeamento de números ao quadrado
let map = fn(arr, f) {
  let iter = fn(arr, accumulated) {
    if (len(arr) == 0) {
      accumulated
    } else {
      iter(rest(arr), push(accumulated, f(first(arr))));
    }
  };
  iter(arr, []);
};

let numeros = [1, 2, 3, 4, 5];
let quadrados = map(numeros, fn(x) { x * x });

puts(quadrados); // [1, 4, 9, 16, 25]

"Criar uma linguagem de programação é a melhor forma de entender como elas realmente funcionam." 🛠️🐵
