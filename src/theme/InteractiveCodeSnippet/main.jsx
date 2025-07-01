import React, { useState, useEffect, useLayoutEffect, useRef } from 'react';
import styles from './styles.module.css';

import Editor from 'react-simple-code-editor';
import Prism from 'prismjs';

import 'prismjs/components/prism-go';
import { usePrismTheme } from '@docusaurus/theme-common';


const LAMBDA_URL = 'https://il5jdhvn5li6mbvknix7sq7u7e0mqtwv.lambda-url.us-east-1.on.aws/';

function getInitialCode(initialCode, src) {
  if (src) {
    try {
      return require(`!raw-loader!@site/static/${src}`).default;
    } catch (e) {
      console.error(`Falha ao carregar o arquivo durante o build: ${src}`, e);
      return `Erro ao carregar ${src}.`;
    }
  }
  return initialCode || '';
}

export default function InteractiveCodeSnippet({
  initialCode,
  src,
  fileName = "main.go",
  allowExecute = true,
  allowEdit = true,
}) {
  const [code, setCode] = useState(() => getInitialCode(initialCode, src));
  const prismTheme = usePrismTheme();
  const textareaRef = useRef(null);
  
  const [output, setOutput] = useState('');
  const [error, setError] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [isEditable, setIsEditable] = useState(false);
  const [copyText, setCopyText] = useState('⎘ Copiar');

  useLayoutEffect(() => {
    const textarea = textareaRef.current;
    if (!textarea) {
      return;
    }

    const resize = () => {
      requestAnimationFrame(() => {
        if (textarea.offsetHeight > 0) {
          textarea.style.height = 'auto';
          textarea.style.height = `${textarea.scrollHeight}px`;
        }
      });
    };

    resize();
    const observer = new ResizeObserver(resize);
    observer.observe(textarea);

    return () => {
      observer.disconnect();
    };
  }, [code]);
  
  const handleRun = async () => {
    setIsLoading(true);
    setOutput('');
    setError('');
    try {
      const response = await fetch(LAMBDA_URL, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ code: code }),
      });

      if (!response.ok) {
        const errorText = await response.text();
        throw new Error(`API Error (${response.status}): ${errorText}`);
      }

      const result = await response.json();
      if (result.error) setError(result.error);
      else setOutput(result.output);
    } catch (err) {
      setError(`Falha na comunicação com o ambiente de execução: ${err.message}`);
    } finally {
      setIsLoading(false);
    }
  };

  const handleCopy = () => {
    navigator.clipboard.writeText(code).then(() => {
      setCopyText('✓ Copiado!');
      setTimeout(() => {
        setCopyText('⎘ Copiar');
      }, 2000);
    }).catch(err => {
      console.error('Falha ao copiar código: ', err);
    });
  };

  const handleDownload = () => {
    const blob = new Blob([code], { type: 'text/plain;charset=utf-8' });
    const url = URL.createObjectURL(blob);
    const link = document.createElement('a');
    link.href = url;
    link.download = fileName;
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
    URL.revokeObjectURL(url);
  };

  return (
    <div className={styles.snippetContainer}>
      <div className={styles.codeWrapper}>
        <div className={styles.codeHeader}>Código Fonte (Go)</div>
        <Editor
          value={code}
          onValueChange={newCode => setCode(newCode)}
          highlight={code => Prism.highlight(code, Prism.languages.go, 'go')}
          padding={16}
          style={prismTheme}
          className={styles.codeEditor}
          readOnly={!allowEdit || !isEditable}
        />
      </div>

      <div className={styles.commandsWrapper}>
        {allowExecute && (
          <button onClick={handleRun} disabled={isLoading} className={styles.commandButton}>
            ▶ {isLoading ? 'Executando...' : 'Run'}
          </button>
        )}
        {allowEdit && (
          <button onClick={() => setIsEditable(!isEditable)} className={styles.commandButton}>
            {isEditable ? '✓ Pronto' : '✎ Editar'}
          </button>
        )}
        <button onClick={handleCopy} className={styles.commandButton}>
          {copyText}
        </button>
        <button onClick={handleDownload} className={styles.commandButton}>
          ⤓ Download
        </button>
      </div>
      
      {(isLoading || output || error) && (
        <div className={styles.previewWrapper}>
            <div className={styles.previewHeader}>Resultado</div>
            {isLoading && <div className={styles.loading}>Aguardando...</div>}
            {output && <pre className={styles.output}>{output}</pre>}
            {error && <pre className={styles.error}>{error}</pre>}
        </div>
      )}
    </div>
  );
}