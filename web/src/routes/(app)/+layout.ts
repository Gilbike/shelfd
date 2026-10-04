import { api, apiRoute } from '$lib/api';
import { isApiError } from '$lib/api/error';
import type { UserCurrentResponse } from '$lib/api/types';
import { redirect } from '@sveltejs/kit';
import type { LayoutLoad } from './$types';

export const load: LayoutLoad = async () => {
	const user = await api<UserCurrentResponse>(apiRoute('user.current'));
	if (isApiError(user)) {
		redirect(307, '/auth');
	}
	return { user };
};
