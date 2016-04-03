'use strict';

var Reflux = require("reflux");
var React = require("react");
var ReactDOM = require("react-dom");
var UActions = require("../actions/UserActions.js");
var UStore = require("../stores/UserStore.js");

var userActions = UActions.UserActions;
var userStore = UStore.UserStore;

var csrfName = null;
var csrfValue = null;

function mountLogin(){
    ReactDOM.render(React.createElement(LoginForm, null), document.getElementById('content'));
}


function renderMainView(){
	$.ajax({
		url: "api/amIConnected",
		dataType: 'json',
		cache: false,
		success: function(data) {
			//
		},
		error: function(xhr, status, err) {
            //csrfName = xhr.getResponseHeader('X-CSRF-PARAM');
            //csrfValue = xhr.getResponseHeader('X-CSRF-TOKEN');
			mountLogin();
		}
	});
}

var LoginForm = React.createClass({
	render: function(){
		return(
			<input ref="login" type="text" placeholder="login" name="login"/>
		)
	}
})

exports.render = function(){
  renderMainView();
};