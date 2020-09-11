import React from 'react';
import '../css/PasswordResetRequest.css';
import { requestAPIForm } from '../requestAPI';
import { hideAPIError, showAPIError } from '../apiError';

interface PasswordResetRequestState {
	formGood: boolean,
	submitClicked: boolean,
	resetRequestSuccess: boolean
}

export default class PasswordResetRequest extends React.Component<any, PasswordResetRequestState> {
	constructor(props: any) {
		super(props);

		this.state = {
			formGood: false,
			submitClicked: false,
			resetRequestSuccess: false
		};
	}

	public render() {
		if (!this.state.resetRequestSuccess) {
			return (
				<div className="PasswordResetRequest">
					<h1 className="mb-3">Password Reset</h1>
					<p>If you have forgotten your password, simply enter your email address below. We will send you an email with a link to let you reset your password. Please use the link quickly, as it will expire after one hour. Do not share your link with anyone, or they may be able to gain access to your account!</p>
					<form onSubmit={event => { this.requestPasswordReset(event); return false; }} className="mb-3">
						<div className="form-group">
							<label htmlFor="email">Email</label>
							<input type="email" className="form-control" id="email" name="email" maxLength={63} onChange={() => this.checkForm()} autoFocus />
						</div>
						<button type="submit" className="btn btn-pink" disabled={!this.state.formGood || this.state.submitClicked}>Reset password</button>
					</form>
				</div>
			);
		} else {
			return (
				<div className="PasswordResetRequest">
					<h1 className="mb-3">Password Reset</h1>
					<p>You will receive an email in the next few minutes containing a link that will allow you to reset your password. Please use the link quickly, as it will expire after one hour. Do not share your link with anyone, or they may be able to gain access to your account!</p>
				</div>
			);
		}
	}

	private async requestPasswordReset(e: React.FormEvent<HTMLFormElement>): Promise<void> {
		e.preventDefault();

		this.setState({
			submitClicked: true
		});

		const formData = new FormData(e.currentTarget);
		const res = await requestAPIForm('/requestPasswordReset', formData, ['email']);

		if (res.error === null) {
			hideAPIError();
			this.setState({
				resetRequestSuccess: true
			});
		} else {
			this.setState({
				submitClicked: false
			});
			showAPIError(res.error);
		}
	}

	private async checkForm(): Promise<void> {
		const email = (document.getElementById('email') as HTMLInputElement).value;

		this.setState({
			formGood: email.match(/^([A-Za-z0-9.]+)@([a-z0-9]+)\.([a-z]+)$/g) !== null
		});
	}
}
