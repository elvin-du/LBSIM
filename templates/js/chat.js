var sock = null;
var wsuri = "ws://127.0.0.1:8888/wsChat"
function chatInit(){
	console.log("onload");
	sock = new WebSocket(wsuri);
	sock.onopen = function(){
		console.log("connected to "+ wsuri);
	}
	sock.onclose = function(e){
		console.log("connection closed(" + e.code +")");
	}
	sock.onmessage = function(e){
		console.log("message received:" + e.data);
		if(e.data == "UO"){
			onlineUserRefresh();
		}else{
			$('#log').append('<p> others says: '+e.data+'</p>');
		}
	}
};

function send(){
	var msg = document.getElementById('message').value;
	$('#log').append('<p style="color:red;">I say: '+msg+'</p>');
	$('#log').get(0).scrollTop = $('#log').get(0).scrollHeight;
	$('#message').val('');
	var ToWho = "";//{{.Name}} + ":";
	sock.send(ToWho + msg);
};

function onlineUserRefresh(){
	window.location.reload();
};

