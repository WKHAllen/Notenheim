import React from 'react';
import '../css/List.css';
import { requestAPI } from '../requestAPI';
import { getCookie } from '../cookie';
import { hideAPIError, showAPIError } from '../apiError';

interface ListState {
	refreshClicked: boolean,
	editingName: boolean,
	editNameFormGood: boolean,
	editNameSubmitClicked: boolean,
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
			editingName: false,
			editNameFormGood: false,
			editNameSubmitClicked: false,
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
					<h1 className="mb-3">List</h1>
					<p className="loading">Getting list data...</p>
				</div>
			);
		} else {
			return (
				<div className="List">
					<h1 className="mb-3">{this.state.listInfo.title}</h1>
					{this.state.editingName ?
						<div>
							<h3>Rename list</h3>
							<form onSubmit={event => { this.editListName(event); return false; }} className="mb-3">
								<div className="form-group">
									<label htmlFor="list-name">List name</label>
									<input type="text" className="form-control" id="list-name" name="list-name" maxLength={255} onChange={() => this.checkEditNameForm()} defaultValue={this.state.listInfo.title} />
								</div>
								<button type="submit" className="btn btn-primary btn-pink mr-2" disabled={!this.state.editNameFormGood || this.state.editNameSubmitClicked}>Change name</button>
								<button type="button" className="btn btn-primary btn-purple" onClick={() => this.disableNameEditing()}>Cancel</button>
							</form>
						</div>
					:
						<div className="hidden"></div>
					}
					
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
										<button type="button" className="btn btn-primary btn-pink btn-icon" onClick={() => this.enableNameEditing()} disabled={this.state.editingName}>
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
											<div className="form-check">
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

	private async enableNameEditing(): Promise<void> {
		this.setState({
			editingName: true
		});
	}

	private async disableNameEditing(): Promise<void> {
		this.setState({
			editingName: false
		});
	}

	private async editListName(e: React.FormEvent<HTMLFormElement>): Promise<void> {
		e.preventDefault();

		this.setState({
			editNameSubmitClicked: true
		});

		const formData = new FormData(e.currentTarget);
		const res = await requestAPI('/renameList', {
			listID: this.props.match.params.listID,
			newName: formData.get('list-name')
		});

		this.setState({
			editNameSubmitClicked: false
		});

		if (res.error === null) {
			hideAPIError();
			this.setState({
				editingName: false
			});
			this.getListInfo();
		} else {
			showAPIError(res.error);
		}
	}

	private async checkEditNameForm(): Promise<void> {
		const listName = (document.getElementById('list-name') as HTMLInputElement).value;

		this.setState({
			editNameFormGood: listName.length >= 1
		});
	}
}
