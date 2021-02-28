package basic

import (
	"flag"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

/*
websocket 是http 服务的升级，通过upgrader.Upgrage 可得到升级包
*/

var wsserverAddr = flag.String("wsserverAddr", "localhost:2003", "http service address")

//为什么要定义这个方法
var upgrader = websocket.Upgrader{
	//这是心跳方法吗
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebsocketTest() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", home)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/other", otherInfo)
	log.Println("Starting websocket server at " + *wsserverAddr)
	log.Fatal(http.ListenAndServe(*wsserverAddr, nil))
}

//疑问，这个时候访问http://localhost:2003/echo会怎样
func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
	}

	defer conn.Close()
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func otherInfo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "otherInfo")
}

//访问home 就会跳转到自定义websocket页面
//同时点击open ，就会连接websocket
func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {

    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;

    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };

    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
		var web_url=document.getElementById("web_url").value
        ws = new WebSocket(web_url);
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };

    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };

    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };

});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="web_url" type="text" value="{{.}}">
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))
