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

