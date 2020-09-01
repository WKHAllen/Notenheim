import React from 'react';
import { Link } from 'react-router-dom';
import '../css/Header.css';

export default class Header extends React.Component {
	render() {
		return (
			<div className="Header">
				<Link to="/">
					<img src="logo192.png" alt="logo" width="48" height="48" />
					<h4 className="Home-Label">PGTR App</h4>
				</Link>
			</div>
		);
	}
}
