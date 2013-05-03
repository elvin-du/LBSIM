var sock = null;
var wsuri = "ws://127.0.0.1:8888/wsChat"
function chatInit(){
	sock = new WebSocket(wsuri);
	sock.onopen = function(){
		console.log("connected to "+ wsuri);
	}
	sock.onclose = function(e){
		console.log("connection closed(" + e.code +")");
	}
	sock.onmessage = function(e){
		console.log("message received:" + e.data);
		$('#log').append('<p> others says: '+e.data+'</p>');
	}
};

function send(){
	var msg = document.getElementById('msg').value;
	$('#log').append('<p style="color:red;">I say: '+msg+'</p>');
	$('#log').get(0).scrollTop = $('#log').get(0).scrollHeight;
	$('#msg').val('');
	var ToWho = {{.Name}};
	if ("" == ToWho){
		alert("select who you want to talk");
		return
	}
	ToWho = ToWho + ":";
	sock.send(ToWho + msg);
};
