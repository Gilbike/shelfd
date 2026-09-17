import type { ApiError } from './types';

const apiRoutes = {
	'user.auth': '/v1/auth/'
};

type Route = keyof typeof apiRoutes;

export async function api<T>(route: Route): Promise<T | ApiError> {
	const response = await fetch(`/api/${apiRoutes[route]}`, {
		headers: { 'Content-Type': 'application/json' }
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
