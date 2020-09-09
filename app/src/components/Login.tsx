import React, { useEffect } from 'react';
import { Link } from 'react-router-dom';
import '../css/Login.css';
import { requestAPIForm } from '../requestAPI';
import { showAPIError, hideAPIError } from '../apiError';

interface LoginState {
	formGood: boolean,
	submitClicked: boolean
}

export default class Login extends React.Component<any, LoginState> {
	constructor(props: any) {
		super(props);

		this.state = {
			formGood: false,
			submitClicked: false
		};
	}

	public render() {
		return (
			<div className="Login">
				<h1 className="mb-3">Login</h1>
				<form onSubmit={event => { this.login(event); return false; }} className="mb-3">
					<div className="form-group">
						<label htmlFor="email">Email</label>
						<input type="email" className="form-control" id="email" name="email" maxLength={63} onChange={() => this.checkForm()} />
					</div>
					<div className="form-group">
						<label htmlFor="password">Password</label>
						<input type="password" className="form-control" id="password" name="password" maxLength={255} onChange={() => this.checkForm()} />
					</div>
					<button type="submit" className="btn btn-primary btn-pink" disabled={!this.state.formGood || this.state.submitClicked}>Login</button>
				</form>
				<small>If you do not have an account, please <Link to="/register">register here</Link>.</small>
				<br />
				<small>Forgot your password? <Link to="/resetPassword">Reset it here</Link>.</small>
			</div>
		);
	}

	private async login(e: React.FormEvent<HTMLFormElement>): Promise<void> {
		e.preventDefault();

		this.setState({
			submitClicked: true
		});

		const formData = new FormData(e.currentTarget);
		const res = await requestAPIForm('/login', formData, ['email', 'password']);

		if (res.error === null) {
			hideAPIError();
			const after = new URL(window.location.href).searchParams.get('after');
			if (after !== null) {
				this.props.history.push(after);
			} else {
				this.props.history.push('/');
			}
			window.location.reload();
		} else {
			this.setState({
				submitClicked: false
			});
			showAPIError(res.error);
		}
	}

	private async checkForm(): Promise<void> {
		const email    = (document.getElementById('email')    as HTMLInputElement).value;
		const password = (document.getElementById('password') as HTMLInputElement).value;

		this.setState({
			formGood: email.match(/^([A-Za-z0-9.]+)@([a-z0-9]+)\.([a-z]+)$/g) !== null && password.length >= 8
		});
	}
}
