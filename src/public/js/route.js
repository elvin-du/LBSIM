var directionsDisplay;
var directionsService;
var gmap;
var localLat;
var localLng;
var marker;

function routeInit(){
	getLocation();
};

function calcRoute(endLat, endLng) {
	if(endLat < 0 || endLng < 0){
		console.log(String(endLat)+":"+String(endLng) +"is not right GPS");
		alert("you do not select people who you want to see");
		return
	}
	var start =  new google.maps.LatLng(localLat, localLng);
	//var end = new google.maps.LatLng(endLat, endLng);
	//for test
	var end = new google.maps.LatLng(31.186172,121.427414);
	var request = {
		origin: start,
		destination: end,
		travelMode: google.maps.TravelMode.DRIVING
	};

	directionsService = new google.maps.DirectionsService();
	directionsService.route(request, showRoute);
};

function showRoute(result,stat){
	console.log(stat);
	if (stat == google.maps.DirectionsStatus.OK){
		console.log("showroute");
		path = result.routes[0].overview_path;
		if(path){
			window.setInterval(function(){
				if(!marker){
					marker = new google.maps.Marker({position:path[0], map:gmap});
				}else{
					marker.setPosition(path[0]);
				}
			}, 1);
		}
		
		directionsDisplay = new google.maps.DirectionsRenderer();
		directionsDisplay.setMap(gmap);
		directionsDisplay.setDirections(result);
	}
};

function savePosition(position){
	localLng = position.coords.longitude;
	localLat = position.coords.latitude;
	console.log("you are in where lat is" + String(localLat) + ","+"lng is"+String(localLng));
	var local = new google.maps.LatLng(localLat, localLng);
	var mapOptions = {
		center: local, 
		zoom: 18,
		mapTypeId: google.maps.MapTypeId.ROADMAP
	};

	gmap = new google.maps.Map(document.getElementById("mapCanvas"), mapOptions);
	if(!marker){
		marker = new google.maps.Marker({position:local, map:gmap});
	}else{
		marker.setPosition(path[0]);
	}
};

function getLocation(){
	if (navigator.geolocation){
		navigator.geolocation.getCurrentPosition(savePosition);
	}else{
		alert("The browser can not support geolocation");
	}
};
