import { api, apiRoute } from '$lib/api';
import { isApiError } from '$lib/api/error';
import type { Book } from '$lib/api/types';
import { redirect } from '@sveltejs/kit';
import type { LayoutLoad } from './$types';

export const load: LayoutLoad = async ({ params }) => {
	const book = await api<Book>(apiRoute('books.get', { id: params.bid }));

	if (isApiError(book)) {
		redirect(307, '/books');
	}

	return { book };
};
