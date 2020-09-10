import React from 'react';
import { ReactSortable } from 'react-sortablejs';
import '../css/List.css';
import { requestAPI } from '../requestAPI';
import { getCookie } from '../cookie';
import { hideAPIError, showAPIError } from '../apiError';

interface ListItem {
	listItemID: string,
	content: string,
	position: number,
	checked: boolean
}

interface ListItemWithID {
	id: string,
	listItemID: string,
	content: string,
	position: number,
	checked: boolean
}

interface ListInfo {
	title: string,
	items: ListItem[]
}

interface ListInfoWithIDs {
	title: string,
	items: ListItemWithID[]
}

interface ListState {
	refreshClicked: boolean,
	newItemFormGood: boolean,
	newItemSubmitClicked: boolean,
	editNameFormGood: boolean,
	editNameSubmitClicked: boolean,
	deleteListClicked: boolean,
	listInfo: ListInfoWithIDs | null
}

export default class List extends React.Component<any, ListState> {
	constructor(props: any) {
		super(props);

		this.state = {
			refreshClicked: false,
			newItemFormGood: false,
			newItemSubmitClicked: false,
			editNameFormGood: true,
			editNameSubmitClicked: false,
			deleteListClicked: false,
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
										<button type="button" className="btn btn-primary btn-pink btn-icon" data-toggle="modal" data-target="#new-list-item-modal" onClick={() => { (document.getElementById('new-item-content') as HTMLInputElement).value = '' }}>
											<i className="fas fa-plus" />
										</button>
									</div>
									<div>
										<button type="button" className="btn btn-primary btn-pink btn-icon" data-toggle="modal" data-target="#edit-list-name-modal">
											<i className="fas fa-edit" />
										</button>
									</div>
									<div>
										<button type="button" className="btn btn-primary btn-pink btn-icon" data-toggle="modal" data-target="#delete-list-modal">
											<i className="fas fa-trash" />
										</button>
									</div>
								</div>
							</li>
							<ReactSortable list={this.state.listInfo.items} setList={newList => this.updateList(newList)}>
								{this.state.listInfo.items.map(item => 
									<li key={item.listItemID}>
										<div className="d-flex ListItem">
											<div className="p-2">
												<div className="form-check">
													<input className="form-check-input ListItem-Check" type="checkbox" id={`checked-${item.listItemID}`} onChange={() => this.checkListItem(item.listItemID)} checked={item.checked} />
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
							</ReactSortable>
						</ul>
					</div>
					{/* New list item modal */}
					<div className="modal fade ListModal" id="new-list-item-modal" tabIndex={-1} role="dialog" aria-labelledby="new-list-item-label" aria-hidden="true">
						<div className="modal-dialog" role="document">
							<div className="modal-content">
								<div className="modal-header">
									<h5 className="modal-title" id="new-list-item-label">New item</h5>
									<button type="button" className="close" data-dismiss="modal" aria-label="Close">
										<span aria-hidden="true" className="times">&times;</span>
									</button>
								</div>
								<div className="modal-body">
									<form onSubmit={event => { this.newListItem(event); return false; }} className="mb-3">
										<div className="form-group">
											<label htmlFor="list-name">Item</label>
											<input type="text" className="form-control" id="new-item-content" name="new-item-content" maxLength={1023} onChange={() => this.checkNewItemForm()} />
										</div>
									</form>
								</div>
								<div className="modal-footer">
									<button type="button" className="btn btn-purple" data-dismiss="modal" id="cancel-new-item-button">Cancel</button>
									<button type="submit" className="btn btn-pink" data-dismiss="modal" disabled={!this.state.newItemFormGood || this.state.newItemSubmitClicked}>Create item</button>
								</div>
							</div>
						</div>
					</div>
					{/* Edit list name modal */}
					<div className="modal fade ListModal" id="edit-list-name-modal" tabIndex={-1} role="dialog" aria-labelledby="edit-list-name-label" aria-hidden="true">
						<div className="modal-dialog" role="document">
							<div className="modal-content">
								<div className="modal-header">
									<h5 className="modal-title" id="edit-list-name-label">Rename list</h5>
									<button type="button" className="close" data-dismiss="modal" aria-label="Close">
										<span aria-hidden="true" className="times">&times;</span>
									</button>
								</div>
								<div className="modal-body">
									<form onSubmit={event => { this.editListName(event); return false; }} className="mb-3">
										<div className="form-group">
											<label htmlFor="list-name">List name</label>
											<input type="text" className="form-control" id="list-name" name="list-name" maxLength={255} onChange={() => this.checkEditNameForm()} defaultValue={this.state.listInfo.title} />
										</div>
									</form>
								</div>
								<div className="modal-footer">
									<button type="button" className="btn btn-purple" data-dismiss="modal" id="cancel-list-rename-button">Cancel</button>
									<button type="submit" className="btn btn-pink" data-dismiss="modal" disabled={!this.state.editNameFormGood || this.state.editNameSubmitClicked}>Change name</button>
								</div>
							</div>
						</div>
					</div>
					{/* Delete list modal */}
					<div className="modal fade ListModal" id="delete-list-modal" tabIndex={-1} role="dialog" aria-labelledby="delete-list-modal-label" aria-hidden="true">
						<div className="modal-dialog" role="document">
							<div className="modal-content">
								<div className="modal-header">
									<h5 className="modal-title" id="delete-list-modal-label">Delete list</h5>
									<button type="button" className="close" data-dismiss="modal" aria-label="Close">
										<span aria-hidden="true" className="times">&times;</span>
									</button>
								</div>
								<div className="modal-body">
									Are you sure you want to delete this list?
								</div>
								<div className="modal-footer">
									<button type="button" className="btn btn-purple" data-dismiss="modal">Cancel</button>
									<button type="button" className="btn btn-pink" onClick={() => this.deleteList()} data-dismiss="modal">Delete</button>
								</div>
							</div>
						</div>
					</div>
				</div>
			);
		}
	}

	private async getListInfo(): Promise<void> {
		this.setState({
			refreshClicked: true
		});

		const res = await requestAPI('/listInfo', {
			listID: this.props.match.params.listID
		});

		if (res.error === null) {
			hideAPIError();
			const listInfo = this.addListInfoIDs(res.info);
			this.setState({
				refreshClicked: false,
				listInfo
			});
		} else {
			this.setState({
				refreshClicked: false
			});
			showAPIError(res.error);
		}
	}

	private addListInfoIDs(listInfo: ListInfo): ListInfoWithIDs {
		let listInfoWithIDs: ListInfoWithIDs = {
			title: listInfo.title,
			items: []
		};

		for (const listItem of listInfo.items) {
			listInfoWithIDs.items.push({
				id: listItem.listItemID,
				listItemID: listItem.listItemID,
				content: listItem.content,
				position: listItem.position,
				checked: listItem.checked
			});
		}

		return listInfoWithIDs;
	}

	private async updateList(newList: ListItemWithID[]): Promise<void> {
		if (this.state.listInfo !== null) {
			this.setState({
				listInfo: {
					title: this.state.listInfo.title,
					items: newList
				}
			});
		}
	}

	private async newListItem(e: React.FormEvent<HTMLFormElement>): Promise<void> {
		document.getElementById('cancel-new-item-button')?.click();

		e.preventDefault();

		this.setState({
			newItemSubmitClicked: true
		});

		const formData = new FormData(e.currentTarget);
		const res1 = await requestAPI('/newListItem', {
			listID: this.props.match.params.listID
		});

		if (res1.error === null) {
			hideAPIError();
			
			const res2 = await requestAPI('/editListItem', {
				listItemID: res1.listItemID,
				newContent: formData.get('new-item-content')
			});

			this.setState({
				newItemSubmitClicked: false
			});
			if (res2.error === null) {
				this.getListInfo();
			} else {
				showAPIError(res2.error);
			}
		} else {
			showAPIError(res1.error);
		}
	}

	private async checkNewItemForm(): Promise<void> {
		const itemContent = (document.getElementById('new-item-content') as HTMLInputElement).value;

		this.setState({
			newItemFormGood: itemContent.length >= 1
		});
	}

	private async editListName(e: React.FormEvent<HTMLFormElement>): Promise<void> {
		document.getElementById('cancel-list-rename-button')?.click();

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

	private async deleteList(): Promise<void> {
		this.setState({
			deleteListClicked: true
		});

		const res = await requestAPI('/deleteList', {
			listID: this.props.match.params.listID
		});

		this.setState({
			deleteListClicked: false
		});

		if (res.error === null) {
			hideAPIError();
			this.props.history.push('/');
		} else {
			showAPIError(res.error);
		}
	}

	private async checkListItem(listItemID: string): Promise<void> {
		const checked = (document.getElementById(`checked-${listItemID}`) as HTMLInputElement).checked;
		this.setCheckState(listItemID, checked);
		
		const res = await requestAPI('/checkListItem', {
			listItemID,
			checked
		});

		if (res.error === null) {
			hideAPIError();
		} else {
			showAPIError(res.error);
		}
	}

	private async setCheckState(listItemID: string, checked: boolean): Promise<void> {
		if (this.state.listInfo !== null) {
			const listInfo = this.state.listInfo;
			for (let i = 0; i < listInfo.items.length; i++) {
				if (listInfo.items[i].listItemID === listItemID) {
					listInfo.items[i].checked = checked;
				}
			}

			this.setState({
				listInfo
			});
		}
	}
}
