var app = angular.module("ddesktop", [
	'ngRoute',
	'ngResource',
	'noVNC'
])

app.config(['$routeProvider', function ($routeProvider) {

	'use strict';

	$routeProvider.when('/login', {
		templateUrl: '/static/partials/login.html',
		controller: 'LoginController'
	})
	$routeProvider.when('/connect', {
		templateUrl: '/static/partials/connect.html',
		controller: 'ConnectController'
	})
	$routeProvider.when('/error/:code', {
		templateUrl: '/static/partials/error.html',
		controller: 'ErrorController'
	})
	$routeProvider.when('/', {redirectTo: '/login'})

	$routeProvider.otherwise({redirectTo: '/error/404'});
}]);
