var Header = React.createClass({

  render: function() {
    return (
        <nav className="navbar navbar-default">
          <div className="container-fluid">
            <div className="navbar-header">
              <a className="navbar-brand" href="/">Bit-A-Bytes</a>
            </div>

            <div className="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
              <ul className="nav navbar-nav">
                <li><a href="/create">Create</a></li>
                <li><a href="/about">About</a></li>
              </ul>
            </div>
          </div>
        </nav>
    );
  }
});

ReactDOM.render(
  <Header />,
  document.getElementById('header')
);
