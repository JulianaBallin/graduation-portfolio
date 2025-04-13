# Projeto Sistemas Paralelos e Distribuídos

## Equipe GoSync

**Integrantes:**
- Ana Beatriz Maciel Nunes  
- Fernando Luiz Da Silva Freire  
- Juliana Ballin Lima  
- Lucas Carvalho Dos Santos  
- Marcelo Heitor De Almeida Lira  

## 📌 Descrição do Projeto
Este projeto implementa o problema clássico **Produtor x Consumidor** utilizando a linguagem **Go**.  
A solução utiliza **goroutines** e **canais** para gerenciar a concorrência e simular um **buffer com tamanho limitado**.  
O número de produtores, consumidores e o tamanho do buffer podem ser facilmente configurados no código.

---

## ▶️ Como compilar e executar o código

### ✅ Opção 1: Executar localmente (recomendado)
Certifique-se de ter o Go instalado. Para executar o projeto, use o comando:

```bash
go run main.go
```

### 🌐 Opção 2: Copiar e colar em um compilador online
Você pode colar o código no [Go Playground](https://go.dev/play/) ou em outro compilador online compatível com Go.

---

## ⚙️ Parâmetros de teste sugeridos
- Tamanho do buffer: `3`
- Número de produtores: `3`
- Número de consumidores: `2`

---

## 🧪 Casos de teste

### 🟢 Caso básico
- **Buffer**: 3  
- **Produtores**: 3  
- **Consumidores**: 2  
- Cada produtor produz 3 itens; cada consumidor consome 4 itens.  
- Espera-se que todos os itens produzidos sejam consumidos com controle de concorrência.

### 🟡 Buffer maior
- **Buffer**: 5  
- **Produtores**: 2  
- **Consumidores**: 3  
- Permite maior acúmulo de itens no buffer antes de serem consumidos.

### 🔴 Muitos produtores
- **Buffer**: 2  
- **Produtores**: 5  
- **Consumidores**: 2  
- Testa sobrecarga no buffer e bloqueio dos produtores enquanto o buffer estiver cheio.

### 🔵 Muitos consumidores
- **Buffer**: 2  
- **Produtores**: 2  
- **Consumidores**: 5  
- Testa consumo mais rápido que a produção, provocando espera nos consumidores.

---

## 📤 Saída esperada

A saída esperada mostra o fluxo de produção e consumo dos itens, por exemplo:

```
Produtor 1 produziu item 1  
Produtor 2 produziu item 2  
Consumidor 1 consumiu item 1  
Buffer atual: 1 itens  
...
```

---
## ▶️ Link Video Explicação

https://drive.google.com/file/d/1U3_LlKzsqyL_JWzxOA9EdcvKkrFvOMcg/view?usp=sharing
