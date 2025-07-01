import React from 'react';
import BrowserOnly from '@docusaurus/BrowserOnly';

import MainInteractiveCodeSnippet from './main';

export default function InteractiveCodeSnippetWrapper(props) {
  return (
    <BrowserOnly fallback={<div>Carregando editor interativo...</div>}>
      {() => {
        return <MainInteractiveCodeSnippet {...props} />;
      }}
    </BrowserOnly>
  );
}