function storeChatLog(msg){
	var name = msg.substr(0, msg.indexOf(":"));
	var content = msg.substr(msg.indexOf(":")+1);
	var currentDate = new Date();
	var strDate = currentDate.toString();
	sessionStorage.setItem(name+"@"+strDate,content);
};
