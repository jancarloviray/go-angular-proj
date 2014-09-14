window.app = angular.module('app',[
	'ngAnimate',
	'angular-growl'
])

app.config(['growlProvider', function(growlProvider) {
	growlProvider.globalPosition('bottom-right');
	growlProvider.globalTimeToLive(3000);
}]);

app.controller('todo/home', [
	'$scope', 'growl',
	function($scope, growl){
		var _this = this;

		_this.mock = {
			tasks: [
				{id:0, title:"Get Milk", done: true, created_date: new Date(), owner: "Jan Carlo" },
				{id:1, title:"Drink Milk", done: false, created_date: new Date(), owner: "Jan Carlo" }
			]
		};

		_this.model = {
			tasks: []
		};

		_this.refreshTasks = function(){
			_this.model.tasks = _this.mock.tasks;
		};

		_this.createTask = function(){
			var task = {}
			$scope.currentTask.id = newID();
			$scope.currentTask.created_date = new Date();

			// prevent passed by reference bug
			angular.extend(task, $scope.currentTask);

			_this.mock.tasks.push(task);
			_this.refreshTasks();

			growl.success('Successfully added a task: ' + task.title, {title:'Success'});
		};

		_this.updateTask = function(id, task){

		};

		_this.deleteTask = function(index){
			_this.mock.tasks.splice(index,1)
			_this.refreshTasks()

			growl.success('Successfully deleted a task.', {title:'Success'})
		};

		// HELPERS

		function newID(){
			return _this.mock.tasks[_this.mock.tasks.length - 1].id + 1;
		}

		_this.refreshTasks();

		$scope.taskFilter = "";

		$scope.currentTask = {
				id: null,
				title: '',
				owner: '',
				done: false,
				created_date: null,
		};

		$scope.home = _this;
	}
])
