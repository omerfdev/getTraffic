package main

import (
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/capture", captureTraffic)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func captureTraffic(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Sadece GET istekleri kabul edilir", http.StatusMethodNotAllowed)
        return
    }

    // URL parametresini al
    url := r.URL.Query().Get("url")

    // URL'yi tamamlamak için "http://" ön eki ekleyin
    if !startsWithHTTP(url) {
        url = "http://" + url
    }

    // URL'den veri al
    resp, err := http.Get(url)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    // Response body'sini oku
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Response body'sini geri döndür
    w.Header().Set("Content-Type", "text/plain")
    w.Write(body)
}

// Verilen URL'nin "http://" veya "https://" ile başlayıp başlamadığını kontrol eder
func startsWithHTTP(url string) bool {
    return len(url) >= 7 && (url[:7] == "http://" || url[:8] == "https://")
}
