import type { ApiError } from './types';

type RouteData = {
	path: string;
	method: 'GET' | 'POST' | 'PUT' | 'DELETE';
};

const apiRoutes = {
	'user.auth': { path: '/v1/auth', method: 'POST' },
	'user.signup': { path: '/v1/users', method: 'POST' },
	'user.current': { path: '/v1/users/me', method: 'GET' },
	'user.logout': { path: '/v1/auth/logout', method: 'POST' },
	'books.list': { path: '/v1/books', method: 'GET' },
	'books.create': { path: '/v1/books', method: 'POST' },
	'books.get': { path: '/v1/books/{id}', method: 'GET' },
	'books.delete': { path: '/v1/books/{id}', method: 'DELETE' },
	'books.cover.update': { path: '/v1/books/{id}/cover', method: 'PUT' }
} as const satisfies Record<string, RouteData>;

export type Route = keyof typeof apiRoutes;

export function apiRoute(route: Route, params: Record<string, unknown> = {}): RouteData {
	return {
		path: (apiRoutes[route].path as string).replace(/{([a-zA-Z0-9_-]+)}/g, (_, key) =>
			key in params ? String(params[key]) : 'undefined'
		),
		method: apiRoutes[route].method
	};
}

// TODO: make type safe
export async function api<T>(
	route: RouteData,
	options: RequestInit = {},
	params?: Record<string, unknown>
): Promise<T | ApiError> {
	const method = route.method;

	const urlParams =
		params === undefined
			? ''
			: `?${Object.entries(params)
					.map((entry) => `${entry[0]}=${entry[1]}`)
					.join('&')}`;

	const headers = new Headers(options.headers);

	if (!headers.has('Content-Type') && !(options.body instanceof FormData)) {
		headers.set('Content-Type', 'application/json');
	}

	const response = await fetch(`/api${route.path}${urlParams}`, {
		...options,
		method: method,
		body: method === 'GET' ? undefined : options.body,
		headers
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
