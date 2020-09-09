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

export function setCookie(cookieName: string, value: string, expireDays: number = 1): void {
	const d = new Date();
	d.setTime(d.getTime() + (expireDays * 24 * 60 * 60 * 1000));
	const expires = "expires="+ d.toUTCString();
	document.cookie = cookieName + "=" + value + ";" + expires + ";path=/";
}

export function deleteCookie(cookieName: string): void {
	document.cookie = cookieName + "=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
}
