import React from 'react';
import '../css/Register.css';
import { requestAPIForm } from '../requestAPI';
import { showAPIError, hideAPIError } from '../apiError';

interface RegisterState {
	formGood: boolean,
	submitClicked: boolean,
	registerSuccess: boolean
}

export default class Register extends React.Component<any, RegisterState> {
	constructor(props: any) {
		super(props);

		this.state = {
			formGood: false,
			submitClicked: false,
			registerSuccess: false
		};
	}

	public render() {
		if (!this.state.registerSuccess) {
			return (
				<div className="Register">
					<h1>Register</h1>
					<form onSubmit={event => { this.register(event); return false; }}>
						<div className="form-group">
							<label htmlFor="email">Email</label>
							<input type="email" className="form-control" id="email" name="email" onChange={() => this.checkPasswords()} />
						</div>
						<div className="form-group">
							<label htmlFor="password">Password</label>
							<input type="password" className="form-control" id="password" name="password" onChange={() => this.checkPasswords()} />
						</div>
						<div className="form-group">
							<label htmlFor="confirm-password">Confirm password</label>
							<input type="password" className="form-control" id="confirm-password" name="confirm-password" onChange={() => this.checkPasswords()} />
						</div>
						<button type="submit" className="btn btn-primary btn-pink" id="register-button" disabled={!this.state.formGood || this.state.submitClicked}>Register</button>
					</form>
				</div>
			);
		} else {
			return (
				<div className="Register">
					<h1>Register</h1>
					<p>Congratulations! You have registered for Notenheim. In the next few minutes you will receive an email containing a link that will confirm your registration. All you need to do is click the link in the email and your account will be ready! Please do so soon, as the link will expire after one hour.</p>
				</div>
			)
		}
	}

	private async register(e: React.FormEvent<HTMLFormElement>): Promise<void> {
		e.preventDefault();

		this.setState({
			submitClicked: true
		});

		const formData = new FormData(e.currentTarget);
		const res = await requestAPIForm('/register', formData, ['email', 'password']);

		if (res.error === null) {
			hideAPIError();
			this.setState({
				registerSuccess: true
			});
		} else {
			this.setState({
				submitClicked: false
			});
			showAPIError(res.error);
		}
	}

	private async checkPasswords(): Promise<void> {
		const email             = (document.getElementById('email')            as HTMLInputElement).value;
		const password          = (document.getElementById('password')         as HTMLInputElement).value;
		const confirmedPassword = (document.getElementById('confirm-password') as HTMLInputElement).value;

		this.setState({
			formGood: email.length > 0 && password === confirmedPassword && password.length >= 8
		});
	}
}
