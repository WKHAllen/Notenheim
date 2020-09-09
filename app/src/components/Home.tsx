import React from 'react';
import '../css/Home.css';
import { requestAPI } from '../requestAPI';
import { showAPIError } from '../apiError';
import { Link } from 'react-router-dom';
import { getCookie } from '../cookie';

interface HomeState {
	refreshClicked: boolean,
	lists: {
		listID: string,
		title: string,
		updateTimestamp: number
	}[] | null
}

export default class Home extends React.Component<any, HomeState> {
	constructor(props: any) {
		super(props);

		this.state = {
			refreshClicked: false,
			lists: null
		};
	}

	public componentDidMount() {
		if (getCookie('loggedIn') === 'true') {
			this.getLists();
		}
	}

	public render() {
		if (getCookie('loggedIn') !== 'true') {
			return (
				<div className="Home">
					<h1>Notenheim</h1>
					<p>Hello! Welcome to Notenheim, a place where you can keep track of all your notes and lists from one place, across all devices. You're not logged in, but you can do so <Link to="/login">here</Link>. If you do not have an account, you can create one <Link to="/register">here</Link>.</p>
				</div>
			);
		} else if (this.state.lists === null) {
			return (
				<div className="Home">
					<h1>Notenheim</h1>
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
								<button type="button" className="btn btn-primary btn-pink" onClick={() => this.getLists()} disabled={this.state.refreshClicked}>
									<i className="fas fa-sync-alt" />
								</button>
							</li>
							{this.state.lists.map(item => 
								<li key={item.listID}>
									<Link to={`/list/${item.listID}`}>
										<div className="d-flex Home-List">
											<div className="p-2 flex-grow-1">{item.title}</div>
											<div className="p-2 Home-List-Timestamp">{this.formatTimestamp(item.updateTimestamp)}</div>
										</div>
									</Link>
								</li>
							)}
						</ul>
					</div>
				</div>
			);
		}
	}

	private async getLists(): Promise<void> {
		this.setState({
			refreshClicked: true
		});

		requestAPI('/getLists')
			.then(res => {
				this.setState({
					refreshClicked: false
				});

				if (res.error === null) {
					this.setState({
						lists: res.lists
					});
				} else {
					showAPIError(res.error);
				}
			});
	}

	private formatTimestamp(timestamp: number): string {
		const now = Math.floor(new Date().getTime() / 1000);
		if (now - timestamp === 0) {
			return 'now';
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
