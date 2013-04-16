function onlineEvtReg(){
	var list = document.getElementsByName("online");	
	for (var i = 0; i < list.length; i++){
			list[i].onclick=function(){
					alert(this.innerHTML)
			}
	}
};

function onlineInit(){
	onlineEvtReg();
};
