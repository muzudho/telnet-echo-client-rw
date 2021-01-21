package main

import (
	"bufio"
	"os"

	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
)

func startClient() error {
	// return telnet.DialToAndCall("localhost:5555", clientListener{})
	return telnet.DialToAndCall("localhost:9696", clientListener{})
}

type clientListener struct{}

// CallTELNET - 決まった形のメソッド。
func (c clientListener) CallTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {
	go read(r)
	write(w)
}

// 送られてくるメッセージを待ち構えるループです。
func read(r telnet.Reader) {
	var buffer [1]byte // これが満たされるまで待つ。1バイト。
	p := buffer[:]

	for {
		n, err := r.Read(p) // 送られてくる文字がなければ、ここでブロックされます。

		if n > 0 {
			bytes := p[:n]
			print(string(bytes)) // 受け取るたびに表示。
		}

		if nil != err {
			break // 相手が切断したなどの理由でエラーになるので、終了します。
		}
	}
	// ↑ このループから出れない☆（＾～＾）
}

// いつでも書き込んで送信できるようにするループです。
func write(w telnet.Writer) {
	// scanner - 標準入力を監視します。
	scanner := bufio.NewScanner(os.Stdin)
	// 一行読み取ります。
	for scanner.Scan() {
		// 書き込みます。最後に改行を付けます。
		oi.LongWrite(w, scanner.Bytes())
		oi.LongWrite(w, []byte("\n"))
	}
}
