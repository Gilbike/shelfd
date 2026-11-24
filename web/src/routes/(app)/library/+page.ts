import { api, apiRoute } from '$lib/api';
import { handleApiError, isApiError } from '$lib/api/error';
import type { BookListForUserResponse } from '$lib/api/types';
import type { PageLoad } from './$types';

export const load: PageLoad = async () => {
	const books = await api<BookListForUserResponse>(apiRoute('user.books'));
	if (isApiError(books)) {
		handleApiError(books);
		return { books: [], metadata: { totalBooks: 0 } };
	}
	return {
		books: books.data,
		metadata: {
			totalBooks: books.total_books
		}
	};
};
