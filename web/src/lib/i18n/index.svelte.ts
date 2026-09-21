import * as en from './en.json';

export type Locales = 'en';

type JsonTree = { [key: string]: string | JsonTree };

type Leaves<T> = T extends object
	? {
			[
				K in keyof T
			]: `${Exclude<K, symbol>}${Leaves<T[K]> extends never ? '' : `.${Leaves<T[K]>}`}`;
		}[keyof T]
	: never;

export type TranslationKey = Leaves<typeof en>;
export type ErrorKeys = `errors.${Leaves<typeof en.errors>}`;

export const locale = $state<Locales>('en');
const translations: Record<Locales, JsonTree> = { en };

const paramRegexp = /{{([a-zA-Z0-9_]+)}}/g;

function resolveKey(obj: unknown, path: string): string | undefined {
	let current: unknown = obj;
	for (const segment of path.split('.')) {
		if (current == null || typeof current !== 'object') return undefined;
		current = (current as Record<string, unknown>)[segment];
	}
	return typeof current === 'string' ? current : undefined;
}

export function _(key: TranslationKey, params: Record<string, unknown> = {}): string {
	const language = translations[locale];
	let text = resolveKey(language, key);

	// fallback to english
	if (!text && locale != 'en') {
		text = resolveKey('en', key);
	}

	if (!text) {
		console.error(`missing key "${key}" for locale "${locale}"`);
		return key;
	}

	// replace params
	for (const match of text.matchAll(paramRegexp)) {
		const value = match[1] in params ? params[match[1]] : '[no key]';
		text = text.replaceAll(match[0], value as string);
	}

	return text;
}
