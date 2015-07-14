var app = angular.module("ddesktop", [
	'ngRoute',
	'noVNC'
])

app.config(['$routeProvider', function ($routeProvider) {

	'use strict';

	$routeProvider.when('/connect', {
		templateUrl: '/partials/connect.html',
		controller: 'ConnectController'
	})
	$routeProvider.when('/error/:code', {
		templateUrl: '/partials/error.html',
		controller: 'ErrorController'
	})
	$routeProvider.when('/', {redirectTo: '/connect'})

	$routeProvider.otherwise({redirectTo: '/error/404'});
}]);
