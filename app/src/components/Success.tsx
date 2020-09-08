import React from 'react';
import '../css/Success.css';
import { hideAPISuccess } from '../apiSuccess';

interface SuccessProps {
	message?: string
}

export default class Success extends React.Component<SuccessProps> {
	public render() {
		return (
			<div id="success" className="alert alert-primary alert-dismissible fade show Success hidden" role="alert">
				Success: <span id="success-message">{this.props.message}</span>
				<button type="button" className="close" data-dismiss="alert" aria-label="Close" onClick={hideAPISuccess}>
					<span aria-hidden="true">&times;</span>
				</button>
			</div>
		);
	}
}
