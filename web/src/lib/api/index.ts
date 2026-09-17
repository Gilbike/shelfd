import type { ApiError } from './types';

const apiRoutes = {
	'user.auth': '/v1/hello'
};

export type Route = keyof typeof apiRoutes;

export async function api<T>(route: Route, options: RequestInit = {}): Promise<T | ApiError> {
	const response = await fetch(`/api${apiRoutes[route]}`, {
		...options,
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
