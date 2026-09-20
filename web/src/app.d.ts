// See https://svelte.dev/docs/kit/types#app.d.ts

import type { User } from '$lib/api/types';

// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		// interface Locals {}
		interface PageData {
			user: User;
		}
		// interface PageState {}
		// interface Platform {}
	}
}

export {};
