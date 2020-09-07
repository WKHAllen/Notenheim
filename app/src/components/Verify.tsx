import React from 'react';
import { Link } from 'react-router-dom';
import '../css/Verify.css';
import { requestAPI } from '../requestAPI';

interface VerifyState {
	verifyID: string,
	verificationSuccess: boolean | null,
	errorMessage?: string | null
}

export default class Verify extends React.Component<any, VerifyState> {
	constructor(props: any) {
		super(props);

		this.state = {
			verifyID: this.props.match.params.verifyID,
			verificationSuccess: null
		};
	}

	public componentDidMount() {
		requestAPI('/verify', {
			verifyID: this.state.verifyID
		}).then(res => {
			this.setState({
				verificationSuccess: res.error === null,
				errorMessage: res.error
			});
		});
	}

	public render() {
		if (this.state.verificationSuccess === null) {
			return (
				<div className="Verify">
					<h1>Verify</h1>
					<p className="loading">Verifying your account...</p>
				</div>
			);
		} else if (this.state.verificationSuccess) {
			return (
				<div className="Verify">
					<h1>Verify</h1>
					<p>Success! Your account has been verified. You may now proceed to <Link to="/login">the login page</Link>. You may also delete the email we sent you, as it will serve no further purpose.</p>
				</div>
			);
		} else {
			return (
				<div className="Verify">
					<h1>Verify</h1>
					<p>We were unable to verify your account: <code>{this.state.errorMessage}</code>. This may be because you have already verified your email. It may also be because the link expired, in which case you should <Link to="/register">register again</Link>.</p>
				</div>
			)
		}
	}
}
