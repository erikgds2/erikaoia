# erikao :: ai terminal

> Interface de terminal local para rodar modelos de IA via Ollama — 100% no seu PC, zero nuvem, zero custo por token.

[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://go.dev)
[![Ollama](https://img.shields.io/badge/Ollama-local-black?style=flat)](https://ollama.com)
[![License](https://img.shields.io/badge/license-MIT-grey?style=flat)](LICENSE)

---

## O que é

Um servidor HTTP escrito em **Go** que:

- Serve um frontend estilo terminal (inspirado em CLI hacker)
- Conecta ao **Ollama** rodando localmente para processar mensagens
- Registra **todas as perguntas e respostas** em `queries.json`
- Expõe um **painel admin** em `/admin.html` para visualizar o histórico
- Suporta **troca de modelo** (llama3.2, mistral, codellama, etc.)
- **Mobile-first**: sidebar retrátil, input otimizado para touch

---

## Stack

| Camada    | Tecnologia                          |
|-----------|-------------------------------------|
| Backend   | Go 1.22+ (stdlib pura, zero deps)   |
| IA local  | Ollama (`/api/generate`)            |
| Frontend  | HTML + CSS + JS vanilla             |
| Persistência | JSON file (`queries.json`)       |
| Fontes    | Share Tech Mono (Google Fonts)      |

---

## Pré-requisitos

1. [Go 1.22+](https://go.dev/dl/)
2. [Ollama](https://ollama.com/download) instalado e rodando
3. Pelo menos um modelo baixado:

```bash
ollama pull llama3.2
# ou
ollama pull mistral
# ou
ollama pull codellama
```

---

## Instalação e uso

```bash
# clone o repositório
git clone https://github.com/erikgds2/erikaoia.git
cd erikaoia

# inicie o Ollama (em outro terminal)
ollama serve

# rode o servidor Go
go run main.go

# acesse
# Terminal: http://localhost:8080
# Admin:    http://localhost:8080/admin.html
```

---

## Variáveis de ambiente

| Variável       | Padrão        | Descrição                          |
|----------------|---------------|------------------------------------|
| `PORT`         | `8080`        | Porta do servidor HTTP             |
| `OLLAMA_MODEL` | `llama3.2`    | Modelo padrão do Ollama            |

Exemplo:
```bash
PORT=3000 OLLAMA_MODEL=mistral go run main.go
```

---

## Endpoints da API

| Método | Endpoint      | Descrição                                      |
|--------|---------------|------------------------------------------------|
| POST   | `/api/chat`   | Envia mensagem para o Ollama e salva o log     |
| GET    | `/api/logs`   | Retorna todos os logs em JSON                  |
| GET    | `/api/models` | Lista modelos disponíveis no Ollama            |
| GET    | `/api/status` | Status do servidor e do Ollama                 |

### POST /api/chat

**Request:**
```json
{
  "message": "explique recursão em Go",
  "model": "llama3.2"
}
```

**Response:**
```json
{
  "result": "Recursão é quando uma função chama a si mesma...",
  "model": "llama3.2",
  "duration_ms": 2341
}
```

---

## Estrutura do projeto

```
erikaoia/
├── main.go              # Servidor Go (HTTP + Ollama proxy + logs)
├── go.mod               # Módulo Go
├── queries.json         # Log de queries (gerado em runtime, gitignored)
├── frontend/
│   ├── index.html       # Terminal UI (mobile-first)
│   └── admin.html       # Painel admin de logs
├── .claude/
│   └── settings.json    # Permissões Claude Code (sem prompts de aprovação)
├── .gitignore
├── CRONOGRAMA.md        # Roadmap do projeto
└── README.md
```

---

## Painel Admin

Acesse `http://localhost:8080/admin.html` para ver:

- **Estatísticas**: total de queries, queries hoje, tempo médio de resposta
- **Status do Ollama**: online/offline em tempo real
- **Tabela de logs**: ID, timestamp, duração, modelo, pergunta, resposta
- **Filtros**: busca por texto, filtro por modelo
- **Exportar**: download do histórico em JSON
- **Detalhe**: click em qualquer linha para ver pergunta/resposta completa

---

## Como mudar o modelo em runtime

No terminal web, use o seletor de modelo na barra de input (desktop) ou o botão `◈` na topbar (mobile). Os modelos disponíveis são carregados automaticamente do Ollama.

---

## Build para produção

```bash
# Linux/Mac
go build -o erikaoia main.go
./erikaoia

# Windows
go build -o erikaoia.exe main.go
.\erikaoia.exe

# Cross-compile (Linux → Windows)
GOOS=windows GOARCH=amd64 go build -o erikaoia.exe main.go
```

---

## Contribuindo

1. Fork o repositório
2. Crie sua branch: `git checkout -b feat/minha-feature`
3. Commit: `git commit -m "feat: adiciona X"`
4. Push: `git push origin feat/minha-feature`
5. Abra um Pull Request

---

## Licença

MIT — veja [LICENSE](LICENSE).

---

> Feito pelo erikao às 4am com cafeína e desespero.
