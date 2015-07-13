app.controller('ConnectController', ['$scope', function ($scope) {
	'use strict';

	$scope.host = window.location.hostname;
	$scope.port = window.location.port;


	$scope.reconnect = function() {
		location.reload();
  	};

}]);