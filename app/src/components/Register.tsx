import React from 'react';
import '../css/Register.css';

interface RegisterState {
	formGood: boolean
}

export default class Register extends React.Component<any, RegisterState> {
	constructor(props: any) {
		super(props);

		this.state = {
			formGood: false
		};
	}

	public render() {
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
						<input type="text" className="form-control" id="password" name="password" onChange={() => this.checkPasswords()} />
					</div>
					<div className="form-group">
						<label htmlFor="confirm-password">Confirm password</label>
						<input type="text" className="form-control" id="confirm-password" name="confirm-password" onChange={() => this.checkPasswords()} />
					</div>
					<button type="submit" className="btn btn-primary btn-pink" id="register-button" disabled={!this.state.formGood}>Register</button>
				</form>
			</div>
		);
	}

	private async register(e: React.FormEvent<HTMLFormElement>): Promise<void> {
		e.preventDefault();

		const data = new FormData(e.currentTarget);
		fetch('/api/register?' + new URLSearchParams({
			email: data.get('email') as string,
			password: data.get('password') as string
		}))
			.then(res => res.json())
			.then(console.log);
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
