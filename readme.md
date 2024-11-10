## Pagbank: O banco do futuro

<div align="center">
    <img src="/logo2.png" alt="EmployEd Logo">
</div>

---

# AIA: Assistente Financeiro

> \_TEAM 5: Hackathon Pagbank

![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen)
![Platform](https://img.shields.io/badge/Platform-Mobile-blue)
![License](https://img.shields.io/badge/License-MIT-green)

---

## 🌐 Introdução

A AIA é um assistente financeiro integrado ao aplicativo da PagBank, criado para simplificar a vida financeira dos usuários e auxiliar na organização de suas finanças de forma prática e inteligente. Através de um fluxo de uso intuitivo e integração com o WhatsApp, a AIA oferece uma série de funcionalidades que permitem ao usuário ter controle total sobre suas finanças.

### 📹 App Preview

[![Preview the App](assets/readme/prototype1.png)](https://www.youtube.com/)

---

## 🛠 Instalação (Golang)

1. _Pré-requisitos_

   - Certifique-se de ter o Golang instalado na sua máquina.

2. _Clonar o Repositório_

   bash
   git clone https://github.com/bellujrb/hackathon_bank

3. _Instalar Dependências_

   bash
   go mod tidy

4. _Executar a Aplicação_

   bash
   go run main.go

---

### 🔄 Fluxograma

1. _Início no Aplicativo PagBank_

   - O usuário abre o aplicativo da PagBank.
   - Encontra e clica no ícone da _AIA_ dentro do aplicativo.

2. _Aceite dos Termos e Personalização_

   - O usuário aceita os termos de uso da AIA.
   - Personaliza as configurações da AIA conforme suas preferências.

3. _Redirecionamento para o WhatsApp_

   - Após a personalização, o usuário é automaticamente direcionado para o WhatsApp.

4. _Início do Uso da AIA_
   - No WhatsApp, o usuário começa a interagir e usar a AIA, aproveitando suas funcionalidades para as necessidades desejadas.

---

## 📂 Arvore do Projeto

hackathon_pagbank
├── db
│ └── ...
├── events
│ └── invest.go
│ └── questions.go
│ └── recharge.go
│ └── slipers.go
│ └── weekly.go
├── mocks
│ └── ...
│ └── ...
├── examplestore.db
├── go.sum
└── main.go
├── README.md

---

#### back-end/mocks

- creating.data.go
  - Juntar os nomes e valores para gerar boletos e investimentos.
- names.go
  - Gerar nomes de lugares e investimento.
- value.go
  - Gerar valores de investimento, rendimentos deles, e o valor dos boletos.
- text.go
  - Todos os textos que são usadas na aplicação.

## 🛡 Segurança de Dados e Privacidade

- Todos os dados são rigorosamente processados ​​e anonimizados para garantir a privacidade.
- Medidas de segurança avançadas protegem contra acesso não autorizado.

---

## 🛠 Tech Stack (Back-end)

### Design Pattern

- Clean Code

### External Packages

- Whatsapp
- Web Socket

### Architecture

- Clear Archicheture

---

## 🌈 Roadmap Futuro

Planejamos expandir nosso software para ser multiplataforma, visando atingir um público mais amplo.

---

## 🙏 Agradecimentos

Um agradecimento especial ao Pagbank por oferecer essa oportunidade incrivel.

---
