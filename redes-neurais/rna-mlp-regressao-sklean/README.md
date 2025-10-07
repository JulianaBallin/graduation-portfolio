# 🧠 Redes Neurais Artificiais - MLP para Previsão de Consumo de Combustível

Este repositório contém a atividade desenvolvida por **Juliana Ballin Lima** para a disciplina **Redes Neurais Artificiais** (UEA – Escola Superior de Tecnologia), utilizando um modelo de **Perceptron Multicamadas (MLP)** com a biblioteca **Scikit-Learn** para prever o consumo de combustível de veículos a partir do conjunto de dados **Auto MPG**.

---

## 📁 Estrutura do Repositório

```
📦 redes-neurais-mlp-autompg
├─ RNA-2025.2-MLP-SkLearn.ipynb # Notebook principal com implementação do modelo MLP
├─ autompg.csv # Base de dados utilizada no treinamento e teste
├─ requirements.txt # Lista de dependências necessárias ao projeto
└─ README.md # Documentação do projeto
```

---

## 📚 Descrição do Projeto

O objetivo desta atividade é aplicar os conceitos de **Redes Neurais Artificiais (RNA)** por meio de um modelo do tipo **MLP (Multilayer Perceptron)** para resolver um problema de **regressão**: prever o consumo de combustível (mpg – *miles per gallon*) com base em características de veículos, como número de cilindros, peso, aceleração, entre outras.

---

## 🧪 Tecnologias Utilizadas

- 🐍 **Python 3.10+**
- 📊 **Pandas** – manipulação e análise de dados
- 📈 **Matplotlib / Seaborn** – visualização de dados
- 🧠 **Scikit-Learn** – criação e treinamento do modelo MLP
- 📓 **Jupyter Notebook** – ambiente interativo para desenvolvimento

---

## ⚙️ Como Utilizar

Para executar este projeto em sua máquina local, siga os passos abaixo:

### 1️⃣ Clonar o repositório
```bash
git clone https://github.com/seu-usuario/redes-neurais-mlp-autompg.git
cd redes-neurais-mlp-autompg
```

### 2️⃣ Criar e ativar um ambiente virtual (venv)

**Linux/macOS:**
```bash
python3 -m venv venv
source venv/bin/activate
```
**Windows (PowerShell):**
```bash
python -m venv venv
.\venv\Scripts\activate
```

### 3️⃣ Instalar as dependências

```bash
pip install -r requirements.txt
```

### 4️⃣ Executar o Jupyter Notebook

```bash
jupyter notebook
```

---
## 📊 Base de Dados

O dataset utilizado é o **Auto MPG**, um conjunto de dados clássico disponível no repositório da UCI Machine Learning.  
Cada linha representa um veículo e inclui atributos como:

- `mpg` – milhas por galão (variável alvo)
- `cylinders` – número de cilindros
- `displacement` – deslocamento do motor
- `horsepower` – potência
- `weight` – peso
- `acceleration` – aceleração
- `model year` – ano do modelo
- `origin` – origem do veículo

---

## 🧠 Modelo Utilizado – MLPRegressor

O modelo escolhido foi o **MLPRegressor**, um Perceptron Multicamadas aplicado a problemas de regressão.  
Durante a atividade foram realizados os seguintes passos:

1. 📥 **Carregamento e exploração dos dados**
2. 🧹 **Pré-processamento** – limpeza de dados faltantes e normalização
3. ✂️ **Divisão em treino e teste**
4. 🧠 **Criação e treinamento do modelo MLP**
5. 📊 **Avaliação de desempenho** – análise de métricas como R² e erro médio

---

## 📈 Resultados Esperados

O modelo treinado é capaz de prever com boa precisão o consumo de combustível a partir dos atributos fornecidos.  
Os gráficos de erro e a comparação entre valores reais e preditos são apresentados no notebook principal.

---

## 👩‍🎓 Informações da Autora

- **Nome:** Juliana Ballin Lima  
- **Matrícula:** 2315310011  
- **E-mail:** [jbl.snf23@uea.edu.br](mailto:jbl.snf23@uea.edu.br)  
- **Disciplina:** Redes Neurais Artificiais – 2025.2  
- **Instituição:** Universidade do Estado do Amazonas – Escola Superior de Tecnologia (UEA/EST)

---

## 📜 Licença

Este projeto foi desenvolvido apenas para fins acadêmicos no contexto da disciplina **Redes Neurais Artificiais** e não possui fins comerciais.

---

✍️ *Desenvolvido por Juliana Ballin Lima – 2025.2*
