app.controller('ErrorController', ['$scope', '$routeParams',
	function ($scope, $routeParams) {

	'use strict';

	$scope.code = 404;
	if (typeof $routeParams.code !== "undefined") {
		if (parseInt($routeParams.code) > 0 &&  parseInt($routeParams.code) < 1000){
			$scope.code  = $routeParams.code;
		}
	}
}]);