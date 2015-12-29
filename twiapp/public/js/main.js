/** @jsx React.DOM */
var app = app || {};


var Tweet = React.createClass({

	/**
   	 * @return {object}
     */
	render: function(){
		return (
			<div id={this.props.id}>
			   <title> {this.props.title} </title>
			   <a href={this.props.link} </a>
			   <p> {this.props.message} </p>
			</div>
		)
	}
})

var App = React.createClass({

	getInitialState: function(){
		this.ws = new WebSocket("ws://" + location.host + "/sockets/" + getRandomId(1000,999999));
		return {
			tweets: [],
		}
	},
	componentDidMount: function(){
		var ws = new WebSocket();
		this.ws.addEventListener('message', function(e){
			var result = JSON.parse(e.data);
			if(result.event == "loadlist"){
				// Load list of tweets
			}
		});
	}

	/**
   	 * @return {object}
     */
	render : function(){

	};
});

ReactDOM.render(
	<App />,
	document.getElementById("dashboard")
	)