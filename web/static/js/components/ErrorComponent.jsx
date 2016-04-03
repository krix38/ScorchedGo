'use strict';

var React = require("react");

var ErrorMessage = React.createClass({
	render: function(){
		return(
			<div>Connection with server failed</div>
		)
	}
})

exports.errorMessage = ErrorMessage;