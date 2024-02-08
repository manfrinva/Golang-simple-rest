# Golang-simple-rest
RESTful básica em Go (Golang) 

Para rodar você precisará instalar o pacote gorilla/mux para lidar com as rotas. Você pode instalar utilizando o seguinte comando:
go get -u github.com/gorilla/mux

Este código cria um servidor RESTful com operações CRUD para contatos. Ele também implementa uma autenticação básica com usuário e senha.

Note que esta é apenas uma implementação básica e não deve ser usada em produção sem considerar questões de segurança, como armazenamento seguro de senhas e proteção contra ataques de injeção de código. Além disso, o armazenamento de contatos é mantido na memória e será reiniciado sempre que o servidor for reiniciado. Em uma aplicação real, você usaria um banco de dados para armazenamento persistente.
