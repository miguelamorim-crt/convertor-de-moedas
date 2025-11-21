# Conversor de Moedas em Go

Este programa realiza conversão de valores de Dólar ou Euro para Real brasileiro usando taxas fixas.

## Como funciona

- Menu interativo com opções para converter dólar ou euro para real.
- Permite múltiplas conversões até o usuário escolher sair.
- Taxas de câmbio usadas:
  - 1 Dólar = 5.41 BRL
  - 1 Euro = 6.22 BRL

## Como usar

1. Execute o programa.
2. Quando perguntado se deseja calcular, responda com "S" para iniciar.
3. Escolha a moeda e informe a quantidade para converter.
4. Veja o resultado.
5. Para sair, selecione "0" no menu ou responda "N" inicialmente.

## Exemplo de execução
deseja calcular(S/N): S
=== ESCOLHA ===
1 - dolar
2 - euro
0 - sair
1
quantos dolares quer converter pra real:
10
lembre-se: 1 dolar custa 5.41
resultado: R$54.10
=== ESCOLHA ===
1 - dolar
2 - euro
0 - sair
0
saindo...


## Observações técnicas

- Utiliza `fmt` para entrada e saída.
- Taxas fixas no código (altere se necessário).
- Tratamento simples de opções inválidas.

