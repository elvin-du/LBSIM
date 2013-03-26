function login() {
		//window.location.href='http://127.0.0.1:8888/chat';
		var username= document.getElementById('username').value;
		var password= document.getElementById('password').value;
		sock.send(username + password);
}

