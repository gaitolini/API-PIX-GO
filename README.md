# API-PIX-GO 🚀
## Middleware API PIX in GO 🐿️

---

Automatic deploy test via GitHub Actions

---

## About the Project 🚀

**API-PIX-GO** is a **GoLang** application that implements an API for financial systems using charges via **PIX**. This project was developed specifically to interact with the **Banco do Brasil** (BB) system and operates both in **real** environments and the **BB Developers** **sandbox**.

The main focus of this API is to facilitate the generation and management of PIX charges for companies and systems that need to integrate this payment method into their financial automation processes.

## Technologies Used 🛠️

This project runs in a **Debian Linux** environment using the following technologies and tools:

- **GoLang**: The main language of the project, known for its efficiency and simplicity.
- **Docker**: Used to containerize the application, ensuring it runs consistently in any environment.
- **Cloudflare Tunnel**: Facilitates external access to the application securely without the need to open firewall ports.
- **Proxmox**: The project runs in a virtualized environment with **Proxmox**, providing flexibility and scalability.
- **Git**: Used for version control and continuous integration.
- **GitHub Actions**: We implemented a **CI/CD** pipeline to automate the build and deployment process.

## CI/CD Structure ⚙️

The **CI/CD** pipeline for this project is configured as follows:

1. **GitHub Actions**: When a change is made to the `main` branch, GitHub Actions automatically builds the application.
2. **Webhook Deploy**: After the build, a webhook triggered via GitHub sends a request to the **Debian** server, which runs the **deploy script**.
3. **Automatic Deploy with Docker**: The deploy script pulls the latest version of the application, builds the Docker image, and starts a new container on the server.
4. **Cloudflare Tunnel**: The **Cloudflare Tunnel** exposes the API securely, allowing the application to be accessed externally without the need to open additional ports.

### Webhook Deploy Repository 📦

You can check the repository responsible for the **automated deploy** at the following link:

🔗 [Deploy Webhook](https://github.com/gaitolini/deploy-web-hook)

## Execution Environment 🖥️

The application runs on a **VM** with **Debian Linux**, provisioned in a virtualized environment via **Proxmox**. The entire infrastructure is managed with modern tools, including:

- **Docker**: To manage the application's containers.
- **Cloudflare Tunnel**: To provide secure and efficient access to the API.

## About PIX and Banco do Brasil 💳🏦

This project is developed for integration with **PIX**, the instant payment system in Brazil. It implements all the functionalities for generating and managing charges using the **Banco do Brasil** API.

- The **real** environment allows for the issuance of PIX charges in production for actual use.
- The **sandbox** environment offers a controlled testing environment for development and simulation.

### Integration with BB Developers

The API interacts directly with **BB Developers**, a platform provided by Banco do Brasil to facilitate integration with its financial solutions. Through the **sandbox**, it is possible to test and simulate operations before putting them into production.

## About GoLang 🐹

**GoLang**, also known as **Go**, is a programming language developed by Google. It is ideal for developing high-performance applications, especially web services and distributed systems.

- **Simplicity**: Go is a minimalist language and easy to learn.
- **Performance**: Being a compiled language, Go offers excellent performance, comparable to languages like C and C++.
- **Concurrency**: Go has native support for concurrent programming, which is essential for scalable applications.

For more information about **GoLang**, visit the official website: [https://golang.org/](https://golang.org/)

## How to Run the Project Locally 🔧

Follow the instructions below to run **API-PIX-GO** locally using Docker:

1. Clone the repository:

   ```bash
   git clone https://github.com/gaitolini/API-PIX-GO.git
   cd API-PIX-GO
   ```
2. Build the Docker image:

 ```bash
docker build -t api-pix-go .

```
3. Run the container:

```bash
docker run -d --name api-pix-go -p 8080:8080 api-pix-go
```

4. Access the API:
```bash
http://localhost:8080

```

## Contact 📬
If you want to know more about this project or discuss future collaborations, feel free to contact me:

- LinkedIn: [Anderson Gaitolini](https://www.linkedin.com/in/andersongaitolini/)
- WhatsApp: [Entre em contato](https://youtu.be/IGP38bz-K48?si=62Khct2-dAFR3qn5)
 
## License 📜
This project is licensed under the terms of the MIT License. See the LICENSE file for more details.

### 🚀 Made with lots of Go, Docker, and PIX by Anderson Gaitolini.

### Thanks 🙌
A big thank you to everyone who contributed to the development of this project and helped improve automation and continuous deployment! 🎉

# Feel free to contribute, test, and suggest improvements! 😄
  
---
### Pt-BR
---
# API-PIX-GO 🚀
## Middleware API PIX em GO 🐿️

---

Teste de deploy automático via GitHub Actions

---

## Sobre o Projeto 🚀

**API-PIX-GO** é uma aplicação **GoLang** que implementa uma API para sistemas financeiros que utilizam cobranças via **PIX**. Este projeto foi desenvolvido especificamente para interagir com o sistema do **Banco do Brasil** (BB) e opera tanto em ambientes **reais** quanto no **sandbox** de desenvolvedores do **BB Developers**.

O foco principal desta API é facilitar a geração e o gerenciamento de cobranças via PIX para empresas e sistemas que precisam integrar esse método de pagamento em seus processos de automação financeira.

## Tecnologias Utilizadas 🛠️

Este projeto está rodando em um ambiente **Debian Linux** utilizando as seguintes tecnologias e ferramentas:

- **GoLang**: A linguagem principal do projeto, conhecida por sua eficiência e simplicidade.
- **Docker**: Utilizado para containerizar a aplicação, garantindo que ela rode de forma consistente em qualquer ambiente.
- **Cloudflare Tunnel**: Facilita o acesso externo à aplicação de maneira segura, sem a necessidade de abrir portas de firewall.
- **Proxmox**: O projeto está rodando em um ambiente virtualizado com **Proxmox**, permitindo flexibilidade e escalabilidade.
- **Git**: Usado para controle de versão e integração contínua.
- **GitHub Actions**: Implementamos um pipeline de **CI/CD** para automatizar o processo de build e deploy.

## Estrutura do CI/CD ⚙️

O pipeline de **CI/CD** para este projeto foi configurado da seguinte forma:

1. **GitHub Actions**: Quando uma alteração é feita na branch `main`, o GitHub Actions automaticamente faz o build da aplicação.
2. **Webhook Deploy**: Após o build, um webhook acionado via GitHub envia uma requisição para o servidor **Debian**, que executa o **script de deploy**.
3. **Deploy Automático com Docker**: O script de deploy puxa a última versão da aplicação, faz o build da imagem Docker e inicia um novo container no servidor.
4. **Cloudflare Tunnel**: O **túnel Cloudflare** expõe a API de forma segura, permitindo que a aplicação seja acessada externamente sem necessidade de abrir portas adicionais.

### Repositório do Webhook Deploy 📦

Você pode conferir o repositório responsável pelo **deploy automatizado** no seguinte link:

🔗 [Deploy Webhook](https://github.com/gaitolini/deploy-web-hook)

## Ambiente de Execução 🖥️

A aplicação está rodando em uma **VM** com **Debian Linux**, provisionada em um ambiente virtualizado via **Proxmox**. Toda a infraestrutura é gerenciada com ferramentas modernas, incluindo:

- **Docker**: Para gerenciar os containers da aplicação.
- **Cloudflare Tunnel**: Para fornecer acesso seguro e eficiente à API.

## Sobre o PIX e o Banco do Brasil 💳🏦

Este projeto é desenvolvido para integração com o **PIX**, o sistema de pagamentos instantâneos do Brasil. Ele implementa todas as funcionalidades para geração e gerenciamento de cobranças, utilizando a API do **Banco do Brasil**.

- O ambiente **real** permite a emissão de cobranças PIX em produção, para uso real.
- O ambiente **sandbox** oferece um ambiente de testes controlado para desenvolvimento e simulação.

### Integração com BB Developers

A API interage diretamente com o **BB Developers**, uma plataforma oferecida pelo Banco do Brasil para facilitar a integração com suas soluções financeiras. Através do **sandbox**, é possível testar e simular operações antes de colocá-las em produção.

## Sobre GoLang 🐹

**GoLang**, também conhecido como **Go**, é uma linguagem de programação desenvolvida pela Google. Ela é ideal para desenvolvimento de aplicações de alto desempenho, especialmente serviços web e sistemas distribuídos.

- **Simplicidade**: Go é uma linguagem minimalista e fácil de aprender.
- **Desempenho**: Por ser compilada, Go oferece um excelente desempenho, comparável a linguagens como C e C++.
- **Concorrência**: Go tem suporte nativo à programação concorrente, o que é essencial para aplicações escaláveis.

Para mais informações sobre **GoLang**, visite o site oficial: [https://golang.org/](https://golang.org/)

## Como Rodar o Projeto Localmente 🔧

Siga as instruções abaixo para rodar o **API-PIX-GO** localmente usando Docker:

1. Clone o repositório:

```bash
git clone https://github.com/gaitolini/API-PIX-GO.git
cd API-PIX-GO
```

### Build da imagem Docker:

~~~~bash
docker build -t api-pix-go .
Rode o container:
~~~~

~~~~bash
docker run -d --name api-pix-go -p 8080:8080 api-pix-go
Acesse a API via:
~~~~

~~~bash
http://localhost:8080
Contato 📬
~~~
---

## Se você quiser saber mais sobre este projeto ou discutir futuras colaborações, entre em contato comigo:

- LinkedIn: [Anderson Gaitolini](https://www.linkedin.com/in/andersongaitolini/)
- WhatsApp: [Entre em contato](https://youtu.be/IGP38bz-K48?si=62Khct2-dAFR3qn5)

### Licença 📜
Este projeto é licenciado sob os termos da MIT License. Veja o arquivo LICENSE para mais detalhes.

🚀 Feito com muito Go, Docker, e PIX por Anderson Gaitolini.

## Agradecimentos 🙌
Um grande obrigado a todos que contribuíram para o desenvolvimento deste projeto e ajudaram a aprimorar a automação e o deploy contínuo! 🎉

## Fique à vontade para contribuir, testar e sugerir melhorias! 😄

---

