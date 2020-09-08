import React from 'react';
import '../css/Logout.css';
import { requestAPI } from '../requestAPI';
import { showAPIError } from '../apiError';
import { getCookie } from '../cookie';

export default class Logout extends React.Component<any> {
	public componentDidMount() {
		if (getCookie('loggedIn') === 'true') {
			requestAPI('/logout')
				.then(res => {
					if (res.error === null) {
						this.props.history.push('/');
						window.location.reload();
					} else {
						showAPIError(res.error);
					}
				});
		} else {
			this.props.history.push('/');
			window.location.reload();
		}
	}

	public render() {
		return (
			<div className="Logout">
				<h1>Log Out</h1>
				<p>Logging you out...</p>
			</div>
		);
	}
}
