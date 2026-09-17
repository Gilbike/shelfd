import type { ApiError } from './types';

export const ERROR_CODES = {
	INVALID_CREDENTIALS: 'INVALID_CREDENTIALS'
};

export function isApiError(obj: unknown): obj is ApiError {
	return (
		typeof obj === 'object' && obj !== null && 'status' in obj && 'code' in obj && 'message' in obj
	);
}
