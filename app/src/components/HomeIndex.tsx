import React from "react";
import "../css/HomeIndex.css";
import { Link } from "react-router-dom";

export default class HomeIndex extends React.Component {
  public render() {
    return (
      <div className="HomeIndex">
        <div>
          <h1 className="HomeIndex-Title">Notenheim</h1>
          <p className="HomeIndex-Subtitle">
            Keep track of all your notes and lists from one place, across all
            devices.
          </p>
          <Link to="/login" className="btn btn-pink">
            Get started
          </Link>
        </div>
      </div>
    );
  }
}
