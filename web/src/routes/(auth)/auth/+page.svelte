<script lang="ts">
	import AuthForm from '$lib/components/auth/AuthForm.svelte';
	import SignupForm from '$lib/components/auth/SignupForm.svelte';
	import Button from '$lib/components/shared/Button.svelte';
	import { _, type TranslationKey } from '$lib/i18n/index.svelte';

	type FunctionState = {
		ctaKey: TranslationKey;
		buttonKey: TranslationKey;
	};

	// $state because of https://svelte.dev/docs/svelte/runtime-warnings#Client-warnings-state_proxy_equality_mismatch
	const authFunction: FunctionState = $state({
		ctaKey: 'auth.signup_cta',
		buttonKey: 'auth.signup'
	});

	// $state because of https://svelte.dev/docs/svelte/runtime-warnings#Client-warnings-state_proxy_equality_mismatch
	const signupFunction: FunctionState = $state({
		ctaKey: 'auth.signin_cta',
		buttonKey: 'auth.signin'
	});

	let currentFunction = $state(authFunction);

	let isAuthCurrent = $derived(currentFunction == authFunction);

	function switchFunction() {
		if (isAuthCurrent) currentFunction = signupFunction;
		else currentFunction = authFunction;
	}
</script>

<div class="panel mx-auto w-11/12 self-center sm:w-8/12 md:w-6/12 lg:w-1/4">
	{#if isAuthCurrent}
		<AuthForm />
	{:else}
		<SignupForm />
	{/if}
	<div class="mx-auto mt-4 flex w-fit flex-row items-center gap-1 text-sm text-foreground/60">
		{_(currentFunction.ctaKey)}
		<Button onclick={switchFunction} raw class="underline">{_(currentFunction.buttonKey)}</Button>
	</div>
</div>
