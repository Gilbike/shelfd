import type { ApiError } from './types';

type RouteData = {
	path: string;
	method: 'GET' | 'POST' | 'PUT' | 'DELETE';
};

const apiRoutes = {
	'user.auth': { path: '/v1/auth', method: 'POST' },
	'user.current': { path: '/v1/users/me', method: 'GET' },
	'user.logout': { path: '/v1/auth/logout', method: 'POST' },
	'books.list': { path: '/v1/books', method: 'GET' }
} as const satisfies Record<string, RouteData>;

export type Route = keyof typeof apiRoutes;

// TODO: make type safe
export async function api<T>(
	route: Route,
	options: RequestInit = {},
	params?: Record<string, unknown>
): Promise<T | ApiError> {
	const method = apiRoutes[route].method;

	const urlParams =
		params === undefined
			? ''
			: `?${Object.entries(params)
					.map((entry) => `${entry[0]}=${entry[1]}`)
					.join('&')}`;

	const response = await fetch(`/api${apiRoutes[route].path}${urlParams}`, {
		...options,
		method: method,
		body: method === 'GET' ? undefined : options.body,
		headers: {
			'Content-Type': 'application/json',
			...options.headers
		}
	});

	// not the best one
	if (response.status == 204) {
		return undefined as T;
	}

	const contentType = response.headers.get('content-type');
	if (!contentType || !contentType.includes('application/json')) {
		throw new Error(`Expected JSON but received ${contentType || 'unknown type'}`);
	}

	const result = await response.json();
	if (!response.ok) {
		return { status: response.status, ...result } as Promise<ApiError>;
	}

	return result as Promise<T>;
}
