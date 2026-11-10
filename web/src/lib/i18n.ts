import en, { type MessageKey } from "@/lib/locales/en";
import zh from "@/lib/locales/zh";
import tr from "@/lib/locales/tr";
import de from "@/lib/locales/de";
import fr from "@/lib/locales/fr";
import es from "@/lib/locales/es";
import it from "@/lib/locales/it";
import pt from "@/lib/locales/pt";
import pt_BR from "@/lib/locales/pt_BR";
import nl from "@/lib/locales/nl";
import pl from "@/lib/locales/pl";
import ru from "@/lib/locales/ru";
import uk from "@/lib/locales/uk";
import cs from "@/lib/locales/cs";
import hu from "@/lib/locales/hu";
import sv from "@/lib/locales/sv";
import ca from "@/lib/locales/ca";
import lt from "@/lib/locales/lt";
import ja from "@/lib/locales/ja";
import ko from "@/lib/locales/ko";
import zh_TW from "@/lib/locales/zh_TW";

export type { MessageKey };

export type Locale =
  | "en"
  | "tr"
  | "de"
  | "fr"
  | "es"
  | "it"
  | "pt"
  | "pt_BR"
  | "nl"
  | "pl"
  | "ru"
  | "uk"
  | "cs"
  | "hu"
  | "sv"
  | "ca"
  | "lt"
  | "ja"
  | "ko"
  | "zh"
  | "zh_TW";

export const locales: { value: Locale; label: string }[] = [
  { value: "en", label: "English" },
  { value: "tr", label: "Türkçe" },
  { value: "de", label: "Deutsch" },
  { value: "fr", label: "Français" },
  { value: "es", label: "Español" },
  { value: "it", label: "Italiano" },
  { value: "pt", label: "Português" },
  { value: "pt_BR", label: "Português (Brasil)" },
  { value: "nl", label: "Nederlands" },
  { value: "pl", label: "Polski" },
  { value: "ru", label: "Русский" },
  { value: "uk", label: "Українська" },
  { value: "cs", label: "Čeština" },
  { value: "hu", label: "Magyar" },
  { value: "sv", label: "Svenska" },
  { value: "ca", label: "Català" },
  { value: "lt", label: "Lietuvių" },
  { value: "ja", label: "日本語" },
  { value: "ko", label: "한국어" },
  { value: "zh", label: "简体中文" },
  { value: "zh_TW", label: "繁體中文" },
];

const dictionaries: Record<Locale, Partial<Record<MessageKey, string>>> = {
  en,
  tr,
  de,
  fr,
  es,
  it,
  pt,
  pt_BR,
  nl,
  pl,
  ru,
  uk,
  cs,
  hu,
  sv,
  ca,
  lt,
  ja,
  ko,
  zh,
  zh_TW,
};

export function normalizeLocale(value?: string | null): Locale | null {
  if (!value) {
    return null;
  }
  const lower = value.trim().toLowerCase().replace("-", "_");
  if (lower === "zh_cn" || lower === "zh_hans") return "zh";
  if (lower === "zh_tw" || lower === "zh_hant" || lower === "zh_hk") return "zh_TW";
  if (lower === "pt_br") return "pt_BR";

  const direct = locales.find((l) => l.value.toLowerCase() === lower);
  if (direct) return direct.value;

  const base = lower.split("_")[0];
  const baseHit = locales.find((l) => l.value.toLowerCase() === base);
  if (baseHit) return baseHit.value;

  return null;
}

export type Translate = (key: MessageKey, vars?: Record<string, string | number>) => string;

export function makeT(locale: Locale): Translate {
  const dict = dictionaries[locale] ?? en;
  return (key, vars) => {
    let text: string = dict[key] ?? en[key] ?? key;
    if (vars) {
      for (const [name, value] of Object.entries(vars)) {
        text = text.replaceAll(`{${name}}`, String(value));
      }
    }
    return text;
  };
}
