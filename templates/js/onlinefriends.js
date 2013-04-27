var sock = null;
var wsuri = "ws://127.0.0.1:8888/wsOnlineFriends"
var localLng;
var localLat;
function onlineInit(){
	wsInit();
	getCurLocation();
	//onlineEvtReg();
};
function wsInit(){
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
			onlineFriendsRefresh();
		}else{
			$('#log').append('<p> others says: '+e.data+'</p>');
		}
	}
};
function onlineFriendsRefresh(){
	window.location.reload();
};
function onlineEvtReg(){
	var list = document.getElementsByName("online");	
	for (var i = 0; i < list.length; i++){
		list[i].onclick=function(){
			alert(this.innerHTML)
		}
	}
};
function savePosition(position){
	localLng = position.coords.longitude;
	localLat = position.coords.latitude;
	//console.log("savePosition() LNG:"+String(localLng));
	sock.send(String(localLng) + ":"+String(localLat));
};
function getCurLocation(){
	if (navigator.geolocation){
		navigator.geolocation.getCurrentPosition(savePosition);
	}else{
		alert("The browser can not support geolocation");
	}
};
