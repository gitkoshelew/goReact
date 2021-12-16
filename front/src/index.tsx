import React from "react";
import ReactDOM from "react-dom";
import "./index.css";
import App from "./main/ui/pages/App";
import { BrowserRouter as Router } from "react-router-dom";
import { Provider } from "react-redux";
import { store } from "./main/bll/store/store";

ReactDOM.render(
  <Router>
    <Provider store={store}>
      <App />
    </Provider>
  </Router>,
  document.getElementById("root")
);
