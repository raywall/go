import type {ReactNode} from 'react';
import clsx from 'clsx';
import Heading from '@theme/Heading';
import styles from './styles.module.css';
import Translate, {translate} from '@docusaurus/Translate';

type FeatureItem = {
  title: string;
  Svg: React.ComponentType<React.ComponentProps<'svg'>>;
  description: ReactNode;
};

const FeatureList: FeatureItem[] = [
  {
    title: translate({
      id: 'homepage.feature.simple.title',
      message: 'Simples',
      description: 'Title for the Simple feature on the homepage',
    }),
    Svg: require('@site/static/img/gopher-simple-sm.svg').default,
    description: (
      <Translate
        id="homepage.feature.simple.description"
        description="Description for the Simple feature on the homepage">
        Sintaxe minimalista e intuitiva, projetada para ser extremamente fácil de
        aprender e ler. Escreva código limpo e direto ao ponto, aumentando sua
        produtividade desde o início.
      </Translate>
    ),
  },
  {
    title: translate({
      id: 'homepage.feature.fast.title',
      message: 'Rápido',
      description: 'Title for the Fast feature on the homepage',
    }),
    Svg: require('@site/static/img/gopher-flash-sm.svg').default,
    description: (
      <Translate
        id="homepage.feature.fast.description"
        description="Description for the Fast feature on the homepage">
        Possui alta performance e compila diretamente para código de máquina,
        entregando aplicações com inicialização quase instantânea e desempenho
        excepcional para qualquer tarefa.
      </Translate>
    ),
  },
  {
    title: translate({
      id: 'homepage.feature.powerful.title',
      message: 'Poderoso',
      description: 'Title for the Powerful feature on the homepage',
    }),
    Svg: require('@site/static/img/gopher-strong-sm.svg').default,
    description: (
      <Translate
        id="homepage.feature.powerful.description"
        description="Description for the Powerful feature on the homepage">
        Possui concorrência nativa. Com goroutines e channels, você pode criar
        sistemas complexos e eficientes que aproveitam ao máximo o hardware
        moderno de forma descomplicada.
      </Translate>
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