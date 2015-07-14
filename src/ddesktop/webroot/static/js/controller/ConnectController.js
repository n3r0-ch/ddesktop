app.controller('ConnectController', ['$scope', function ($scope) {
	'use strict';

	var port = window.location.port;
	if (!port) {
		if (window.location.protocol != "https") {
			$scope.port = 443;
		}
		else {
			$scope.port = 80;
		}
	}
	else{
		$scope.port = port;
	}
	$scope.host = window.location.hostname;

	$scope.reconnect = function() {
		location.reload();
  	};

}]);