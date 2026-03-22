# API Simples

_API_ Simples para demostrar conhecimento na construção, manutenção de uma _API REST_ na linguagem go.

## Sobre

<div align="center">
    <img src="https://img.shields.io/static/v1?label=Golang&message=1.25.7&color=7159c1&style=for-the-badge&logo=golang"/>
    <img src="https://img.shields.io/static/v1?label=Versao&message=0.0.0&color=7159c1&style=for-the-badge&logo=circle"/>
    <img src="https://img.shields.io/static/v1?label=Status&message=Desenvolvimento&color=7159c1&style=for-the-badge&logo=circle"/>
    <img src="https://img.shields.io/static/v1?label=Licenca&message=Proprietaria&color=7159c1&style=for-the-badge&logo=circle"/>
</div>

## Iniciando

- Tenha o golang instalado, verificar com o comando abaixo(recomendado versão 1.25):
```shell
    go version
```
- Clone o projeto.
```shell
    git clone https://github.com/BrunoPessoa097/api-go-simples.git
```
- Entre na Pasta.
```shell
    cd api-go-simples
```

- Dentro da pasta iniciar esse comando:
```shell
    go mod tidy
```

## Métodos para executar
Segue abaixo 3 métodos para que rode a _API_, usando o `air`, `go run .` e o `docker`. segue abaixo como proceder em cada situação:

### Air
- Configure o air manualmente, ou usando o seguinte abaixo:
```shell
    air init
```
- Execute o comando abaixo:
```shell
    air
```
### GO RUN 
- Entre na pasta onde tem a main principal, `main.go`:
```shell
    cd cmd/api
```
- Execute o comando abaixo:
```shell
    go run .
```
### DOCKER
- Verificar se tem o docker instalado:
```shell
    docker -V
```
- Execute o comando abaixo para criar a imagem(escolha o nome):
```shell
    docker build -t <nome> .
```
- Execute a imagem:
```shell
    docker container run -d -p 8080:8080 <nome>
```
## Features

- [x] CRUD de Usuário
- [ ] CRUD de Post
- [ ] CRUD de Nivel
- [ ] Segurança
- [ ] Autentificação JWT
- [ ] Helmet
- [ ] CORs
- [ ] DotEnv

## Criado Por
- **Nome**: Bruno Pessoa
- **Área**: Desenvolver NodeJs|Typescript|Javascript|Go(Golang)
- **Formado**: UNIGRANDE - Centro Universitário da grande Fortaleza.
- **Curso**: Sistemas para _Internet_.
- **Git Hub**: [github.com/BrunoPessoa097](https://github.com/BrunoPessoa097/api-agenda.git)
- **LinkedIn**: [www.linkedin.com/in/bruno-pesoa-097](https://www.linkedin.com/in/bruno-pessoa-097/)
- **Portifólio**: [https://bruno-portifolio-eight.vercel.app/](https://bruno-portifolio-eight.vercel.app/)

## _License_
Esse projeto esta sobre a licença `Proprietária` ©Bruno Pessoa - 2026.
