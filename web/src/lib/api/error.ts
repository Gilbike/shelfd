import { redirect } from '@sveltejs/kit';
import type { ApiError } from './types';

export const ERROR_CODES = {
	INVALID_CREDENTIALS: 'INVALID_CREDENTIALS',
	UNAUTHENTICATED: 'UNAUTHENTICATED'
};

export function isApiError(obj: unknown): obj is ApiError {
	return (
		typeof obj === 'object' && obj !== null && 'status' in obj && 'code' in obj && 'message' in obj
	);
}

// TODO: improve usability of the function as it may cause redirect and side effects
export function handleApiError(err: ApiError) {
	if (err.code === ERROR_CODES.UNAUTHENTICATED && err.status === 401) {
		redirect(307, '/auth');
	}
}
