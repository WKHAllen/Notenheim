import React from 'react';
import { Link } from 'react-router-dom';
import '../css/PasswordReset.css';
import { requestAPI, requestAPIForm } from '../requestAPI';
import { hideAPIError, showAPIError } from '../apiError';

interface PasswordResetState {
	resetID: string,
	validResetID: boolean | null,
	resetSuccess: boolean | null,
	submitClicked: boolean,
	formGood: boolean,
	errorMessage?: string | null
}

export default class PasswordReset extends React.Component<any, PasswordResetState> {
	constructor(props: any) {
		super(props);

		this.state = {
			resetID: this.props.match.params.resetID,
			validResetID: null,
			resetSuccess: null,
			submitClicked: false,
			formGood: false
		};
	}

	public componentDidMount() {
		requestAPI('/validPasswordResetID', {
			resetID: this.state.resetID
		}).then(res => {
			this.setState({
				validResetID: res.error === null,
				errorMessage: res.error
			});
		});
	}

	public render() {
		if (this.state.validResetID === null) {
			return (
				<div className="PasswordReset">
					<h1 className="mb-3">Password Reset</h1>
					<p className="loading">Checking your password reset request...</p>
				</div>
			);
		} else if (!this.state.validResetID) {
			return (
				<div className="PasswordReset">
					<h1 className="mb-3">Password Reset</h1>
					<p>We are unable to reset your password: <code>{this.state.errorMessage}</code>. This may be because the link we sent you expired, in which case you should <Link to="/resetPassword">try again</Link>.</p>
				</div>
			);
		} else if (!this.state.resetSuccess) {
			return (
				<div className="PasswordReset">
					<h1 className="mb-3">Password Reset</h1>
					<p>Please enter a new password below. You will only be able to reset your password while this link is valid. Do not share this link with anyone.</p>
					<form onSubmit={event => { this.resetPassword(event); return false; }} className="mb-3">
						<div className="form-group">
							<label htmlFor="password">New password</label>
							<input type="password" className="form-control" id="password" name="password" maxLength={255} onChange={() => this.checkForm()} autoFocus />
						</div>
						<div className="form-group">
							<label htmlFor="confirm-password">Confirm new password</label>
							<input type="password" className="form-control" id="confirm-password" name="confirm-password" maxLength={255} onChange={() => this.checkForm()} />
						</div>
						<button type="submit" className="btn btn-pink" disabled={!this.state.formGood || this.state.submitClicked}>Reset password</button>
					</form>
				</div>
			);
		} else {
			return (
				<div className="PasswordReset">
					<h1 className="mb-3">Password Reset</h1>
					<p>Success! Your password has been reset. You may now proceed to <Link to="/login">the login page</Link>. You may also delete the email we sent you, as it will serve no further purpose.</p>
				</div>
			);
		}
	}

	private async resetPassword(e: React.FormEvent<HTMLFormElement>): Promise<void> {
		e.preventDefault();

		this.setState({
			submitClicked: true
		});

		const formData = new FormData(e.currentTarget);
		const newPassword = formData.get('password');
		const res = await requestAPI('/resetPassword', {
			resetID: this.state.resetID,
			newPassword
		});

		if (res.error === null) {
			hideAPIError();
			this.setState({
				resetSuccess: true
			});
		} else {
			this.setState({
				submitClicked: false
			});
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
