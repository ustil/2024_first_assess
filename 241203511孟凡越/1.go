package main
import ("fmt"
       "net/http")
func index_handler(w http.Responsewriter,r *http.Request)
{
	fmt.Fprintf(w,"这是一个网页")
}
func index_handler(w http.Responsewriter,r *http.Request)
{
	fmt.Fprintf(w,"这是第二个网页")
}
func main()
{
http.HandleFunc("/", index_handler)
http.HandleFunc("/next/", next_handler)
http.ListenAndServe(":8080",nil)
}