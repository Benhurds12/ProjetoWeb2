package middleware

import "net/http"

// SecurityHeaders adiciona cabeçalhos de segurança em todas as respostas,
// mitigando ataques comuns (clickjacking, MIME sniffing, vazamento de referrer).
// Corresponde à categoria OWASP A05 – Security Misconfiguration.
func SecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Impede que o navegador "adivinhe" o content-type (MIME sniffing).
		w.Header().Set("X-Content-Type-Options", "nosniff")
		// Impede que a API seja embutida em iframes (clickjacking).
		w.Header().Set("X-Frame-Options", "DENY")
		// Não vaza a URL de origem em requisições para outros sites.
		w.Header().Set("Referrer-Policy", "no-referrer")
		// Como é uma API JSON, restringe o que pode ser carregado/embutido.
		w.Header().Set("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none'")

		next.ServeHTTP(w, r)
	})
}
