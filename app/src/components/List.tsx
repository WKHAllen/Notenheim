import React from 'react';
import { arrayMove, SortableContainer, SortableElement, SortEnd, SortEvent } from 'react-sortable-hoc';
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

interface ListInfo {
	title: string,
	items: ListItem[]
}

interface ListState {
	refreshClicked: boolean,
	newItemFormGood: boolean,
	newItemSubmitClicked: boolean,
	editNameFormGood: boolean,
	editNameSubmitClicked: boolean,
	deleteListClicked: boolean,
	editItemFormGood: boolean,
	editItemSubmitClicked: boolean,
	editingItemID: string,
	deleteListItemClicked: boolean,
	listInfo: ListInfo | null
}

type CheckboxUpdateHandler = (listItemID: string) => void;
type EditDeleteOnClickHandler = (listItemID: string) => void;

const SortableListItem = SortableElement(({ item, checkboxUpdate, editDeleteOnClick }: { item: ListItem, checkboxUpdate: CheckboxUpdateHandler, editDeleteOnClick: EditDeleteOnClickHandler }) => (
	<li>
		<div className="d-flex ListItem">
			<div className="p-2">
				<label className="checkbox-container">
					<input type="checkbox" id={`checked-${item.listItemID}`} onChange={() => checkboxUpdate(item.listItemID)} checked={item.checked} />
					<span className="checkmark" />
				</label>
			</div>
			<div className="p-2 flex-grow-1">
				{item.content}
			</div>
			<div className="p-2 ListItem-Control">
				<button type="button" className="btn btn-pink btn-icon" data-toggle="modal" data-target="#edit-delete-item-modal" onClick={() => editDeleteOnClick(item.listItemID)}>
					<i className="fas fa-ellipsis-h" />
				</button>
			</div>
		</div>
	</li>
));

const SortableList = SortableContainer(({ items, checkboxUpdate, editDeleteOnClick }: { items: ListItem[], checkboxUpdate: CheckboxUpdateHandler, editDeleteOnClick: EditDeleteOnClickHandler }) => {
	return (
		<ul>
			{items.map((item, index) => 
				<SortableListItem
					key={`list-item-${index}`}
					index={index}
					item={item}
					checkboxUpdate={checkboxUpdate}
					editDeleteOnClick={editDeleteOnClick} />
			)}
		</ul>
	);
});

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
			editItemFormGood: true,
			editItemSubmitClicked: false,
			editingItemID: '',
			deleteListItemClicked: false,
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
										<button type="button" className="btn btn-pink btn-icon" onClick={() => this.props.history.push('/')}>
											<i className="fas fa-chevron-left" />
										</button>
									</div>
									<div>
										<button type="button" className="btn btn-pink btn-icon" onClick={() => this.getListInfo()} disabled={this.state.refreshClicked}>
											<i className="fas fa-sync-alt" />
										</button>
									</div>
									<div>
										<button type="button" className="btn btn-pink btn-icon" data-toggle="modal" data-target="#new-list-item-modal" onClick={() => this.prepareNewItemModal()}>
											<i className="fas fa-plus" />
										</button>
									</div>
									<div>
										<button type="button" className="btn btn-pink btn-icon" data-toggle="modal" data-target="#edit-list-name-modal" onClick={() => this.focusInput('list-name')}>
											<i className="fas fa-edit" />
										</button>
									</div>
									<div>
										<button type="button" className="btn btn-pink btn-icon" data-toggle="modal" data-target="#delete-list-modal">
											<i className="fas fa-trash" />
										</button>
									</div>
								</div>
							</li>
							<SortableList
								items={this.state.listInfo.items}
								onSortEnd={(end, event) => this.updateList(end, event)}
								distance={5}
								useWindowAsScrollContainer={true}
								checkboxUpdate={listItemID => this.checkListItem(listItemID)}
								editDeleteOnClick={listItemID => { this.focusInput('item-content'); this.setState({ editingItemID: listItemID }) }} />
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
									<div className="form-group">
										<label htmlFor="list-name">Item</label>
										<input type="text" className="form-control" id="new-item-content" name="new-item-content" maxLength={1023} onChange={() => this.checkNewItemForm()} onKeyDown={event => this.detectSubmit(event, 'new-item-button')} />
									</div>
								</div>
								<div className="modal-footer">
									<button type="button" className="btn btn-purple" data-dismiss="modal">Cancel</button>
									<button type="button" className="btn btn-pink" data-dismiss="modal" id="new-item-button" onClick={() => this.newListItem()} disabled={!this.state.newItemFormGood || this.state.newItemSubmitClicked}>Create item</button>
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
									<div className="form-group">
										<label htmlFor="list-name">List name</label>
										<input type="text" className="form-control" id="list-name" name="list-name" maxLength={255} onChange={() => this.checkEditNameForm()} defaultValue={this.state.listInfo.title} onKeyDown={event => this.detectSubmit(event, 'list-rename-button')} />
									</div>
								</div>
								<div className="modal-footer">
									<button type="button" className="btn btn-purple" data-dismiss="modal">Cancel</button>
									<button type="button" className="btn btn-pink" data-dismiss="modal" id="list-rename-button" onClick={() => this.editListName()} disabled={!this.state.editNameFormGood || this.state.editNameSubmitClicked}>Change name</button>
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
									<button type="button" className="btn btn-pink" onClick={() => this.deleteList()} data-dismiss="modal" disabled={this.state.deleteListClicked}>Delete</button>
								</div>
							</div>
						</div>
					</div>
					{/* Edit/delete list item modal */}
					<div className="modal fade ListModal" id="edit-delete-item-modal" tabIndex={-1} role="dialog" aria-labelledby="edit-delete-item-modal-label" aria-hidden="true">
						<div className="modal-dialog" role="document">
							<div className="modal-content">
								<div className="modal-header">
									<h5 className="modal-title" id="edit-delete-item-modal-label">Edit item</h5>
									<button type="button" className="close" data-dismiss="modal" aria-label="Close">
										<span aria-hidden="true" className="times">&times;</span>
									</button>
								</div>
								<div className="modal-body">
									<div className="form-group">
										<label htmlFor="item-content">Item content</label>
										<input type="text" className="form-control" id="item-content" name="item-content" maxLength={1023} onChange={() => this.checkEditItemForm()} defaultValue={this.getItemContent(this.state.editingItemID)} onKeyDown={event => this.detectSubmit(event, 'edit-item-button')} />
									</div>
								</div>
								<div className="modal-footer">
									<button type="button" className="btn btn-blue" data-dismiss="modal">Cancel</button>
									<button type="button" className="btn btn-purple" data-dismiss="modal" id="edit-item-button" onClick={() => this.editListItem()} disabled={!this.state.editItemFormGood || this.state.editItemSubmitClicked}>Edit</button>
									<button type="button" className="btn btn-pink" data-dismiss="modal" onClick={() => this.deleteListItem()}>Delete</button>
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
	}

	// private async updateList(newList: ListItem[]): Promise<void> {
	private async updateList(end: SortEnd, event: SortEvent): Promise<void> {
		if (this.state.listInfo !== null) {
			const oldList = this.state.listInfo.items;
			const newList = arrayMove(this.state.listInfo.items, end.oldIndex, end.newIndex);

			this.setState({
				listInfo: {
					title: this.state.listInfo.title,
					items: newList
				}
			});

			const listItemID = oldList[end.oldIndex].listItemID;
			const res = await requestAPI('/moveListItem', {
				listItemID,
				newPosition: end.newIndex
			});

			if (res.error === null) {
				hideAPIError();
			} else {
				showAPIError(res.error);
			}
		}
	}

	private async newListItem(): Promise<void> {
		this.setState({
			newItemSubmitClicked: true
		});

		const newItemContent = (document.getElementById('new-item-content') as HTMLInputElement).value;
		const res1 = await requestAPI('/newListItem', {
			listID: this.props.match.params.listID
		});

		if (res1.error === null) {
			hideAPIError();
			
			const res2 = await requestAPI('/editListItem', {
				listItemID: res1.listItemID,
				newContent: newItemContent
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

	private async editListName(): Promise<void> {
		this.setState({
			editNameSubmitClicked: true
		});

		const listName = (document.getElementById('list-name') as HTMLInputElement).value;
		const res = await requestAPI('/renameList', {
			listID: this.props.match.params.listID,
			newName: listName
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

		if (res.error === null) {
			hideAPIError();
			this.props.history.push('/');
		} else {
			showAPIError(res.error);
			this.setState({
				deleteListClicked: false
			});
		}
	}

	private async prepareNewItemModal(): Promise<void> {
		this.setState({ newItemFormGood: false });
		(document.getElementById('new-item-content') as HTMLInputElement).value = '';
		this.focusInput('new-item-content');
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

	private getItemContent(listItemID: string): string {
		if (this.state.listInfo !== null) {
			for (const item of this.state.listInfo.items) {
				if (item.listItemID === listItemID) {
					return item.content;
				}
			}
			return '';
		} else {
			return '';
		}
	}

	private async checkEditItemForm(): Promise<void> {
		const itemContent = (document.getElementById('item-content') as HTMLInputElement).value;

		this.setState({
			editItemFormGood: itemContent.length >= 1
		});
	}

	private async editListItem(): Promise<void> {
		this.setState({
			editItemSubmitClicked: true
		});

		const itemContent = (document.getElementById('item-content') as HTMLInputElement).value;
		const res = await requestAPI('/editListItem', {
			listItemID: this.state.editingItemID,
			newContent: itemContent
		});

		this.setState({
			editItemSubmitClicked: false
		});

		if (res.error === null) {
			hideAPIError();
			this.getListInfo();
		} else {
			showAPIError(res.error);
		}
	}

	private async deleteListItem(): Promise<void> {
		this.setState({
			deleteListItemClicked: true
		});

		const res = await requestAPI('/deleteListItem', {
			listItemID: this.state.editingItemID
		});
	
		this.setState({
			deleteListItemClicked: false
		});

		if (res.error === null) {
			hideAPIError();
			this.getListInfo();
		} else {
			showAPIError(res.error);
		}
	}

	private detectSubmit(event: React.KeyboardEvent<HTMLInputElement>, submitButtonID: string): void {
		const submitButton = (document.getElementById(submitButtonID) as HTMLButtonElement);
		if (event.keyCode === 13 && !submitButton.disabled) {
			submitButton.click();
		}
	}

	private focusInput(inputID: string): void {
		setTimeout(() => (document.getElementById(inputID) as HTMLInputElement).focus(), 500);
	}
}
