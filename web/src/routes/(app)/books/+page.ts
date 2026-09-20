import { api } from '$lib/api';
import { handleApiError, isApiError } from '$lib/api/error';
import type { BookListResponse } from '$lib/api/types';
import type { PageLoad } from './$types';

export const load: PageLoad = async () => {
	const books = await api<BookListResponse>('books.list');
	if (isApiError(books)) {
		handleApiError(books);
		return { books: [] };
	}
	return { books: books.data };
};
