window.app = angular.module('app',[])

app.controller('todo/home', [
	'$scope',
	function($scope){
		var _this = this;

		_this.model = {
			tasks : [
				{id:"16PT3437", title:"Get Milk", done: true, created_date: new Date(), owner: "Jan Carlo" },
				{id:"64314X2", title:"Drink Milk", done: false, created_date: new Date(), owner: "Jan Carlo" }
			]
		}

		$scope.home = _this;
	}
])
