import React from 'react';
import '../css/NotFound.css';

export default class NotFound extends React.Component {
	public render() {
		return (
			<div>
				<h1 className="ErrorHeader">Page Not Found</h1>
				<p className="ErrorMessage">The page you are looking for does not exist. <a href="/">Click here</a> to go home.</p>
			</div>
		);
	}
}
