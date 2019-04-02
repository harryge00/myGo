package main

import (
    "log"
    "time"
    "github.com/miekg/dns"
)

func main() {

    target := "test2.marathon.slave.mesos.zj.chinamobile.com"
    server := "172.31.17.186"
    qps := 3

    c := dns.Client{}
    m := dns.Msg{}
    m.SetQuestion(target+".", dns.TypeA)
    ticker := time.NewTicker(1000 * time.Millisecond)
    for t := range ticker.C {
        log.Println("Tick at", t)
        for i := 0; i < qps; i++ {
            go func() {
                r, t, err := c.Exchange(&m, server+":53")
                if err != nil {
                    log.Println(err)
                    return
                }
                log.Printf("Took %v", t)
                if len(r.Answer) == 0 {
                    log.Println("No results")
                    return
                }
                for _, ans := range r.Answer {
                    Arecord := ans.(*dns.A)
                    log.Printf("%s", Arecord.A)
                }
            } ()
        }
    }
}
