/** @jsx React.DOM */

var app = app ||{};

var buttonStyle = {
	background: "#f7f7f7",
    fontsize: "14px",
    fontweight: "700",
    lineheight: "38px",
    padding: "0 15px 0 14px"
};

var Auth = React.createClass({
	render: function(){
		return (
			<form acton="/authorize" method="POST">
			   <button style={buttonStyle} onClick={this.openAuthPage}> Twitter Authorization </button>
			</form>
			)
	},

	openAuthPage: function(e) {
		
	}
});

ReactDOM.render(
	<Auth />,
	document.getElementById('auth')
);