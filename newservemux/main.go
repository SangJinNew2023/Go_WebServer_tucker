func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/". func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Hello Bar")
	})

	http.ListeAndServe(":3000", mux)
}