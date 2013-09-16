function mainInit(){
	tabEvtReg();
};

function tabEvtReg(){
	var listTab= document.getElementsByName("tab");
	for(var i = 0; i<listTab.length; i++){
			listTab[i].onclick = function(){
					var listContent = document.getElementsByName("menuContent");
					var suffix = this.id.substr(3,this.id.length);
					
					for(var w=0; w < listContent.length; w++){
							if (suffix == listContent[w].id.substr(4, listContent[w].id.length)){
									listContent[w].style.display = "block";	
									continue;
							}
							listContent[w].style.display = "none";	
					}

					for(var j=0; j<listTab.length; j++){
							if(i != j){
									listTab[j].style.backgroundColor = "LightSlateBlue";
							}
					}
					this.style.backgroundColor = "MediumBlue";
					this.style.display = "block";
					//alert(this.innerHTML);
			}
	}
};

