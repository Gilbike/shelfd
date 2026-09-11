import * as en from './en.json';

export type Locales = 'en';
type Translation = Record<string, Record<string, string> | string>;

export const locale = $state<Locales>('en');

const translations: Record<Locales, Translation> = { en };

function translate(locale: Locales, key: string) {
	const levels = key.split('.');

	let currentLevel = translations[locale];
	let text: string | null = null;
	for (const level of levels) {
		if (typeof currentLevel[level] === 'string') {
			text = currentLevel[level];
		} else {
			currentLevel = currentLevel[level];
		}
	}

	if (!text) throw new Error(`no translation found for ${locale}.${key}`);

	return text;
}

const t = $derived((key: string) => {
	return translate(locale, key);
});

export const _ = (key: string) => t(key);
