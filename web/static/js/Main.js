'use strict';

var UActions = require("./actions/UserActions.js");
var UStore = require("./stores/UserStore.js");

var LoginComponent = require("./components/LoginComponent.jsx");
var ErrorComponent = require("./components/ErrorComponent.jsx");

var UserActions = UActions.UserActions;

UserActions.renderApp({
	ifSignedIn:    ErrorComponent.errorMessage,
	ifNotSignedIn: LoginComponent.loginForm,
	ifError:       ErrorComponent.errorMessage,
	handler:       "content"
});