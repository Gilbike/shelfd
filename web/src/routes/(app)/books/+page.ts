import { api } from '$lib/api';
import { handleApiError, isApiError } from '$lib/api/error';
import type { BookListResponse } from '$lib/api/types';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ url }) => {
	const page = parseInt(url.searchParams.get('page') || '1') || 1;
	const books = await api<BookListResponse>('books.list', undefined, { page: page });
	if (isApiError(books)) {
		handleApiError(books);
		return { books: [], metadata: { totalBooks: 0, totalPages: 0, page: 1, perPage: 0 } };
	}
	return {
		books: books.data,
		metadata: {
			totalBooks: books.total_books,
			totalPages: books.total_pages,
			page: books.page,
			perPage: books.per_page
		}
	};
};
