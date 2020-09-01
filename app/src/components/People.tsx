import React from 'react';
import '../css/People.css';

interface PeopleState {
	person: string
}

interface PeopleJSON {
	person: string
	people: string[]
	error: string
}

export default class People extends React.Component<{}, PeopleState> {
	constructor(props: any) {
		super(props);
		this.state = {
			person: ''
		};
	}

	async getJSON(url: string): Promise<PeopleJSON> {
		const res = await fetch(url);
		try {
			return await res.json();
		} catch (_) {
			return { person: '', people: [], error: '' };
		}
	}

	getPersonById(event: React.ChangeEvent<HTMLInputElement>) {
		this.getJSON(`/api/person/id/${event.target.value}`)
			.then(data => {
				this.setState({ person: data.person });
			});
	}

	getRandomPerson() {
		this.getJSON('/api/person/random')
			.then(data => {
				this.setState({ person: data.person });
			});
	}

	getAllPeople() {
		this.getJSON('/api/people')
			.then(data => {
				this.setState({ person: data.people.map(person => `"${person}"`).join(', ') });
			});
	}

	render() {
		return (
			<div className="People">
				<div className="form-group">
					<label htmlFor="person-id">Person ID</label>
					<input id="person-id"
						   className="form-control"
						   type="number"
						   onChange={(event) => this.getPersonById(event)} />
				</div>
				<div className="form-group">
					<button className="btn btn-primary"
							type="button"
							onClick={() => this.getRandomPerson()}>
						Random person
					</button>
				</div>
				<div className="form-group">
					<button className="btn btn-primary"
							type="button"
							onClick={() => this.getAllPeople()}>
						All people
					</button>
				</div>
				<div className="Person">{this.state.person}</div>
			</div>
		);
	}
}
