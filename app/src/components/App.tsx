import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import '../css/App.css';
import Header from './Header';
import Home          from './Home';
import Register      from './Register';
import Login         from './Login';
import Logout        from './Logout';
import Verify        from './Verify';
import PasswordReset from './PasswordReset';
import Profile       from './Profile';
import NewList       from './NewList';
import List          from './List';
import NotFound from './errors/NotFound';

export default class App extends React.Component {
	render() {
		return (
			<Router>
				<div className="App">
					<Header />
					<div className="App-Body">
						<Switch>
							<Route exact path="/"              component={Home} />
							<Route exact path="/register"      component={Register} />
							<Route exact path="/login"         component={Login} />
							<Route exact path="/logout"        component={Logout} />
							<Route       path="/verify"        component={Verify} />
							<Route       path="/resetPassword" component={PasswordReset} />
							<Route exact path="/profile"       component={Profile} />
							<Route exact path="/new"           component={NewList} />
							<Route       path="/list"          component={List} />
							<Route component={NotFound} />
						</Switch>
					</div>
				</div>
			</Router>
		);
	}
}
