app.controller('LoginController', [
	'$scope', 'LoginService', 
	function($scope, loginService) {
	'use strict';

	$scope.submit = function() {
		loginService.login($scope.form);
    };
}]);
