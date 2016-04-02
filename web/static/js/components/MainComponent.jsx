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