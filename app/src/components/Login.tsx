import React, { useEffect } from 'react';
import { Link } from 'react-router-dom';
import '../css/Login.css';
import { requestAPIForm } from '../requestAPI';
import { showAPIError, hideAPIError } from '../apiError';

interface LoginState {
	submitClicked: boolean
}

export default class Login extends React.Component<any, LoginState> {
	constructor(props: any) {
		super(props);

		this.state = {
			submitClicked: false
		};
	}

	public render() {
		return (
			<div className="Login">
				<h1>Log In</h1>
				<form onSubmit={event => { this.login(event); return false; }} className="mb-3">
					<div className="form-group">
						<label htmlFor="email">Email</label>
						<input type="email" className="form-control" id="email" name="email" />
					</div>
					<div className="form-group">
						<label htmlFor="password">Password</label>
						<input type="password" className="form-control" id="password" name="password" />
					</div>
					<button type="submit" className="btn btn-primary btn-pink" disabled={this.state.submitClicked}>Register</button>
				</form>
				<small>If you do not have an account, please <Link to="/register">register here</Link>.</small>
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
			this.props.history.push('/');
			window.location.reload();
		} else {
			this.setState({
				submitClicked: false
			});
			showAPIError(res.error);
		}
	}
}
