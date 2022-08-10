import React from "react";
import "../css/Error.css";
import { hideAPIError } from "../apiError";

interface ErrorProps {
  message?: string;
}

export default class Error extends React.Component<ErrorProps> {
  public render() {
    return (
      <div
        id="error"
        className="alert alert-warning alert-dismissible fade show Error hidden"
        role="alert"
      >
        Error: <span id="error-message">{this.props.message}</span>
        <button
          type="button"
          className="close"
          data-dismiss="alert"
          aria-label="Close"
          onClick={hideAPIError}
        >
          <span aria-hidden="true" className="times">
            &times;
          </span>
        </button>
      </div>
    );
  }
}
