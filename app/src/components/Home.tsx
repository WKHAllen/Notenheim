import React from "react";
import { Link } from "react-router-dom";
import ReactMarkdown from "react-markdown";
import gfm from "remark-gfm";
import "../css/Home.css";
import { requestAPI } from "../requestAPI";
import { showAPIError, hideAPIError } from "../apiError";
import { getCookie } from "../cookie";
import HomeIndex from "./HomeIndex";

const refreshInterval = 60 * 1000; // One minute

interface HomeState {
  refreshClicked: boolean;
  lists:
    | {
        listID: string;
        title: string;
        updateTimestamp: number;
      }[]
    | null;
}

export default class Home extends React.Component<any, HomeState> {
  private refreshTimeout: NodeJS.Timeout;

  constructor(props: any) {
    super(props);

    this.state = {
      refreshClicked: false,
      lists: null,
    };

    this.refreshTimeout = setTimeout(() => {}, 0);
  }

  public componentDidMount() {
    if (getCookie("sessionID")) {
      this.getLists();
    }
  }

  public render() {
    if (!getCookie("sessionID")) {
      return <HomeIndex />;
    } else if (this.state.lists === null) {
      return (
        <div className="Home">
          <h1 className="mb-3">Notenheim</h1>
          <p className="loading">Fetching your lists...</p>
        </div>
      );
    } else {
      return (
        <div className="Home">
          <h1 className="mb-3">Notenheim</h1>
          <div className="Home-Lists">
            <ul>
              <li>
                <h2>Your lists</h2>
                <div className="Home-List-Buttons">
                  <button
                    type="button"
                    className="btn btn-pink btn-icon"
                    onClick={() => this.props.history.push("/new")}
                    aria-label="New list"
                  >
                    <i className="fas fa-plus" />
                  </button>
                  <button
                    type="button"
                    className="btn btn-pink btn-icon"
                    onClick={() => this.getLists()}
                    disabled={this.state.refreshClicked}
                    aria-label="Refresh"
                  >
                    <i className="fas fa-sync-alt" />
                  </button>
                </div>
              </li>
              {this.state.lists.map((item) => (
                <li key={item.listID}>
                  <Link to={`/list/${item.listID}`}>
                    <div className="d-flex Home-List">
                      <div className="p-2 flex-grow-1">
                        <ReactMarkdown
                          plugins={[gfm]}
                          children={item.title}
                          linkTarget="_blank"
                        />
                      </div>
                      <div className="p-2 Home-List-Timestamp">
                        {this.formatTimestamp(item.updateTimestamp)}
                      </div>
                    </div>
                  </Link>
                </li>
              ))}
            </ul>
          </div>
        </div>
      );
    }
  }

  private async getLists(): Promise<void> {
    clearTimeout(this.refreshTimeout);

    this.setState({
      refreshClicked: true,
    });

    const res = await requestAPI("/getLists");

    this.setState({
      refreshClicked: false,
    });

    if (res.error === null) {
      hideAPIError();
      this.setState({
        lists: res.lists,
      });
    } else {
      showAPIError(res.error);
    }

    this.refreshTimeout = setTimeout(() => this.getLists(), refreshInterval);
  }

  private formatTimestamp(timestamp: number): string {
    const now = Math.floor(new Date().getTime() / 1000);
    if (now - timestamp === 0) {
      return "now";
    } else if (now - timestamp < 60) {
      return `${now - timestamp}s`;
    } else if (now - timestamp < 60 * 60) {
      return `${Math.floor((now - timestamp) / 60)}m`;
    } else if (now - timestamp < 60 * 60 * 24) {
      return `${Math.floor((now - timestamp) / (60 * 60))}h`;
    } else if (now - timestamp < 60 * 60 * 24 * 7) {
      return `${Math.floor((now - timestamp) / (60 * 60 * 24))}d`;
    } else {
      return `${Math.floor((now - timestamp) / (60 * 60 * 24 * 7))}w`;
    }
  }
}
