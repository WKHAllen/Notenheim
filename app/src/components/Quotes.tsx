import React from 'react';
import '../css/Quotes.css';

interface QuotesState {
	quote: string
}

interface QuotesJSON {
	quote: string
	quotes: string[]
	error: string
}

export default class Quotes extends React.Component<{}, QuotesState> {
	constructor(props: any) {
		super(props);
		this.state = {
			quote: ''
		};
	}

	async getJSON(url: string): Promise<QuotesJSON> {
		const res = await fetch(url);
		try {
			return await res.json();
		} catch (_) {
			return { quote: '', quotes: [], error: '' };
		}
	}

	getQuoteById(event: React.ChangeEvent<HTMLInputElement>) {
		this.getJSON(`/api/quote/id/${event.target.value}`)
			.then(data => {
				this.setState({ quote: data.quote });
			});
	}

	getRandomQuote() {
		this.getJSON('/api/quote/random')
			.then(data => {
				this.setState({ quote: data.quote });
			});
	}

	getAllQuotes() {
		this.getJSON('/api/quotes')
			.then(data => {
				this.setState({ quote: data.quotes.map(quote => `"${quote}"`).join(', ') });
			});
	}

	render() {
		return (
			<div className="Quotes">
				<div className="form-group">
					<label htmlFor="quote-id">Quote ID</label>
					<input id="quote-id"
						   className="form-control"
						   type="number"
						   onChange={(event) => this.getQuoteById(event)} />
				</div>
				<div className="form-group">
					<button className="btn btn-primary"
							type="button"
							onClick={() => this.getRandomQuote()}>
						Random quote
					</button>
				</div>
				<div className="form-group">
					<button className="btn btn-primary"
							type="button"
							onClick={() => this.getAllQuotes()}>
						All quotes
					</button>
				</div>
				<div className="Quote">{this.state.quote}</div>
			</div>
		);
	}
}
