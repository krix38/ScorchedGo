'use strict';

var React = require("react");

var LoginForm = React.createClass({
	render: function(){
		return(
			<input ref="login" type="text" placeholder="login" name="login"/>
		)
	}
})

exports.loginForm = LoginForm;