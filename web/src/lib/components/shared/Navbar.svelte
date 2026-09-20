<script>
	import { goto } from '$app/navigation';
	import { resolve } from '$app/paths';
	import { page } from '$app/state';
	import { api } from '$lib/api';
	import { isApiError } from '$lib/api/error';
	import { _ } from '$lib/i18n/index.svelte';
	import { LogOut, User } from '@lucide/svelte';
	import { DropdownMenu } from 'bits-ui';
	import { fly } from 'svelte/transition';
	import Button from './Button.svelte';

	const initial = $derived(
		page.data.user.display_name
			.split(' ')
			.map((x) => x.at(0)?.toUpperCase())
			.join('')
			.slice(0, 2)
	);

	async function handleLogout() {
		const result = await api('user.logout');
		if (isApiError(result)) {
			console.error(`Failed to logout: ${result.code}`);
			return;
		}
		goto(resolve('/(auth)/auth'));
	}
</script>

<nav class="flex flex-col justify-between border-r border-r-border bg-surface p-4">
	<div></div>

	<DropdownMenu.Root>
		<DropdownMenu.Trigger class="interactive-secondary cursor-pointer rounded p-2">
			<User />
		</DropdownMenu.Trigger>

		<DropdownMenu.Content side="right" align="end" sideOffset={8} forceMount>
			{#snippet child({ wrapperProps, props, open })}
				{#if open}
					<div {...wrapperProps}>
						<div {...props} class="w-48" transition:fly={{ duration: 160, opacity: 0, x: -20 }}>
							<DropdownMenu.Item class="mb-2 flex flex-row items-center gap-1">
								<div class="flex h-8 w-8 items-center justify-center rounded-xl bg-primary">
									{initial}
								</div>
								<div class="font-bold">{page.data.user.display_name}</div>
							</DropdownMenu.Item>
							<DropdownMenu.Item onSelect={handleLogout}>
								{#snippet child({ props })}
									<Button {...props} secondary class="flex flex-row items-center gap-1 px-2">
										<LogOut size={16} />
										{_('auth.logout')}
									</Button>
								{/snippet}
							</DropdownMenu.Item>
						</div>
					</div>
				{/if}
			{/snippet}
		</DropdownMenu.Content>
	</DropdownMenu.Root>
</nav>
