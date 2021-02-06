import React from "react";
import { Link } from "react-router-dom";
import "../css/Header.css";
import { getCookie } from "../cookie";

export default class Header extends React.Component {
  public render() {
    if (getCookie("loggedIn") !== "true") {
      return (
        <div className="Header">
          <Link to="/">
            <img src="/logo192.png" alt="logo" width="32" height="32" />
            <h5 className="Home-Label">Notenheim</h5>
          </Link>
          <nav>
            <ul className="Nav-Links">
              <li>
                <Link to="/login" aria-label="Login">
                  <i className="fas fa-sign-in-alt"></i>
                </Link>
              </li>
            </ul>
          </nav>
        </div>
      );
    } else {
      return (
        <div className="Header">
          <Link to="/">
            <img src="/logo192.png" alt="logo" width="32" height="32" />
            <h5 className="Home-Label">Notenheim</h5>
          </Link>
          <nav>
            <ul className="Nav-Links">
              <li>
                <Link to="/new" aria-label="New list">
                  <i className="fas fa-plus" />
                </Link>
              </li>
              <li>
                <Link to="/profile" aria-label="Profile">
                  <i className="fas fa-user"></i>
                </Link>
              </li>
              <li>
                <Link to="/logout" aria-label="Logout">
                  <i className="fas fa-sign-out-alt"></i>
                </Link>
              </li>
            </ul>
          </nav>
        </div>
      );
    }
  }
}
