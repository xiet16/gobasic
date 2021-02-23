<template>
  <div class="hello">
      {{ msg }}
	<h2>WebSocket Test</h2>  
	<input type="text" id="input" />
	<button @click="sendBtnClick" >send</button>
	<button @click="closeBtnClick" >close</button>
	<div id="output"></div> 	
  </div>

</template>

<script>
var wsUri = 'ws://127.0.0.1:7777/ws'
var output
export default {
  name: 'HelloWorld',
  data () {
    return {
      msg: 'WelcoFIRST',
      websocket: ''
    }
  },
  created () {

  },
  methods: {
    init () {
	    output = document.getElementById('output')
        this.testWebSocket()
        window.addEventListener('load', this.init(), false)
	},

	     testWebSocket () {
	        this.websocket = new WebSocket(wsUri)
	        this.websocket.onopen = function (evt) {
	            this.onOpen(evt)
	    	 }
	        this.websocket.onclose = function (evt) {
	            this.onClose(evt)
	        }
	        this.websocket.onmessage = function (evt) {
	            this.onMessage(evt)
	        }
	        this.websocket.onerror = function (evt) {
	            this.onError(evt)
	        }
	    },

	     onOpen (evt) {
	        this.writeToScreen('CONNECTED')
	       // doSend("WebSocket rocks");
	    },

	     onClose (evt) {
	        this.writeToScreen('DISCONNECTED')
	    },

	     onMessage (evt) {
	        this.writeToScreen('<span style="color: blue;">RESPONSE: ' + evt.data + '</span>')
	       // websocket.close();
	    },

	     onError (evt) {
	        this.writeToScreen('<span style="color: red;">ERROR:</span> ' + evt.data)
	    },

	     doSend (message) {
	        this.writeToScreen('SENT: ' + message)
	        this.websocket.send(message)
	    },

	     writeToScreen (message) {
	        var pre = document.createElement('p')
	        pre.style.wordWrap = 'break-word'
	        pre.innerHTML = message
	        output.appendChild(pre)
	    },

	     sendBtnClick () {
			 debugger
	    	var msg = document.getElementById('input').value
	    	this.doSend(msg)
	    	document.getElementById('input').value = ''
	    },
	     closeBtnClick () {
	    	this.websocket.close()
	    }
  }
}
</script>
