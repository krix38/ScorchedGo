'use strict';

var React = require("react");
var ReactDOM = require("react-dom");
var Reflux = require("reflux");
var Actions = require("../actions/UserActions.js");
var UserActions = Actions.UserActions;

var UserStore = Reflux.createStore({

            CSRF: null,

            listenables: [UserActions],

            getAllChannels: function(){
				
            },
			
			unmount: function(component){
				ReactDOM.unmountComponentAtNode(document.getElementById(handler));
			},
			
			render: function(component, handler){
				ReactDOM.render(React.createElement(component, null), document.getElementById(handler));
			},
			
			renderApp: function(renderData){
				$.ajax({
					url: "api/connectionStatus/get",
					dataType: 'json',
					cache: false,
					success: function(sessionInfo) {
						if(sessionInfo.SignedIn){
							this.render(renderData.ifSignedIn, renderData.handler);
						}else{
							this.render(renderData.ifNotSignedIn, renderData.handler);
						}
					}.bind(this),
					error: function(xhr, status, err) {
						this.render(renderData.ifError, renderData.handler);
					}.bind(this)
				})
			},

            getInitialState: function(){
            }

    })

exports.UserStore = UserStore;
