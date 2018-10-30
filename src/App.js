import React, { Component } from "react";
import "./App.css";

import newApp from "./k8s-store.png";
import oldApp from "./k8s-store-old.png";
import errorPage from "./500-error.png";

class App extends Component {
  constructor() {
    super();
    this.state = { url: undefined };
  }
  componentDidMount() {
    fetch("/api/mode")
      .then(res => res.text())
      .then(res => {
        let url;
        switch (res.toLowerCase()) {
          case "new":
            url = newApp;
            break;
          case "error":
            url = errorPage;
            break;
          case "old":
          default:
            url = oldApp;
        }
        this.setState({ url });
      });
  }
  render() {
    return (
      <div className="App">
        {this.state.url && <img src={this.state.url} alt="main" />}
      </div>
    );
  }
}

export default App;
