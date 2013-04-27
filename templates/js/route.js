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
	var start =  new google.maps.LatLng(localLat, localLng);
	var endLat = endLat;//{{.Latitude}}; 
	var endLng = endLng;//{{.Longitude}}; 
	var end = new google.maps.LatLng(endLat, endLng);
	var request = {
		origin: start,
		destination: end,
		travelMode: google.maps.TravelMode.DRIVING
	};

	directionsService = google.maps.DirectionsService();
	directionsService.route(request, showRoute);
};

function showRoute(result,stat){
	if (stat == google.maps.DirectionsStatus.OK){
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
