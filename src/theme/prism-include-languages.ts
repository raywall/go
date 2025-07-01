import siteConfig from '@generated/docusaurus.config';

export default function prismIncludeLanguages(PrismObject) {
  const {
    themeConfig: { prism },
  } = siteConfig;
  const { additionalLanguages } = prism;

  // Importar Go manualmente
  require('prismjs/components/prism-go');

  globalThis.Prism = PrismObject;

  additionalLanguages.forEach((lang) => {
    if (lang === 'go') {
      // Go já foi importado acima
      return;
    }
    require(`prismjs/components/prism-${lang}`);
  });

  delete globalThis.Prism;
}