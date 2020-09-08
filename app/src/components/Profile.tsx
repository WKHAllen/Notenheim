import React from 'react';
import '../css/Profile.css';
import { requestAPI } from '../requestAPI';
import { hideAPIError, showAPIError } from '../apiError';
import { getCookie } from '../cookie';

interface ProfileState {
	formGood: boolean,
	submitClicked: boolean,
	profileInfo: {
		email: string,
		joinTimestamp: number
	} | null
}

export default class Profile extends React.Component<any, ProfileState> {
	constructor(props: any) {
		super(props);

		this.state = {
			formGood: false,
			submitClicked: false,
			profileInfo: null
		};
	}

	public componentDidMount() {
		requestAPI('/getProfileInfo')
			.then(res => {
				if (res.error === null) {
					hideAPIError();
					this.setState({
						profileInfo: res.info
					});
				} else {
					showAPIError(res.error);
				}
			});
	}

	public render() {
		if (getCookie('loggedIn') !== 'true') {
			this.props.history.push('/login');
		} else if (this.state.profileInfo === null) {
			return (
				<div className="Profile">
					<h1>Profile</h1>
					<p className="loading">Getting profile data...</p>
				</div>
			);
		} else {
			return (
				<div className="Profile">
					<h1>Profile</h1>
					<ul>
						<li>{this.state.profileInfo.email}</li>
						<li>{new Date(this.state.profileInfo.joinTimestamp * 1000).toDateString()}</li>
					</ul>
					<h3>Change your password</h3>
					<form onSubmit={event => { this.changePassword(event); return false; }} className="mb-3">
						<div className="form-group">
							<label htmlFor="password">New password</label>
							<input type="password" className="form-control" id="password" name="password" maxLength={255} onChange={() => this.checkForm()} />
						</div>
						<div className="form-group">
							<label htmlFor="confirm-password">Confirm new password</label>
							<input type="password" className="form-control" id="confirm-password" name="confirm-password" maxLength={255} onChange={() => this.checkForm()} />
						</div>
						<button type="submit" className="btn btn-primary btn-pink" disabled={!this.state.formGood || this.state.submitClicked}>Change password</button>
					</form>
				</div>
			);
		}
	}

	private async changePassword(e: React.FormEvent<HTMLFormElement>): Promise<void> {
		e.preventDefault();

		this.setState({
			submitClicked: true
		});

		const formData = new FormData(e.currentTarget);
		const res = await requestAPI('/changePassword', {
			newPassword: formData.get('password')
		});

		this.setState({
			submitClicked: false
		});

		if (res.error === null) {
			hideAPIError();
		} else {
			showAPIError(res.error);
		}
	}

	private async checkForm(): Promise<void> {
		const password          = (document.getElementById('password')         as HTMLInputElement).value;
		const confirmedPassword = (document.getElementById('confirm-password') as HTMLInputElement).value;

		this.setState({
			formGood: password === confirmedPassword && password.length >= 8
		});
	}
}
