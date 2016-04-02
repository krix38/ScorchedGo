'use strict';

var Reflux = require("reflux");
var Actions = require("../actions/UserActions.js");
var UserActions = Actions.UsersActions;

var UserStore = Reflux.createStore({

            CSRF: null,

            listenables: [UserActions],

            getAllChannels: function(){
            },


            getInitialState: function() {
            }

    })

exports.UserStore = UserStore;
