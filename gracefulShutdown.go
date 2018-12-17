package main
import (
       "os"
       "os/signal"
       "syscall"
       "fmt"
       "time"
       "net/http"
)
func main() {

       http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
              fmt.Fprint(w,"Server is running")
       })

       var gracefulStop = make(chan os.Signal)
       signal.Notify(gracefulStop, syscall.SIGTERM)
       signal.Notify(gracefulStop, syscall.SIGINT)
       go func() {
              sig := <-gracefulStop
              fmt.Printf("caught sig: %+v", sig)
              fmt.Println("Wait for 2 second to finish processing")
              time.Sleep(2*time.Second)
              os.Exit(0)
       }()
       http.ListenAndServe(":7080",nil)
}
