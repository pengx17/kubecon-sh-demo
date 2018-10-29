import React, { Component } from "react";
import "./App.css";

import newMain from "./k8s-store.png";
import oldMain from "./k8s-store-old.png";

class App extends Component {
  constructor() {
    super();
    this.state = { url: undefined };
  }
  componentDidMount() {
    fetch("/api/mode")
      .then(res => res.text())
      .then(res => {
        const url = res.toLowerCase() === "new" ? newMain : oldMain;
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
