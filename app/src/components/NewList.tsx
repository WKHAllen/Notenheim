import React from "react";
import "../css/NewList.css";
import { requestAPI } from "../requestAPI";
import { showAPIError, hideAPIError } from "../apiError";
import { getCookie } from "../cookie";

interface NewListState {
  formGood: boolean;
  submitClicked: boolean;
}

export default class NewList extends React.Component<any, NewListState> {
  constructor(props: any) {
    super(props);

    this.state = {
      formGood: false,
      submitClicked: false,
    };
  }

  public componentWillMount() {
    if (getCookie("loggedIn") !== "true") {
      this.props.history.push("/login?after=/new");
    }
  }

  public render() {
    return (
      <div className="NewList">
        <h1 className="mb-3">New List</h1>
        <form
          onSubmit={(event) => {
            this.createNewList(event);
            return false;
          }}
          className="mb-3"
        >
          <div className="form-group">
            <label htmlFor="list-name">List name</label>
            <input
              type="text"
              className="form-control"
              id="list-name"
              name="list-name"
              maxLength={255}
              onChange={() => this.checkForm()}
              autoFocus
            />
            <small className="form-text">You can change this later.</small>
          </div>
          <button
            type="submit"
            className="btn btn-pink"
            disabled={!this.state.formGood || this.state.submitClicked}
          >
            Create list
          </button>
        </form>
      </div>
    );
  }

  private async createNewList(
    e: React.FormEvent<HTMLFormElement>
  ): Promise<void> {
    e.preventDefault();

    this.setState({
      submitClicked: true,
    });

    const formData = new FormData(e.currentTarget);
    const res = await requestAPI("/newList", {
      title: formData.get("list-name"),
    });

    this.setState({
      submitClicked: false,
    });

    if (res.error === null) {
      hideAPIError();
      this.props.history.push(`/list/${res.listID}`);
    } else {
      showAPIError(res.error);
    }
  }

  private async checkForm(): Promise<void> {
    const listName = (document.getElementById("list-name") as HTMLInputElement)
      .value;

    this.setState({
      formGood: listName.length >= 1,
    });
  }
}
