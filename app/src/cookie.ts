export function getCookie(cookieName: string): string | null {
	const name = cookieName + "=";
	const decodedCookie = decodeURIComponent(document.cookie);
	const splitCookie = decodedCookie.split(';');
	for (let i = 0; i < splitCookie.length; i++) {
		let c = splitCookie[i];
		while (c.charAt(0) === ' ') {
			c = c.substring(1);
		}
		if (c.indexOf(name) === 0) {
			return c.substring(name.length, c.length);
		}
	}
	return null;
}
