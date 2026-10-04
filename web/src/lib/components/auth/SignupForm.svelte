<script lang="ts">
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { api, apiRoute } from '$lib/api';
	import { ERROR_CODES, isApiError } from '$lib/api/error';
	import type { ApiError, ApiValidationError } from '$lib/api/types';
	import { enhance } from '$lib/enhance';
	import { _, type ErrorKeys } from '$lib/i18n/index.svelte';
	import Button from '../shared/Button.svelte';
	import FormInput from '../shared/FormInput.svelte';

	let username = $state('');
	let password = $state('');

	let errors: Record<string, ErrorKeys[]> = $state({});

	function handleError(error: ApiError) {
		if (error.code === ERROR_CODES.INVALID_INPUT) {
			errors = Object.fromEntries(
				Object.entries(error.details!).map((err) => [
					err[0],
					(err[1] as ApiValidationError[]).map((e) => e.code)
				])
			);
		}
	}

	async function handleSuccess() {
		const result = await api(apiRoute('user.auth'), {
			body: JSON.stringify({ username, password })
		});
		if (isApiError(result)) {
			// TODO: implement feedback
			return;
		}
		goto(resolve('/(app)/books'));
	}
</script>

<h1 class="font-semibold">{_('auth.signup')}</h1>
<form
	use:enhance={{
		route: 'user.signup',
		onError: handleError,
		onSuccess: () => handleSuccess()
	}}
	class="flex flex-col gap-1"
>
	<FormInput
		id="username"
		name="username"
		label={_('auth.username')}
		type="text"
		placeholder={_('auth.username')}
		autocomplete="username"
		required
		aria-required
		errors={errors['username']}
		bind:value={username}
	/>
	<FormInput
		id="display_name"
		name="display_name"
		label={_('auth.display_name')}
		type="text"
		placeholder={_('auth.display_name')}
		required
		aria-required
		errors={errors['display_name']}
	/>
	<FormInput
		id="password"
		name="password"
		label={_('auth.password')}
		type="password"
		placeholder={_('auth.password')}
		autocomplete="current-password"
		required
		aria-required
		errors={errors['password']}
		bind:value={password}
	/>
	<Button type="submit">{_('auth.signup')}</Button>
</form>
