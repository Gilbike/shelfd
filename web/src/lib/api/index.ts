import type { ApiError } from './types';

type RouteData = {
	path: string;
	method: 'GET' | 'POST' | 'PUT' | 'DELETE';
};

const apiRoutes: { [k: string]: RouteData } = {
	'user.auth': { path: '/v1/auth', method: 'POST' }
};

export type Route = keyof typeof apiRoutes;

export async function api<T>(route: Route, options: RequestInit = {}): Promise<T | ApiError> {
	const method = apiRoutes[route].method;

	const response = await fetch(`/api${apiRoutes[route].path}`, {
		...options,
		method: method,
		body: method === 'GET' ? undefined : options.body,
		headers: {
			'Content-Type': 'application/json',
			...options.headers
		}
	});

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
