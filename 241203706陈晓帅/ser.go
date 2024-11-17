package main

import (
	"net"
)

func Jiami(text, key string) string {

	keyLength := len(key)
	textLength := len(text)
	JiamiText := make([]byte, textLength)

	for i := 0; i < textLength; i++ {

		JiamiText[i] = text[i] ^ key[i%keyLength]
	}

	return string(JiamiText)
}

func Jiemi(JiamiText, key string) string {

	keyLength := len(key)
	JiamiTextLength := len(JiamiText)
	Jiemitext := make([]byte, JiamiTextLength)

	for i := 0; i < JiamiTextLength; i++ {
		Jiemitext[i] = JiamiText[i] ^ key[i%keyLength]

	}

	return string(Jiemitext)
}

func choose(conn net.Conn) {
	defer conn.Close()
	EorD := make([]byte, 1000)

	for {
		n, _ := conn.Read(EorD)

		EORD := string(EorD[:n])

		if EORD == "Jia" {
			text := make([]byte, 1000)
			key := make([]byte, 1000)

			for {
				t, _ := conn.Read(text)
				k, _ := conn.Read(key)

				JiamiText := Jiami(string(text[:t]), string(key[:k]))
				conn.Write([]byte(JiamiText))
			}
		} else if EORD == "Jie" {
			JiamiText := make([]byte, 1024)
			key := make([]byte, 1024)

			for {
				t, _ := conn.Read(JiamiText)
				k, _ := conn.Read(key)
				text := Jiemi(string(JiamiText[:t]), string(key[:k]))
				conn.Write([]byte(text))

			}
		}

	}
}

func main() {
	listener, _ := net.Listen("tcp", "127.0.0.1:8888")

	defer listener.Close()

	for {
		conn, _ := listener.Accept()

		go choose(conn)
	}
}
