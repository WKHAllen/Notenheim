import React from 'react';
import '../css/Home.css';
import { requestAPI } from '../requestAPI';
import { showAPIError } from '../apiError';
import { Link } from 'react-router-dom';

interface HomeState {
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
			lists: null
		};
	}

	public componentDidMount() {
		requestAPI('/getLists')
			.then(res => {
				if (res.error === null) {
					this.setState({
						lists: res.lists
					});
				} else {
					showAPIError(res.error);
				}
			});
	}

	public render() {
		if (this.state.lists === null) {
			return (
				<div className="Home">
					<h1>Your Lists</h1>
					<p className="loading">Fetching your lists...</p>
				</div>
			);
		} else {
			return (
				<div className="Home">
					<h1 className="mb-3">Your Lists</h1>
					<div className="Lists">
						<ul>
							{this.state.lists.map(item => 
								<li key={item.listID}>
									<Link to={`/list/${item.listID}`}>
										{item.title}
									</Link>
								</li>
							)}
						</ul>
					</div>
				</div>
			);
		}
	}
}
