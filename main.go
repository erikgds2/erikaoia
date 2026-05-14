package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

const (
	defaultPort  = "8080"
	defaultModel = "llama3.2:latest"
	ollamaBase   = "http://localhost:11434"
	logPath      = "queries.json"
	systemPrompt = `Você é uma IA criada pelo erikao. Responda sempre em português brasileiro, de forma concisa e direta. Seja levemente irônico e bem-humorado, mas genuinamente útil. Quando não souber algo, admita com humor. Evite respostas muito longas — prefira precisão. Você roda localmente no PC do erikao via Ollama.`
)

// ── Ollama types ──────────────────────────────────────────────────────────────

type ollamaGenerateReq struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	System string `json:"system"`
	Stream bool   `json:"stream"`
}

type ollamaGenerateResp struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

type ollamaTagsResp struct {
	Models []struct {
		Name    string `json:"name"`
		Size    int64  `json:"size"`
		Details struct {
			Family string `json:"family"`
		} `json:"details"`
	} `json:"models"`
}

// ── API types ─────────────────────────────────────────────────────────────────

type ChatRequest struct {
	Message string `json:"message"`
	Model   string `json:"model,omitempty"`
}

type ChatResponse struct {
	Result     string `json:"result"`
	Model      string `json:"model"`
	DurationMs int64  `json:"duration_ms"`
	Error      string `json:"error,omitempty"`
}

// ── Log types ─────────────────────────────────────────────────────────────────

type QueryEntry struct {
	ID         int    `json:"id"`
	Timestamp  string `json:"timestamp"`
	Question   string `json:"question"`
	Answer     string `json:"answer"`
	Model      string `json:"model"`
	DurationMs int64  `json:"duration_ms"`
}

// ── State ─────────────────────────────────────────────────────────────────────

var (
	queries []QueryEntry
	mu      sync.RWMutex
	nextID  = 1
)

// ── Main ──────────────────────────────────────────────────────────────────────

func main() {
	loadLogs()

	port := envOr("PORT", defaultPort)
	model := envOr("OLLAMA_MODEL", defaultModel)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/chat", cors(chatHandler(model)))
	mux.HandleFunc("/api/logs", cors(logsHandler))
	mux.HandleFunc("/api/models", cors(modelsHandler))
	mux.HandleFunc("/api/status", cors(statusHandler(model)))
	mux.Handle("/", http.FileServer(http.Dir("frontend")))

	log.Printf("erikao-ai iniciado em http://localhost:%s", port)
	log.Printf("modelo padrao: %s | admin: http://localhost:%s/admin.html", model, port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}

// ── Handlers ──────────────────────────────────────────────────────────────────

func chatHandler(defaultModel string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			jsonErr(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req ChatRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Message == "" {
			jsonErr(w, "mensagem inválida", http.StatusBadRequest)
			return
		}

		model := req.Model
		if model == "" {
			model = defaultModel
		}

		start := time.Now()
		result, err := callOllama(model, req.Message)
		elapsed := time.Since(start).Milliseconds()

		resp := ChatResponse{Model: model, DurationMs: elapsed}
		if err != nil {
			resp.Error = err.Error()
			resp.Result = "erro ao conectar com ollama. certifique-se que `ollama serve` está rodando."
		} else {
			resp.Result = result
		}

		logQuery(QueryEntry{
			Timestamp:  time.Now().Format("2006-01-02 15:04:05"),
			Question:   req.Message,
			Answer:     resp.Result,
			Model:      model,
			DurationMs: elapsed,
		})

		writeJSON(w, resp)
	}
}

func logsHandler(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()
	writeJSON(w, queries)
}

func modelsHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(ollamaBase + "/api/tags")
	if err != nil {
		jsonErr(w, "ollama inacessível", http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, resp.Body)
}

func statusHandler(model string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ollamaOK := false
		resp, err := http.Get(ollamaBase + "/api/tags")
		if err == nil {
			resp.Body.Close()
			ollamaOK = resp.StatusCode == 200
		}

		mu.RLock()
		total := len(queries)
		mu.RUnlock()

		writeJSON(w, map[string]any{
			"ollama_ok":    ollamaOK,
			"model":        model,
			"total_queries": total,
			"server":       "erikao-ai",
		})
	}
}

// ── Ollama ────────────────────────────────────────────────────────────────────

func callOllama(model, prompt string) (string, error) {
	payload, _ := json.Marshal(ollamaGenerateReq{
		Model:  model,
		Prompt: prompt,
		System: systemPrompt,
		Stream: false,
	})

	resp, err := http.Post(ollamaBase+"/api/generate", "application/json", bytes.NewReader(payload))
	if err != nil {
		return "", fmt.Errorf("ollama unreachable: %w", err)
	}
	defer resp.Body.Close()

	var result ollamaGenerateResp
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("decode error: %w", err)
	}

	return result.Response, nil
}

// ── Persistence ───────────────────────────────────────────────────────────────

func logQuery(q QueryEntry) {
	mu.Lock()
	defer mu.Unlock()
	q.ID = nextID
	nextID++
	queries = append(queries, q)
	data, _ := json.MarshalIndent(queries, "", "  ")
	os.WriteFile(logPath, data, 0644)
}

func loadLogs() {
	data, err := os.ReadFile(logPath)
	if err != nil {
		return
	}
	mu.Lock()
	defer mu.Unlock()
	if err := json.Unmarshal(data, &queries); err == nil && len(queries) > 0 {
		nextID = queries[len(queries)-1].ID + 1
		log.Printf("%d queries carregadas do histórico", len(queries))
	}
}

// ── Helpers ───────────────────────────────────────────────────────────────────

func cors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		h(w, r)
	}
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func jsonErr(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
