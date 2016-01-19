var Cell = React.createClass({
  render: function() {
    return(
      <td>
        {this.props.value}
      </td>
    );
  }
});

var Row = React.createClass({
  render: function() {
    var cellValues = this.props.cells.map(function(cell) {
      return(<Cell key={cell} value={cell} />)
    });
    return(<tr>{cellValues}</tr>)
  }
});

function values(obj) {
    var vals = [];
    for( var key in obj ) {
        if ( obj.hasOwnProperty(key) ) {
            vals.push(obj[key]);
        }
    }
    return vals;
}

var Rows = React.createClass({
  render: function() {
    var rowNodes = this.props.body.map(function(node) {
      return(<Row key={node.name} cells={values(node)} />);
    });
    return(<tbody>
      <Header headers={this.props.headers} />
      {rowNodes}
    </tbody>);
  }
});

var Header = React.createClass({
  render: function() {
    var headerNodes = []


    for (var i in this.props.headers) {
      headerNodes.push(<th>{this.props.headers[i]}</th>)
    }

    return (
        <thead><tr>{headerNodes}</tr></thead>
    );
  }
});

var Table = React.createClass({
  render: function() {
    return (
        <table className="table table-bordered table-striped">
            <Rows headers={columns} body={this.props.body} />
        </table>
    );
  }
});

var columns = ["name", "description", "homepage", "installed"]

var data = [
  {"name":"LaTeX","description":"Compiles Tex","homepage":"https://www.latex-project.org","installed":"true"},
  {"name":"pdflatex","description":"Converts .tex to .pdf","homepage":"https://www.latex-project.org","installed":"true"},
  {"name":"convert","description":"Converts .pdf to .png","homepage":"http://www.imagemagick.org/script/index.php","installed":"true"}
];

ReactDOM.render(
  <Table headers={columns} body={data} />,
  document.getElementById('sanity-check')
);
