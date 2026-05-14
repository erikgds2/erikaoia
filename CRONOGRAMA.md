# Cronograma — erikao :: ai terminal

> Roadmap de desenvolvimento por sprints semanais.

---

## Sprint 0 — Fundação ✅ (concluído em 2026-05-14)

**Objetivo:** Migrar de Cloudflare Worker + Netlify para stack local com Go + Ollama.

| Tarefa | Status |
|--------|--------|
| Baixar source do site Netlify | ✅ |
| Criar repositório público no GitHub | ✅ |
| Implementar servidor HTTP em Go (stdlib pura) | ✅ |
| Proxy para Ollama (`/api/generate`) | ✅ |
| Logging de queries em JSON | ✅ |
| API endpoints: `/api/chat`, `/api/logs`, `/api/models`, `/api/status` | ✅ |
| Frontend mobile-first (sidebar drawer) | ✅ |
| Seletor de modelo em runtime | ✅ |
| Painel admin com filtros e exportação | ✅ |
| README detalhado | ✅ |
| `.claude/settings.json` sem prompts de aprovação | ✅ |

---

## Sprint 1 — UX & Qualidade (2026-05-15 → 2026-05-21)

**Objetivo:** Polimento da experiência, acessibilidade, testes.

| Tarefa | Status |
|--------|--------|
| Histórico de conversa (contexto multi-turn no Ollama) | ⬜ |
| Atalhos de teclado: `↑↓` para histórico de comandos | ⬜ |
| Comando `/clear` para limpar o terminal | ⬜ |
| Comando `/model <nome>` para trocar modelo inline | ⬜ |
| Comando `/help` para listar comandos disponíveis | ⬜ |
| Suporte a Markdown na resposta (negrito, código, listas) | ⬜ |
| Scroll to bottom automático e suave | ⬜ |
| PWA: `manifest.json` + service worker (offline shell) | ⬜ |
| Meta tags Open Graph para compartilhamento | ⬜ |
| Testes básicos no backend Go (`go test`) | ⬜ |

---

## Sprint 2 — Features Avançadas (2026-05-22 → 2026-05-28)

**Objetivo:** Tornar o terminal realmente útil para o dia a dia.

| Tarefa | Status |
|--------|--------|
| Streaming de resposta (SSE / chunked) para feedback em tempo real | ⬜ |
| Sistema de system-prompt customizável via `/persona <texto>` | ⬜ |
| Salvar conversas como arquivos `.txt` | ⬜ |
| Modo "foco" (esconde sidebar, maximiza log) | ⬜ |
| Histórico de sessão persistido em `localStorage` | ⬜ |
| Admin: gráfico de queries por dia (canvas puro) | ⬜ |
| Admin: destacar queries mais longas / mais lentas | ⬜ |
| Admin: busca por intervalo de datas | ⬜ |
| Suporte a múltiplas sessões simultâneas (goroutines seguras) | ⬜ |

---

## Sprint 3 — Integração & Automação (2026-05-29 → 2026-06-04)

**Objetivo:** Conectar a IA com o ambiente local.

| Tarefa | Status |
|--------|--------|
| Comando `/exec <shell>` para rodar comandos locais (opt-in, com aviso) | ⬜ |
| RAG simples: indexar arquivos de texto locais como contexto | ⬜ |
| Integração com clipboard (`/copy` para copiar última resposta) | ⬜ |
| Webhook de notificação quando query termina (para queries longas) | ⬜ |
| Suporte a imagens no prompt (Ollama vision models) | ⬜ |
| Hot-reload do servidor Go em dev (via `go generate` ou Air) | ⬜ |
| Makefile com targets: `make run`, `make build`, `make test` | ⬜ |

---

## Sprint 4 — Deploy & Distribuição (2026-06-05 → 2026-06-11)

**Objetivo:** Facilitar instalação e uso por outras pessoas.

| Tarefa | Status |
|--------|--------|
| Dockerfile para rodar em container | ⬜ |
| Script de instalação one-liner (`curl \| bash`) | ⬜ |
| GitHub Actions: build + release automático de binários | ⬜ |
| Release binários para Windows, Mac, Linux no GitHub Releases | ⬜ |
| Página de landing no Netlify (marketing do projeto) | ⬜ |
| Video demo (30s) para o README | ⬜ |

---

## Backlog (sem data definida)

- [ ] Temas visuais: amber, green phosphor, blue-white, dracula
- [ ] Modo multiplayer (shared terminal via WebSocket)
- [ ] Plugin system para comandos customizados
- [ ] Integração com n8n para automações via IA
- [ ] Suporte a OpenAI/Anthropic como fallback (quando Ollama offline)
- [ ] Dashboard de uso de GPU/CPU durante inferência
- [ ] Autenticação básica (senha) para expor o server na rede local

---

## Convenção de commits

```
feat:    nova funcionalidade
fix:     correção de bug
refactor: refatoração sem mudar comportamento
style:   mudanças visuais / CSS
docs:    documentação
test:    testes
chore:   tarefas de manutenção (deps, ci, etc)
```

---

## Links úteis

- [Ollama API docs](https://github.com/ollama/ollama/blob/main/docs/api.md)
- [Go net/http docs](https://pkg.go.dev/net/http)
- [Share Tech Mono font](https://fonts.google.com/specimen/Share+Tech+Mono)
