# Project-go-with-gin

Projeto em Go para explorar o framework [Gin](https://github.com/gin-gonic/gin), aplicando boas práticas de design e desenvolvimento de APIs RESTful.

---

## 📌 Estrutura do Projeto

- **main.go** – ponto de entrada da aplicação.
- **router/** – configuração das rotas da API.
- **handler/** – implementação dos handlers das rotas (CRUD).
- **schemas/** – definição das estruturas de dados (DTOs, Models).
- **docs/** – documentação (Swagger/OpenAPI).
- **config/** – arquivos de configuração (ex.: conexão com banco).
- **Makefile** – comandos úteis para desenvolvimento.

---

## ⚡ Funcionalidades

A API implementa operações **CRUD** para gerenciar vagas (`openings`):

| Método | Endpoint                  | Descrição                     |
|--------|---------------------------|-------------------------------|
| POST   | `/api/v1/opening`         | Criar uma nova vaga           |
| GET    | `/api/v1/openings`        | Listar todas as vagas         |
| GET    | `/api/v1/opening?id={id}` | Exibir uma vaga específica    |
| PUT    | `/api/v1/opening?id={id}` | Atualizar uma vaga existente  |
| DELETE | `/api/v1/opening?id={id}` | Deletar uma vaga              |

---

## 🛠️ Tecnologias

- [Go](https://go.dev/) — linguagem principal  
- [Gin](https://github.com/gin-gonic/gin) — framework web leve e rápido  
- [GORM](https://gorm.io/) — ORM para Go  
- [Swagger](https://swagger.io/) — documentação da API via comentários  

---

## ▶️ Como Rodar o Projeto

1. **Clonar o repositório**
   ```bash
   git clone https://github.com/DaviRodrigues/Project-go-with-gin.git
   cd Project-go-with-gin
   ```

2. **Configurar variáveis de ambiente**
   - Defina a string de conexão com o banco de dados (ex.: PostgreSQL, SQLite, MySQL).  

3. **Executar a aplicação**
   ```bash
   go run main.go
   ```

4. **(Opcional) Usar Makefile**
   ```bash
   make run     # rodar aplicação
   make test    # rodar testes (se configurados)
   ```

5. **Acessar API**
   - Exemplo: [http://localhost:8080/api/v1/openings](http://localhost:8080/api/v1/openings)

6. **Acessar documentação Swagger** (se disponível)
   ```
   http://localhost:8080/swagger/index.html
   ```

---

## ✅ Exemplos de Uso

### Criar vaga
```bash
curl -X POST http://localhost:8080/api/v1/opening   -H "Content-Type: application/json"   -d '{
        "role": "Backend Developer",
        "company": "OpenAI",
        "location": "Remote",
        "remote": true,
        "link": "https://example.com/job",
        "salary": 8000
      }'
```

### Listar vagas
```bash
curl http://localhost:8080/api/v1/openings
```

---

## 🚀 Melhorias Futuras

- [ ] Usar rotas RESTful com `/:id` em vez de query params (`?id=`).  
- [ ] Implementar autenticação (ex.: JWT).  
- [ ] Adicionar validação com [`go-playground/validator`](https://github.com/go-playground/validator).  
- [ ] Criar testes unitários e de integração.  
- [ ] Adicionar migrações de banco (ex.: `golang-migrate`).  
- [ ] Configurar CI/CD (GitHub Actions, GitLab CI, etc.).  
