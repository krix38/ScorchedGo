'use strict';

var Reflux = require("reflux");

var UserActions = Reflux.createActions([
	"renderApp",
    "getAllChannels"
	
]);

exports.UserActions = UserActions;