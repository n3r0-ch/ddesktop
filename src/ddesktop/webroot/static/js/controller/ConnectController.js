app.controller('ConnectController', ['$scope', function ($scope) {
	'use strict';

	$scope.reconnect = function() {
		location.reload();
  	};

}]);