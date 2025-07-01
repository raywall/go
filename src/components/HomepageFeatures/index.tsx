import type {ReactNode} from 'react';
import clsx from 'clsx';
import Heading from '@theme/Heading';
import styles from './styles.module.css';
// import Translate, {translate} from '@docusaurus/Translate';

type FeatureItem = {
  title: string;
  Svg: React.ComponentType<React.ComponentProps<'svg'>>;
  description: ReactNode;
};

const FeatureList: FeatureItem[] = [
  {
    title: 'Simples',
    Svg: require('@site/static/img/gopher-simple-sm.svg').default,
    description: (
      <>
        Sintaxe minimalista e intuitiva, projetada para ser extremamente f&aacute;cil
        de aprender e ler. Escreva c&oacute;digo limpo e direto ao ponto,
        aumentando sua produtividade desde o in&iacute;cio.
      </>
    ),
  },
  {
    title: 'Rápido',
    Svg: require('@site/static/img/gopher-flash-sm.svg').default,
    description: (
      <>
        Possui alta performance e compila diretamente para código de m&aacute;quina,
        entregando aplica&ccedil;&otilde;es com inicializa&ccedil;&atilde;o quase instant&acirc;nea e
        desempenho excepcional para qualquer tarefa.
      </>
    ),
  },
  {
    title: 'Poderoso',
    Svg: require('@site/static/img/gopher-strong-sm.svg').default,
    description: (
      <>
        Possui concorr&ecirc;ncia nativa. Com goroutines e channels, voc&ecirc; pode criar
        sistemas complexos e eficientes que aproveitam ao m&aacute;ximo o hardware
        moderno de forma descomplicada.
      </>
    ),
  },
];

function Feature({title, Svg, description}: FeatureItem) {
  return (
    <div className={clsx('col col--4')}>
      <div className="text--center">
        <Svg className={styles.featureSvg} role="img" />
      </div>
      <div className="text--center padding-horiz--md">
        <Heading as="h3">{title}</Heading>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures(): ReactNode {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}