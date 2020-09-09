import React from 'react';
import '../css/List.css';
import { requestAPI } from '../requestAPI';
import { getCookie } from '../cookie';
import { hideAPIError, showAPIError } from '../apiError';

interface ListState {
	refreshClicked: boolean,
	listInfo: {
		title: string,
		items: {
			listItemID: string,
			content: string,
			position: number,
			checked: boolean
		}[]
	} | null
}

export default class List extends React.Component<any, ListState> {
	constructor(props: any) {
		super(props);

		this.state = {
			refreshClicked: false,
			listInfo: null
		};
	}

	public componentWillMount() {
		if (getCookie('loggedIn') !== 'true') {
			this.props.history.push(`/login?after=/list/${this.props.match.params.listID}`);
		}
	}

	public componentDidMount() {
		this.getListInfo();
	}

	public render() {
		if (this.state.listInfo === null) {
			return (
				<div className="List">
					<h1>List</h1>
					<p className="loading">Getting list data...</p>
				</div>
			);
		} else {
			return (
				<div className="List">
					<h1 className="mb-3">{this.state.listInfo.title}</h1>
					<div className="ListItems">
						<ul>
							<li>
								<div className="List-Buttons">
									<div>
										<button type="button" className="btn btn-primary btn-pink btn-icon" onClick={() => this.props.history.push('/')}>
											<i className="fas fa-chevron-left" />
										</button>
									</div>
									<div>
										<button type="button" className="btn btn-primary btn-pink btn-icon" onClick={() => this.getListInfo()} disabled={this.state.refreshClicked}>
											<i className="fas fa-sync-alt" />
										</button>
									</div>
									<div>
										<button type="button" className="btn btn-primary btn-pink btn-icon">
											<i className="fas fa-edit" />
										</button>
									</div>
									<div>
										<button type="button" className="btn btn-primary btn-pink btn-icon">
											<i className="fas fa-trash" />
										</button>
									</div>
								</div>
							</li>
							{this.state.listInfo.items.map(item => 
								<li key={item.listItemID}>
									<div className="d-flex ListItem">
										<div className="p-2">
											<div className="form-check ListItem-Check-Container">
												<input className="form-check-input ListItem-Check" type="checkbox" id={`checked-${item.listItemID}`} defaultChecked={item.checked} />
											</div>
										</div>
										<div className="p-2 flex-grow-1">
											{item.content}
										</div>
										<div className="p-2 ListItem-Control">
											<button type="button" className="btn btn-primary btn-pink btn-icon">
												<i className="fas fa-ellipsis-h" />
											</button>
										</div>
									</div>
								</li>
							)}
						</ul>
					</div>
				</div>
			);
		}
	}

	private async getListInfo(): Promise<void> {
		this.setState({
			refreshClicked: true
		});

		requestAPI('/listInfo', {
			listID: this.props.match.params.listID
		}).then(res => {
			if (res.error === null) {
				hideAPIError();
				this.setState({
					refreshClicked: false,
					listInfo: res.info
				});
			} else {
				this.setState({
					refreshClicked: false
				});
				showAPIError(res.error);
			}
		});
	}
}
